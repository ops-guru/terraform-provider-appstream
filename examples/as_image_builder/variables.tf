variable "access_key" {
  description = "AWS access key"
}
variable "secret_key" {
  description = "AWS secret key"
}
variable "token" {
  default = ""
}
variable "region" {
  description = "AWS default region"
  default = "us-east-1"
}

variable "image_builder_name" {
  description = "Appstream 2.0 image builder name"
  default = "default-image-builder"
}

variable "appstream_agent_version" {
  description = "Apstream 2.0 image builder agent version"
  default = "LATEST"
}

variable "image_builder_description" {
  description = "Appstream 2.0 image builder description"
  default = "Default image builder"
}

variable "image_builder_display_name" {
  description = "Appstream 2.0 image builder display name"
  default = "default-image-builder"
}

variable "enable_default_internet_access" {
  description = "Enable internet access from image builder instance? (boolean)"
  default = true
}

variable "iam_role_arn" {
  description = "Appstream 2.0 image builder aim role to use"
  default = ""
}

variable "image_name" {
  description = "Appstream 2.0 image builder base image name"
  default = "Amazon-AppStream2-Sample-Image-02-04-2019"
}

variable "instance_type" {
  description = "Appstream 2.0 image builder instance type"
  default = "stream.standard.large"
}

variable "security_group_ids" {
  description = "Appstream 2.0 image builder security group"
}

variable "subnet_ids" {
  description = "Appstream 2.0 image builder subnet id"
}

variable "state" {
  description = "Appstream 2.0 image builder state (valid values: STOPPED, RUNNING)"
  default = "RUNNING"
}