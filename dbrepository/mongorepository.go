package dbrepository

import (
	domain "../domain"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"encoding/json"
	"bufio"
	"os"
)

//MongoRepository mongodb repo
type MongoRepository struct {
	mongoSession *mgo.Session
	db           string
}

var collectionName = "restaurant"

//NewMongoRepository create new repository
func NewMongoRepository(mongoSession *mgo.Session, db string) *MongoRepository {
	return &MongoRepository{
		mongoSession: mongoSession,
		db:           db,
	}
}

//Find a Restaurant
func (r *MongoRepository) Get(id domain.ID) (*domain.Restaurant, error) {
	result := domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"_id": id}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//Store a Restaurantrecord
func (r *MongoRepository) Store(b *domain.Restaurant) (domain.ID, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	if domain.ID(0) == b.DBID {
		b.DBID = domain.NewID()
	}

	_, err := coll.UpsertId(b.DBID, b)

	if err != nil {
		return domain.ID(0), err
	}
	return b.DBID, nil
}
//Import a restaurant record from a file to the db it returns noofrecords gets inserted into db and error if any 
func (r *MongoRepository) Import(filename string) (int,error){
	file,err:=os.Open(filename)
	if err!=nil{
		return 0,err
	}
	defer file.Close()
	scanner:=bufio.NewScanner(file)
	var rest=&domain.Restaurant{}
	recordcount:=0
	for scanner.Scan(){
			recordcount+=1
		json.Unmarshal([]byte(scanner.Text()),rest)
		rest.DBID=domain.NewID()
		_,err:=r.Store(rest)
		if	err!=nil{
			return recordcount,err
		}
	}
	if err:=scanner.Err();err!=nil {
		return recordcount,err
	}
	return recordcount,nil
}

//get all records from restaurant
func (r *MongoRepository) GetAll() ([]domain.Restaurant, error){
	result := []domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}
//Find all records by name
func (r *MongoRepository) FindByName(name string) ([]domain.Restaurant, error){
	result := []domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"name":name}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//delete record by id
func (r *MongoRepository) Delete(id domain.ID)  error {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Remove(bson.M{"_id":id})
	switch err {
	case nil:
		return nil
	case mgo.ErrNotFound:
		return domain.ErrNotFound
	default:
		return err
	}
}

//find all records by type of food
func (r *MongoRepository) FindByTypeOfFood(foodType string) ([]domain.Restaurant, error){
	result := []domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"type_of_food":foodType}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//find all records by postcode
func (r *MongoRepository) FindByTypeOfPostCode(postCode string) ([]domain.Restaurant, error){
	result := []domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"postcode":postCode}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}
