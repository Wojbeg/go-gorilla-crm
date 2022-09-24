package routes

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Wojbeg/go-gorilla-crm/database"
	"github.com/Wojbeg/go-gorilla-crm/models"
	"github.com/Wojbeg/go-gorilla-crm/models/dto"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var formDecoder = schema.NewDecoder()

type staticRouter struct {
	r         *mux.Router
	rep       database.Repository
	templates map[string]*template.Template
}

func CreateNewStaticRouter(router *mux.Router, repository database.Repository) *staticRouter {
	staticRouter := staticRouter{
		r:         router,
		rep:       repository,
		templates: make(map[string]*template.Template),
	}
	return &staticRouter
}

func (rout *staticRouter) Init() {
	rout.createTemplates()
	rout.handleStaticFiles()

	rout.r.HandleFunc("/", rout.GetHomePage).Methods("GET")
	rout.r.HandleFunc("/create", rout.GetCreatePage).Methods("GET")
	rout.r.HandleFunc("/createInfo", rout.PostCreatePage).Methods("POST")
	rout.r.HandleFunc("/update/{id:[0-9]+}", rout.UpdateInfo).Methods("POST")
	rout.r.HandleFunc("/delete/{id:[0-9]+}", rout.DeleteInfo).Methods("GET")
	rout.r.HandleFunc("/edit/{id:[0-9]+}", rout.EditInfo).Methods("GET")
}

func (rout *staticRouter) createTemplates() {

	rout.templates["home.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/home.html"))
	rout.templates["create.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/create.html"))
	rout.templates["edit.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/edit.html"))
	rout.templates["error.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/error.html"))

}

func (rout *staticRouter) handleStaticFiles() {
	fileServer := http.FileServer(http.Dir("./static"))
	rout.r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))
}

func (rout *staticRouter) GetHomePage(w http.ResponseWriter, r *http.Request) {

	infos := models.GetAllInfos(rout.rep)

	inf := dto.InfoModelsToResponses(&infos)

	err := rout.templates["home.html"].ExecuteTemplate(w, "base", inf)

	if err != nil {
		rout.executeError(w, err)
		return
	}
}

func (rout *staticRouter) GetCreatePage(w http.ResponseWriter, r *http.Request) {

	err := rout.templates["create.html"].ExecuteTemplate(w, "base", nil)

	if err != nil {
		rout.executeError(w, err)
		return
	}
}

func (rout *staticRouter) PostCreatePage(w http.ResponseWriter, r *http.Request) {

	info := &dto.PersonInfoRequest{}
	r.ParseForm()
	err := formDecoder.Decode(info, r.PostForm)

	if err != nil {
		rout.executeError(w, err)
		return
	}

	_, err = models.CreatePersonInfo(info.InfoRequestToModel(), rout.rep)

	if err != nil {
		rout.executeError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (rout *staticRouter) UpdateInfo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	infoId := vars["id"]
	ID, err := strconv.ParseInt(infoId, 0, 0)

	if err != nil {
		rout.executeError(w, err)
		return
	}

	info := &dto.PersonInfoRequest{}
	r.ParseForm()
	err = formDecoder.Decode(info, r.PostForm)

	if err != nil {
		rout.executeError(w, err)
		return
	}

	infoToUpdate, err := models.GetInfoById(ID, rout.rep)

	if err != nil {
		rout.executeError(w, err)
		return
	}

	info.CopyInfoRequestToModel(infoToUpdate)

	_, err = models.Save(infoToUpdate, rout.rep)

	if err != nil {
		rout.executeError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (rout *staticRouter) DeleteInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	infoId := vars["id"]
	ID, err := strconv.ParseInt(infoId, 0, 0)

	if err != nil {
		rout.executeError(w, err)
		return
	}

	_, err = models.DeleteInfo(ID, rout.rep)

	if err != nil {
		rout.executeError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (rout *staticRouter) EditInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	infoId := vars["id"]
	ID, err := strconv.ParseInt(infoId, 0, 0)

	if err != nil {
		rout.executeError(w, err)
		return
	}

	infoDetails, err := models.GetInfoById(ID, rout.rep)

	if err != nil {
		rout.executeError(w, err)
		return
	}

	infoDto := dto.InfoModelToResponse(infoDetails)
	
	err = rout.templates["edit.html"].ExecuteTemplate(w, "base", infoDto)

	if err != nil {
		rout.executeError(w, err)
		return
	}

}

func (rout *staticRouter) executeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	rout.templates["error.html"].ExecuteTemplate(w, "base", err.Error())
}
