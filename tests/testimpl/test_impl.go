package testimpl

import (
	"context"
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestComposableComplete verifies the deployed AppConfig deployment strategy and exercises a reversible tag write.
func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	client, arn := verifyDeploymentStrategy(t, ctx)
	exerciseTagWrite(t, client, arn)
}

// TestComposableCompleteReadOnly verifies the deployed AppConfig deployment strategy using read-only AWS API calls.
func TestComposableCompleteReadOnly(t *testing.T, ctx types.TestContext) {
	verifyDeploymentStrategy(t, ctx)
}

func verifyDeploymentStrategy(t *testing.T, ctx types.TestContext) (*appconfig.Client, string) {
	opts := ctx.TerratestTerraformOptions()
	region := terraform.Output(t, opts, "region")
	id := terraform.Output(t, opts, "id")
	arn := terraform.Output(t, opts, "arn")
	name := terraform.Output(t, opts, "name")
	growthType := terraform.Output(t, opts, "growth_type")
	replicateTo := terraform.Output(t, opts, "replicate_to")
	deploymentDuration := int32Output(t, ctx, "deployment_duration_in_minutes")
	finalBakeTime := int32Output(t, ctx, "final_bake_time_in_minutes")
	growthFactor, err := strconv.ParseFloat(terraform.Output(t, opts, "growth_factor"), 32)
	require.NoError(t, err)

	require.NotEqual(t, "", id)
	assert.Equal(t, terraform.Output(t, opts, "expected_name"), name)
	assert.Equal(t, terraform.Output(t, opts, "expected_growth_type"), growthType)
	assert.Equal(t, terraform.Output(t, opts, "expected_replicate_to"), replicateTo)
	assert.Equal(t, int32Output(t, ctx, "expected_deployment_duration_in_minutes"), deploymentDuration)
	assert.Equal(t, int32Output(t, ctx, "expected_final_bake_time_in_minutes"), finalBakeTime)
	assert.InEpsilon(t, float32Output(t, ctx, "expected_growth_factor"), growthFactor, 0.001)

	client := appConfigClient(t, region)
	strategy, err := client.GetDeploymentStrategy(context.Background(), &appconfig.GetDeploymentStrategyInput{DeploymentStrategyId: aws.String(id)})
	require.NoError(t, err)

	assert.Equal(t, id, aws.ToString(strategy.Id))
	assert.Equal(t, name, aws.ToString(strategy.Name))
	assert.Equal(t, growthType, string(strategy.GrowthType))
	assert.Equal(t, replicateTo, string(strategy.ReplicateTo))
	assert.Equal(t, deploymentDuration, strategy.DeploymentDurationInMinutes)
	assert.Equal(t, finalBakeTime, strategy.FinalBakeTimeInMinutes)
	assert.InEpsilon(t, growthFactor, aws.ToFloat32(strategy.GrowthFactor), 0.001)

	return client, arn
}

func appConfigClient(t *testing.T, region string) *appconfig.Client {
	t.Helper()

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	require.NoError(t, err)

	return appconfig.NewFromConfig(cfg)
}

func exerciseTagWrite(t *testing.T, client *appconfig.Client, resourceARN string) {
	t.Helper()

	const tagKey = "codex-functional-test"
	_, err := client.TagResource(context.Background(), &appconfig.TagResourceInput{
		ResourceArn: aws.String(resourceARN),
		Tags:        map[string]string{tagKey: "true"},
	})
	require.NoError(t, err)

	_, err = client.UntagResource(context.Background(), &appconfig.UntagResourceInput{
		ResourceArn: aws.String(resourceARN),
		TagKeys:     []string{tagKey},
	})
	require.NoError(t, err)
}

func int32Output(t *testing.T, ctx types.TestContext, name string) int32 {
	t.Helper()

	value, err := strconv.ParseInt(terraform.Output(t, ctx.TerratestTerraformOptions(), name), 10, 32)
	require.NoError(t, err)

	return int32(value)
}

func float32Output(t *testing.T, ctx types.TestContext, name string) float32 {
	t.Helper()

	value, err := strconv.ParseFloat(terraform.Output(t, ctx.TerratestTerraformOptions(), name), 32)
	require.NoError(t, err)

	return float32(value)
}
