package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	// "net/http"

	"github.com/cgansen/health-near-me/healthnearme"
)

func main() {
	f, err := ioutil.ReadFile("./data/condom-distribution-sites.json")
	if err != nil {
		log.Printf("err loading file: %s", err)
	}

	var cds []healthnearme.CondomDistributionSite

	if err := json.Unmarshal(f, &cds); err != nil {
		log.Printf("err loading json: %s", err)
	}

	for _, item := range cds {
		log.Printf("cds: %#v", item)
	}

	return
}
