package routes

import (
	"encoding/json"
	"net/http"
	"rekiny/models"
)

func Update(data *models.SharkAttacks) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        date := r.URL.Query().Get("date")
        var updatedData models.SharkAttack
        if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
            http.Error(w, "Bad request", http.StatusBadRequest)
            return
        }
        for i, attack := range *data {
            if attack.Date == date {
                (*data)[i] = updatedData
                w.WriteHeader(http.StatusOK)
                return
            }
        }
        http.Error(w, "No data found for the given date", http.StatusNotFound)
    }
}