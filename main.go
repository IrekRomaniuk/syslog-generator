package main

import "github.com/IrekRomaniuk/syslog-generator/generator"

func main() {

	//generator.Send("udp","10.34.1.100","11514")
	generator.Send("tcp","10.34.1.100","11666")

}
