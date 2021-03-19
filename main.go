package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/cbrgm/clickbaiter/clickbaiter"
	"github.com/go-chi/chi"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Config struct {
	HttpAddr string `arg:"env:HTTP_ADDR"`
}

type Response struct {
	Headline string `json:"headline"`
}

func main() {
	rand.Seed(time.Now().Unix())

	c := Config{
		HttpAddr: ":8080",
	}

	arg.MustParse(&c)

	cbg := clickbaiter.NewGenerator()

	r := chi.NewRouter()
	r.Get("/", HandleFunc(file("./web/index.html")))
	r.Handle("/web/*", http.StripPrefix("/web", http.FileServer(http.Dir("./web"))))
	r.Post("/generate", HandleFunc(generate(cbg)))

	if err := http.ListenAndServe(c.HttpAddr, r); err != nil {
		log.Fatal(err)
	}
}

type HandlerFunc func(http.ResponseWriter, *http.Request) (int, error)

func HandleFunc(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode, err := h(w, r)
		if err != nil {
			http.Error(w, err.Error(), statusCode)
			fmt.Println(err)
			return
		}
		if statusCode != http.StatusOK {
			w.WriteHeader(statusCode)
		}
	}
}

func file(filename string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		http.ServeFile(w, r, filename)
		return http.StatusOK, nil
	}
}

func generate(cbg clickbaiter.ClickbaitGenerator) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		var resp = Response{
			Headline: cbg.RandomSentence(),
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))

		return http.StatusOK, nil
	}
}
