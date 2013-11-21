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
	geo "github.com/kellydunn/golang-geo"
)

func SearchHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("%#v", req)
	slat, slon, sdist, styp := req.FormValue("lat"), req.FormValue("lon"), req.FormValue("dist"), req.FormValue("searchType")

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

	var queryMatch string
	if styp == "all" {
		queryMatch = `"match_all": {}`
	} else {
		typ, err := strconv.Atoi(styp)
		if err != nil {
			http.Error(w, "searchType is required and must be an integer", 400)
			return
		}
		queryMatch = fmt.Sprintf(`"term": { "provider_type": "%d"}`, typ)
	}

	query := fmt.Sprintf(`{
		"size": 100,		
		"query":{
			%s
		},
		"sort": [
			{
				"_geo_distance":{
					"location.lat_lon": {
						"lat": %f,
						"lon": %f						
					},
					"order": "asc",
					"unit": "mi"
				}
			}
		],
		"filter": {
			"geo_distance": {
				"distance": "%dm",
				"location.lat_lon": {
					"lat": %f,
					"lon": %f
				}
			}
		}
	}`, queryMatch, lat, lon, dist, lat, lon)

	log.Print(query)

	result, err := core.SearchRequest(true, "health-near-me", "health-provider", query, "", 0)
	if err != nil {
		log.Printf("error searching: %s", err)
		http.Error(w, "error searching index", 503)
		return
	}

	origin := geo.NewPoint(lat, lon)
	var hits []*healthnearme.HealthProvider

	for _, hit := range result.Hits.Hits {
		// unmarshal to a struct
		hp := &healthnearme.HealthProvider{}
		jsn, _ := hit.Source.MarshalJSON()
		if err := json.Unmarshal(jsn, hp); err != nil {
			log.Printf("could not translate to struct: %s", err)
			http.Error(w, "error translating search results", 500)
			return
		}

		hp.Distance = hp.CalcDistance(origin)
		hp.TypeName = hp.FriendlyTypeName()
		log.Printf("dist: %f", hp.Distance)
		hits = append(hits, hp)
	}

	jsn, err := json.MarshalIndent(hits, "", "  ")
	if err != nil {
		log.Print(err)
		http.Error(w, "error dumping search results to json", 500)
		return
	}

	w.Header().Add("Content-type", "application/json")
	// delim := ")]}',\n"
	delim := ""
	resp := fmt.Sprintf("%s%s(%s);", delim, req.FormValue("callback"), string(jsn))

	_, err = w.Write([]byte(resp))
	return

}

func main() {
	api.Domain = "localhost"

	http.HandleFunc("/search", SearchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
