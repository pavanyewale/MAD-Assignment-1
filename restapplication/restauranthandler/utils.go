package restauranthandler

import (
	"pavan/MAD-Assignment-1/domain"
)

func TransObjListToResponse(rests []*domain.Restaurant) RestaurantGetListRespDTO {
	resp := RestaurantGetListRespDTO{}
	for _, obj := range rests {
		restObj := RestaurantGetRespDTO{
			ID:           string(obj.DBID),
			Name:         obj.Name,
			Address:      obj.Address,
			AddressLine2: obj.AddressLine2,
			URL:          obj.URL,
			Outcode:      obj.Outcode,
			Postcode:     obj.Postcode,
			Rating:       obj.Rating,
			TypeOfFood:   obj.TypeOfFood,
		}
		resp.Restaurants = append(resp.Restaurants, restObj)
	}
	resp.Count = len(resp.Restaurants)
	return resp
}
