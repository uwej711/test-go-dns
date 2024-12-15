package main

import (
	"fmt"
	"github.com/miekg/dns"
	"time"
)

func main() {
	msg := new(dns.Msg)

	client := &dns.Client{Net: "udp", Timeout: time.Second * 1}

	msg.Id = dns.Id()
	msg.Opcode = dns.OpcodeQuery

	msg.Question = make([]dns.Question, 1)
	msg.Question[0] = dns.Question{Name: dns.Fqdn("www.google.de."), Qtype: dns.TypeA, Qclass: dns.ClassINET}

	conn, err := client.Dial("1.1.1.1:53")
	time.Sleep(500 * time.Microsecond)
	r, _, err := client.ExchangeWithConn(msg, conn)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)
}
