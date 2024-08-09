package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const TARGET_URL = "https://character-service.dndbeyond.com/character/v5/character/"

var (
	PORT                  = os.Getenv("PORT")
	ALLOWED_CHARACTER_IDS = strings.Split(os.Getenv("ALLOWED_CHARACTER_IDS"), ",")
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	port, err := strconv.Atoi(PORT)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	allowedIDS := map[string]bool{}
	for _, id := range ALLOWED_CHARACTER_IDS {
		if len(strings.TrimSpace(id)) > 0 {
			allowedIDS[strings.TrimSpace(id)] = true
		}
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		characterID := r.URL.Query().Get("character")
		if len(characterID) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Printf("%+v", allowedIDS)

		if len(allowedIDS) > 0 {
			if _, ok := allowedIDS[characterID]; !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		finalURL := TARGET_URL + characterID
		resp, err := http.DefaultClient.Get(finalURL)
		if err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(body)
	})

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
