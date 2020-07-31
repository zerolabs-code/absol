package absol

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type header struct {
	k string
	v string
}

var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusTeapot)
	_, _ = fmt.Fprint(w, "teapot")
})

var fooMiddleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Foo", "foo-test")
		next.ServeHTTP(w, r)
	})
}

var barMiddleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Bar", "bar-test")
		next.ServeHTTP(w, r)
	})
}

func TestMux_Head(t *testing.T) {
	var testMux = NewMux()
	testMux.Head("/test", testHandler)
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

func TestMux_Get(t *testing.T) {
	var testMux = NewMux()
	testMux.Get("/test", testHandler)
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

func TestMux_Post(t *testing.T) {
	var testMux = NewMux()
	testMux.Post("/test", testHandler)
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

func TestMux_Put(t *testing.T) {
	var testMux = NewMux()
	testMux.Put("/test", testHandler)
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

func TestMux_Delete(t *testing.T) {
	var testMux = NewMux()
	testMux.Delete("/test", testHandler)
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

	expBody := "absol: request path not registered\n"
	if body := rec.Body.String(); body != expBody {
		t.Errorf("wrong body: got %v expected %v", body, expBody)
	}
}

func TestMux_ServeHTTP_MethodNotAllowed(t *testing.T) {
	var testMux = NewMux()
	testMux.Get("/test", testHandler)
	r, _ := http.NewRequest(http.MethodPost, "/test", nil)
	rec := httptest.NewRecorder()
	testMux.ServeHTTP(rec, r)

	expStatus := http.StatusMethodNotAllowed
	if status := rec.Code; status != expStatus {
		t.Errorf("wrong status code: got %v expected %v", status, expStatus)
	}

	expBody := "absol: request method not registered\n"
	if body := rec.Body.String(); body != expBody {
		t.Errorf("wrong body: got %v expected %v", body, expBody)
	}
}

func TestMux_MultipleMethods(t *testing.T) {
	var testMux = NewMux()
	testMux.Get("/test", testHandler)
	testMux.Post("/test", testHandler)

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

func TestMux_Use(t *testing.T) {
	var testMux = NewMux()
	testMux.Get("/test", testHandler)
	testMux.Use(fooMiddleware)
	r, _ := http.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()
	testMux.ServeHTTP(rec, r)

	expStatus := http.StatusTeapot
	if status := rec.Code; status != expStatus {
		t.Errorf("wrong status code: got %v expected %v", status, expStatus)
	}

	expHeader := "foo-test"
	if header := rec.Header().Get("X-Foo"); header != expHeader {
		t.Errorf("wrong header value: got %v expected %v", header, expHeader)
	}
}

func TestMux_Use_Multiple(t *testing.T) {
	var testMux = NewMux()
	testMux.Get("/test", testHandler)
	testMux.Use(fooMiddleware)
	testMux.Use(barMiddleware)
	r, _ := http.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()
	testMux.ServeHTTP(rec, r)

	expStatus := http.StatusTeapot
	if status := rec.Code; status != expStatus {
		t.Errorf("wrong status code: got %v expected %v", status, expStatus)
	}

	for _, h := range []header{{"X-Foo", "foo-test"}, {"X-Bar", "bar-test"}} {
		expHeader := h.v
		if header := rec.Header().Get(h.k); header != expHeader {
			t.Errorf("wrong header value: got %v expected %v", header, expHeader)
		}
	}
}
