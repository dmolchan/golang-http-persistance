package tweets

import (
	"encoding/json"
	"net/http"

	"github.com/aren55555/learn/dec7/webapp/models"
)

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	result, err := models.CreateTweet(h.PG, "Sample body")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	// Use 201 code for creation
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
