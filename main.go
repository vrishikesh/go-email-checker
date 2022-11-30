package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("could not read from stdin:", err)
	}
}

func checkDomain(domain string) {
	var (
		hasMX, hasSPF, hasDMARC bool
		spfRecord, dmarcRecord  string
	)

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Println(err)
	}

	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Println(err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Println(err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	log.Println(domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
