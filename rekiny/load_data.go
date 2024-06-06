package main

import (
	"encoding/json"
	"fmt"
	"os"
	"rekiny/models"
)

func Load_data() (models.SharkAttacks, error) {
    file, err := os.Open("global-shark-attack.json")
    if err != nil {
        return nil, fmt.Errorf("error opening file: %w", err)
    }
    defer file.Close()

    var data models.SharkAttacks
    dec := json.NewDecoder(file)
    if err := dec.Decode(&data); err != nil {
        return nil, fmt.Errorf("error decoding JSON: %w", err)
    }

    return data, nil
}