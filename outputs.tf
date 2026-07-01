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
  value       = aws_appconfig_deployment_strategy.deployment_strategy.id
}

output "arn" {
  description = "The ARN of the deployment strategy."
  value       = aws_appconfig_deployment_strategy.deployment_strategy.arn
}

output "name" {
  description = "The name of the deployment strategy."
  value       = aws_appconfig_deployment_strategy.deployment_strategy.name
}

output "deployment_duration_in_minutes" {
  description = "The deployment duration in minutes."
  value       = aws_appconfig_deployment_strategy.deployment_strategy.deployment_duration_in_minutes
}

output "final_bake_time_in_minutes" {
  description = "The final bake time in minutes."
  value       = aws_appconfig_deployment_strategy.deployment_strategy.final_bake_time_in_minutes
}

output "growth_factor" {
  description = "The growth factor."
  value       = aws_appconfig_deployment_strategy.deployment_strategy.growth_factor
}

output "growth_type" {
  description = "The growth type."
  value       = aws_appconfig_deployment_strategy.deployment_strategy.growth_type
}

output "replicate_to" {
  description = "The replication destination."
  value       = aws_appconfig_deployment_strategy.deployment_strategy.replicate_to
}
