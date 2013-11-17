package healthnearme

type WICClinic struct {
	Address              string   `json:"street_address"`
	State                string   `json:"state"`
	City                 string   `json:"city"`
	PhoneNumber          string   `json:"phone_number"`
	ZipCode              int      `json:"zipcode,string"`
	Location             Location `json:"physical_address"`
	Hours                string   `json:"hours_of_operation"`
	Name                 string   `json:"site_name"`
	WIC                  string   `json:"wic"`
	SiteNumber           int      `json:"site_number,string"`
	FamilyCaseManagement string   `json:'family_case_management'`
}

// {
//   "location" : {
//     "needs_recoding" : false,
//     "longitude" : "-87.66739075899966",
//     "latitude" : "41.85227227900049",
//     "human_address" : "{\"address\":\"1643 W. Cermak St.\",\"city\":\"Chicago\",\"state\":\"IL\",\"zip\":\"60608\"}"
//   },
//   "hours_of_operation" : "Monday - Friday (8:00 a.m. - 4:00 p.m.)",
//   "wic" : "Y",
//   "zipcode" : "60608",
//   "state" : "IL",
//   "site_number" : "17",
//   "site_name" : "Lower West Side NHC",
//   "fax_1" : "312-747-1648",
//   "city" : "Chicago",
//   "phone_2" : "312-747-1651",
//   "family_case_management" : "Y",
//   "phone_1" : "312-747-1650",
//   "phone_4" : "312-747-1653",
//   "phone_3" : "312-747-1652",
//   "street_address" : "1643 W. Cermak St. "
// }
