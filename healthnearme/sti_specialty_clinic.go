package healthnearme

type STISpecialtyClinic struct {
	Address     string   `json:"address"`
	State       string   `json:"state"`
	City        string   `json:"city"`
	PhoneNumber string   `json:"phone_number"`
	ZipCode     int      `json:"zip_code,string"`
	Location    Location `json:"physical_address"`
	Hours       string   `json:"hours_of_operation"`
	Name        string   `json:"site_name"`
}

func (m *STISpecialtyClinic) FormatLocation() {
	m.Location.FormatLocation()
}

// {
//   "zip" : "60644",
//   "phone" : "(312) 746-4871",
//   "fax" : "(312) 746-4637",
//   "location" : {
//     "needs_recoding" : false,
//     "longitude" : "-87.74961559423011",
//     "latitude" : "41.880519084171674",
//     "human_address" : "{\"address\":\"4958 W. Madison\",\"city\":\"Chicago\",\"state\":\"IL\",\"zip\":\"60644\"}"
//   },
//   "address" : "4958 W. Madison",
//   "hours_of_operation" : "Mon. and Wed., 8 a.m. - 4 p.m.; Tue., Thu.: 10 a.m. - 6 p.m.",
//   "state" : "IL",
//   "site_name" : "South Austin STI Specialty Clinic",
//   "city" : "Chicago"
// }
