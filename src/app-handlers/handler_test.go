package apphandlers

import (
	"github.com/Oleg-Skalozub/testtask/src/mock"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Request(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fetch := mocks.NewMockFetcher(ctrl)

	h := handler{
		service: fetch,
	}

	handler := http.HandlerFunc(h.Request)
	req, err := http.NewRequest(http.MethodGet, "request?day=4&month=4", nil)
	if err != nil {
		t.Fatal(err)
	}

	fetch.EXPECT().FetchData(gomock.Any(), gomock.Any()).Return(mocks.ArrayDataResponse, nil)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected response code %d, but got %d", http.StatusOK, rr.Code)
	}
}
