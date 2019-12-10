package swis

import (
    "encoding/json"
    "regexp"
    "log"

    "github.com/mrxinu/gosolar"
)

// Node struct holds query results
type Node struct {
    IPNODEID int `json:"ipnodeid"` //not used
    SUBNETID int `json:"subnetid"`
    STATUS int `json:"status"`
    IPADDRESS string `json:"ipaddress"`
    COMMENTS string `json:"comments"` //not used
    URI string `json:"uri"`
}


func queryOrionServer(client * gosolar.Client, query string) []*Node {
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

func updateIPNodeStatus(client * gosolar.Client, uri string, status string) {
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
