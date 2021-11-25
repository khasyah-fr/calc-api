package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

//Entah kenapa testing nya gagal padahal value expected nya sama, mungkin karena response isinya string buffer
// sementara expected isinya string biasa (???)

func TestAddHandler(t *testing.T) {
	var jsonStr = []byte(`{"num1":"10", "num2": "5"}`)
    req, err := http.NewRequest("POST", "/api/v1/add", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(addHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

	expected := "15"
	if rr.Body.String() != expected {
		t.Fatalf("handler returned: got %v want %v", rr.Body, expected)
	}
}

func TestSubtractHandler(t *testing.T) {
	var jsonStr = []byte(`{"num1":"-6", "num2": "-4"}`)
    req, err := http.NewRequest("POST", "/api/v1/subtract", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(subtractHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

	expected := "-2"
	if rr.Body.String() != expected {
		t.Fatalf("handler returned: got %v want %v", rr.Body, expected)
	}
}

func TestMultiplyHandler(t *testing.T) {
	var jsonStr = []byte(`{"num1":"-3", "num2": "-2"}`)
    req, err := http.NewRequest("POST", "/api/v1/multiply", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(multiplyHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

	expected := "6"
	if rr.Body.String() != expected {
		t.Fatalf("handler returned: got %v want %v", rr.Body, expected)
	}
}

func TestDivideHandler(t *testing.T) {
	var jsonStr = []byte(`{"num1":"-6", "num2": "0"}`)
    req, err := http.NewRequest("POST", "/api/v1/divide", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(divideHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

	expected := "Infinity"
	if rr.Body.String() != expected{
		t.Fatalf("handler returned: got %v want %#v", rr.Body, expected)
	}
}