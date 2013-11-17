package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/cgansen/elastigo/api"
	"github.com/cgansen/elastigo/core"
	"github.com/cgansen/health-near-me/healthnearme"
)

const INDEX_NAME = "health-near-me"
const INDEX_TYPE = "health-provider"

func main() {
	api.Domain = "localhost"

	f, err := ioutil.ReadFile("./data/condom-distribution-sites.json")
	if err != nil {
		log.Printf("err loading file: %s", err)
	}

	var cds []healthnearme.CondomDistributionSite

	if err := json.Unmarshal(f, &cds); err != nil {
		log.Printf("err loading json: %s", err)
	}

	for _, item := range cds {
		item.FormatLocation()
		_, err := core.Index(true, INDEX_NAME, INDEX_TYPE, "", item)
		if err != nil {
			log.Print(err)
		}
	}

	//--------------------------------------

	f, err = ioutil.ReadFile("./data/licensed-substance-abuse-providers.json")
	if err != nil {
		log.Printf("err loading file: %s", err)
	}

	var saps []healthnearme.SubstanceAbuseProvider

	if err := json.Unmarshal(f, &saps); err != nil {
		log.Printf("err loading json: %s", err)
	}

	for _, item := range saps {
		item.FormatLocation()
		_, err := core.Index(true, INDEX_NAME, INDEX_TYPE, "", item)
		if err != nil {
			log.Print(err)
		}
	}

	//--------------------------------------

	f, err = ioutil.ReadFile("./data/mental-health-clinics.json")
	if err != nil {
		log.Printf("err loading file: %s", err)
	}

	var mhcs []healthnearme.MentalHealthClinic

	if err := json.Unmarshal(f, &mhcs); err != nil {
		log.Printf("err loading json: %s", err)
	}

	for _, item := range mhcs {
		item.FormatLocation()
		_, err := core.Index(true, INDEX_NAME, INDEX_TYPE, "", item)
		if err != nil {
			log.Print(err)
		}
	}

	//--------------------------------------

	f, err = ioutil.ReadFile("./data/sti-specialty-clinics.json")
	if err != nil {
		log.Printf("err loading file: %s", err)
	}

	var sticscs []healthnearme.STISpecialtyClinic

	if err := json.Unmarshal(f, &sticscs); err != nil {
		log.Printf("err loading json: %s", err)
	}

	for _, item := range sticscs {
		item.FormatLocation()
		_, err := core.Index(true, INDEX_NAME, INDEX_TYPE, "", item)
		if err != nil {
			log.Print(err)
		}
	}

	//--------------------------------------

	f, err = ioutil.ReadFile("./data/wic-clinics.json")
	if err != nil {
		log.Printf("err loading file: %s", err)
	}

	var wiccs []healthnearme.WICClinic

	if err := json.Unmarshal(f, &wiccs); err != nil {
		log.Printf("err loading json: %s", err)
	}

	for _, item := range wiccs {
		item.FormatLocation()
		_, err := core.Index(true, INDEX_NAME, INDEX_TYPE, "", item)
		if err != nil {
			log.Print(err)
		}
	}

	return
}
