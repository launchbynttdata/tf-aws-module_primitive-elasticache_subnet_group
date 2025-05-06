# Basic

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.14 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.95.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_vpc"></a> [vpc](#module\_vpc) | terraform.registry.launch.nttdata.com/module_primitive/vpc/aws | ~> 1.0 |
| <a name="module_subnet"></a> [subnet](#module\_subnet) | terraform.registry.launch.nttdata.com/module_primitive/subnet/aws | ~> 1.0 |
| <a name="module_subnet_group"></a> [subnet\_group](#module\_subnet\_group) | ../.. | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_default_security_group.default](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_security_group) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_vpc_cidr_block"></a> [vpc\_cidr\_block](#input\_vpc\_cidr\_block) | The IPv4 CIDR block for the VPC. | `string` | n/a | yes |
| <a name="input_subnet_cidr_block"></a> [subnet\_cidr\_block](#input\_subnet\_cidr\_block) | The IPv4 CIDR block for the subnet. | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resource. | `map(string)` | n/a | yes |
| <a name="input_subnet_group_name"></a> [subnet\_group\_name](#input\_subnet\_group\_name) | Name for the cache subnet group. ElastiCache converts this name to lowercase. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_elasticache_subnet_group_name"></a> [elasticache\_subnet\_group\_name](#output\_elasticache\_subnet\_group\_name) | n/a |
| <a name="output_elasticache_subnet_vpc_id"></a> [elasticache\_subnet\_vpc\_id](#output\_elasticache\_subnet\_vpc\_id) | n/a |
| <a name="output_elasticache_subnet_subnets_ids"></a> [elasticache\_subnet\_subnets\_ids](#output\_elasticache\_subnet\_subnets\_ids) | n/a |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
