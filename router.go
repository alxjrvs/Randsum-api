package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
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
		"/v1/roll/{roll}/d/{sides}",
		Index,
	},
	Route{
		"Pun",
		"GET",
		"/v1/puns",
		PunIndex,
	},
	Route{
		"default",
		"GET",
		"/v1/roll",
		DefaultIndex,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	for _, route := range routes {
		handler := c.Handler(route.HandlerFunc)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func PunIndex(w http.ResponseWriter, r *http.Request) {
	result := CriticalHit()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func DefaultIndex(w http.ResponseWriter, r *http.Request) {
	params := RollParams{1, 20}
	result := RollResult(params)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func Index(w http.ResponseWriter, r *http.Request) {
	params, err, valid := paramsParser(r)
	if valid == false {
		json.NewEncoder(w).Encode(err)
	} else {
		result := RollResult(params)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func paramsParser(r *http.Request) (RollParams, map[string][]string, bool) {
	rawParams, errors, valid := paramsValidator(r)
	roll := rawParams["roll"]
	sides := rawParams["sides"]
	params := RollParams{roll, sides}

	return params, errors, valid
}

func getParams(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func paramsValidator(r *http.Request) (map[string]int64, map[string][]string, bool) {
	rawParams := getParams(r)
	errors := map[string][]string{}
	errors, valid := checkRequired(rawParams, errors)
	if valid == false {
		return nil, errors, valid
	}
	parsedParams, errors, valid := formatParams(rawParams, errors)
	if valid == false {
		return nil, errors, valid
	}
	return parsedParams, errors, true
}

func formatParams(rawParams map[string]string, errors map[string][]string) (map[string]int64, map[string][]string, bool) {
	valid := true
	parsedParams := map[string]int64{}
	for k, v := range rawParams {
		value, err := strconv.Atoi(v)
		if value <= 0 {
			errors, valid = logError(errors, "cannot_be_below_zero", k)
		}
		if err != nil {
			errors, valid = logError(errors, "not_a_number", k)
		}
		parsedParams[k] = int64(value)
	}

	return parsedParams, errors, valid
}

func checkRequired(rawParams map[string]string, errors map[string][]string) (map[string][]string, bool) {
	valid := true
	for _, p := range requiredParams() {
		val := rawParams[p]
		if val == "" {
			errors, valid = logError(errors, "required_param", p)
		}
	}
	return errors, valid
}

func logError(errors map[string][]string, key string, val string) (map[string][]string, bool) {
	errors[key][len(errors)] = val
	return errors, false
}

func requiredParams() []string {
	return []string{"roll", "sides"}
}
