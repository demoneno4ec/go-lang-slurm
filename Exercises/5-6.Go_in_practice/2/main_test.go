package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
go test -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
*/

func TestClient_data(t *testing.T) {
	ds := &dummyServer{}
	ts := httptest.NewServer(http.HandlerFunc(ds.Dummy))
	defer ts.Close()

	tests := []struct {
		name     string
		dummyKey string
		url      string
		want     string
	}{
		{
			"тест 200 ОК",
			"ok",
			ts.URL,
			fmt.Sprintf(dataTemplate, "pass", "MBPadmincity101", "pass"),
		},
		{
			"тест No data, 500 error",
			"__internal_error",
			ts.URL,
			noData,
		},
		{
			"тест No data, broker json",
			"__broken_json",
			ts.URL,
			noData,
		},
		{
			"тест No data, http.Get error",
			"",
			"",
			noData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds.key = tt.dummyKey
			c := NewClient(tt.url)
			if got := c.data(); got != tt.want {
				t.Errorf("Client.data() = %v, want %v", got, tt.want)
			}
		})
	}
}

type dummyServer struct {
	key string
}

func (ds *dummyServer) Dummy(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/health" {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, `allowed only health`)
	}
	//
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, `allowed only GET`)
	}
	//

	switch ds.key {
	case "ok":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status":"pass","service_id":"MBPadmincity101","checks":{"ping_mysql":{"component_id":"mysql","component_type":"db","status":"pass"}}}`)
	case "__broken_json":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status": 400`) // broken json
	case "__internal_error":
		fallthrough
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
