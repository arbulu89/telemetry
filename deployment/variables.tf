# AWS provider

variable "region" {
  type        = string
  description = "AWS region to use"
  default     = "eu-central-1"
}

# Influxdb variables

variable "influxdb_url" {
  type        = string
  description = "Influxdb database url (it doesn't containt the organization id). Example: https://influxdb.com"
}

variable "influxdb_token" {
  type        = string
  description = "Influxdb connection token. It must have bucket and api creation authorizations"
}

variable "influxdb_org" {
  type        = string
  description = "The organization id of the deployed Influxdb"
}

variable "influxdb_bucket" {
  type        = string
  description = "Created Influxdb bucket name. This bucket will have write/read authorizations through an API token"
  default     = "trento-telemetry"
}

# ECS variables

variable "name" {
  type        = string
  description = "Given prefix name of the deployed resources"
  default     = "trento-telemetry"
}

variable "environment" {
  type        = string
  description = "Given environment name to group the resources names"
  default     = "default"
}

variable "container_image" {
  type        = string
  description = "Deployed container name"
  default     = "ghcr.io/trento-project/telemetry:rolling"
}

variable "lb_certificate_arn" {
  type        = string
  description = "The TLS certificate ARN to use for the TLS termination in ELB"
}
