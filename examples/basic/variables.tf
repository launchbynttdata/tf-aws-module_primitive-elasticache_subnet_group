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

variable "vpc_cidr_block" {
  type        = string
  description = "The IPv4 CIDR block for the VPC."
}

variable "subnet_cidr_block" {
  type        = string
  description = "The IPv4 CIDR block for the subnet."
}

variable "tags" {
  type        = map(string)
  description = "A map of tags to assign to the resource."
}

variable "subnet_group_name" {
  type        = string
  description = "Name for the cache subnet group. ElastiCache converts this name to lowercase."
}
