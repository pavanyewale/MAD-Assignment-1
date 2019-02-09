package restauranthandler

type RestaurantCreateReqDTO struct {
	Name         string  `json:"name" `
	Address      string  `json:"address" `
	AddressLine2 string  `json:"addressLine2" `
	URL          string  `json:"url"`
	Outcode      string  `json:"outcode" `
	Postcode     string  `json:"postcode"`
	Rating       float32 `json:"rating"`
	TypeOfFood   string  `json:"type_of_food"`
}

type RestaurantUpdateReqDTO struct {
	ID           string  `json:"id" `
	Name         string  `json:"name" `
	Address      string  `json:"address" `
	AddressLine2 string  `json:"addressLine2" `
	URL          string  `json:"url"`
	Outcode      string  `json:"outcode" `
	Postcode     string  `json:"postcode"`
	Rating       float32 `json:"rating"`
	TypeOfFood   string  `json:"type_of_food"`
}
