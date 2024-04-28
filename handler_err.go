package main

import (
	"net/http"
)

func handlerError(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, 400, "something went wrong")
}
