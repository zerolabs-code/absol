package absol

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusTeapot)
	_, _ = fmt.Fprint(w, "teapot")
})

func TestMux_HEAD(t *testing.T) {
	var testMux = NewMux()
	testMux.HEAD("/test", testHandler)
	r, _ := http.NewRequest(http.MethodHead, "/test", nil)
	rec := httptest.NewRecorder()
	testMux.ServeHTTP(rec, r)

	expStatus := http.StatusTeapot
	if status := rec.Code; status != expStatus {
		t.Errorf("wrong status code: got %v expected %v", status, expStatus)
	}

	expBody := "teapot"
	if body := rec.Body.String(); body != expBody {
		t.Errorf("wrong body: got %v expected %v", body, expBody)
	}
}

func TestMux_GET(t *testing.T) {
	var testMux = NewMux()
	testMux.GET("/test", testHandler)
	r, _ := http.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()
	testMux.ServeHTTP(rec, r)

	expStatus := http.StatusTeapot
	if status := rec.Code; status != expStatus {
		t.Errorf("wrong status code: got %v expected %v", status, expStatus)
	}

	expBody := "teapot"
	if body := rec.Body.String(); body != expBody {
		t.Errorf("wrong body: got %v expected %v", body, expBody)
	}
}

func TestMux_POST(t *testing.T) {
	var testMux = NewMux()
	testMux.POST("/test", testHandler)
	r, _ := http.NewRequest(http.MethodPost, "/test", nil)
	rec := httptest.NewRecorder()
	testMux.ServeHTTP(rec, r)

	expStatus := http.StatusTeapot
	if status := rec.Code; status != expStatus {
		t.Errorf("wrong status code: got %v expected %v", status, expStatus)
	}

	expBody := "teapot"
	if body := rec.Body.String(); body != expBody {
		t.Errorf("wrong body: got %v expected %v", body, expBody)
	}
}

func TestMux_PUT(t *testing.T) {
	var testMux = NewMux()
	testMux.PUT("/test", testHandler)
	r, _ := http.NewRequest(http.MethodPut, "/test", nil)
	rec := httptest.NewRecorder()
	testMux.ServeHTTP(rec, r)

	expStatus := http.StatusTeapot
	if status := rec.Code; status != expStatus {
		t.Errorf("wrong status code: got %v expected %v", status, expStatus)
	}

	expBody := "teapot"
	if body := rec.Body.String(); body != expBody {
		t.Errorf("wrong body: got %v expected %v", body, expBody)
	}
}

func TestMux_DELETE(t *testing.T) {
	var testMux = NewMux()
	testMux.DELETE("/test", testHandler)
	r, _ := http.NewRequest(http.MethodDelete, "/test", nil)
	rec := httptest.NewRecorder()
	testMux.ServeHTTP(rec, r)

	expStatus := http.StatusTeapot
	if status := rec.Code; status != expStatus {
		t.Errorf("wrong status code: got %v expected %v", status, expStatus)
	}

	expBody := "teapot"
	if body := rec.Body.String(); body != expBody {
		t.Errorf("wrong body: got %v expected %v", body, expBody)
	}
}

func TestMux_ServeHTTP_NotFound(t *testing.T) {
	var testMux = NewMux()
	r, _ := http.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()
	testMux.ServeHTTP(rec, r)

	expStatus := http.StatusNotFound
	if status := rec.Code; status != expStatus {
		t.Errorf("wrong status code: got %v expected %v", status, expStatus)
	}

	expBody := "request path not registered\n"
	if body := rec.Body.String(); body != expBody {
		t.Errorf("wrong body: got %v expected %v", body, expBody)
	}
}

func TestMux_ServeHTTP_MethodNotAllowed(t *testing.T) {
	var testMux = NewMux()
	testMux.GET("/test", testHandler)
	r, _ := http.NewRequest(http.MethodPost, "/test", nil)
	rec := httptest.NewRecorder()
	testMux.ServeHTTP(rec, r)

	expStatus := http.StatusMethodNotAllowed
	if status := rec.Code; status != expStatus {
		t.Errorf("wrong status code: got %v expected %v", status, expStatus)
	}

	expBody := "handler not registered for the given HTTP verb\n"
	if body := rec.Body.String(); body != expBody {
		t.Errorf("wrong body: got %v expected %v", body, expBody)
	}
}

func TestMux_MultipleMethods(t *testing.T) {
	var testMux = NewMux()
	testMux.GET("/test", testHandler)
	testMux.POST("/test", testHandler)

	for _, m := range []string{http.MethodPost, http.MethodPost} {
		r, _ := http.NewRequest(m, "/test", nil)
		rec := httptest.NewRecorder()
		testMux.ServeHTTP(rec, r)

		expStatus := http.StatusTeapot
		if status := rec.Code; status != expStatus {
			t.Errorf("wrong status code: got %v expected %v", status, expStatus)
		}

		expBody := "teapot"
		if body := rec.Body.String(); body != expBody {
			t.Errorf("wrong body: got %v expected %v", body, expBody)
		}
	}
}
