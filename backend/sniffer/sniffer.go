package sniffer

import (
	"fmt"
	"time"

	"../baseLib"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func GetPcapVersion() (version string) {
	version = pcap.Version()
	return
}

func GetDevices() (devices []pcap.Interface) {
	devices, _ = pcap.FindAllDevs()
	return
}

type Sniffer struct {
	DeviceID    string
	Handle      *pcap.Handle
	Transporter transporter
}

func (sn *Sniffer) Open() {
	Handle, err := pcap.OpenLive(
		sn.DeviceID,
		int32(65535),
		false,
		-1*time.Second,
	)
	baseLib.ErrorIf(err, "Open sniffer failed.")
	sn.Handle = Handle
}

func (sn *Sniffer) Stop() {
	if sn.Handle != nil {
		sn.Handle.Close()
		sn.Handle = nil
	}
}

func (sn *Sniffer) StartCapture() {
	if sn.Handle == nil {
		fmt.Println("Handle not exists!")
		return
	}
	packetSource := gopacket.NewPacketSource(
		sn.Handle,
		sn.Handle.LinkType())

	for packet := range packetSource.Packets() {
		if tempLayer, ipLayer := packet.Layer(packetLayer), packet.Layer(layers.LayerTypeIPv4); tempLayer != nil && ipLayer != nil {
			protocolStruct := GetProtocolStruct(tempLayer)

			ip, _ := ipLayer.(*layers.IPv4)
			dstIP := ip.DstIP.String()
			protocol := ip.Protocol.String()
			dstPort := int(protocolStruct.DstPort)
			portType := portNames[protocolStruct.DstPort]

			if extraData, retBool := Filter(dstPort, portType, protocolStruct.BaseLayer.Payload); retBool != false {
				sn.Transporter(dstIP, protocol, dstPort, portType, extraData)
			}
		}
	}
}
