package healthnearme

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/cgansen/elastigo/core"
	geo "github.com/kellydunn/golang-geo"
)

// DoSearch performs a search for HealthProvider objects in the ES index
func DoSearch(lat, lon float64, dist int64, typ string) (result core.SearchResult, err error) {
	var queryMatch string

	if typ == "all" {
		queryMatch = `"match_all": {}`
	} else {
		ityp, serr := strconv.Atoi(typ)
		if serr != nil {
			return result, serr
		}

		queryMatch = fmt.Sprintf(`"term": { "provider_type": "%d"}`, ityp)
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

	result, err = core.SearchRequest(true, "health-near-me", "health-provider", query, "", 0)

	return
}

func LoadResults(result core.SearchResult, origin *geo.Point) (hits []*HealthProvider, err error) {
	for _, hit := range result.Hits.Hits {
		// unmarshal to a struct
		hp := &HealthProvider{}
		jsn, _ := hit.Source.MarshalJSON()
		if err = json.Unmarshal(jsn, hp); err != nil {
			log.Printf("could not translate to struct: %s", err)
			return
		}

		hp.Distance = hp.CalcDistance(origin)
		hp.TypeName = hp.FriendlyTypeName()
		hits = append(hits, hp)
	}

	return
}
