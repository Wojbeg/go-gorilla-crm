package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Wojbeg/go-gorilla-crm/database"
	"github.com/Wojbeg/go-gorilla-crm/models"
	"github.com/Wojbeg/go-gorilla-crm/models/dto"
	"github.com/gorilla/mux"
)

type infoRouter struct {
	r   *mux.Router
	rep database.Repository
}

func CreateNewInfoRouter(router *mux.Router, repository database.Repository) *infoRouter {
	infosRouter := infoRouter{
		r:   router,
		rep: repository,
	}
	return &infosRouter
}

func (rout *infoRouter) Init() {
	infos := rout.r.PathPrefix("/infos").Subrouter()
	infos.HandleFunc("/all", rout.GetInfo).Methods("GET")
	infos.HandleFunc("/byId/{infoId}", rout.GetInfoById).Methods("GET")
	infos.HandleFunc("/create", rout.CreateInfo).Methods("POST")
	infos.HandleFunc("/update/{infoId}", rout.UpdateInfo).Methods("PUT")
	infos.HandleFunc("/delete/{infoId}", rout.DeleteInfo).Methods("DELETE")
}

func (rout *infoRouter) GetInfo(w http.ResponseWriter, r *http.Request) {
	infos := models.GetAllInfos(rout.rep)

	inf := dto.InfoModelsToResponses(&infos)

	sendJSONResponse(w, inf, http.StatusOK)
}

func (rout *infoRouter) GetInfoById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	infoId := vars["infoId"]
	ID, err := strconv.ParseInt(infoId, 0, 0)

	if err != nil {
		sendJSONErrorResponse(w, "Invalid Id - should be a number", http.StatusBadRequest)
		return
	}

	infoDetails, err := models.GetInfoById(ID, rout.rep)

	if err != nil {
		sendJSONErrorResponse(w, "Resource Not Found - Incorrect Id", http.StatusNotFound)
		return
	}

	infoDto := dto.InfoModelToResponse(infoDetails)

	sendJSONResponse(w, infoDto, http.StatusOK)
}

func (rout *infoRouter) CreateInfo(w http.ResponseWriter, r *http.Request) {
	info := &dto.PersonInfoRequest{}

	err := json.NewDecoder(r.Body).Decode(info)
	if err != nil {
		sendJSONErrorResponse(w, "Invalid Info struct", http.StatusBadRequest)
		return
	}

	insertedInfo, err := models.CreatePersonInfo(info.InfoRequestToModel(), rout.rep)

	if err != nil {
		sendJSONErrorResponse(w, "Could not create info", http.StatusInternalServerError)
		return
	}

	infoDto := dto.InfoModelToResponse(insertedInfo)

	sendJSONResponse(w, infoDto, http.StatusOK)
}

func (rout *infoRouter) DeleteInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	infoId := vars["infoId"]
	ID, err := strconv.ParseInt(infoId, 0, 0)

	if err != nil {
		sendJSONErrorResponse(w, "Invalid Id - should be a number", http.StatusBadRequest)
		return
	}

	deleted, err := models.DeleteInfo(ID, rout.rep)

	if err != nil {
		sendJSONErrorResponse(w, "Could not delete info", http.StatusNotFound)
		return
	}

	infoDto := dto.InfoModelToResponse(deleted)

	sendJSONResponse(w, infoDto, http.StatusOK)
}

func (rout *infoRouter) UpdateInfo(w http.ResponseWriter, r *http.Request) {

	info := &dto.PersonInfoRequest{}

	err := json.NewDecoder(r.Body).Decode(info)

	if err != nil {
		sendJSONErrorResponse(w, "Invalid Info struct", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	infoId := vars["infoId"]
	ID, err := strconv.ParseInt(infoId, 0, 0)

	if err != nil {
		sendJSONErrorResponse(w, "Invalid Id - should be a number", http.StatusBadRequest)
		return
	}

	infoToUpdate, err := models.GetInfoById(ID, rout.rep)

	if err != nil {
		sendJSONErrorResponse(w, "Resource Not Found - Incorrect Id", http.StatusNotFound)
		return
	}

	info.CopyInfoRequestToModel(infoToUpdate)

	infoToUpdate, err = models.Save(infoToUpdate, rout.rep)

	if err != nil {
		sendJSONErrorResponse(w, "Could not update this info", http.StatusInternalServerError)
		return
	}

	infoDto := dto.InfoModelToResponse(infoToUpdate)

	sendJSONResponse(w, infoDto, http.StatusOK)
}

func sendJSONErrorResponse(w http.ResponseWriter, errorString string, status int) {
	resp := make(map[string]string)
	resp["error"] = errorString

	sendJSONResponse(w, resp, status)
}

func sendJSONResponse(w http.ResponseWriter, data interface{}, status int) {
	res, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	_, err = w.Write(res)

	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
		return
	}
}
