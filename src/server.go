package src

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Server struct {
	storage *Storage
}

type PriceBody struct {
	Company int    `json:"Company"`
	Price   int    `json:"Price"`
	Origin  string `json:"Origin"`
	Date    string `json:"Date"`
}

func NewServer() *Server {
	return &Server{
		storage: NewStorage(),
	}
}

func (s *Server) HandleGet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.storage.GetAveragePrices())
}

func (s *Server) HandlePost(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close() // Close the body

	// Unmarshal the JSON into a struct
	priceBody := PriceBody{}
	err = json.Unmarshal(bodyBytes, &priceBody)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	s.storage.InsertPrice(priceBody)
}

func (s *Server) Start(port string) {
	if port == "" {
		port = "3142"
	}
	log.Println("Starting server on port", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			s.HandleGet(w, r)
		} else if r.Method == "POST" {
			s.HandlePost(w, r)
		} else {
			log.Fatal("Method not allowed")
		}
	})

	http.ListenAndServe(":"+port, nil)
}
