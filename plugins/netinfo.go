package plugins

import (
	"event-log-collect/utils"
	"fmt"

	"github.com/gookit/color"
	"github.com/shirou/gopsutil/net"
)

// NetInterfaces 网卡信息
func NetInterfaces(path, timestramp string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		color.Error.Println(err)
		return
	}
	var data = fmt.Sprintf("%v", interfaces)
	utils.Save_data_to_file(path, "netinfo", data, timestramp)
}

// 网络连接信息
func NetConnections(path, timestramp string) {
	conns, err := net.Connections("inet")
	if err != nil {
		color.Error.Println(err)
		return
	}
	var data = fmt.Sprintf("%v", conns)
	utils.Save_data_to_file(path, "netconnect", data, timestramp)
}

//
// func UDPAndTCPMap(value uint32) string {
// 	var name string = ""
// 	switch value {
// 	case syscall.SOCK_STREAM:
// 		name = "TCP"
// 		break
// 	case syscall.SOCK_DGRAM:
// 		name = "UDP"
// 		break
// 	}
// 	return name
// }
