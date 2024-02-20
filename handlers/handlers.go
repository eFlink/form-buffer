package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	pkg "github.com/eFlink/form-buffer/pkg"
)

var (
	forms    sync.Map
	mapMutex sync.Mutex
)

type FormData struct {
	Data  map[string]interface{}
	Timer *time.Timer
}

func PartialFormHandler(w http.ResponseWriter, r *http.Request) {
	var form map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, emailOk := form["email"].(string)
	_, urlOk := form["url"].(string)
	if !emailOk || !urlOk {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	handleFormSubmission(form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Form Received")
}

func handleFormSubmission(form map[string]interface{}) error {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	delay, _ := os.LookupEnv("DELAY_IN_SECONDS")
	delayInSeconds, err := strconv.Atoi(delay)
	if err != nil {
		return errors.New("delay in seconds environment variable couldn't be converted to int")
	}
	delayDuration := time.Duration(delayInSeconds) * time.Second

	id := pkg.GenerateUniqueID(form["email"].(string), form["url"].(string))

	if existingFormData, ok := forms.Load(id); ok {
		existingFormData.(*FormData).Timer.Stop()
		existingFormData.(*FormData).Data = form
	} else {
		formData := &FormData{Data: form}
		forms.Store(id, formData)
	}

	timer := time.AfterFunc(delayDuration, func() {
		sendFormToZap(form)
		forms.Delete(id)
	})

	if existingFormData, ok := forms.Load(id); ok {
		existingFormData.(*FormData).Timer = timer
	} else {
		return errors.New("form doesn't exist, this should happen at all")
	}

	return nil
}

func sendFormToZap(form map[string]interface{}) {
	formJSON, _ := json.Marshal(form)
	url, _ := os.LookupEnv("URL")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(formJSON))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}
