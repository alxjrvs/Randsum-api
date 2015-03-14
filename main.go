package main

import (
	"github.com/codegangsta/negroni"
	"github.com/rs/cors"
	"os"
)

func newNegroniApi() *negroni.Negroni {
	return negroni.New(negroni.NewRecovery(), negroni.NewLogger())
}
func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	mux := NewRouter()
	n := newNegroniApi()
	n.Use(c)
	n.UseHandler(mux)
	n.Run(":" + os.Getenv("PORT"))
}
