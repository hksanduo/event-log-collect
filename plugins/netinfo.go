package plugins

import (
	"event-log-collect/utils"
	"syscall"

	"github.com/gookit/color"
	"github.com/shirou/gopsutil/net"
)

//NetInterfaces 网卡信息
func NetInterfaces(path, timestramp string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		color.Error.Println(err)
		return
	}

	color.Info.Printf("%-15s|%-20s|%-10s|%-35s|%s\n", "Name", "Mac", "MTU", "Status", "Ip")
	var data = ""
	for i := range interfaces {
		inter := interfaces[i]
		status := ""
		for index, flag := range inter.Flags {
			status = status + flag
			if index != len(inter.Flags)-1 {
				status = status + ","
			}
		}
		color.Info.Printf("%-15s|%-20s|%-10d|%-35s|%s\n", inter.Name, inter.HardwareAddr, inter.MTU, status, inter.Addrs)
		// netinfoData := strings.Join([]string{inter.Name, inter.HardwareAddr, strconv.Itoa(inter.MTU), status, inter.Addrs}, ",")

		// data += string(json.Marshal(inter))
		// data = json.Marshal(inter)
		data += inter
	}
	// save file
	utils.Save_data_to_file(path, "netinfo", data, timestramp)
	// utils

}

func NetConnections() {
	conns, err := net.Connections("inet")
	if err != nil {
		color.Error.Println(err)
		return
	}

	color.Info.Println("PID", "SRC_IP", "DIST_IP", "STATUS", "TYPE")
	for _, conn := range conns {
		laddr := conn.Laddr.IP + ":" + utils.Int2Str(int(conn.Laddr.Port))
		raddr := conn.Raddr.IP + ":" + utils.Int2Str(int(conn.Raddr.Port))
		color.Info.Println(conn.Pid, laddr, raddr, conn.Status, UDPAndTCPMap(conn.Type))
	}

}

func UDPAndTCPMap(value uint32) string {
	var name string = ""
	switch value {
	case syscall.SOCK_STREAM:
		name = "TCP"
		break
	case syscall.SOCK_DGRAM:
		name = "UDP"
		break
	}
	return name
}
