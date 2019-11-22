resource "swis_server" "orion" {
  address = var.ORION_HOST
  port = var.ORION_PORT
  username = var.ORION_USERNAME
  password = var.ORION_USERPASS
}

resource "swis_query" "orion" {
  query = "select sth"
}


variable "ORION_HOST" {
  default = "10.50.8.10"
}

variable "ORION_PORT" {
  default = "17778"
}

variable "ORION_USERNAME" {
  default = "skp"
}
variable "ORION_USERPASS" {
  default = "password"
}
