package generator

import (
	"fmt"
	"time"
	"reflect"
	"net"
	"strings"
	"os"
)

func (p PanThreatLogs) Send(protocol,ip,port string) error {
	l := "2006/02/01 03:04:05"
	name, err := os.Hostname()
	if err != nil {
		return err
	}
	now := time.Now()
	t := now.Format(l)  //yyyy/dd/mm HH:mm:ss
	//fmt.Println("Time", t, )
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
type PanTrafficLogs struct {
	Domain,ReceiveTime,SerialNum,Type,Subtype,ConfigVersion,GeneratedTime,SourceIP,DestinationIP,
NATSourceIP,NATDestinationIP,RuleName,SourceUser,DestinationUser,Application,VirtualSystem,SourceZone,DestinationZone,
InboundInterface,OutboundInterface,LogAction,TimeLogged,SessionID,RepeatCount,SourcePort,DestinationPort,NATSourcePort,
NATDestinationPort,Flags,Protocol,Action,Bytes,BytesSent,BytesReceived,Packets, StartTime,ElapsedTime,Category, Padding,
SequenceNumber,ActionFlags,SourceLocation,DestinationLocation,Cpadding,PacketsSent,PacketsReceived,
SessionEndReason string
}
