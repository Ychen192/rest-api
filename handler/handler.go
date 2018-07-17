package handler

import (
	. "rest-api/db"
	. "rest-api/model"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

var dao = CompaniesDAO{}

func AllCompaniesEndPoint(w http.ResponseWriter) {
	companies, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, companies)
}

// POST a new movie
func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var company Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Insert(company); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, company)
}

func FindCompanyEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for e := range params {
		println(e)
	}
	companyName, err := dao.FindByName(params["name"])
	if (err != nil) {
		respondWithError(w, http.StatusBadRequest, "Invalid company name")
		return
	}
	respondWithJson(w, http.StatusOK, companyName)
}

func DeleteCompanyEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var company Company

	if err := json.NewDecoder(r.Body).Decode(company); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request paylaod")
		return
	}

	if err := dao.Delete(company); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, company)
}

func UpdateCompanyEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var company Company

	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request paylaod")
		return
	}

	if err := dao.Update(company); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, company)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
