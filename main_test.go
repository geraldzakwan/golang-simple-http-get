package main

import (
  "testing"
  "net/http"
  "net/http/httptest"
)

// Case 1: Request without parameter id
func TestCase1(t *testing.T){
  data, dataMap := loadData()

  req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(DataHandler(data, dataMap))
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"code":200,"data":[{"id":1,"name":"A"},{"id":2,"name":"B"},{"id":3,"name":"C"}]}`
	if rec.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

// Case 2: Request with single id
func TestCase2(t *testing.T){
  data, dataMap := loadData()

  req, err := http.NewRequest("GET", "/?id=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(DataHandler(data, dataMap))
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"code":200,"data":[{"id":2,"name":"B"}]}`
	if rec.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

// Case 3: Request with multiple ids
func TestCase3(t *testing.T){
  data, dataMap := loadData()

  req, err := http.NewRequest("GET", "/?id=1,3,4", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(DataHandler(data, dataMap))
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"code":200,"data":[{"id":1,"name":"A"},{"id":3,"name":"C"}]}`
	if rec.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

// Case 4: Request with invalid ID
func TestCase4(t *testing.T){
  data, dataMap := loadData()

  req, err := http.NewRequest("GET", "/?id=xxx", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(DataHandler(data, dataMap))
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := `{"code":400,"message":"Invalid or empty ID: \"xxx\""}`
	if rec.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

// Case 5: Request with ID not found
func TestCase5(t *testing.T){
  data, dataMap := loadData()

  req, err := http.NewRequest("GET", "/?id=4", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(DataHandler(data, dataMap))
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	expected := `{"code":404,"message":"Resource with ID: 4 doesn't exist"}`
	if rec.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
