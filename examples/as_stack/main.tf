resource "appstream_stack" "appstream-stack" {
    name                    = var.stack_name
    description             = var.stack_description
    display_name            = var.stack_display_name
    feedback_url            = var.feedback_url
    redirect_url            = var.redirect_url
    storage_connectors {
        connector_type      = var.connector_type
    }
    tags                    = var.tags
}


resource "appstream_fleet" "appstream-fleet" {
    name                            = var.fleet_name
    stack_name                      = appstream_stack.appstream-stack.name
    compute_capacity {
        desired_instances           = var.desired_instances
    }
    description                     = var.fleet_description
    disconnect_timeout              = var.disconnect_timeout
    display_name                    = var.fleet_display_name
    enable_default_internet_access  = var.enable_default_internet_access
    fleet_type                      = var.fleet_type
    iam_role_arn                    = var.iam_role_arn
    idle_disconnect_timeout         = var.idle_disconnect_timeout
    image_arn                       = var.image_arn
    image_name                      = var.image_name
    instance_type                   = var.instance_type
    max_user_duration               = var.max_user_duration
    vpc_config {
        security_group_ids          = var.security_group_ids
        subnet_ids                  = var.subnet_ids
    }
    tags                            = var.tags
    state                           = var.state
}