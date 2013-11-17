package healthnearme

type MentalHealthClinic struct {
	PhoneNumber string   `json:"phone"`
	Location    Location `json:"location"`
	Hours       string   `json:"hours_of_operation"`
	ZipCode     int      `json:"zipcode,string"`
	Address     string   `json:"street_address"`
	State       string   `json:"state"`
	City        string   `json:"city"`
	Name        string   `json:"site_name"`
}

//   "phone" : "(312) 747-7496",
//   "location" : {
//     "needs_recoding" : false,
//     "longitude" : "-87.64157740963851",
//     "latitude" : "41.779780316087795",
//     "human_address" : "{\"address\":\"641 63rd St\",\"city\":\"Chicago\",\"state\":\"IL\",\"zip\":\"60621\"}"
//   },
//   "hours_of_operation" : "Mon - Fri: 8:30 am – 4:30 pm",
//   "zipcode" : "60621",
//   "state" : "IL",
//   "street_address" : "641 W. 63rd St",
//   "site_name" : "Englewood MHC",
//   "city" : "Chicago"
// }
