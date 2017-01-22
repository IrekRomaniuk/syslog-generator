package main

import (
	"github.com/IrekRomaniuk/syslog-generator/generator"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	IP = flag.String("ip", "10.34.1.100", "Syslog server IP address")
	PORT = flag.String("port", "11666", "Port")
	PROTOCOL = flag.String("protocol", "tcp", "Protocol")
	SLEEP = flag.Int("sleep", 1, "Sleep time between syslog messages in sec")
	COUNT = flag.Int("count", 1, "Number of syslog messages to send")
	version = flag.Bool("v", false, "Prints current version")
	//PRINT = flag.Bool("print", true, "print to console")
)
var (
	Version = "No Version Provided"
	BuildTime = ""
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2017 @IrekRomaniuk. All rights reserved.\n")
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

	p := generator.PanThreatLogs{"<141>Nov  3 12:53:35 syslog.generator 1","2017/20/01 13:53:35","001901000999",
		"THREAT", "file","1","2017/20/01 13:53:35","1.1.1.1","2.2.2.2","0.0.0.0","0.0.0.0","G0s9J4jAU3",
		"me","you", "App test","vsys1","src","dst","ae1.100","ae2.200","LF-elk",
		"2017/20/01 13:53:35","33891243","1","11111","22222","0","0","0x0","tcp","test",
		"Test","This is test only","any","low","server-to-client","5210010","0x0",
		"10.10.10.0-10.255.255.255","10.20.20.20-10.255.255.255","0","","","","","","", "","","","","","",""}
	/*q := generator.PanTrafficLogs{"2016-10-28T08:14:04+00:00 10.34.2.21 syslog.generator 1","2016/10/28 08:14:04","001901000999",
		"TRAFFIC","end","1","2016/10/28 08:14:04","1.1.1.1","2.2.2.2","0.0.0.0","0.0.0.0","G0s9J4jAU3",
		"","","App test","vsys1","src","dst","ae1.100","ae2.200","LF-elk",
		"2016/10/28 08:14:04","305917","1","11111","22222","0","0","0x401b","tcp","test",
		"1024","528","496","10","2016/10/28 08:13:10","52","any","0","2805290265","0x0",
		"10.0.0.0-10.255.255.255","10.0.0.0-10.255.255.255","0","6","4","tcp-rst-from-server","38","0","0","0",
		"PROD","DC-FW","from-policy"}*/

	for i:=0; i < *COUNT; i++ {
		p.Send(*PROTOCOL,*IP,*PORT)
		time.Sleep(time.Duration(*SLEEP)*time.Second)
	}

}
