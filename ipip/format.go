package ipip

import (
	"encoding/json"
	"strings"
)

type LocationInfo struct {
	Country   string `json:"country"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Org       string `json:"org"`
	Isp       string `json:"isp"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	TimeZone  string `json:"timeZone"`
	UTC       string `json:"UTC"`
	ChinaCode string `json:"chinaCode"`
	PhoneCode string `json:"phoneCode"`
	ISO2      string `json:"ISO2"`
	Continent string `json:"continent"`
}

func JsonLocationInfo(s string) string {
	info := &LocationInfo{
		Country:   na,
		Province:  na,
		City:      na,
		Org:       na,
		Isp:       na,
		Latitude:  na,
		Longitude: na,
		TimeZone:  na,
		UTC:       na,
		ChinaCode: na,
		PhoneCode: na,
		ISO2:      na,
		Continent: na,
	}

	arr := strings.Split(s, field_drt)
	if len(arr) < 4 {
		panic("s error" + s)
	}

	info.Country = arr[0]
	info.Province = arr[1]
	info.City = arr[2]
	info.Org = arr[3]
	info.Isp = arr[4]
	info.Latitude = arr[5]
	info.Longitude = arr[6]
	info.TimeZone = arr[7]
	info.UTC = arr[8]
	info.ChinaCode = arr[9]
	info.PhoneCode = arr[10]
	info.ISO2 = arr[11]
	info.Continent = arr[12]

	data, err := json.Marshal(info)
	if err != nil {
	}
	return string(data)
}
