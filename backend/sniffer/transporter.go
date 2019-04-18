package sniffer

type transporter func(string, string, int, string, []byte)
type transportFunc func(interface{})
type TransportInfo struct {
	IP       string
	Protocol string
	Port     int
	Type     string
}

func Transporter(fn transportFunc) transporter {
	return func(dstIP string, protocol string, dstPort int, portType string, payload []byte) {
		transportInfo := &TransportInfo{
			IP:       dstIP,
			Protocol: protocol,
			Port:     dstPort,
			Type:     portType,
		}
		fn(transportInfo)
	}
}
