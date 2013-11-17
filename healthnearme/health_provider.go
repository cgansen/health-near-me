package healthnearme

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
