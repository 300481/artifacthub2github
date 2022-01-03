package artifacthub2github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Payload struct {
	EventType     string `json:"event_type"`
	ClientPayload string `json:"client_payload"`
}

// Handle will handle the cloud function call
func Handle(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("TOKEN")
	owner := os.Getenv("OWNER")
	repo := os.Getenv("REPO")

	hookBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	payload := Payload{
		EventType:     "artifacthub",
		ClientPayload: string(hookBody),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Panic(err)
	}

	log.Println(string(jsonPayload))

	bytesPayload := bytes.NewReader(jsonPayload)

	req, err := http.NewRequest("POST", "https://api.github.com/repos/"+owner+"/"+repo+"/dispatches", bytesPayload)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "token "+token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic(err)
	}

	fmt.Fprint(w, "OK")
	log.Println(response)
}
