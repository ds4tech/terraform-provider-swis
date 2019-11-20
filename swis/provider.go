package swis

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
  return &schema.Provider{
    ResourcesMap: map[string]*schema.Resource{
      "example_server": resourceServer(),
    },
  }
}

/*
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{

			"user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user name of the AD Server",
				DefaultFunc: schema.EnvDefaultFunc("AD_USER", nil),
			},
		},

	}
}
*/
