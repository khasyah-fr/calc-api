package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var err error

type Calc struct {
	Num1 string `json:"num1" validate:"required"`
	Num2 string `json:"num2" validate:"required"`
}

func handleRequests() {
	log.Println("Start the development server at http://127.0.0.1:9999")

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", homePage)

	s := r.PathPrefix("/api/v1/").Subrouter()

	s.HandleFunc("/add", addHandler)
	s.HandleFunc("/subtract", subtractHandler)
	s.HandleFunc("/multiply", multiplyHandler)
	s.HandleFunc("/divide", divideHandler)

	http.ListenAndServe(":9999", r)
}

func main() {
	handleRequests()
}

func homePage(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Welcome to the homepage")
}

func addHandler(rw http.ResponseWriter, r *http.Request) {
	payload, _ := ioutil.ReadAll(r.Body)

	var calc Calc
	json.Unmarshal(payload, &calc)

	num1, _ := strconv.ParseFloat(calc.Num1, 64) 
	num2, _ := strconv.ParseFloat(calc.Num2, 64) 

	res := num1 + num2

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(res)
}

func subtractHandler(rw http.ResponseWriter, r *http.Request){
	payload, _ := ioutil.ReadAll(r.Body)

	var calc Calc
	json.Unmarshal(payload, &calc)

	num1, _ := strconv.ParseFloat(calc.Num1, 64) 
	num2, _ := strconv.ParseFloat(calc.Num2, 64) 

	res := num1 - num2

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(res)
}

func multiplyHandler(rw http.ResponseWriter, r *http.Request){
	payload, _ := ioutil.ReadAll(r.Body)

	var calc Calc
	json.Unmarshal(payload, &calc)

	num1, _ := strconv.ParseFloat(calc.Num1, 64) 
	num2, _ := strconv.ParseFloat(calc.Num2, 64) 

	res := num1 * num2

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(res)
}

func divide(x float64, y float64) (float64, error) {
	if y == 0 {
    	return -1, errors.New("Cannot divide by 0!")
  	}
  	return x/y, nil
}

func divideHandler(rw http.ResponseWriter, r *http.Request){
	payload, _ := ioutil.ReadAll(r.Body)

	var calc Calc
	json.Unmarshal(payload, &calc)

	num1, _ := strconv.ParseFloat(calc.Num1, 64) 
	num2, _ := strconv.ParseFloat(calc.Num2, 64) 

	res, err := divide(num1, num2)

	if(err != nil){
		rw.Header().Set("Content-Type", "application/json")
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(rw).Encode("Infinity")
	} else {
		rw.Header().Set("Content-Type", "application/json")
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(rw).Encode(res)
	}
}