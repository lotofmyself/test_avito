package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var (
	redisClient *redis.Client
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}

func SetKeyHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := DecodeJSONBody(r, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for key, val := range data {
		err := redisClient.Set(context.Background(), key, val, 0).Err()
		if err != nil {
			http.Error(w, "403", http.StatusForbidden)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func GetKeyHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	val, err := redisClient.Get(context.Background(), key).Result()
	if err != nil {
		http.Error(w, "404", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Value: %s\n", val)
}

func DeleteKeyHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := DecodeJSONBody(r, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key := data["key"]
	err := redisClient.Del(context.Background(), key).Err()
	if err != nil {
		http.Error(w, "403", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/set_key", SetKeyHandler).Methods("POST")
	r.HandleFunc("/get_key", GetKeyHandler).Methods("GET")
	r.HandleFunc("/del_key", DeleteKeyHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func DecodeJSONBody(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
