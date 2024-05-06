package utils

import (
	"net/http"
	"strings"
	"text/template"
)

type homePageStruct struct {
	CompanyList []ApiData
}

type mapPageStruct struct {
	ImagePath   string
	CompanyList []ApiData
	MapAction   string
}

const IMAGEPATH = "./assets/nameofimage.png"

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	// Query all users from the database
	details := FetchCompanyData(db_conn)
	homePageData := homePageStruct{
		CompanyList: details,
	}
	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/index.html")
	ThrowHttpError(w, err)

	// Execute the template
	err = tmpl.Execute(w, struct {
		HomePageData homePageStruct
	}{homePageData})
	ThrowHttpError(w, err)
}

func MapHandler(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Query().Get("action")
	// Check if the request is a POST request
	if r.Method == http.MethodPost {
		const IMAGEPATHUPDATE = "./assets/updatenameofimage.png"

		initValue := r.Form.Get("changeMap")
		var changedMapArr1 string
		for _, value := range strings.Split(initValue, ",") {
			// value = 45
			changedMapArr1 += value + ","
		}

		details1 := FetchCompanyDataFromSelectedId(db_conn, changedMapArr1)

		CreateMap(details1, IMAGEPATHUPDATE)

		http.Redirect(w, r, "/map?action=1", http.StatusSeeOther)
	}

	// Query all users from the database
	details := FetchCompanyData(db_conn)
	mapPageData := mapPageStruct{
		ImagePath:   IMAGEPATH,
		CompanyList: details,
		MapAction:   action,
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/map.html")
	ThrowHttpError(w, err)

	// Execute the template
	err = tmpl.Execute(w, struct {
		MapPageData mapPageStruct
	}{mapPageData})
	ThrowHttpError(w, err)
}
