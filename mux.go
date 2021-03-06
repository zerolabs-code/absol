package absol

import (
	"net/http"
)

type Mux struct {
	handlers   map[string]map[string]http.Handler
	middleware Middleware
}

func NewMux() *Mux {
	mux := new(Mux)
	mux.handlers = make(map[string]map[string]http.Handler)
	mux.middleware = nil
	return mux
}

func (mux *Mux) getOrCreatePathHandlers(path string) map[string]http.Handler {
	if handlers, ok := mux.handlers[path]; !ok {
		emptyHandlers := make(map[string]http.Handler)
		mux.handlers[path] = emptyHandlers
		return emptyHandlers
	} else {
		return handlers
	}
}

func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handlers, ok := mux.handlers[r.URL.Path]; !ok {
		http.Error(w, "absol: request path not registered", http.StatusNotFound)
	} else {
		if handler, found := handlers[r.Method]; !found {
			http.Error(w, "absol: request method not registered", http.StatusMethodNotAllowed)
		} else {
			if mux.middleware != nil {
				mux.middleware(handler).ServeHTTP(w, r)
			} else {
				handler.ServeHTTP(w, r)
			}
		}
	}
}

func (mux *Mux) Head(path string, handler http.Handler) {
	pathHandlers := mux.getOrCreatePathHandlers(path)
	pathHandlers[http.MethodHead] = handler
}

func (mux *Mux) Get(path string, handler http.Handler) {
	pathHandlers := mux.getOrCreatePathHandlers(path)
	pathHandlers[http.MethodGet] = handler
}

func (mux *Mux) Post(path string, handler http.Handler) {
	pathHandlers := mux.getOrCreatePathHandlers(path)
	pathHandlers[http.MethodPost] = handler
}

func (mux *Mux) Put(path string, handler http.Handler) {
	pathHandlers := mux.getOrCreatePathHandlers(path)
	pathHandlers[http.MethodPut] = handler
}

func (mux *Mux) Delete(path string, handler http.Handler) {
	pathHandlers := mux.getOrCreatePathHandlers(path)
	pathHandlers[http.MethodDelete] = handler
}

func (mux *Mux) Use(m Middleware) {
	if mux.middleware == nil {
		mux.middleware = m
	} else {
		mux.middleware = compose(m, mux.middleware)
	}
}
