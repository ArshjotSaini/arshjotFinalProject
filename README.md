ARSHJOT SAINI

FILE STRUCTURE
- Db:-        *arshjot_final_project.db
- Assets:-    *updatenameofimage.png
              *nameofimage.png
- Templates:- *index.html
              *map.html
- Utils:-     *errorhandling.go
              *handler.go
              *helper.go
              *maps.go
              *sqlstatment.go
			  *arshjot_test.go
- HOW TO RUN THE APPLICATION
              * To run the project you need to install these dependencies in the terminal
                go get .
              * After running these dependencies type go run main.go in terminal to run the project
       (or you an just type these dependencies in the import and hover over them to import the dependencies)

- Key Points
    - Read data API Call:- On load, application reads data from SERAPI and insert it into sqllite3 database to avoid redundant api calls
    - INSERT into database:- Insert data read from api into sqllite3 database to avoid redundant api calls
    - Map generation:- Map gets generated on application load based upon the locations stored in database after fecthing from SERAPI
    - On demand Map generation:- Map can be generated on ad-hoc basis based upon the supplied location
    - Automate test cases:- 3 test cases are added to the arshjot_test.go file and these test cases gets executed on gitlab during pipeline execution
