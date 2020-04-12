package models

import (
	"net"

	"github.com/likiiiiii/foxpot_backend/utils"
)

// GeoIP2Info IP基本信息
type GeoIP2Info struct {
	Latitude  float64
	Longitude float64
	Region    string
}

// GetGeoIP2Info 获取IP基本信息
func GetGeoIP2Info(ip string) (*GeoIP2Info, error) {
	if !utils.IsVaild(ip) {
		return nil, net.InvalidAddrError("invalid ipv4 addr.")
	}
	record, err := GeoDB.City(net.ParseIP(ip))
	if err != nil {
		return nil, err
	}
	info := &GeoIP2Info{
		Latitude:  record.Location.Latitude,
		Longitude: record.Location.Longitude,
		Region:    record.Country.Names["zh-CN"] + "·" + record.City.Names["zh-CN"],
	}
	return info, nil
}
