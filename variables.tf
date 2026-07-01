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

# -----------------------------------------------------------------------------
# Required
# -----------------------------------------------------------------------------

variable "name" {
  description = "Name of the AppConfig deployment strategy. Must be 1 to 64 characters."
  type        = string

  validation {
    condition     = length(var.name) >= 1 && length(var.name) <= 64
    error_message = "Name must be between 1 and 64 characters."
  }
}

variable "deployment_duration_in_minutes" {
  description = "Total number of minutes for deployments using this strategy. Must be between 0 and 1440."
  type        = number
  validation {
    condition     = var.deployment_duration_in_minutes >= 0 && var.deployment_duration_in_minutes <= 1440
    error_message = "deployment_duration_in_minutes must be between 0 and 1440."
  }
}

variable "growth_factor" {
  description = "Percentage of targets to receive the configuration during each interval. Must be between 1.0 and 100.0."
  type        = number
  validation {
    condition     = var.growth_factor >= 1 && var.growth_factor <= 100
    error_message = "growth_factor must be between 1.0 and 100.0."
  }
}

variable "replicate_to" {
  description = "Where AWS AppConfig replicates the deployment strategy. Valid values are NONE or SSM_DOCUMENT."
  type        = string
  validation {
    condition     = contains(["NONE", "SSM_DOCUMENT"], var.replicate_to)
    error_message = "replicate_to must be NONE or SSM_DOCUMENT."
  }
}

# -----------------------------------------------------------------------------
# Optional
# -----------------------------------------------------------------------------

variable "description" {
  description = "Description of the AppConfig deployment strategy. Must be at most 1024 characters."
  type        = string
  default     = null

  validation {
    condition     = var.description == null ? true : length(var.description) <= 1024
    error_message = "Description must be at most 1024 characters."
  }
}

variable "final_bake_time_in_minutes" {
  description = "Number of minutes AWS AppConfig monitors alarms after deployment reaches 100 percent. Must be between 0 and 1440."
  type        = number
  default     = 0
  validation {
    condition     = var.final_bake_time_in_minutes >= 0 && var.final_bake_time_in_minutes <= 1440
    error_message = "final_bake_time_in_minutes must be between 0 and 1440."
  }
}

variable "growth_type" {
  description = "Growth algorithm for the deployment strategy. Valid values are LINEAR or EXPONENTIAL."
  type        = string
  default     = "LINEAR"
  validation {
    condition     = contains(["LINEAR", "EXPONENTIAL"], var.growth_type)
    error_message = "growth_type must be LINEAR or EXPONENTIAL."
  }
}

variable "region" {
  description = "AWS Region where this resource is managed. Defaults to the provider-configured Region."
  type        = string
  default     = null
}

variable "tags" {
  description = "Map of tags to assign to the resource. Up to 50 tags are allowed; tag keys must be 1 to 128 characters and values must be at most 256 characters."
  type        = map(string)
  default     = {}

  validation {
    condition = length(var.tags) <= 50 && alltrue([
      for key, value in var.tags : length(key) >= 1 && length(key) <= 128 && length(value) <= 256
    ])
    error_message = "Tags must contain at most 50 entries. Tag keys must be 1 to 128 characters and values must be at most 256 characters."
  }
}
