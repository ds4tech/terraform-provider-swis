package swis

import (
    "log"
    "fmt"

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
    ipaddress := d.Get("ipaddress").(string) //ipaddress
    client := m.(*gosolar.Client)
    querySubnet := ""
    comment := ""
    var response []*Node

    log.Printf("[DEBUG] Reserve: IPADDRESS: %v", ipaddress)
/* funkcja powinna dzialac odwrotnie, czyli jest jes polanczenie do trzeba sprawdzic negacje: !
    if( checkConnectivity(ipaddress) ) {
        log.Printf("[DEBUG] checkConnectivity finished")
    }
*/

    if(len(ipaddress)!=0) {
        querySubnet = fmt.Sprintf("SELECT IpNodeId,IPAddress,Comments,Status,Uri FROM IPAM.IPNode WHERE IPAddress='%s'", ipaddress)
        response = queryOrionServer(client, querySubnet)
        log.Printf("[DEBUG] Reserve: IPADDRESS: %v, COMMENTS: %v\n", response[0].IPADDRESS, response[0].COMMENTS)

        if(2 == response[0].STATUS) {
          updateIPNodeStatus(client, response[0].URI, "1", comment) // '1' == ip used
        } else {
          log.Fatalf("[ERROR] Ip address is not available. It has status %v.", response[0].STATUS)
        }
    } else {
        subnetIpAddr := findIP(vsphere_vlan)
        cidr := string(vsphere_vlan[len(vsphere_vlan)-2:]) //substring last 2 chars

        log.Printf("[TRACE] vsphere_vlan: %s, subnetIpAddr: %s, cidr: %s", vsphere_vlan, subnetIpAddr, cidr)

        querySubnet = fmt.Sprintf("SELECT  SubnetId,CIDR,Address,AddressMask,GroupTypeText FROM IPAM.Subnet WHERE  Address='%s' AND CIDR=%s", subnetIpAddr, cidr)
        response = queryOrionServer(client, querySubnet)

        queryIpnode := fmt.Sprintf("SELECT TOP 1 IpNodeId,IPAddress,Comments,Status,Uri FROM IPAM.IPNode WHERE SubnetId='%d' and status=2 AND IPOrdinal BETWEEN 11 AND 254", response[0].SUBNETID)
        response = queryOrionServer(client, queryIpnode)
        if(2 == response[0].STATUS) {
          updateIPNodeStatus(client, response[0].URI, "1", "") // '1' == ip used
          ipaddress = response[0].IPADDRESS
        } else {
          log.Printf("[ERROR] Ip address is not available for specified subnet.")
          return fmt.Errorf("Error while reserving ip address. None is available for specified subnet.")
        }
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
    response = queryOrionServer(client, query)

    log.Printf("[DEBUG] response: %s, %s", response[0], status)
    updateIPNodeStatus(client, response[0].URI, status, "")

    //d.Set("ipaddress", ipaddress)
    d.SetId(ipaddress)
    return nil
}

func resourceIpDelete(d *schema.ResourceData, m interface{}) error {
    ipaddress := d.Get("ipaddress").(string)
    client := m.(*gosolar.Client)
    var response []*Node

    query := fmt.Sprintf("SELECT IpNodeId,IPAddress,Comments,Status,Uri FROM IPAM.IPNode WHERE IPAddress='%s'", ipaddress)
    response = queryOrionServer(client, query)
    updateIPNodeStatus(client, response[0].URI, "2", "") // '2' == ip available

    d.SetId("")
    return nil
}
