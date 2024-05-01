package utils

import (
	"fmt"
	"net/http"
)

// This function prints an error message if the given error is not nil.
func ThrowError(err error) {
	if err != nil {
		fmt.Println("Something went wrong !!! Please try again.", err)
		return
	}
}

// This function responds with Internal Server Error if the given error is not nil.
func ThrowHttpError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
