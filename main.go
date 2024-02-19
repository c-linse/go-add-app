package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Numbers struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type Result struct {
	Sum int `json:"sum"`
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	var nums Numbers
	err := json.NewDecoder(r.Body).Decode(&nums)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sum := nums.Num1 + nums.Num2

	result := Result{Sum: sum}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsonResult); err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/add", addHandler)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
