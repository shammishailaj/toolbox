package network

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/richardwilkes/toolbox/xio"
)

var sites = []string{
	"http://whatismyip.akamai.com/",
	"https://myip.dnsomatic.com/",
	"http://icanhazip.com/",
	"http://diagnostic.opendns.com/myip",
	"https://myexternalip.com/raw",
	"http://ifconfig.io/ip",
	"http://api.ipify.org/",
	"http://checkip.amazonaws.com/",
	"http://ident.me/",
	"https://canihazip.com/s",
	"https://tnx.nl/ip",
}

// ExternalIP returns your IP address as seen by external sites. It does this
// by iterating through a list of websites that will return your IP address as
// they see it. The first response with a valid IP address will be returned.
// timeout sets the maximum amount of time for each attempt.
func ExternalIP(timeout time.Duration) string {
	client := &http.Client{
		Timeout: timeout,
	}
	for _, site := range sites {
		if ip := externalIP(client, site); ip != "" {
			return ip
		}
	}
	return ""
}

func externalIP(client *http.Client, site string) string {
	if resp, err := client.Get(site); err == nil {
		defer xio.CloseIgnoringErrors(resp.Body)
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			if ip := net.ParseIP(strings.TrimSpace(string(body))); ip != nil {
				return ip.String()
			}
		}
	}
	return ""
}
