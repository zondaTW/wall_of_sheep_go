package sniffer

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"../baseLib"
	"github.com/google/gopacket/layers"
)

var (
	packetLayer     = layers.LayerTypeTCP
	portNames       = layers.TCPPortNames
	filterPorts     = []int{80, 443}
	filterPortTypes = []string{"http", "https"}
)

type WallOfSheepInfo struct {
	Method   string
	User     string
	Password string
}

func GetProtocolStruct(tempLayer interface{}) (tcp *layers.TCP) {
	tcp, _ = tempLayer.(*layers.TCP)
	return
}

func findUserAndPassword(sourceStr string) (user string, password string) {
	for _, subString := range strings.Split(sourceStr, "&") {
		temp := strings.Split(subString, "=")
		if len(temp) != 2 {
			continue
		}
		title, value := temp[0], temp[1]
		if match, _ := regexp.MatchString("(u|U)(s|S)(e|E)(r|R)", title); match {
			user = value
		} else if match, _ := regexp.MatchString("(p|P)(a|A)(s|S)(s|S)", title); match {
			password = value
		}
	}
	return
}

func Filter(dstPort int, portType string, payload []byte) (extraData interface{}, ret bool) {
	ret = false

	reader := bufio.NewReader(bytes.NewReader(payload))
	httpReq, _ := http.ReadRequest(reader)
	if httpReq != nil {
		ret = baseLib.FindIntArray(filterPorts, dstPort)
		ret = ret || baseLib.FindStringArray(filterPortTypes, portType)

		user := ""
		password := ""
		body, _ := ioutil.ReadAll(httpReq.Body)
		user, password = findUserAndPassword(string(body))
		if user == "" && password == "" {
			temp := strings.Split(httpReq.URL.String(), "?")
			if len(temp) == 2 {
				user, password = findUserAndPassword(temp[1])
			}
		}

		if user != "" || password != "" {
			extraData = &WallOfSheepInfo{
				Method:   httpReq.Method,
				User:     user,
				Password: password}
		} else {
			ret = false
		}
	}
	return
}
