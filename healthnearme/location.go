package healthnearme

type Location struct {
	Latitude     float64 `json:"latitude,string"`
	Longitude    float64 `json:"longitude,string"`
	HumanAddress string  `json:"human_address"`
}
