#!/bin/bash

go clean
rm terraform.tfstate* .terraform/plugins/darwin_amd64/lock.json
go build -o terraform-provider-swis
terraform init
terraform apply
