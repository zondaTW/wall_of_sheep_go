package main

import (
	"fmt"
	"os"

	"./sniffer"
	"./websocketServer"
	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("env.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	deviceID := cfg.Section("").Key("device_id").String()

	ws := &websocketServer.WebsocketServer{Addr: ":3000"}
	ws.Start()

	snifferObj := &sniffer.Sniffer{
		DeviceID:    deviceID,
		Handle:      nil,
		Transporter: sniffer.Transporter(ws.Send)}
	snifferObj.Open()
	defer snifferObj.Stop()
	snifferObj.StartCapture()
}
