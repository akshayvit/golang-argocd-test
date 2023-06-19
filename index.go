package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type user struct {
	Name string `json:"Name"`
	Id int `json:"Id"`
}
var tmp user

func main() {
	r:=mux.NewRouter();
	r.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		//var User user
			fmt.Fprint(w, "{Name: "+tmp.Name+",Id:"+strconv.Itoa(tmp.Id)+"}")
		
	}).Methods("GET")
	r.HandleFunc("/post",func(w http.ResponseWriter,r *http.Request){
	//	var User user
		//err:=json.NewEncoder(r.Body).Encode(&User)
		//var tmp user
		err:=json.NewDecoder(r.Body).Decode(&tmp)
		fmt.Println(tmp.Id)
		fmt.Println(tmp.Name)
		if err!=nil  {
			fmt.Fprint(w,"Done");
		}
		
	}).Methods("POST")
	log.Fatal(http.ListenAndServe(":3002",r))
	fmt.Println("Connecting...running at port 3000")
}