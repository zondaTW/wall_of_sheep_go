package sniffer

import (
	"encoding/json"
	"strconv"
)

type transporter func(string, string, int, string, interface{})
type transportFunc func(interface{})
type TransportInfo struct {
	IP       string
	Protocol string
	Port     string
	Type     string
}

func Transporter(fn transportFunc) transporter {
	return func(dstIP string, protocol string, dstPort int, portType string, extraData interface{}) {
		transportInfo := &TransportInfo{
			IP:       dstIP,
			Protocol: protocol,
			Port:     strconv.Itoa(dstPort),
			Type:     portType,
		}

		var mergeData map[string]string

		jsonTransportInfo, _ := json.Marshal(transportInfo)
		json.Unmarshal(jsonTransportInfo, &mergeData)

		if extraData != nil {
			jsonExtraData, _ := json.Marshal(extraData)
			json.Unmarshal(jsonExtraData, &mergeData)
		}

		fn(mergeData)
	}
}
