package healthnearme

type Location struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	HumanAddress string  `json:"human_address"`
}
