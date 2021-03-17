package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/entities"
)

type linkUsecaseInteface interface {
	GetLinkByFrom(from string) (string, error)
	CreateLink(requestData entities.Link) error
}

type HttpServer struct {
	LinkUseCases linkUsecaseInteface
}

func (hs *HttpServer) GetLinkByFrom(from string) (string, error) {
	to, err := hs.LinkUseCases.GetLinkByFrom(from)
	return to, err
}

func (hs *HttpServer) CreateLink(requestData entities.Link) error {
	err := hs.LinkUseCases.CreateLink(requestData)
	return err
}

func (hs *HttpServer) get(rw http.ResponseWriter, r *http.Request) {
	from := r.URL.Path
	to, err := hs.GetLinkByFrom(from)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	rw.Header().Add("Location", to)
	rw.WriteHeader(http.StatusFound)
}

func (hs *HttpServer) post(rw http.ResponseWriter, r *http.Request) {
	var err error
	var requestData entities.Link

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&requestData)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = hs.CreateLink(requestData)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (hs *HttpServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		hs.get(rw, r)
	case http.MethodPost:
		hs.post(rw, r)
	}
}

func (hs *HttpServer) Serve(addr string) {
	http.Handle("/", hs)

	s := &http.Server{
		Addr:           addr,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server is running...")
	log.Fatal(s.ListenAndServe())
}
