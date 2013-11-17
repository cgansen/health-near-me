package healthnearme

type SubstanceAbuseProvider struct {
	Address           string   `json:"address"`
	State             string   `json:"state"`
	City              string   `json:"city"`
	PhoneNumber       string   `json:"phone_number"`
	ZipCode           int      `json:"zip_code,string"`
	ServiceModalities string   `json:"service_modalities"`
	TreatmentCenter   string   `json:"treatment_center"`
	PopulationServed  string   `json:"population_served"`
	PaymentMethod     string   `json:"payment_method"`
	Location          Location `json:"physical_address"`
}

// {
//   "phone_number" : "312-372-6707 x122",
//   "zip_code" : "60645",
//   "address" : "2049 W Jarvis St",
//   "service_modalities" : "RH",
//   "treatment_center" : "A Safe Haven",
//   "state" : "Illinois",
//   "population_served" : "Adults",
//   "payment_method" : "Sliding Scale",
//   "physical_address" : {
//     "needs_recoding" : false,
//     "longitude" : "-87.68244598055738",
//     "latitude" : "42.015078518290736",
//     "human_address" : "{\"address\":\"2049 Jarvis St\",\"city\":\"Chicago\",\"state\":\"ILLINOIS\",\"zip\":\"60645\"}"
//   },
//   "city" : "Chicago"
// }
