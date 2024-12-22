package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/timurgulov/calc_go/api"
)

func newRequest(t *testing.T, method, url, body string) *http.Request {
	req := httptest.NewRequest(method, url, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}

// Тест успешного вычисления
func TestCalcHandler_Success(t *testing.T) {
	reqbody := `{"expression":"2+2"}`

	req := newRequest(t, "POST", "/api/v1/calculate", reqbody)
	rr := httptest.NewRecorder()

	api.CalcHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var result api.Response
	err := json.NewDecoder(rr.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}

	expected := 4.0
	if result.Result != expected {
		t.Errorf("expected result %f, got %f", expected, result.Result)
	}

	if result.Error != "" {
		t.Errorf("expected no error, got %s", result.Error)
	}
}

// Тест для случая, когда выражение пустое
func TestCalcHandler_EmptyExpression(t *testing.T) {
	// Пример запроса с пустым выражением
	reqbody := `{"expression":""}`

	req := newRequest(t, "POST", "/api/v1/calculate", reqbody)
	rr := httptest.NewRecorder()

	api.CalcHandler(rr, req)

	if rr.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status code %d, got %d", http.StatusUnprocessableEntity, rr.Code)
	}

	var result api.Response
	err := json.NewDecoder(rr.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}

	expectedError := "Expression cannot be empty. Please provide a valid mathematical expression."
	if result.Error != expectedError {
		t.Errorf("expected error message '%s', got '%s'", expectedError, result.Error)
	}
}

// Тест для случая, когда выражение некорректно
func TestCalcHandler_InvalidExpression(t *testing.T) {
	// Пример запроса с некорректным выражением
	reqbody := `{"expression":"2+("}` 

	req := newRequest(t, "POST", "/api/v1/calculate", reqbody)
	rr := httptest.NewRecorder()

	api.CalcHandler(rr, req)

	if rr.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status code %d, got %d", http.StatusUnprocessableEntity, rr.Code)
	}

	var result api.Response
	err := json.NewDecoder(rr.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}

	expectedError := "Expression is not valid. Please check the syntax."
	if result.Error != expectedError {
		t.Errorf("expected error message '%s', got '%s'", expectedError, result.Error)
	}
}