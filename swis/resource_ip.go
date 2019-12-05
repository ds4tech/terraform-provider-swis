package swis

import (
    "encoding/json"
    "log"
    "fmt"
    "regexp"

    "github.com/hashicorp/terraform/helper/schema"
    "github.com/mrxinu/gosolar"
)

func resourceIp() *schema.Resource {
    return &schema.Resource{
            Create: resourceIpCreate,
            Read:   resourceIpRead,
            Update: resourceIpUpdate,
            Delete: resourceIpDelete,
            Schema: map[string]*schema.Schema{
                "vsphere_vlan": &schema.Schema{
                        Type:     schema.TypeString,
                        Required: true,
                },
                "ipaddress": &schema.Schema{
                        Type:     schema.TypeString,
                        Optional: true,
                },
                "status": &schema.Schema{
                        Type:     schema.TypeString,
                        Optional: true,
                },
            },
    }
}

func resourceIpCreate(d *schema.ResourceData, m interface{}) error {
    vsphere_vlan := d.Get("vsphere_vlan").(string) //VLAN100_10.141.16.0m24
    ipaddress := "0.0.0.0"
    client := m.(*gosolar.Client)
    var response []*Node

    subnetIpAddr := findIP(vsphere_vlan)
    cidr := string(vsphere_vlan[len(vsphere_vlan)-2:]) //substring last 2 chars

    log.Printf("[DEBUG] vsphere_vlan: %s, subnetIpAddr: %s, cidr: %s", vsphere_vlan, subnetIpAddr, cidr)

    querySubnet := fmt.Sprintf("SELECT  SubnetId,CIDR,Address,AddressMask,GroupTypeText FROM IPAM.Subnet WHERE  Address='%s' AND CIDR=%s", subnetIpAddr, cidr)
    response = requestAndReply(client, querySubnet)

    queryIpnode := fmt.Sprintf("SELECT TOP 1 IpNodeId,IPAddress,Comments,Status,Uri FROM IPAM.IPNode WHERE SubnetId='%d' and status=2 AND IPOrdinal BETWEEN 11 AND 254", response[0].SUBNETID)
    response = requestAndReply(client, queryIpnode)
    if(2 == response[0].STATUS) {
      updateStatus(client, response[0].URI, "1") // '1' == ip used
      ipaddress = response[0].IPADDRESS
    } else {
      log.Printf("[DEBUG] Ip address is not available for specified subnet.")
      return fmt.Errorf("Error while reserving ip address. None is available for specified subnet.")
    }

    d.Set("ipaddress", ipaddress)
    d.SetId(ipaddress)
    return nil
}

func resourceIpRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceIpUpdate(d *schema.ResourceData, m interface{}) error {
    status := d.Get("status").(string)
    ipaddress := d.Get("ipaddress").(string)
    client := m.(*gosolar.Client)
    var response []*Node

    if(""==ipaddress){
      log.Printf("[DEBUG] Ip address is missing in resource parameters.")
      return fmt.Errorf("Error. IP address is missing in resource parameters." )
    }

    query := fmt.Sprintf("SELECT IpNodeId,IPAddress,Comments,Status,Uri FROM IPAM.IPNode WHERE IPAddress='%s'", ipaddress)
    response = requestAndReply(client, query)

    log.Printf("[DEBUG] response: %s, %s", response[0], status)
    updateStatus(client, response[0].URI, status)

    //d.Set("ipaddress", ipaddress)
    d.SetId(ipaddress)
    return nil
}

func resourceIpDelete(d *schema.ResourceData, m interface{}) error {
    ipaddress := d.Get("ipaddress").(string)
    client := m.(*gosolar.Client)
    var response []*Node

    query := fmt.Sprintf("SELECT IpNodeId,IPAddress,Comments,Status,Uri FROM IPAM.IPNode WHERE IPAddress='%s'", ipaddress)
    response = requestAndReply(client, query)
    updateStatus(client, response[0].URI, "2") // '2' == ip available

    d.SetId("")
    return nil
}


// Node struct holds query results
type Node struct {
    IPNODEID int `json:"ipnodeid"` //not used
    SUBNETID int `json:"subnetid"`
    STATUS int `json:"status"`
    IPADDRESS string `json:"ipaddress"`
    COMMENTS string `json:"comments"` //not used
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

func findIP(input string) string {
   numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
   regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock

   regEx := regexp.MustCompile(regexPattern)
   return regEx.FindString(input)
}
