package healthnearme

import geo "github.com/kellydunn/golang-geo"

type HealthProvider struct {
	Name        string             `json:"name"`
	State       string             `json:"state"`
	Address     string             `json:"address"`
	City        string             `json:"city"`
	ZipCode     int                `json:"zip_code,string"`
	Location    Location           `json:"location"`
	PhoneNumber string             `json:"phone"`
	Hours       string             `json:"hours_of_operation"`
	Type        HealthProviderType `json:"provider_type"`
	Distance    float64            `json:"distance,string,omitempty"`
}

type HealthProviderType int

const (
	CondomDistributionSite HealthProviderType = iota
	SubstanceAbuseProvider
	MentalHealthClinic
	STISpecialtyClinic
	WICClinic
)

func (hp *HealthProvider) FormatLocation() {
	hp.Location.FormatLocation()
}

// Return the distance, in miles, between the HealthProvider and a given Point
func (hp HealthProvider) CalcDistance(p *geo.Point) float64 {
	return p.GreatCircleDistance(geo.NewPoint(hp.Location.Latitude, hp.Location.Longitude)) * 0.621371
}
