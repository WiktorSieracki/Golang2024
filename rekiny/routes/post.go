package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rekiny/models"
)

func Post(data *models.SharkAttacks) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var newData models.SharkAttack
        if err := json.NewDecoder(r.Body).Decode(&newData); err != nil {
            http.Error(w, "Bad request", http.StatusBadRequest)
            return
        }
        *data = append(*data, newData)
		fmt.Println((*data)[len(*data)-1])
        w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Data added successfully"))
    }
}