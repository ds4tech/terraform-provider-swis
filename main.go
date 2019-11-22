package main

import (
	"github.com/ds4tech/terraform-provider-swis/swis"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts {
		ProviderFunc: swis.Provider})
}
