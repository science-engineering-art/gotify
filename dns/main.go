package main

import (
	lightdns "github.com/openmohan/lightdns"
)

var records = map[string]string{
	"mail.amazon.com":  "192.162.1.2",
	"paste.amazon.com": "191.165.0.3",
}

func lookupFunc(string) (string, error) {
	//Do some action
	//Get data from DB
	//Process it further more
	return "192.2.2.1", nil
}

func main() {
	var googleRecords = map[string]string{
		"mail.google.com":  "192.168.0.2",
		"paste.google.com": "192.168.0.3",
	}
	var microsoftRecords = map[string]string{
		"mail.microsoft.com":  "192.168.0.78",
		"paste.microsoft.com": "192.168.0.25",
	}
	dns := lightdns.NewDNSServer(8900)
	dns.AddZoneData("google.com", googleRecords, nil, lightdns.DNSForwardLookupZone)
	dns.AddZoneData("microsoft.com", microsoftRecords, nil, lightdns.DNSForwardLookupZone)

	/* Incase if the records are not static or to be taken from DB or from any other sources
	lookupFunc method can be used.append*/
	dns.AddZoneData("amazon.com", nil, lookupFunc, lightdns.DNSForwardLookupZone)
	dns.StartAndServe()
}
