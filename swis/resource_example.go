package swis

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceQuery() *schema.Resource {
        return &schema.Resource{
                Create: resourceExampleCreate,
                Read:   resourceExampleRead,
                Update: resourceExampleUpdate,
                Delete: resourceExampleDelete,

                Schema: map[string]*schema.Schema{
                        "query": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func resourceExampleCreate(d *schema.ResourceData, m interface{}) error {
    query := d.Get("query").(string)
    d.SetId(query + " from czlowiek")
    return resourceServerRead(d, m)
}

func resourceExampleRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceExampleUpdate(d *schema.ResourceData, m interface{}) error {
    query := d.Get("query").(string)
    d.SetId(query)
    return resourceServerRead(d, m)
}

func resourceExampleDelete(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}
