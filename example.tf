resource "swis_server" "ip1" {
  vsphere_vlan = var.vsphere_vlan
  ipaddress = "10.141.16.14"
#  status = var.status
}


variable "vsphere_vlan" {
  default = "VLAN100_10.141.16.0m24"
}


#status and ipddress are required for Update action only

variable "status" {
  type = string
  default = "4"
}
