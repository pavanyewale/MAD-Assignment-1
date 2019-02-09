package restauranthandler

import (
	"encoding/json"
	"io/ioutil"
	logger "log"
	"net/http"
	"pavan/MAD-Assignment-1/dbrepository"
	"pavan/MAD-Assignment-1/domain"
	customerrors "pavan/MAD-Assignment-1/restapplication/packages/errors"
	"pavan/MAD-Assignment-1/restapplication/packages/httphandlers"
	mthdroutr "pavan/MAD-Assignment-1/restapplication/packages/mthdrouter"
	"pavan/MAD-Assignment-1/restapplication/packages/resputl"

	"github.com/gorilla/mux"
)

type RestaurantHandler struct {
	httphandlers.BaseHandler
	mongoSession *dbrepository.MongoRepository
}

func NewRestaurantHandler(mongoSession *dbrepository.MongoRepository) *RestaurantHandler {
	return &RestaurantHandler{mongoSession: mongoSession}
}

func (p *RestaurantHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Printf("UserCrudHandler.ServeHTTP")
	response := mthdroutr.RouteAPICall(p, r)
	response.RenderResponse(w)
}

func (p *RestaurantHandler) Get(r *http.Request) resputl.SrvcRes {
	pathparams := mux.Vars(r)
	rsID := pathparams["id"]
	if rsID == "" {
		resp, err := p.mongoSession.GetAll()
		if err != nil {
			return resputl.ResponseCustomError(err)
		}
		res := TransObjListToResponse(resp)
		return resputl.Response200OK(res)
	}
	ID := domain.StringToID(rsID)
	resp, err := p.mongoSession.Get(ID)
	if err != nil {
		return resputl.ResponseCustomError(err)
	}
	return resputl.Response200OK(resp)
}

func (p *RestaurantHandler) Post(r *http.Request) resputl.SrvcRes {
	//logger.Printf("UserCrudHandler.post")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resputl.ResponseCustomError(err)
	}
	e, err := ValidateRestaurantCreateUpdateRequest(string(body))
	if e == false {
		return resputl.ProcessError(err, body)
		return resputl.SimpleBadRequest("Invalid Input Data")

	}
	logger.Printf("Received POST request to Create schedule %s ", string(body))
	var rd *RestaurantCreateReqDTO
	err = json.Unmarshal(body, &rd)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}

	restObj := &domain.Restaurant{Name: rd.Name, Address: rd.Address, AddressLine2: rd.AddressLine2, URL: rd.URL, Outcode: rd.Outcode, Postcode: rd.Postcode, Rating: rd.Rating, TypeOfFood: rd.TypeOfFood}

	//userObj := f.NewUser(requestdata.FirstName, requestdata.LastName, requestdata.Age)
	id, err := p.mongoSession.Store(restObj)
	if err != nil {
		//logger.Fatalf("Error while creating in DB: %v", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in writing to DB"), "")
	}
	return resputl.Response200OK(RestaurantCreateRespDTO{ID: string(id)})
}

func (p *RestaurantHandler) Delete(r *http.Request) resputl.SrvcRes {
	pathparams := mux.Vars(r)
	rsID := pathparams["id"]
	if rsID == "" {
		return resputl.SimpleBadRequest("Invalid Parameters")
	}
	ID := domain.StringToID(rsID)
	err := p.mongoSession.Delete(ID)
	if err != nil {
		return resputl.ResponseCustomError(err)
	}
	return resputl.Response200OK("deleted")

}

func (p *RestaurantHandler) Put(r *http.Request) resputl.SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resputl.ResponseCustomError(err)
	}
	e, err := ValidateRestaurantCreateUpdateRequest(string(body))
	if e == false {
		return resputl.ProcessError(err, body)
		return resputl.SimpleBadRequest("Invalid Input Data")

	}
	logger.Printf("Received POST request to Create schedule %s ", string(body))
	var rd *RestaurantUpdateReqDTO
	err = json.Unmarshal(body, &rd)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}

	restObj := &domain.Restaurant{DBID: domain.StringToID(rd.ID), Name: rd.Name, Address: rd.Address, AddressLine2: rd.AddressLine2, URL: rd.URL, Outcode: rd.Outcode, Postcode: rd.Postcode, Rating: rd.Rating, TypeOfFood: rd.TypeOfFood}

	//userObj := f.NewUser(requestdata.FirstName, requestdata.LastName, requestdata.Age)
	err = p.mongoSession.Update(restObj)
	if err != nil {
		//logger.Fatalf("Error while creating in DB: %v", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in writing to DB"), "")
	}
	return resputl.Response200OK("updated")
}
