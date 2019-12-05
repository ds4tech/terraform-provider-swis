package swis

import (
	"log"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/mrxinu/gosolar"
)

func Provider() terraform.ResourceProvider {
  return &schema.Provider{
		Schema: map[string]*schema.Schema{
					"hostname": &schema.Schema{
									Type:     schema.TypeString,
									Required: true,
									Description: "Orion server ip address",
									DefaultFunc: schema.EnvDefaultFunc("ORION_IP", nil),
					},
					"username": &schema.Schema{
									Type:     schema.TypeString,
									Required: true,
									Description: "The user name to login to Orion server",
									DefaultFunc: schema.EnvDefaultFunc("ORION_USER", nil),
					},
					"password": &schema.Schema{
									Type:     schema.TypeString,
									Required: true,
									Sensitive:   true,
									Description: "The user password to Orion server",
									DefaultFunc: schema.EnvDefaultFunc("ORION_PASSWORD", nil),
					},
		},
    ResourcesMap: map[string]*schema.Resource{
					"swis_server": resourceIp(),
//			"swis_example": resourceExample(),
    },
		ConfigureFunc: providerConfigure,
		}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
    host := d.Get("hostname").(string)
    username := d.Get("username").(string)
    password := d.Get("password").(string)
		var client * gosolar.Client = gosolar.NewClient(host, username, password, true)

		log.Printf("[TRACE] Connecting to Orion")
		return client, nil
}
