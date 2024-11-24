# go-ipapi

A Go package to query IP addresses using the [ipquery.io](https://ipquery.io) API.

## Overview

`go-ipapi` is a lightweight Go package that allows you to fetch detailed information about IP addresses, including ISP details, location data, and risk assessment, using the `ipquery.io` API. You can use this package to query individual IP addresses, perform bulk lookups, and even find out your own public IP address.

## Features

- Query detailed information for a specific IP address.
- Fetch geographical and ISP data.
- Check if an IP address is associated with VPNs, proxies, or data centers.
- Get the public IP address of your own machine.

## Installation

To install the package, run:

```bash
go get github.com/ipqwery/ipapi-go
```

## Importing the Package

```go
import "github.com/ipqwery/ipapi-go"
```

## Usage Examples

### Query a Specific IP Address

You can fetch detailed information about a specific IP address using the `QueryIP` function:

```go
package main

import (
    "fmt"
    "github.com/ipqwery/ipapi-go"
)

func main() {
    ipInfo, err := ipapi.QueryIP("8.8.8.8")
    if err != nil {
        fmt.Println("Error querying IP:", err)
        return
    }
    fmt.Printf("IP Info: %+v\n", ipInfo)
}
```

#### Example Output
```
IP Info: &{IP:8.8.8.8 ISP:{ASN:AS15169 Org:Google LLC ISP:Google LLC} Location:{Country:United States CountryCode:US City:Mountain View State:California ZipCode:94035 Latitude:37.386 Longitude:-122.0838 Timezone:America/Los_Angeles Localtime:2024-11-09T12:45:32} Risk:{IsMobile:false IsVPN:false IsTor:false IsProxy:false IsDatacenter:true RiskScore:0}}
```

### Fetch Your Own Public IP Address

To fetch the public IP address of the machine running your code, use the `QueryOwnIP` function:

```go
package main

import (
    "fmt"
    "github.com/ipqwery/go-ipapi"
)

func main() {
    ip, err := ipapi.QueryOwnIP()
    if err != nil {
        fmt.Println("Error fetching own IP:", err)
        return
    }
    fmt.Printf("Your IP: %s\n", ip)
}
```

#### Example Output
```
Your IP: 203.0.113.45
```

### Struct Definitions

Below are the definitions of the key data structures used in this package:

- **ISPInfo**: Contains information about the Internet Service Provider (ISP).
- **LocationInfo**: Holds geographical data such as country, city, and coordinates.
- **RiskInfo**: Provides information on potential risks associated with the IP (e.g., if it's a VPN or proxy).
- **IPInfo**: The main structure that combines all the information.

#### Example:
```go
type ISPInfo struct {
    ASN string `json:"asn,omitempty"`
    Org string `json:"org,omitempty"`
    ISP string `json:"isp,omitempty"`
}

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

type RiskInfo struct {
    IsMobile     bool `json:"is_mobile,omitempty"`
    IsVPN        bool `json:"is_vpn,omitempty"`
    IsTor        bool `json:"is_tor,omitempty"`
    IsProxy      bool `json:"is_proxy,omitempty"`
    IsDatacenter bool `json:"is_datacenter,omitempty"`
    RiskScore    int  `json:"risk_score,omitempty"`
}

type IPInfo struct {
    IP       string        `json:"ip"`
    ISP      *ISPInfo      `json:"isp,omitempty"`
    Location *LocationInfo `json:"location,omitempty"`
    Risk     *RiskInfo     `json:"risk,omitempty"`
}
```

## Testing

The package includes unit tests to ensure functionality. To run the tests:

```bash
go test -v
```

### Example Tests

Here's an example of how the tests are structured:

```go
package ipapi

import "testing"

func TestQueryIP(t *testing.T) {
    ipInfo, err := QueryIP("8.8.8.8")
    if err != nil {
        t.Fatalf("Failed to query IP: %v", err)
    }
    if ipInfo.IP != "8.8.8.8" {
        t.Errorf("Expected IP to be '8.8.8.8', got %s", ipInfo.IP)
    }
}

func TestQueryOwnIP(t *testing.T) {
    ip, err := QueryOwnIP()
    if err != nil {
        t.Fatalf("Failed to fetch own IP: %v", err)
    }
    if ip == "" {
        t.Error("Expected non-empty IP")
    }
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- This package uses the [ipquery.io](https://ipquery.io) API for IP information.
- Inspired by various IP geolocation services.

## Links

- [GitHub Repository](https://github.com/ipqwery/ipapi-go)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/ipqwery/ipapi-go)
