package swis

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceQuery() *schema.Resource {
        return &schema.Resource{
                Create: resourceQueryCreate,
                Read:   resourceQueryRead,
                Update: resourceQueryUpdate,
                Delete: resourceQueryDelete,

                Schema: map[string]*schema.Schema{
                        "query": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func resourceQueryCreate(d *schema.ResourceData, m interface{}) error {
    query := d.Get("query").(string)
    d.SetId(query + " from czlowiek")
    return resourceServerRead(d, m)
}

func resourceQueryRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceQueryUpdate(d *schema.ResourceData, m interface{}) error {
    query := d.Get("query").(string)
    d.SetId(query)
    return resourceServerRead(d, m)
}

func resourceQueryDelete(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}
