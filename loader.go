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

	to_load := map[healthnearme.HealthProviderType]string{
		healthnearme.CondomDistributionSite: "condom-distribution-sites.json",
		healthnearme.SubstanceAbuseProvider: "licensed-substance-abuse-providers.json",
		healthnearme.MentalHealthClinic:     "mental-health-clinics.json",
		healthnearme.STISpecialtyClinic:              "sti-specialty-clinics.json",
		healthnearme.WICClinic:              "wic-clinics.json",
	}

	for typ, filename := range to_load {
		f, err := ioutil.ReadFile("./data/" + filename)
		if err != nil {
			log.Printf("err loading file: %s", err)
		}

		var sites []healthnearme.HealthProvider

		if err := json.Unmarshal(f, &sites); err != nil {
			log.Printf("err loading json: %s", err)
		}

		for _, item := range sites {
			item.Type = typ
			item.FormatLocation()
			_, err := core.Index(true, INDEX_NAME, INDEX_TYPE, "", item)
			if err != nil {
				log.Print(err)
			}
		}
	}

	return
}
