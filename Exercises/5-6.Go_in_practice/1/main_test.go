package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type mockFailRW struct{}

func (mockFailRW) Read(p []byte) (n int, err error) {
	return 0, errors.New("mock error")
}

func (mockFailRW) Write(p []byte) (n int, err error) {
	return 0, errors.New("mock error")
}

func Test_logHTTPHandler(t *testing.T) {
	tests := []struct {
		name           string
		body           []byte
		w              io.ReadWriter
		wantStatusCode int
		wantBody       []byte
	}{
		{
			"тест 200 ОК",
			[]byte("data"),
			&bytes.Buffer{},
			http.StatusOK,
			[]byte("OK"),
		},
		{
			"тест 500 InternalServerError",
			[]byte("data"),
			mockFailRW{},
			http.StatusInternalServerError,
			[]byte("OK"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "http://example.com/log"
			r := httptest.NewRequest("POST", url, bytes.NewReader(tt.body))
			w := httptest.NewRecorder()

			h := newHandler(tt.w)
			h.logHTTPHandler(w, r)

			if w.Code != tt.wantStatusCode {
				t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, tt.wantStatusCode)
			}

			resp := w.Result()
			gotBody, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("read response body error, %s", err)
				return
			}
			switch tt.wantStatusCode {
			case http.StatusOK:
				if !reflect.DeepEqual(tt.wantBody, gotBody) {
					t.Errorf("Got response=%s, want=%s", gotBody, tt.wantBody)
				}
			case http.StatusInternalServerError:
				if reflect.DeepEqual([]byte("OK"), gotBody) {
					t.Errorf("Expected response <> OK, got=%s", gotBody)
				}
			default:
				t.Errorf("undefined http status code %d", tt.wantStatusCode)
			}

			if tt.wantStatusCode != http.StatusOK {
				return
			}
			// проверяем содержимое "файла"
			gotBts, err := io.ReadAll(tt.w)
			if err != nil {
				t.Errorf("read test writer content error, %s", err)
				return
			}
			if !bytes.Contains(gotBts, tt.body) {
				t.Errorf("logHTTPHandler written = %s, want %s", gotBts, tt.body)
			}
		})
	}
}
