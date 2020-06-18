package middleware

import "net/http"

// TODO: Create a middleware that adds "Access-Control-Allow-Origin: *" header for every method
// Write the middleware in three techniques
// 1) Wrapping around the whole Mux
// 2) Wrapping around an individual handler function
// 3) As an adapter, returning a handler

// TODO: Technique 1: Wrapping around the whole Mux
// Fill in the required field(s) and method body
type CorsMW_1 struct {
	mux http.ServeMux
}

func (c *CorsMW_1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

}

// TODO: Technique 2: Wrapping around an individual handler function
// Fill in the required field(s) and method body
type CorsMW_2 struct {
	handler http.Handler
}

func (c *CorsMW_2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := GetAuthenticatedUser(r)
	if err != nil {
		http.Error(w, "please sign-in", http.StatusUnauthorized)
		return
	}
	????????????????????????????????????????????????????
	c.handler.ServeHTTP(w, r)
}

// TODO: Technique 3: As an adapter, returning a handler
// Fill in the required field(s) and method body
func CorsMW_3(h http.Handler) http.Handler {
	return nil
}
