resource "swis_server" "orion" {
  hostname = var.orion_host
  ipaddress = var.ipaddress
  username = var.orion_username
  password = var.orion_userpass
  status = var.status
}


variable "orion_host" {
  default = "10.50.8.10"
}
variable "orion_username" {
  default = "centrala\\159435"
}
variable "orion_userpass" {
}

variable "ipaddress" {
  default = "10.141.16.1"
}
variable "status" {
  default = "2"
}

output "ipaddress_status" {
    value = "${swis_server.orion.status}"
}
