package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/IrekRomaniuk/syslog-generator/generator"
)

var (
	// IP : Syslog server IP address
	IP = flag.String("ip", "10.34.1.100", "Syslog server IP address")
	// PORT : Syslog server port
	PORT = flag.String("port", "11666", "Port")
	// PROTOCOL : Syslog server protocol
	PROTOCOL = flag.String("protocol", "tcp", "Protocol")
	// TYPE : threat or traffic
	TYPE     = flag.String("type", "Threat", "Type: Traffic or Threat")
	/*SLEEP = flag.Int("sleep", 1, "Sleep time between syslog messages in sec")*/
	// FREQ : Frequency of syslog messages per sec
	FREQ    = flag.Uint("freq", 2, "Frequency of syslog messages/sec")
	COUNT   = flag.Uint64("count", 10000, "Number of syslog messages to send")
	SRC     = flag.String("src", "1.2.3.4", "Source IP address in syslog")
	SEV     = flag.String("sev", "low", "Severity")
	version = flag.Bool("v", false, "Prints current version")
	//PRINT = flag.Bool("print", true, "print to console")
)
var (
	// Version : Program version
	Version = "No Version Provided"
	// BuildTime : Program build time
	BuildTime = ""
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2017 @IrekRomaniuk. All rights resgit brerved.\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *version {
		fmt.Printf("App Version: %s\nBuild Time : %s\n", Version, BuildTime)
		os.Exit(0)
	}
}

func main() {
	threat := generator.PanThreatLogs{"<141>Nov  3 12:53:35 syslog.generator 1", "2017/20/01 13:53:35", "001901000999",
		"threat", "file", "1", "2017/20/01 13:53:35", "1.2.3.4", "2.2.2.2", "0.0.0.0", "0.0.0.0", "G0s9J4jAU3",
		"me", "you", "App test", "vsys1", "src", "dst", "ae1.100", "ae2.200", "LF-elk",
		"2017/20/01 13:53:35", "33891243", "1", "11111", "22222", "0", "0", "0x0", "tcp", "test",
		"Test", "This is test only", "any", "low", "server-to-client", "5210010", "0x0",
		"10.10.10.0-10.255.255.255", "10.20.20.20-10.255.255.255", "0", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	traffic := generator.PanTrafficLogs{"2016-10-28T08:14:04+00:00 10.34.2.21 syslog.generator 1", "2016/10/28 08:14:04", "001901000999",
		"TRAFFIC", "end", "1", "2016/10/28 08:14:04", "1.1.1.1", "2.2.2.2", "0.0.0.0", "0.0.0.0", "G0s9J4jAU3",
		"", "", "App test", "vsys1", "src", "dst", "ae1.100", "ae2.200", "LF-elk",
		"2016/10/28 08:14:04", "305917", "1", "11111", "22222", "0", "0", "0x401b", "tcp", "test",
		"1024", "528", "496", "10", "2016/10/28 08:13:10", "52", "any", "0", "2805290265", "0x0",
		"10.0.0.0-10.255.255.255", "10.0.0.0-10.255.255.255", "0", "6", "4", "tcp-rst-from-server"}

	dur, _ := time.ParseDuration(strconv.Itoa(1000/int(*FREQ)) + "ms")
	signal := time.Tick(dur)
	var counter uint64 = 0
	var mutex = &sync.Mutex{}
	for range signal {
		switch *TYPE {
		case "Threat":
			go func() {
				mutex.Lock()
				counter++
				if counter > *COUNT {
					os.Exit(1)
				}
				mutex.Unlock()
				if *SRC == "random" {
					r := strconv.Itoa(rand.Intn(254) + 1)
					send(threat, *PROTOCOL, *IP, *PORT, r+"."+r+"."+r+"."+r, *SEV)
				} else {
					send(threat, *PROTOCOL, *IP, *PORT, *SRC, *SEV)
				}

			}()			
		case "Traffic":
			go send(traffic, *PROTOCOL, *IP, *PORT, *SRC, *SEV)
		}
		//time.Sleep(time.Duration(*SLEEP)*time.Second)
	}
}


func send(syslog generator.Syslog, prot, ip, port, src, sev string) {
	syslog.Send(prot, ip, port, src, sev) 
}
