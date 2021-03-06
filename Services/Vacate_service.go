package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"net/http"
)

func Getvacatedetails(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllVacatedetails()

	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type Vacatesingle struct {
	Vacateid int
}
type Tenantvacatesingle struct {
	Tenantid int
}
func Getsinglevacate(w http.ResponseWriter, r *http.Request) {

	var GetVacate Vacatesingle
	err := json.NewDecoder(r.Body).Decode(&GetVacate)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	Data := Db.GetSinglevacate_Db(GetVacate.Vacateid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
func GetsingleTenantvacate(w http.ResponseWriter, r *http.Request) {

	var GetVacate Tenantvacatesingle
	err := json.NewDecoder(r.Body).Decode(&GetVacate)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	Data := Db.GetSinglevacateTenant_Db(GetVacate.Tenantid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func InsertVacate(w http.ResponseWriter, r *http.Request) {

	var Insvacate Model.Vacate
	err := json.NewDecoder(r.Body).Decode(&Insvacate)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	//Insert Vacate Details
	Db.Insertvacate_Db(Insvacate)

}

func UpdatetVacate(w http.ResponseWriter, r *http.Request) {

	var Upvacate Model.Vacatestatus
	err := json.NewDecoder(r.Body).Decode(&Upvacate)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	//Insert Vacate Details
	Db.Updatevacate_Db(Upvacate)

}
