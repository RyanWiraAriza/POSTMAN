package main

import (
	"log"
	"net/http"
        "encoding/json"
)
type Response struct {
	Nama   string `json:"Nama"`
	Nim    string `json:"Nim"`
	Alamat string `json:"Alamat"`
}

var data []Response = []Response{
	{
		Nama:   "Aple",
		Nim:    "111",
		Alamat: "Tree",
	},
	{
		Nama:   "Banana",
		Nim:    "222",
		Alamat: "Tree",
	},
}
func main() {
	http.HandleFunc("/post", PostHandle)
	http.HandleFunc("/get", GetHandle)
	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	
	if err != nil {
		log.Fatal(err)
	}
}
func PostHandle(w http.ResponseWriter, r *http.Request) {
	var newData Response

	newData = Response{
		Nama:   "Ryan",
		Nim:    "333",
		Alamat: "Malang",
	}

	data = append(data, newData)

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newData)
}
func GetHandle(w http.ResponseWriter, r *http.Request) {
	decodedJson,_:=json.Marshal(data)
	w.Header().Set("Content-Type","Application/json")
	w.Write(decodedJson)
}