package utils

import (
	"database/sql"
	"os"
	"time"
)

const FILEPATH = "./db/arshjot_final_project.db"

var db_conn *sql.DB

type ApiData struct {
	ID          int
	Title       string
	CompanyName string
	Location    string
	Createddate time.Time
}

func Dbinit() {

	// make direct database connection, if db file already exists
	if IsFileExist(FILEPATH) {
		// initiate database connection only
		db_conn = Make_db_connection()

	} else {
		// if db file doesn't exist then, create file, db tables and populate data to tables from scratch
		Create_db_file()

		// initiate database connection only
		db_conn = Make_db_connection()

		// initiating table creation
		CreateCompanyTable(db_conn)

		i := 0
		for i < 5 {
			load_api_data := Load_data_from_api(i * 10)
			for _, v := range load_api_data {
				PopulateCompanyTable(db_conn, v)
			}
			i++
		}

	}
}

// method to create db file
func Create_db_file() {
	// creating new database file
	createDb, err := os.Create(FILEPATH)
	ThrowError(err)

	// closing the file open connection
	createDb.Close()
}

// method to create and open sqlite3 database connection
func Make_db_connection() (db *sql.DB) {
	// opening sqlite3 database connection
	open_db, err := sql.Open("sqlite3", FILEPATH)
	ThrowError(err)
	return open_db
}

// method to create new company table
func CreateCompanyTable(db *sql.DB) {
	sqlStatement := `CREATE TABLE COMPANY (
		"ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"TITLE" VARCHAR250,
		"COMPANYNAME" VARCHAR250,
		"LOCATION" VARCHAR250,
		"CREATEDDATE" DATETIME
	);`

	prepareStatement, err := db.Prepare(sqlStatement)
	ThrowError(err)
	prepareStatement.Exec()
}

// method to insert data to company table
func PopulateCompanyTable(db *sql.DB, details ApiData) {
	sqlStatement := `INSERT INTO COMPANY (
		TITLE, COMPANYNAME, LOCATION, CREATEDDATE
	) VALUES (
		?, ?, ?, ?
	);`
	prepareStatement, err := db.Prepare(sqlStatement)
	ThrowError(err)
	_, err = prepareStatement.Exec(details.Title, details.CompanyName, GetStateFromLocationApi(details.Location), time.Now())
	ThrowError(err)
}

// fetching data from comoany table
func FetchCompanyData(db *sql.DB) (details []ApiData) {
	execStatement, err := db.Query("SELECT * FROM COMPANY ORDER BY CREATEDDATE DESC")
	ThrowError(err)
	var detailsList []ApiData
	for execStatement.Next() {
		var details ApiData
		execStatement.Scan(&details.ID, &details.Title, &details.CompanyName, &details.Location, &details.Createddate)

		detailsList = append(detailsList, ApiData{ID: details.ID, Title: details.Title, CompanyName: details.CompanyName, Location: details.Location, Createddate: details.Createddate})
	}
	return detailsList
}

func FetchCompanyDataFromSelectedId(db *sql.DB, id string) (locationList []string) {
	execStatement, err := db.Query("SELECT LOCATION FROM COMPANY WHERE ID IN (" + id[:len(id)-1] + ")")
	ThrowError(err)
	var location []string
	for execStatement.Next() {
		var details ApiData
		execStatement.Scan(&details.Location)

		location = append(location, details.Location)
	}

	return location
}
