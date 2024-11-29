package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func main() {
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		// Проверка метода
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Извлечение параметров запроса
		group := r.URL.Query().Get("group")
		song := r.URL.Query().Get("song")

		if group == "" || song == "" {
			http.Error(w, "Missing required query parameters", http.StatusBadRequest)
			return
		}

		// Пример ответа
		response := SongDetail{
			ReleaseDate: "16.07.2006",
			Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
			Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
