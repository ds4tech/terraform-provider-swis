resource "swis_server" "ip1" {
  vsphere_vlan = var.vsphere_vlan
#  ipaddress = var.ipaddress
#  status = var.status
}


variable "vsphere_vlan" {
  default = "VLAN100_10.141.16.0m24"
}


#status and ipddress are required for Update action only

variable "ipaddress" {
  type = string
  default = "10.141.16.11"
}
variable "status" {
  type = string
  default = "4"
}
