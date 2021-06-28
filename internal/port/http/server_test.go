package http

import (
	"bytes"
	"encoding/json"
	"github.com/dihmuzikien/smallurl/domain"
	"github.com/dihmuzikien/smallurl/domain/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandleListUrl(t *testing.T){
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	retList := []domain.Url {
		{
			ID: "test1",
			Destination: "https://google.com",
			Created: time.Now().Add(-100 * time.Second),
		},
		{
			ID: "test2",
			Destination: "https://yahoo.com",
			Created: time.Now().Add(-200 * time.Second),
		},
		{
			ID: "test3",
			Destination: "https://github.com",
			Created: time.Now().Add(-300 * time.Second),
		},
	}

	m := mocks.NewMockUrlUseCase(ctrl)
	m.EXPECT().List(gomock.Any()).Return(retList, nil)
	type expectedResponse struct {
		ID string `json:"id"`
		Destination string `json:"destination"`
	}
	sut, _ := New(m)
	req := httptest.NewRequest("GET", "/v1", nil)
	w := httptest.NewRecorder()
	handler := sut.handleListUrl()
	handler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK{
		t.Errorf("want %v got %v", http.StatusOK, resp.StatusCode)
	}
	var body []expectedResponse
	err := json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		t.Errorf("unexpected error decoding response %v", err)
	}
	if len(body) != len(retList){
		t.Errorf("response length not match expected result")
	}
}

func TestHandleCreateUrl(t *testing.T){
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	retList := []domain.Url {
		{
			ID: "test1",
			Destination: "https://google.com",
			Created: time.Now().Add(-100 * time.Second),
		},
		{
			ID: "test2",
			Destination: "https://yahoo.com",
			Created: time.Now().Add(-200 * time.Second),
		},
		{
			ID: "test3",
			Destination: "https://github.com",
			Created: time.Now().Add(-300 * time.Second),
		},
	}

	m := mocks.NewMockUrlUseCase(ctrl)
	m.EXPECT().List(gomock.Any()).Return(retList, nil)
	type expectedResponse struct {
		ID string `json:"id"`
		Destination string `json:"destination"`
	}
	sut, _ := New(m)
	req := httptest.NewRequest("GET", "/v1", nil)
	w := httptest.NewRecorder()
	handler := sut.handleListUrl()
	handler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK{
		t.Errorf("want %v got %v", http.StatusOK, resp.StatusCode)
	}
	var body []expectedResponse
	err := json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		t.Errorf("unexpected error decoding response %v", err)
	}
	if len(body) != len(retList){
		t.Errorf("response length not match expected result")
	}
}


func TestHandleCreateUrlWithAlias(t *testing.T){
	t.Run("create URL with alias succeed", func(t *testing.T){
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)

		testSubj := domain.Url {
			ID: "test1",
			Destination: "https://google.com",
			Created: time.Now().Add(-100 * time.Second),
		}
		m := mocks.NewMockUrlUseCase(ctrl)
		m.EXPECT().CreateWithId(gomock.Any(), gomock.Eq(testSubj.ID), gomock.Eq(testSubj.Destination)).Return(testSubj, nil)
		type expectedResponse struct {
			ID string `json:"id"`
		}
		sut, _ := New(m)

		requestBody := struct {
			Alias string `json:"alias"`
			Destination string `json:"destination"`
		}{
			Alias: testSubj.ID,
			Destination: testSubj.Destination,
		}
		b := new(bytes.Buffer)
		err := json.NewEncoder(b).Encode(requestBody)
		if err != nil {
			t.Fatalf("failed to marshal test request %v", requestBody)
		}

		req := httptest.NewRequest("POST", "/v1", b)
		w := httptest.NewRecorder()
		handler := sut.handleCreateUrlWithAlias()
		handler(w, req)
		resp := w.Result()
		if resp.StatusCode != http.StatusOK{
			t.Errorf("want %v got %v", http.StatusOK, resp.StatusCode)
		}

		var body expectedResponse
		err = json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			t.Errorf("unexpected error decoding response %v", err)
		}
		if body.ID != testSubj.ID{
			t.Errorf("want %v got %v", testSubj.ID, body.ID)
		}
	})

	t.Run("CreateUrlWithAlias not specify alias failed", func(t *testing.T){
		testcases := []struct {
			Alias string `json:"alias"`
			Destination string `json:"destination"`
		}{
			{
				Alias: "",
				Destination: "dest",
			},
			{
				Alias:       "aaa",
				Destination: "google.com",
			},

		}
		for _, tc := range testcases {
			ctrl := gomock.NewController(t)
			t.Cleanup(ctrl.Finish)

			m := mocks.NewMockUrlUseCase(ctrl)
			m.EXPECT().CreateWithId(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			sut, _ := New(m)
			b := new(bytes.Buffer)
			err := json.NewEncoder(b).Encode(tc)
			if err != nil {
				t.Fatalf("failed to marshal test request %v", tc)
			}
			req := httptest.NewRequest("POST", "/v1", b)
			w := httptest.NewRecorder()
			handler := sut.handleCreateUrlWithAlias()
			handler(w, req)
			resp := w.Result()
			if resp.StatusCode != http.StatusBadRequest{
				t.Errorf("want %v got %v", http.StatusBadRequest, resp.StatusCode)
			}

		}
	})
}
