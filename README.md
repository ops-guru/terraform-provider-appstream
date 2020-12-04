# terraform-provider-appstream
Terraform AWS Appstream 2.0 custom provider

## Requirements

Tested with:

- terraform 0.12
- go 1.14

# Provider usage

Deploy custom appstream provider 
```
$ do get .
$ go build -o ~/.terraform.d/plugins/terraform-provider-appstream_v0.1.0
```

GOTO terraform module and run commands 

```
$ terraform init
$ terraform plan
$ terraform apply
```