//go:build integration
// +build integration

package docker_test

import (
	"context"
	"gotemplate/Exercises/7.Docker/1/internal/docker"
	"strings"
	"testing"
)

func TestClient_ListCID(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			"test redis",
			"redis",
			false,
		},
		{
			"test nginx",
			"nginx",
			false,
		},
		{
			"test empty",
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// создаем нового клиента
			cli, err := docker.New()
			if err != nil {
				t.Errorf("NewClient error %s", err)
				return
			}
			// получаем контейнеры
			got, err := cli.ListCID(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListCID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// если имя контейнера не задано, то прекращаем тест
			if tt.want == "" {
				return
			}
			//
			exists := false
			// перебираем весь список и ищем совпадения/содержимое по имени
		loop:
			for _, c := range got {
				if strings.Contains(c.Name, tt.want) {
					exists = true
					break loop
				}
			}
			// ожидаем найти, если нет, то фэйл
			if !exists {
				t.Errorf("expected exists containr with name like %s, but not", tt.want)
			}
		})
	}
}
