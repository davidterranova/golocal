package http

import (
	"net/http"

	"github.com/davidterranova/golocal/pkg/xhttp"
	"github.com/gorilla/mux"
)

const pdfPages = "pdfPages"

// New returns a new contacts API router
func New(app PDFSvc) *mux.Router {
	root := mux.NewRouter()

	mountV1PDF(root, app)
	mountPublic(root)

	return root
}

func mountV1PDF(root *mux.Router, app PDFSvc) {
	pdfHandler := NewPDFHandler(app)

	v1 := root.PathPrefix("/v1/pdf").Subrouter()
	v1.HandleFunc("/pages/{"+pdfPages+"}/delete", pdfHandler.RemovePages).Methods(http.MethodPost)
}

func mountPublic(root *mux.Router) {
	root.HandleFunc("/heartbeat", xhttp.Heartbeat).Methods(http.MethodGet)
	root.PathPrefix("/openapi/").Handler(
		http.StripPrefix(
			"/openapi/",
			http.FileServer(http.Dir("docs/openapi")),
		),
	)
	root.PathPrefix("/").Handler(
		http.StripPrefix(
			"/",
			http.FileServer(http.Dir("web")),
		),
	)
}
