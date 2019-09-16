package apphandlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Oleg-Skalozub/testtask/src/domain/services"
	"github.com/Oleg-Skalozub/testtask/src/infrastructure/errorscan"
)

// Handler ...
type Handler interface {
	Request(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service services.Fetcher
}

// NewHandler ...
func NewHandler() Handler {
	return &handler{
		service: services.NewFetch(),
	}
}

// Request ...
func (h handler) Request(w http.ResponseWriter, r *http.Request) {
	day := r.URL.Query().Get("day")
	month := r.URL.Query().Get("month")

	monthTime, dayTime, err := validation(day, month)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	data, err := h.service.FetchData(monthTime, dayTime)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(res)
}

// todo add normal validator
func validation(day, month string) (int, int, error) {
	if day == "" {
		return 0, 0, errorscan.EmptyDayValueError
	}

	if month == "" {
		return 0, 0, errorscan.EmptyMonthValueError
	}

	monthTime, err := strconv.Atoi(month)
	if err != nil {
		return 0, 0, errorscan.WrongMonthTypeError
	}

	dayTime, err := strconv.Atoi(day)
	if err != nil {
		return 0, 0, errorscan.WrongDayTypeError
	}

	if monthTime > 12 {
		return 0, 0, errorscan.BigMonthValueError
	}

	if dayTime > 31 {
		return 0, 0, errorscan.BigDayValueError
	}

	time.Date(1900, time.Month(monthTime), dayTime, 0, 0, 0, 0, time.Local)
	return monthTime, dayTime, nil
}
