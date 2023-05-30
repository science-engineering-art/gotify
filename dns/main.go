package main

import (
	"fmt"

	"github.com/miekg/dns"
)

func Resolve(domain string, qtype uint16) []dns.RR {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), qtype)
	m.RecursionDesired = true

	c := new(dns.Client)
	in, _, err := c.Exchange(m, "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return in.Answer
}

type dnsHandler struct{}

func (h *dnsHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(r)
	msg.Authoritative = true

	for _, question := range r.Question {
		fmt.Printf("Received query: %s\n", question.Name)
		answers := Resolve(question.Name, question.Qtype)
		msg.Answer = append(msg.Answer, answers...)
	}

	w.WriteMsg(msg)
}

func main() {
	handler := new(dnsHandler)
	server := &dns.Server{
		Addr:      "8900",
		Net:       "udp",
		Handler:   handler,
		UDPSize:   65535,
		ReusePort: true,
	}

	fmt.Println("Starting DNS server on port", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
	}
}
