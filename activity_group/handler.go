package activity_group

import (
	"devcode_2nd/helper"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler interface {
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service}
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAll()

	if err != nil {
		response := helper.FormatResponse("error", err.Error(), nil)
		jsonData, _ := json.Marshal(response)

		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "%s", jsonData)
		return
	}

	response := helper.FormatResponse("Success", "Success", data)
	jsonData, _ := json.Marshal(response)
	fmt.Fprintf(w, "%s", jsonData)
}

func (h *handler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, err := h.service.GetByID(vars["id"])

	if err != nil {
		errorMsg := fmt.Sprintf("Activity with ID %s Not Found", vars["id"])
		response := helper.FormatResponse("error", errorMsg, map[string]interface{}{})
		jsonData, _ := json.Marshal(response)

		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", jsonData)
		return
	}

	response := helper.FormatResponse("Success", "Success", data)
	jsonData, _ := json.Marshal(response)
	fmt.Fprintf(w, "%s", jsonData)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var input CreateActivityGroupInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		response := helper.FormatResponse("Bad Request", "Failed to parse Input", map[string]interface{}{})
		jsonData, _ := json.Marshal(response)

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", jsonData)
		return
	}

	if input.Title == "" {
		errMsg := "title cannot be null"
		response := helper.FormatResponse("Bad Request", errMsg, map[string]interface{}{})
		jsonData, _ := json.Marshal(response)

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", jsonData)
		return
	}

	if input.Email == "" {
		errMsg := "email cannot be null"
		response := helper.FormatResponse("Bad Request", errMsg, map[string]interface{}{})
		jsonData, _ := json.Marshal(response)

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", jsonData)
		return
	}

	activityGroup := ActivityGroup{
		Title: input.Title,
		Email: input.Email,
	}

	newActivityGroup, err := h.service.Create(activityGroup)

	if err != nil {
		response := helper.FormatResponse("error", err.Error(), nil)
		jsonData, _ := json.Marshal(response)

		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "%s", jsonData)
		return
	}

	response := helper.FormatResponse("Success", "Success", newActivityGroup)
	jsonData, _ := json.Marshal(response)

	fmt.Fprintf(w, "%s", jsonData)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	status, err := h.service.Delete(vars["id"])

	if err != nil {
		response := helper.FormatResponse("Error", "Error", map[string]interface{}{})
		jsonData, _ := json.Marshal(response)

		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, "%s", jsonData)

		return
	}

	if !status {
		errMsg := fmt.Sprintf("Activity with ID %s Not Found", vars["id"])
		response := helper.FormatResponse("Not Found", errMsg, map[string]interface{}{})
		jsonData, _ := json.Marshal(response)

		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", jsonData)
		return
	}

	response := helper.FormatResponse("Success", "Success", map[string]interface{}{})
	jsonData, _ := json.Marshal(response)

	fmt.Fprintf(w, "%s", jsonData)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var input UpdateInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		response := helper.FormatResponse("Bad Request", "Failed to parse Input", map[string]interface{}{})
		jsonData, _ := json.Marshal(response)

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", jsonData)
		return
	}

	if input.Title == nil {
		errMsg := "title cannot be null"
		response := helper.FormatResponse("Bad Request", errMsg, map[string]interface{}{})
		jsonData, _ := json.Marshal(response)

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", jsonData)
		return
	}

	activityGroup, err := h.service.GetByID(vars["id"])

	if err != nil {
		errorMsg := fmt.Sprintf("Activity with ID %s Not Found", vars["id"])
		response := helper.FormatResponse("error", errorMsg, map[string]interface{}{})
		jsonData, _ := json.Marshal(response)

		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", jsonData)
		return
	}

	activityGroup.Title = *input.Title

	newActivityGroup, _ := h.service.Update(activityGroup)

	response := helper.FormatResponse("Success", "Success", newActivityGroup)
	jsonData, _ := json.Marshal(response)

	fmt.Fprintf(w, "%s", jsonData)

}
