# Complete Example

This example creates a complete AppConfig deployment strategy deployment with the dependencies required to exercise the primitive module.

## Usage

```hcl
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

data "aws_region" "current" {}


module "resource_names" {
  source  = "terraform.registry.launch.nttdata.com/module_library/resource_name/launch"
  version = "~> 2.0"

  for_each = var.resource_names_map

  logical_product_family  = var.logical_product_family
  logical_product_service = var.logical_product_service
  class_env               = var.class_env
  instance_env            = var.instance_env
  instance_resource       = var.instance_resource
  cloud_resource_type     = each.value.name
  maximum_length          = each.value.max_length

  region = join("", split("-", data.aws_region.current.region))
}

module "deployment_strategy" {
  source = "../.."

  name                           = module.resource_names["deployment_strategy"].standard
  deployment_duration_in_minutes = var.deployment_duration_in_minutes
  description                    = var.description
  final_bake_time_in_minutes     = var.final_bake_time_in_minutes
  growth_factor                  = var.growth_factor
  growth_type                    = var.growth_type
  replicate_to                   = var.replicate_to
  tags                           = var.tags
}
```

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.10 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.100, < 7.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_deployment_strategy"></a> [deployment\_strategy](#module\_deployment\_strategy) | ../.. | n/a |
| <a name="module_resource_names"></a> [resource\_names](#module\_resource\_names) | terraform.registry.launch.nttdata.com/module_library/resource_name/launch | ~> 2.0 |

## Resources

| Name | Type |
|------|------|
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_class_env"></a> [class\_env](#input\_class\_env) | Environment class for generated resource names. | `string` | n/a | yes |
| <a name="input_deployment_duration_in_minutes"></a> [deployment\_duration\_in\_minutes](#input\_deployment\_duration\_in\_minutes) | Deployment duration in minutes. | `number` | `0` | no |
| <a name="input_description"></a> [description](#input\_description) | Deployment strategy description. | `string` | `"Example all-at-once deployment strategy."` | no |
| <a name="input_final_bake_time_in_minutes"></a> [final\_bake\_time\_in\_minutes](#input\_final\_bake\_time\_in\_minutes) | Final bake time in minutes. | `number` | `0` | no |
| <a name="input_growth_factor"></a> [growth\_factor](#input\_growth\_factor) | Deployment growth factor. | `number` | `100` | no |
| <a name="input_growth_type"></a> [growth\_type](#input\_growth\_type) | Deployment growth type. | `string` | `"LINEAR"` | no |
| <a name="input_instance_env"></a> [instance\_env](#input\_instance\_env) | Environment instance number for generated resource names. | `number` | n/a | yes |
| <a name="input_instance_resource"></a> [instance\_resource](#input\_instance\_resource) | Resource instance number for generated resource names. | `number` | n/a | yes |
| <a name="input_logical_product_family"></a> [logical\_product\_family](#input\_logical\_product\_family) | Logical product family for generated resource names. | `string` | n/a | yes |
| <a name="input_logical_product_service"></a> [logical\_product\_service](#input\_logical\_product\_service) | Logical product service for generated resource names. | `string` | n/a | yes |
| <a name="input_replicate_to"></a> [replicate\_to](#input\_replicate\_to) | Deployment strategy replication destination. | `string` | `"NONE"` | no |
| <a name="input_resource_names_map"></a> [resource\_names\_map](#input\_resource\_names\_map) | Resource name configuration keyed by resource role. | <pre>map(object({<br/>    name       = string<br/>    max_length = number<br/>  }))</pre> | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | Map of tags to assign to resources. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the deployment strategy. |
| <a name="output_deployment_duration_in_minutes"></a> [deployment\_duration\_in\_minutes](#output\_deployment\_duration\_in\_minutes) | The deployment duration in minutes. |
| <a name="output_expected_deployment_duration_in_minutes"></a> [expected\_deployment\_duration\_in\_minutes](#output\_expected\_deployment\_duration\_in\_minutes) | Expected deployment\_duration\_in\_minutes. |
| <a name="output_expected_final_bake_time_in_minutes"></a> [expected\_final\_bake\_time\_in\_minutes](#output\_expected\_final\_bake\_time\_in\_minutes) | Expected final\_bake\_time\_in\_minutes. |
| <a name="output_expected_growth_factor"></a> [expected\_growth\_factor](#output\_expected\_growth\_factor) | Expected growth\_factor. |
| <a name="output_expected_growth_type"></a> [expected\_growth\_type](#output\_expected\_growth\_type) | Expected growth type. |
| <a name="output_expected_name"></a> [expected\_name](#output\_expected\_name) | Expected deployment strategy name. |
| <a name="output_expected_replicate_to"></a> [expected\_replicate\_to](#output\_expected\_replicate\_to) | Expected replication destination. |
| <a name="output_final_bake_time_in_minutes"></a> [final\_bake\_time\_in\_minutes](#output\_final\_bake\_time\_in\_minutes) | The final bake time in minutes. |
| <a name="output_growth_factor"></a> [growth\_factor](#output\_growth\_factor) | The growth factor. |
| <a name="output_growth_type"></a> [growth\_type](#output\_growth\_type) | The growth type. |
| <a name="output_id"></a> [id](#output\_id) | The deployment strategy ID. |
| <a name="output_name"></a> [name](#output\_name) | The name of the deployment strategy. |
| <a name="output_region"></a> [region](#output\_region) | The AWS Region where the example resources are deployed. |
| <a name="output_replicate_to"></a> [replicate\_to](#output\_replicate\_to) | The replication destination. |
<!-- END_TF_DOCS -->
