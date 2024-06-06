package routes

import (
	"net/http"
	"rekiny/models"
)

func Delete(data *models.SharkAttacks) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        date := r.URL.Query().Get("date")
        for i, attack := range *data {
            if attack.Date == date {
                *data = append((*data)[:i], (*data)[i+1:]...)
                w.WriteHeader(http.StatusOK)
                return
            }
        }
        http.Error(w, "No data found for the given date", http.StatusNotFound)
    }
}