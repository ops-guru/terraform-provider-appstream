variable "access_key" {
  description = "AWS access key id"
}
variable "secret_key" {
  description = "AWS secret key id"
}
variable "token" {
  description = "AWS token, used for temporary credentials"
  default = ""
}
variable "region" {
  description = "AWS region to use"
  default = "us-east-1"
}

variable "stack_name" {
  description = "Appstream 2.0 stack name"
  default = "default-stack"
}

variable "stack_description" {
  description = "Apstream 2.0 stack description"
  default = "Default appstream stack"
}

variable "stack_display_name" {
  description = "Appstream 2.0 stack display name"
  default = "Appstream default stack"
}

variable "feedback_url" {
  description = "Appstream feedback url"
  default = ""
}

variable "redirect_url" {
  description = "Appstream redirect url"
  default = ""
}

variable "connector_type" {
  description = "Appstream connector type, HOMEFOLDERS | GOOGLE_DRIVE | ONE_DRIVE"
  default = "HOMEFOLDERS"
}

variable "tags" {
  description = "Appstream stack and fleet tags"
  type = map(string)
  default = {
    Env = "dev",
    Role = "Appstream",
    Version = "v1"
  }
}

variable "fleet_name" {
  description = "Apstream fleet name"
  default = "default-fleet"
}

variable "desired_instances" {
  description = "Appstream desired number of instances"
  default = 1
}

variable "fleet_description" {
  description = "Appstream fleet description"
  default = "Default appstream fleet"
}

variable "disconnect_timeout" {
  description = "The amount of time that a streaming session remains active after users disconnect"
  default = 300
}

variable "fleet_display_name" {
  description = "Appstream fleet display name"
  default = "default-fleet"
}

variable "enable_default_internet_access" {
  description = "Enables or disables default internet access for the fleet."
  default = true
}

variable "fleet_type" {
  description = "The fleet type. Valid values ALWAYS_ON | ON_DEMAND"
  default = "ON_DEMAND"
}

variable "iam_role_arn" {
  description = "The Amazon Resource Name (ARN) of the IAM role to apply to the fleet"
  default = ""
}

variable "idle_disconnect_timeout" {
  description = "The amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the DisconnectTimeoutInSeconds time interval begins"
  default = 300
}

variable "image_arn" {
  description = "The ARN of the public, private, or shared image to use"
  default = ""
}

variable "image_name" {
  description = "The name of the image used to create the fleet"
  default = "Amazon-AppStream2-Sample-Image-02-04-2019"
}

variable "instance_type" {
  description = "The instance type to use when launching fleet instances"
  default = "stream.standard.large"
}

variable "max_user_duration" {
  description = "The maximum amount of time that a streaming session can remain active, in seconds"
  default = 600
}

variable "security_group_ids" {
  description = "Appstream fleet security group id to use"
}

variable "subnet_ids" {
  description = "Appstream fleet subnet id to run"
}

variable "state" {
  default = "RUNNING"
  description = "Appstream fleet state"
}