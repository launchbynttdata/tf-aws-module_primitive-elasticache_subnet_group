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

module "vpc" {
  source  = "terraform.registry.launch.nttdata.com/module_primitive/vpc/aws"
  version = "~> 1.0"

  cidr_block = var.vpc_cidr_block
}

// Fix for Regula FG_R00089: VPC default security group should restrict all traffic [Medium]
resource "aws_default_security_group" "default" {
  vpc_id = module.vpc.vpc_id

  ingress = []
  egress  = []
}

module "subnet" {
  source  = "terraform.registry.launch.nttdata.com/module_primitive/subnet/aws"
  version = "~> 1.0"

  vpc_id     = module.vpc.vpc_id
  cidr_block = var.subnet_cidr_block
  tags       = var.tags
}

module "subnet_group" {
  source = "../.."

  name       = var.subnet_group_name
  subnet_ids = [module.subnet.subnet_id]
}
