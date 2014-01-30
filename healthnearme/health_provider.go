package healthnearme

import (
	"errors"
	"strings"

	geo "github.com/kellydunn/golang-geo"
)

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
	TypeName    string             `json:"type_name,omitempty"`
}

type HealthProviderType int

const (
	Dummy HealthProviderType = iota // use this to signify "all" in some use cases
	CondomDistributionSite
	SubstanceAbuseProvider
	MentalHealthClinic
	STISpecialtyClinic
	WICClinic
	CommunityServiceCenter
	CoolingCenter
	SeniorCenter
	WarmingCenter
	Hospital
	MedicationDisposal
)

func (hp *HealthProvider) FormatLocation() {
	hp.Location.FormatLocation()
}

// Return the distance, in miles, between the HealthProvider and a given Point
func (hp HealthProvider) CalcDistance(p *geo.Point) float64 {
	return p.GreatCircleDistance(geo.NewPoint(hp.Location.Latitude, hp.Location.Longitude)) * 0.621371
}

func (hp HealthProvider) FriendlyTypeName() string {
	switch hp.Type {
	case CondomDistributionSite:
		return "Condom Distribution Site"
	case SubstanceAbuseProvider:
		return "Licensed Substance Abuse Provider"
	case MentalHealthClinic:
		return "Mental Health Clinic"
	case STISpecialtyClinic:
		return "STI Specialty Clinic"
	case WICClinic:
		return "WIC Clinic"
	case CommunityServiceCenter:
		return "Community Service Center"
	case CoolingCenter:
		return "Cooling Center"
	case SeniorCenter:
		return "Senior Center"
	case WarmingCenter:
		return "Warming Center"
	case Hospital:
		return "Hospital"
	case MedicationDisposal:
		return "Medication Disposal Site"
	}

	return ""
}

// Given a search term, figure out what kind of HP it refers to
func SearchType(term string) (HealthProviderType, error) {
	cleaned := strings.ToLower(strings.TrimSpace(term))

	switch cleaned {
	case "condom", "condoms", "free condoms":
		return CondomDistributionSite, nil
	case "substance abuse", "substance abuse provider", "licensed substance abuse provider":
		return SubstanceAbuseProvider, nil
	case "mental health", "mental health clinic":
		return MentalHealthClinic, nil
	case "sti", "std", "sti clinic", "std clinic", "sti specialty clinic":
		return STISpecialtyClinic, nil
	case "wic", "wic clinic":
		return WICClinic, nil
	case "community service center", "service center":
		return CommunityServiceCenter, nil
	case "cooling", "cooling center":
		return CoolingCenter, nil
	case "warming", "warming center":
		return WarmingCenter, nil
	case "services", "everything", "all", "anything":
		return Dummy, nil
	case "hospital", "hospitals":
		return Hospital, nil
	case "medication disposal", "drug disposal", "dropoff", "drop off", "pharmaceutical":
		return MedicationDisposal, nil
	default:
		return Dummy, errors.New("unknown search type")
	}
}
