package api

import (
	// "bytes"
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
)

type eventsAPI struct {
	client *http.Client
}

func NewEventsAPI(c *http.Client) domain.EventsAPI {
	newAPI := eventsAPI{
		client: c,
	}

	return &newAPI
}

func (a *eventsAPI) SendEvent(payload model.SendEvent) (int, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	request, err := http.NewRequest("POST", os.Getenv("EVENT_API_URL"), bytes.NewBuffer(payloadJSON))
	// request, err := http.NewRequest("GET", "http://localhost:5000", nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	return http.StatusOK, nil
}
