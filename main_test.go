package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddHandler(t *testing.T) {
	testCases := []struct {
		Name     string
		Payload  []byte
		Expected string
	}{
		{
			Name:     "Valid input",
			Payload:  []byte(`{"num1": 5, "num2": 10}`),
			Expected: `{"sum":15}`,
		},
		{
			Name:     "Negative numbers",
			Payload:  []byte(`{"num1": -5, "num2": -10}`),
			Expected: `{"sum":-15}`,
		},
		{
			Name:     "Zero input",
			Payload:  []byte(`{"num1": 0, "num2": 0}`),
			Expected: `{"sum":0}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(tc.Payload))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			http.HandlerFunc(addHandler).ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			if rr.Body.String() != tc.Expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tc.Expected)
			}
		})
	}
}
