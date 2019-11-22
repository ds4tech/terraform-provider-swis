package swis

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceServer() *schema.Resource {
        return &schema.Resource{
                Create: resourceServerCreate,
                Read:   resourceServerRead,
                Update: resourceServerUpdate,
                Delete: resourceServerDelete,

                Schema: map[string]*schema.Schema{
                        "address": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "port": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "username": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "password": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
    address := d.Get("address").(string)
    d.SetId(address)
    return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
    address := d.Get("address").(string)
    d.SetId(address)
    return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}
