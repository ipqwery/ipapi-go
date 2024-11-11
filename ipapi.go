package ipapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BaseURL = "https://api.ipquery.io/"

// ISPInfo represents information about the ISP of an IP address.
type ISPInfo struct {
	ASN string `json:"asn,omitempty"`
	Org string `json:"org,omitempty"`
	ISP string `json:"isp,omitempty"`
}

// LocationInfo represents geographical information about an IP address.
type LocationInfo struct {
	Country     string  `json:"country,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
	City        string  `json:"city,omitempty"`
	State       string  `json:"state,omitempty"`
	ZipCode     string  `json:"zipcode,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	Timezone    string  `json:"timezone,omitempty"`
	Localtime   string  `json:"localtime,omitempty"`
}

// RiskInfo represents risk information about an IP address.
type RiskInfo struct {
	IsMobile     bool `json:"is_mobile,omitempty"`
	IsVPN        bool `json:"is_vpn,omitempty"`
	IsTor        bool `json:"is_tor,omitempty"`
	IsProxy      bool `json:"is_proxy,omitempty"`
	IsDatacenter bool `json:"is_datacenter,omitempty"`
	RiskScore    int  `json:"risk_score,omitempty"`
}

// IPInfo represents all the information returned by the API.
type IPInfo struct {
	IP       string       `json:"ip"`
	ISP      *ISPInfo     `json:"isp,omitempty"`
	Location *LocationInfo `json:"location,omitempty"`
	Risk     *RiskInfo     `json:"risk,omitempty"`
}

// QueryIP fetches information for a specific IP address.
func QueryIP(ip string) (*IPInfo, error) {
	url := fmt.Sprintf("%s%s", BaseURL, ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch IP info: status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ipInfo IPInfo
	if err := json.Unmarshal(body, &ipInfo); err != nil {
		return nil, err
	}

	return &ipInfo, nil
}

// QueryOwnIP fetches information about the current machine's public IP.
func QueryOwnIP() (string, error) {
	resp, err := http.Get(BaseURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch own IP: status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
