package generator

import (
	"fmt"
	"time"
	"reflect"
	"net"
	"strings"
	"os"
)

func Send(protocol,ip,port string) error {
	p := PanThreatLogs{"<141>Nov  3 12:53:35 syslog.generator 1","2017/20/01 13:53:35","001901000999",
		"THREAT", "file","1","2017/20/01 13:53:35","1.1.1.1","2.2.2.2","0.0.0.0","0.0.0.0","RULE fake",
		"me","you", "App test","vsys1","src","dst","ae1.100","ae2.200","LF-elk",
		"2017/20/01 13:53:35","33891243","1","11111","22222","0","0","0x0","tcp","test",
		"This is test only","any","low","server-to-client","5210010","0x0",
		"10.10.10.0-10.255.255.255","10.20.20.20-10.255.255.255","0","","","","","","", "","","","","","",""}

	l := "2006/02/01 03:04:05"
	name, err := os.Hostname()
	if err != nil {
		return err
	}
	now := time.Now()
	t := now.Format(l)  //yyyy/dd/mm HH:mm:ss
	fmt.Println("Time", t, )
	p.Domain = "<141>" + fmt.Sprintf("%03s %02d %02d:%02d:%02d",now.Month(),now.Day(),
		now.Hour(),now.Minute(),now.Second()) + " " + name + " 1"
	p.ReceiveTime, p.GenerateTime, p.TimeLogged = t, now.Add(1*time.Second).Format(l),
		now.Add(2*time.Second).Format(l)
	v := reflect.ValueOf(p)
	values := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface().(string)
	}

	conn, err := net.Dial(protocol, ip + ":" + port)
	if err != nil {
		return err
	}

	msg := strings.Join(values,",")
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

type PanThreatLogs struct {
	Domain,ReceiveTime,SerialNum,Type,ThreatContentType,ConfigVersion,GenerateTime,SourceIP,DestinationIP,
NATSourceIP,NATDestinationIP,Rule,SourceUser,DestinationUser,Application,VirtualSystem,SourceZone,DestinationZone,
InboundInterface,OutboundInterface,LogAction,TimeLogged,SessionID,RepeatCount,SourcePort,DestinationPort,NATSourcePort,
NATDestinationPort,Flags,IPprotocol,Action,URL,ThreatContentName,Category,Severity,Direction,Seqno,ActionFlags,
SourceLocation,DestinationLocation,Cpadding,ContentType,Pcap_id,Filedigest,Cloud,Url_idx,User_agent,Filetype,Xff,
Referer,Sender,Subject,Recipient,Reportid string
}
