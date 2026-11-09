package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/demo", demoHandler)

	log.Println("Server is strarting ...")
	err := http.ListenAndServe(":8080", nil) //localhost:8080
	if err != nil {
		log.Fatal("Server start error: ", err)
	}

}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r)

	if r.Method != http.MethodGet {
		http.Error(w, "Phuong thuc nay khong duoc ho tro", http.StatusMethodNotAllowed)
		return
	}

	respone := map[string]string{
		"message": "Chao mung cac ban den voi khoa hoc lap trinh Golang",
		"info":    "Giang vien Vo Minh Bao",
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Course", "Lap trinh Golang")

	// data, err := json.Marshal(respone)
	// if err != nil {
	// 	http.Error(w, "Loi ma hoa JSON", http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(data)

	json.NewEncoder(w).Encode(respone)
}
