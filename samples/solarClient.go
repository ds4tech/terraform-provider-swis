package main

import (
        "encoding/json"
        "log"
        "os"
        "fmt"

        "github.com/mrxinu/gosolar"
)

var hostname string = "10.50.8.10"
var username string = "centrala\\159435"
var password string = ""
var ipAddress string = "10.141.16.2"
var status string = "2"


// Node struct holds query results
type Node struct {
        IPNODEID int `json:"ipnodeid"`
        IPADDRESS string `json:"ipaddress"`
        COMMENTS string `json:"comments"`
        STATUS int `json:"status"`
        URI string `json:"uri"`
}

// SolarWinds connection details
func requestAndResponse(client * gosolar.Client, query string) []*Node {
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

func main() {
				var client * gosolar.Client = gosolar.NewClient(hostname, username, password, true)
        var response []*Node
        enc := json.NewEncoder(os.Stdout)

        fmt.Printf("Provided ipAddress: %s\n", ipAddress)

        query := fmt.Sprintf("SELECT IpNodeId,IPAddress,Comments,Status,Uri FROM IPAM.IPNode WHERE IPAddress='%s'", ipAddress)
        response = requestAndResponse(client, query)
        enc.Encode(response[0])

        // update property
        req := map[string]interface{}{
                "Status": status,
        }
        fmt.Println("\nUpdate request: ")
        enc.Encode(req)
        _, err := client.Update(response[0].URI, req)
        if err != nil {
                log.Fatal(err)
        }
        // end update

        fmt.Println("\nUpdated Status: ", status, ", for ipAddress", ipAddress ,"\nJSON  response:")
        response = requestAndResponse(client, query)
        enc.Encode(response[0])
}
