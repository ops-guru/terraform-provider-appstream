//updates not supported by service, only state (RUNNING, STOPPED) operations
resource "appstream_image_builder" "test-image-builder" {
    name                            = var.image_builder_name
    appstream_agent_version         = var.appstream_agent_version
    description                     = var.image_builder_description
    display_name                    = var.image_builder_display_name
    enable_default_internet_access  = var.enable_default_internet_access
    iam_role_arn                    = var.iam_role_arn
    image_name                      = var.image_name
    instance_type                   = var.instance_type
    vpc_config {
        security_group_ids          = var.security_group_ids
        subnet_ids                  = var.subnet_ids
    }
    state                           = var.state
}