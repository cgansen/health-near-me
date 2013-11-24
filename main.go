package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/cgansen/elastigo/api"
	"github.com/cgansen/health-near-me/healthnearme"
	geo "github.com/kellydunn/golang-geo"
)

// Perform a search for a SMS user.
func SMSSearchHandler(w http.ResponseWriter, req *http.Request) {
	// TODO(cgansen):
	// support sessions
	// search regex

	search := req.FormValue("body")
	switch cmd = strings.TrimSpace(strings.ToLower(search)) {
	case "help":
		t, err := template.ParseFiles("./tmpl/help.txt")
		if err != nil {
			// handle
		}

		if err := t.Execute(w, nil); err != nil {
			// handle
		}

		return
	default:
		// split query
		pieces := strings.Split(cmd, "near")
		log.Printf("pieces: %#v", pieces)

		// term := strings.TrimSpace(pieces[0])
		location := strings.TrimSpace(pieces[1])

		// geocode
		geocoder := &geo.GoogleGeocoder{}
		point, err := geocoder.Geocode(pieces[1])
		if err != nil {
			// handle
			log.Printf("error geocoding: %s, location is: %s", err, location)
			http.Error(w, "error geocoding", 500)
			return
		}

		log.Printf("geocoded %s to %#v", location, point)

		// TODO map term to searchType

		// lookup
		result, err := healthnearme.DoSearch(point.Lat(), point.Lng(), 1609, "all")

		// respond
		// FIXME temp for now

		w.Header().Add("Content-type", "text/xml")
		resp := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<Response><Message>Got %d hits for %s</Message></Response>", len(result.Hits.Hits), search)
		w.Write([]byte(resp))
		return
	}

	return
}

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

	if styp == "" {
		http.Error(w, "searchType is required and must be an integer or 'all'", 400)
		return
	}

	result, err := healthnearme.DoSearch(lat, lon, dist, styp)

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

	http.HandleFunc("/sms_search", SMSSearchHandler)
	http.HandleFunc("/search", SearchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
