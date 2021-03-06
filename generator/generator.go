package generator

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"
	"time"
)

// Syslog ...
type Syslog interface {
	Send(protocol, ip, port, src, sev string) error
}

// Send threats
func (p PanThreatLogs) Send(protocol, ip, port, src, sev string) error {
	l := "2006/02/01 15:04:05" //24h format
	name, err := os.Hostname()
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	t := now.Format(l) //yyyy/dd/mm HH:mm:ss
	p.Domain = "<141>" + fmt.Sprintf("%03s %02d %02d:%02d:%02d", now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second()) + " " + name + " 1"
	p.ReceiveTime, p.GenerateTime, p.TimeLogged = t, now.Add(1*time.Second).Format(l),
		now.Add(2*time.Second).Format(l)
	p.SourceIP = src
	p.Severity = sev
	//fmt.Println(p.ReceiveTime, p.GenerateTime, p.TimeLogged)
	v := reflect.ValueOf(p)
	values := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface().(string)
	}

	conn, err := net.Dial(protocol, ip+":"+port)
	if err != nil {
		return err
	}

	msg := strings.Join(values, ",")
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

// Send Traffic
func (p PanTrafficLogs) Send(protocol, ip, port, src, sev string) error {
	l := "2006/02/01 15:04:05" //24h format
	name, err := os.Hostname()
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	t := now.Format(l) //yyyy/dd/mm HH:mm:ss
	p.Domain = "<141>" + fmt.Sprintf("%03s %02d %02d:%02d:%02d", now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second()) + " " + name + " 1"
	p.ReceiveTime, p.GenerateTime, p.TimeLogged = t, now.Add(1*time.Second).Format(l),
		now.Add(2*time.Second).Format(l)
	p.SourceIP = src
	//fmt.Println(p.ReceiveTime, p.GenerateTime, p.TimeLogged)
	v := reflect.ValueOf(p)
	values := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface().(string)
	}

	conn, err := net.Dial(protocol, ip+":"+port)
	if err != nil {
		return err
	}

	msg := strings.Join(values, ",")
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

/*type PanLogs struct {
	Domain,ReceiveTime,SerialNum,Type,Subtype,ConfigVersion,GenerateTime,SourceIP,DestinationIP,
NATSourceIP,NATDestinationIP,Rule,SourceUser,DestinationUser,Application,VirtualSystem,SourceZone,DestinationZone,
InboundInterface,OutboundInterface,LogAction,TimeLogged,SessionID,RepeatCount,SourcePort,DestinationPort,NATSourcePort,
NATDestinationPort,Flags string
}

type PanThreatLogs struct {
	PanLogs
	URL,ThreatContentName,Category,Severity,Direction,Seqno,ActionFlags,
	SourceLocation,DestinationLocation,Cpadding_th,ContentType,Pcap_id,Filedigest,Cloud,Url_idx,User_agent,Filetype,Xff,
	Referer,Sender,Subject,Recipient,Reportid string
}

type PanTrafficLogs struct {
	PanLogs
	Bytes,BytesSent,BytesReceived,Packets,StartTime,ElapsedTime,Category_tr,Padding,
	Seqno,ActionFlags,SourceLocation_tr,DestinationLocation_tr,Cpadding_tr,PacketsSent,PacketsReceived,
	SessionEndReason string
}*/

type PanThreatLogs struct {
	Domain, ReceiveTime, SerialNum, Type, Subtype, ConfigVersion, GenerateTime, SourceIP, DestinationIP,
	NATSourceIP, NATDestinationIP, Rule, SourceUser, DestinationUser, Application, VirtualSystem, SourceZone, DestinationZone,
	InboundInterface, OutboundInterface, LogAction, TimeLogged, SessionID, RepeatCount, SourcePort, DestinationPort, NATSourcePort,
	NATDestinationPort, Flags, Protocol, Action, URL, ThreatContentName, Category, Severity, Direction, Seqno, ActionFlags,
	SourceLocation, DestinationLocation, Cpadding_th, ContentType, Pcap_id, Filedigest, Cloud, Url_idx, User_agent, Filetype, Xff,
	Referer, Sender, Subject, Recipient, Reportid string
}
type PanTrafficLogs struct {
	Domain, ReceiveTime, SerialNum, Type, Subtype, ConfigVersion, GenerateTime, SourceIP, DestinationIP,
	NATSourceIP, NATDestinationIP, RuleName, SourceUser, DestinationUser, Application, VirtualSystem, SourceZone, DestinationZone,
	InboundInterface, OutboundInterface, LogAction, TimeLogged, SessionID, RepeatCount, SourcePort, DestinationPort, NATSourcePort,
	NATDestinationPort, Flags, Protocol, Action, Bytes, BytesSent, BytesReceived, Packets, StartTime, ElapsedTime, Category_tr, Padding,
	Seqno, ActionFlags, SourceLocation_tr, DestinationLocation_tr, Cpadding_tr, PacketsSent, PacketsReceived,
	SessionEndReason string
}
