package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type player struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Spor      string `json:"sport"`
}

type message struct {
	WelcomeMessage string `json:"welcomemessage"`
}

func main() {

	// JSON Parsing (Ayrıştırma)

	// API Request gönderirken, struct'ımızı önce JSON'a dönüştürmeliyiz. Bunun için json.Marshal() fonksiyonu kullanılır.

	//JSON Marshall sadece dışa aktarılmış verileri export eder. O yüzden struck içindeki alanlar büyük harf ile yazılmalıdır.

	http.HandleFunc("/", handlerHomePage)
	http.HandleFunc("/my/player", handlerMyPlayer)
	http.HandleFunc("/players", handlerPlayers)

	http.ListenAndServe(":8080", nil)

}

func handlerMyPlayer(w http.ResponseWriter, r *http.Request) {
	data := player{1, "Emre", "Öney", 24, "Football"}
	result, err := json.Marshal(data) // Marshall fonksiyonu sonrasında , []byte tipinde parse edilir verimiz.
	checkError(err)
	fmt.Fprintf(w, string(result))
}

func handlerHomePage(w http.ResponseWriter, r *http.Request) {
	data := message{"Welcome to Rest API Exercise"}
	result, err := json.Marshal(data)
	checkError(err)
	fmt.Fprintf(w, string(result))
}
func handlerPlayers(w http.ResponseWriter, r *http.Request) {
	data := []player{
		player{1, "Emre", "Öney", 24, "Football"},
		player{5, "Deniz", "Yıldız", 24, "Cricket"},
		player{2, "Berk", "Harmandar", 24, "Baseball"},
		player{3, "Göktuğ", "Yavuz", 24, "Baseball"},
		player{4, "Tolga", "Başarır", 24, "Golf"},
		player{6, "Can", "Ege", 24, "Cricket"},
	}
	result, err := json.Marshal(data)
	checkError(err)
	w.Write(result)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Hata!", err.Error())
	}
}
