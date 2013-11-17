package healthnearme

type CondomDistributionSite struct {
        VenueType string `json:"venue_type"`
        ZipCode int `json:"zip_code"`
        Location Location `json:"location"`
        Address string `json:"address"`
        Name string `json:"name"`
        State string `json:"state"`
        City string `json:"city_"`
}

// {
//   "venue_type" : "CBO",
//   "zip_code" : "60647",
//   "location" : {
//     "needs_recoding" : false,
//     "longitude" : "-87.70113441599966",
//     "latitude" : "41.91742251900047",
//     "human_address" : "{\"address\":\"2957 W. Armitage\",\"city\":\"Chicago\",\"state\":\"IL\",\"zip\":\"60647\"}"
//   },
//   "address" : "2957 W. Armitage",
//   "name" : "Access Armitage",
//   "state" : "IL",
//   "city_" : "Chicago"
// }