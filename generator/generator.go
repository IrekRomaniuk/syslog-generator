package generator

import (
	"fmt"
	"time"
	"reflect"
	"net"
	"strings"
	"encoding/gob"
)

func Send(protocol,ip,port string) error {
	p := PanThreatLogs{"<141>Nov  3 12:53:35 DC-FW.com 1","2017/20/01 13:53:35","001901000999",
		"THREAT", "file","1","2017/20/01 13:53:35","1.1.1.1","2.2.2.2","0.0.0.0","0.0.0.0","RULE fake",
		"me","you", "App test","vsys1","app","dmz","ae1.100","ae2.200","LF-elk",
		"2017/20/01 13:53:35","33891243","1","11111","22222","0","0","0x0","tcp","test",
		"CFN WebFix.exe","Microsoft PE File(52060)","any","low","server-to-client","5210010","0x0",
		"10.10.10.0-10.255.255.255","10.20.20.20-10.255.255.255","0","","","","","","", "","","","","","",""}

	//q := Pan{"aaa","bbb","ccc"}
	current := time.Now()
	t := current.Format("2017/20/01 13:53:35")  //yyyy/MM/dd HH:mm:ss
	//fmt.Println("Time", t)
	fmt.Println(time.Now().Format("2017/20/01 13:53:35"))
	p.ReceiveTime, p.GenerateTime, p.TimeLogged = t, t, t
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
	err = gob.NewEncoder(conn).Encode(msg)
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

type Pan struct {
	A, B, C string
}