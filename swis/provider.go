package swis

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
  return &schema.Provider{
    ResourcesMap: map[string]*schema.Resource{
      "swis_server": resourceServer(),
	    "swis_query": resourceQuery(),
    },
  }
}
