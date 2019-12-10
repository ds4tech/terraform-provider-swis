package swis

import (
  "fmt"
  "log"
  "net"
  "net/http"
  "crypto/tls"
  "time"
   "io/ioutil"
  "github.com/sparrc/go-ping"
)

func pingTest(host string) {
  pinger, err := ping.NewPinger(host)
  if err != nil {
        fmt.Printf("ERROR: %s\n", err.Error())
        return
  }

  pinger.OnRecv = func(pkt *ping.Packet) {
        fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
                pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
  }
  pinger.OnFinish = func(stats *ping.Statistics) {
        fmt.Printf("--- %s ping statistics ---\n", stats.Addr)
        fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
                stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
        fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
                stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
  }

  fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
  pinger.Count = 3
  pinger.Run()
  fmt.Println()
}

func ncPortTest(host string, ports []string) {
  fmt.Println("Trying to connect on ports: ", ports)
    for _, port := range ports {
        timeout := time.Second
        conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
        if err != nil {
            fmt.Println("Connecting error:", err)
        }
        if conn != nil {
            defer conn.Close()
            fmt.Println("Opened", net.JoinHostPort(host, port))
        }
    }
}

func curlPortTest(host string, port string) {
//url := fmt.Sprintf("https://%s:%s", host, port)
url := "https://10.50.8.10:17778/SolarWinds/InformationService/v3/Json/Query?query=SELECT IpNodeId,SubnetId,IPAddress,Comments,Status,Uri FROM IPAM.IPNode WHERE IPAddress='10.141.16.11'"

  transCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
  }

  var netClient = &http.Client{
    Transport: transCfg,
    Timeout: time.Second * 10,
  }

        resp, err := netClient.Get(url)
        if err != nil {
                log.Fatalln(err)
        }

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                log.Fatalln(err)
        }

        log.Println(string(body))
}
