package utils

import (
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
)

var (
	a = InetAtoN("10.255.255.255") >> 24
	b = InetAtoN("172.31.255.255") >> 20
	c = InetAtoN("192.168.255.255") >> 16
)

// IsVaild 有效IP地址
func IsVaild(ip string) bool {
	return net.ParseIP(ip) != nil
}

// InetAtoN IP地址转int64
func InetAtoN(ip string) int64 {
	n := big.NewInt(0)
	n.SetBytes(net.ParseIP(ip).To4())
	return n.Int64()
}

// IsInternal 判断是否为内网IP
func IsInternal(ip string) bool {
	if !IsVaild(ip) {
		return false
	}
	n := InetAtoN(ip)
	return n>>24 == a || n>>20 == b || n>>16 == c
}

// GetExternalIPByHTTP 获取公网IP
func GetExternalIPByHTTP() (string, error) {
	res, err := http.Get("http://whatismyip.akamai.com/")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// GetExternalIPByDNS 获取公网IP
func GetExternalIPByDNS() (string, error) {
	conn, err := net.Dial("tcp", "ns1.dnspod.net:6666")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	var ip [15]byte
	n, err := conn.Read(ip[:])
	if err != nil {
		return "", err
	}
	return string(ip[:n]), nil
}
