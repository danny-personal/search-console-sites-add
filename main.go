package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/webmasters/v3"
)

type Message struct {
	Store string `json:"store"`
}

func SitesAdd(w http.ResponseWriter, r *http.Request) {
	keyJson := os.Getenv("KEY_JSON")
	if keyJson == "" {
		return
	} else {
		fmt.Println("KEY_JSON:OK")
	}

	var message Message
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &message)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	siteUrl := fmt.Sprintf("https://www.trial-net.co.jp/shops/%s/", message.Store)
	fmt.Println(siteUrl)
	ctx := context.Background()
	webmastersService, err := webmasters.NewService(ctx, option.WithCredentialsJSON([]byte(keyJson)))
	if err != nil {
		fmt.Println("Failed to create Client.")
	}
	wmxSite, err := webmastersService.Sites.Get(siteUrl).Do()
	if err != nil {
		fmt.Println("Add site.")
	} else {
		fmt.Println("Site is already registered.")
		fmt.Printf("%v\n", wmxSite)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	}

	err = webmastersService.Sites.Add(siteUrl).Do()
	if err != nil {
		fmt.Println("Failed to add site.")
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/", SitesAdd)
	http.ListenAndServe(":8080", nil)
}
