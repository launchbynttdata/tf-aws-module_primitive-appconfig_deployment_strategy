package testimpl

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestComposableComplete verifies the deployed AppConfig deployment strategy.
func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	verifyDeploymentStrategy(t, ctx)
}

// TestComposableCompleteReadOnly verifies the deployed AppConfig deployment strategy using read-only AWS API calls.
func TestComposableCompleteReadOnly(t *testing.T, ctx types.TestContext) {
	verifyDeploymentStrategy(t, ctx)
}

func verifyDeploymentStrategy(t *testing.T, ctx types.TestContext) {
	opts := ctx.TerratestTerraformOptions()
	region := terraform.Output(t, opts, "region")
	id := terraform.Output(t, opts, "id")
	name := terraform.Output(t, opts, "name")
	growthType := terraform.Output(t, opts, "growth_type")
	replicateTo := terraform.Output(t, opts, "replicate_to")

	require.NotEqual(t, "", id)
	assert.Equal(t, terraform.Output(t, opts, "expected_name"), name)
	assert.Equal(t, terraform.Output(t, opts, "expected_growth_type"), growthType)
	assert.Equal(t, terraform.Output(t, opts, "expected_replicate_to"), replicateTo)

	client := appConfigClient(t, region)
	strategy, err := client.GetDeploymentStrategy(context.Background(), &appconfig.GetDeploymentStrategyInput{DeploymentStrategyId: aws.String(id)})
	require.NoError(t, err)

	assert.Equal(t, id, aws.ToString(strategy.Id))
	assert.Equal(t, name, aws.ToString(strategy.Name))
	assert.Equal(t, growthType, string(strategy.GrowthType))
	assert.Equal(t, replicateTo, string(strategy.ReplicateTo))
}

func appConfigClient(t *testing.T, region string) *appconfig.Client {
	t.Helper()

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	require.NoError(t, err)

	return appconfig.NewFromConfig(cfg)
}
