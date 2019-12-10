# terraform-provider-swis
Terraform provider (plugin) for SolarWings

Issues and Contributing
If you find an issue with this provider, please report an it. Contributions are welcome.

Run
When more than one IP resource is provided, the apply action must be run with param parallelism set to 1
  terraform apply -parallelism=1

Example terraform file:
  resource "swis_server" "ip1" {
    vsphere_vlan = var.vsphere_vlan
  }

  variable "vsphere_vlan" {
    default = "VLAN100_10.141.16.0m24"
  }

When you want to update status of existing ip address (does not matter which state it already has) specify two more params in the resouce:
  resource "swis_server" "ip1" {
    vsphere_vlan = var.vsphere_vlan
    ipaddress =  "10.141.16.11"
    status = "4"
  }
