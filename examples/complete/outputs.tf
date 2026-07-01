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

output "id" {
  description = "The deployment strategy ID."
  value       = module.deployment_strategy.id
}
output "arn" {
  description = "The ARN of the deployment strategy."
  value       = module.deployment_strategy.arn
}
output "name" {
  description = "The name of the deployment strategy."
  value       = module.deployment_strategy.name
}
output "growth_type" {
  description = "The growth type."
  value       = module.deployment_strategy.growth_type
}
output "replicate_to" {
  description = "The replication destination."
  value       = module.deployment_strategy.replicate_to
}
output "expected_name" {
  description = "Expected deployment strategy name."
  value       = module.resource_names["deployment_strategy"].standard
}
output "expected_growth_type" {
  description = "Expected growth type."
  value       = var.growth_type
}
output "expected_replicate_to" {
  description = "Expected replication destination."
  value       = var.replicate_to
}

output "region" {
  description = "The AWS Region where the example resources are deployed."
  value       = data.aws_region.current.region
}
