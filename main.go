package main

import (
	"fmt"
	"os"
	dbrepo "pavan/MAD-Assignment-1/dbrepository"
	domain "pavan/MAD-Assignment-1/domain"
	mongoutils "pavan/MAD-Assignment-1/utils"
	"strings"
)

func main() {
	//pass mongohost through the environment
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	fmt.Println(repoaccess)
	var input string
	var result []*domain.Restaurant
	var err error
	if len(os.Args) > 1 {
		input = os.Args[1]
		arr := strings.Split(input, "=")
		switch arr[0] {
		case "--type_of_food":
			result, err = repoaccess.FindByTypeOfFood(arr[1])
		case "--postcode":
			result, err = repoaccess.FindByTypeOfPostCode(arr[1])
		default:
			fmt.Println("invalid argument")
			return
		}
		if err != nil {
			fmt.Println(err)
			//fatal.log(err)
			return
		}

		for _, res := range result {
			fmt.Println(res)
		}
	}
}
