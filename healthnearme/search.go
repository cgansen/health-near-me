package healthnearme

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cgansen/elastigo/core"
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
