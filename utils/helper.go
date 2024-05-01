package utils

import (
	"os"
	"strconv"
	"strings"

	g "github.com/serpapi/google-search-results-golang"
)

// method to call and fetch serpapi response
func Load_data_from_api(i int) []ApiData {

	parameter := map[string]string{
		"engine":        "google_jobs",
		"q":             "Software Developer",
		"google_domain": "google.com",
		"gl":            "us",
		"hl":            "en",
		"location":      "Boston, MA-Manchester, NH, United States",
		"lrad":          "100",
		"start":         strconv.Itoa(i),
		"api_key":       "43b7bffa6ce52e0f68e903f823cd3f93236bd09f7124ac413d544744d588b417",
	}

	search := g.NewGoogleSearch(parameter, "43b7bffa6ce52e0f68e903f823cd3f93236bd09f7124ac413d544744d588b417")
	results, err := search.GetJSON()
	ThrowError(err)
	jobs_results := results["jobs_results"].([]interface{})

	var apidata []ApiData
	for _, jobList := range jobs_results {
		resultOutput := jobList.(map[string]interface{})
		apidata = append(apidata, ApiData{Title: resultOutput["title"].(string), CompanyName: resultOutput["company_name"].(string), Location: resultOutput["location"].(string)})

	}

	return apidata
}

// check if file exist or not
func IsFileExist(filename string) bool {
	res, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !res.IsDir()
}

// covering white space and removing brackets
func GetStateFromLocationApi(location string) string {
	res := strings.Split(location, "(")
	return strings.TrimSpace(res[0])
}
