package common

import (
	"log"
	"net"
	"os"
)

func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		log.Print(err)
		return ""
	}
	return name
}

func GetHostaddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.To4().String()
			}
		}
	}

	return "0.0.0.0"
}

func EnvString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
