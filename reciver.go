package procnias

import (
	"encoding/json"
	"net/http"

	alertmanager "github.com/prometheus/alertmanager/template"
)

func (a App) reciver(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := alertmanager.Data{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		a.logger.Printf("Error: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := a.logic(data); err != nil {
		a.logger.Printf("Error: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
