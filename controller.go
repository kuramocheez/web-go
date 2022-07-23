package controller


import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Note struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var (
	notes []Note
)

func message(m string) string {
	if m == "POST" {
		return "Data berhasil ditambahkan!"
	}
	if m == "PUT" {
		return "Data berhasil diubah!"
	}
	if m == "DELETE" {
		return "Data berhasil dihapus!"
	}
	return "Data kosong!"
}

func checkErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func checkKey(ok bool, keys []string) {
// 	if !ok || len(keys[0]) < 1 {
// 		log.Println("Url Param 'key' is missing")
// 		return
// 	}
//}

func MethodErr(w http.ResponseWriter) {
	http.Error(w, "400 Status bad request, Method salah!", http.StatusBadRequest)
	// return
}
func GetData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(notes)
		checkErr(w, err)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return
	}
	MethodErr(w)
}
func PostData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		note := Note{}
		err := json.NewDecoder(r.Body).Decode(&note)
		checkErr(w, err)

		notes = append(notes, note)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message(r.Method)))
		return
	}
	MethodErr(w)
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		keys := mux.Vars(r)
		// checkKey(ok, keys)

		id, _ := strconv.Atoi(keys["id"])
		data := Note{}
		err := json.NewDecoder(r.Body).Decode(&data)
		checkErr(w, err)
		for i, v := range notes {
			if v.ID == id {
				notes[i] = data
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message(r.Method)))
		return
	}
	MethodErr(w)
}
func DeleteData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		params := mux.Vars(r)		// checkKey(ok, keys)
		id, err := strconv.Atoi(params["id"])
		checkErr(w, err)
		for i, v := range notes {
			if v.ID == id {
				notes = append(notes[:i], notes[i+1:]...)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message(r.Method)))
		return
	}
	MethodErr(w)
}