package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strconv"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v1/roll",
		Index,
	},
	Route{
		"Pun",
		"GET",
		"/api/v1/puns",
		PunIndex,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

func PunIndex(w http.ResponseWriter, r *http.Request) {
	result := CriticalHit()
	json.NewEncoder(w).Encode(result)
	fmt.Println(result)
}

func Index(w http.ResponseWriter, r *http.Request) {
	rawParams := paramsParser(r)
	params := RollParams{
		rawParams["number"],
		rawParams["sides"],
	}
	result := RollResult(params)
	json.NewEncoder(w).Encode(result)
	fmt.Println(result)
}

func paramsParser(r *http.Request) map[string]int64 {
	vals, _ := url.ParseQuery(r.URL.RawQuery)
	iv := map[string]int64{}
	for k, _ := range vals {
		v, _ := strconv.Atoi(vals[k][0])
		iv[k] = int64(v)
	}
	return iv
}
