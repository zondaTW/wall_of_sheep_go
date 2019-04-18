package sniffer

import (
	"../baseLib"
	"github.com/google/gopacket/layers"
)

var (
	packetLayer     = layers.LayerTypeTCP
	portNames       = layers.TCPPortNames
	filterPorts     = []int{80, 443}
	filterPortTypes = []string{"http", "https"}
)

func GetProtocolStruct(tempLayer interface{}) (tcp *layers.TCP) {
	tcp, _ = tempLayer.(*layers.TCP)
	return
}

func Filter(dstPort int, portType string, payload []byte) (ret bool) {
	ret = baseLib.FindIntArray(filterPorts, dstPort)
	ret = ret || baseLib.FindStringArray(filterPortTypes, portType)
	return
}
