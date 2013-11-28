package main

import (
	"encoding/json"
	"flag"
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

var tmplPath string

func init() {
	flag.StringVar(&tmplPath, "tmpl", "../tmpl/", "path to templates")
	flag.Parse()
}

// Perform a search for a SMS user.
func SMSSearchHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s %s", req.Method, req.RequestURI, req.URL.RawQuery, req.Header.Get("User-Agent"))

	if err := req.ParseForm(); err != nil {
		log.Printf("error parsing form: %s", err)
		http.Error(w, "error parsing form body", 500)
		return
	}

	// TODO(cgansen):
	// support sessions
	// search regex

	search := req.FormValue("Body")
	log.Printf("sms search: %s", search)

	cmd := strings.TrimSpace(strings.ToLower(search))
	switch cmd {
	case "list", "list services":
		t, err := template.ParseFiles(tmplPath + "help.txt")
		if err != nil {
			log.Printf("error loading template: %s", err)
			http.Error(w, "error loading template", 500)
			return
		}

		if err := t.Execute(w, nil); err != nil {
			log.Printf("error executing template: %s", err)
			http.Error(w, "error executing template", 500)
			return
		}

		return
	default:
		// split query
		pieces := strings.Split(cmd, "near")

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
		hits, err := healthnearme.LoadResults(result, point)
		if err != nil {
			log.Print(err)
			http.Error(w, "error processing search results", 500)
			return
		}
		log.Printf("%d results for %s", len(hits), cmd)

		t, err := template.New("nearby_providers.txt").Funcs(template.FuncMap{"round": strconv.FormatFloat}).ParseFiles(tmplPath + "nearby_providers.txt")
		if err != nil {
			log.Print("template error: ", err)
			http.Error(w, "error loading template", 500)
			return
		}

		ctxt := map[string]interface{}{
			"Count":    len(hits),
			"Location": location,
			"Results":  hits,
		}

		if err := t.Execute(w, ctxt); err != nil {
			log.Print(err)
			http.Error(w, "error writing results", 500)
			return
		}

		w.Header().Add("Content-type", "text/xml")
		return
	}

	return
}

func SearchHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s %s", req.Method, req.RequestURI, req.RemoteAddr, req.Header.Get("User-Agent"))
	
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

        log.Printf("http search: %f,%f %d %s", lat, lon, dist, styp)
        
	result, err := healthnearme.DoSearch(lat, lon, dist, styp)
	if err != nil {
		log.Printf("error searching: %s", err)
		http.Error(w, "error searching index", 503)
		return
	}

	hits, err := healthnearme.LoadResults(result, geo.NewPoint(lat, lon))
	if err != nil {
		log.Print(err)
		http.Error(w, "error processing search results", 500)
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

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
	return
}

func main() {
	api.Domain = "localhost"

	http.HandleFunc("/sms_search", SMSSearchHandler)
	http.HandleFunc("/search", SearchHandler)
	http.HandleFunc("/healthcheck", HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
