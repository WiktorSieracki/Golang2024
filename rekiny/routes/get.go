package routes

import (
	"encoding/json"
	"net/http"
	"rekiny/models"
)

func Get(data *models.SharkAttacks) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        date := r.URL.Query().Get("date")
		if date == "" {
			jsonData, err := json.Marshal((*data)[0:10])
			if err != nil {
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
			return
		}
        for _, attack := range *data {
            if attack.Date == date {
                jsonData, err := json.Marshal(attack)
                if err != nil {
                    http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
                    return
                }
                w.Header().Set("Content-Type", "application/json")
                w.Write(jsonData)
                return
            }
        }
        http.Error(w, "No data found for the given date", http.StatusNotFound)
    }
}