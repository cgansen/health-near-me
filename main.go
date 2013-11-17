package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cgansen/elastigo/api"
	"github.com/cgansen/elastigo/core"
	"github.com/cgansen/health-near-me/healthnearme"
)

func SearchHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("%#v", req)
	slat, slon, sdist := req.FormValue("lat"), req.FormValue("lon"), req.FormValue("dist")

	lat, err := strconv.ParseFloat(slat, 64)
	if err != nil {
		http.Error(w, "lat is required and must be a float, e.g. 41.42", 400)
		return
	}

	lon, err := strconv.ParseFloat(slon, 64)
	if err != nil {
		http.Error(w, "lon is required and must be a float, e.g. -87.88", 400)
		return
	}

	dist, err := strconv.ParseInt(sdist, 10, 64)
	if err != nil {
		http.Error(w, "dist is required and must be an integer", 400)
		return
	}

	query := fmt.Sprintf(`{"query":{
	  "filtered": {
	    "query": {
	      "match_all": {}
	    },
	    "filter": {
	      "geo_distance": {
	        "distance": "%dm",
	        "location.lat_lon": {
	          "lat": %f,
	          "lon": %f
	        }
	      }
	    }
	  }
	}
	}`, dist, lat, lon)

	log.Print(query)

	result, err := core.SearchRequest(true, "health-near-me", "health-provider", query, "", 0)
	if err != nil {
		log.Printf("error searching: %s", err)
		http.Error(w, "error searching index", 503)
		return
	}

	type Result struct {
		Name, Address, Phone string
		Location             healthnearme.Location
	}

	var hits []*Result

	for _, hit := range result.Hits.Hits {
		// unmarshal to a struct
		var cds healthnearme.CondomDistributionSite
		jsn, _ := hit.Source.MarshalJSON()
		if err := json.Unmarshal(jsn, &cds); err != nil {
			log.Printf("could not translate to struct: %s", err)
			http.Error(w, "error translating search results", 500)
			return
		}

		r := &Result{
			Name:     cds.Name,
			Address:  cds.Address,
			Location: cds.Location,
		}

		hits = append(hits, r)
	}

	jsn, err := json.MarshalIndent(hits, "", "  ")
	if err != nil {
		log.Print(err)
		http.Error(w, "error dumping search results to json", 500)
		return
	}

	_, err = w.Write(jsn)
	return

}

func main() {
	api.Domain = "localhost"

	http.HandleFunc("/search", SearchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
