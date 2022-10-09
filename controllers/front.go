package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	ucs := newUserController()

	http.Handle("/users", *ucs)
	http.Handle("/users/", *ucs)

}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
