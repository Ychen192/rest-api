package db

import (
	"log"

	"gopkg.in/mgo.v2"
	"rest-api/model"
	"gopkg.in/mgo.v2/bson"
)

type CompaniesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "companies"
)

func (m *CompaniesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *CompaniesDAO) FindAll() ([]model.Company, error) {
	var movies []model.Company
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *CompaniesDAO) Insert(company model.Company) error {
	err := db.C(COLLECTION).Insert(&company)
	return err
}

func (m *CompaniesDAO) FindByName(name string) ([]model.Company, error){
	var companies []model.Company
	err := db.C(COLLECTION).Find(bson.M{"name":name}).All(&companies)
	return companies, err
}

func (m *CompaniesDAO) Delete(company model.Company) error{
	err := db.C(COLLECTION).Remove(bson.M{"name":company.Name})
	return err
}

func (m *CompaniesDAO) Update(company model.Company) error{
	err := db.C(COLLECTION).Update(bson.M{"name":company.Name}, &company)
	return err
}