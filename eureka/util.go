package eureka

import (
	"fmt"
	"net"
)

// get one non-loopback ip from net interface
func getLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// generate default instance-id
func genDefaultInstanceId(app string, port int) string {
	hostname := getLocalIp()
	return fmt.Sprintf("%s:%s:%d", hostname, app, port)
}
