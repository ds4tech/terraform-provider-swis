package swis

import (
    "encoding/json"
    "log"
    "fmt"

		"github.com/hashicorp/terraform/helper/schema"
		"github.com/mrxinu/gosolar"
)

func resourceServer() *schema.Resource {
    return &schema.Resource{
            Create: resourceServerCreate,
            Read:   resourceServerRead,
            Update: resourceServerUpdate,
            Delete: resourceServerDelete,

            Schema: map[string]*schema.Schema{
                    "hostname": &schema.Schema{
                            Type:     schema.TypeString,
                            Required: true,
                    },
                    "ipaddress": &schema.Schema{
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
                    "status": &schema.Schema{
                            Type:     schema.TypeString,
                            Required: true,
                    },
            },
    }
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
    host := d.Get("hostname").(string)
    username := d.Get("username").(string)
    password := d.Get("password").(string)
    ipaddress := d.Get("ipaddress").(string)
    status := "1" // '1' == ip used
    var client * gosolar.Client = gosolar.NewClient(host, username, password, true)
    var response []*Node

    query := fmt.Sprintf("SELECT IpNodeId,IPAddress,Comments,Status,Uri FROM IPAM.IPNode WHERE IPAddress='%s'", ipaddress)
    response = requestAndReply(client, query)
    updateStatus(client, response[0].URI, status)

    d.SetId(status)
    return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
		host := d.Get("hostname").(string)
		username := d.Get("username").(string)
		password := d.Get("password").(string)
		ipaddress := d.Get("ipaddress").(string)
		status := d.Get("status").(string)
		var client * gosolar.Client = gosolar.NewClient(host, username, password, true)
		var response []*Node

		query := fmt.Sprintf("SELECT IpNodeId,IPAddress,Comments,Status,Uri FROM IPAM.IPNode WHERE IPAddress='%s'", ipaddress)
		response = requestAndReply(client, query)
    updateStatus(client, response[0].URI, status)

		d.SetId(status)
    return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
    host := d.Get("hostname").(string)
    username := d.Get("username").(string)
    password := d.Get("password").(string)
    ipaddress := d.Get("ipaddress").(string)
    status := "2" // '2' == ip available
    var client * gosolar.Client = gosolar.NewClient(host, username, password, true)
    var response []*Node

    query := fmt.Sprintf("SELECT IpNodeId,IPAddress,Comments,Status,Uri FROM IPAM.IPNode WHERE IPAddress='%s'", ipaddress)
    response = requestAndReply(client, query)
    updateStatus(client, response[0].URI, status)

    d.SetId(status)
    return nil
}


// Node struct holds query results
type Node struct {
        IPNODEID int `json:"ipnodeid"`
        IPADDRESS string `json:"ipaddress"`
        COMMENTS string `json:"comments"`
        STATUS int `json:"status"`
        URI string `json:"uri"`
}


func requestAndReply(client * gosolar.Client, query string) []*Node {
    var data []*Node
    ipNodeDetails, err := client.Query(query, nil)
    if err != nil {
            log.Fatal(err)
    }
    if err := json.Unmarshal(ipNodeDetails, &data); err != nil {
            log.Fatal(err)
    }
    return data
}

func updateStatus(client * gosolar.Client, uri string, status string) {
  req := map[string]interface{}{
          "Status": status,
  }
  _, err := client.Update(uri, req)
  if err != nil {
          log.Fatal(err)
  }
}
