package healthnearme

import "fmt"

type Location struct {
	Latitude     float64 `json:"latitude,string"`
	Longitude    float64 `json:"longitude,string"`
	HumanAddress string  `json:"human_address"`
	LatLon       string  `json:"lat_lon"`
}

func (l *Location) FormatLocation() {
	l.LatLon = fmt.Sprintf("%f,%f", l.Latitude, l.Longitude)
}
