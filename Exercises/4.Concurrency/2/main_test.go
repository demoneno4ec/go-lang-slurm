package main

import (
	"testing"
	"time"
)

func Test_realMain(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		wantErr  bool
	}{
		{
			"тест менее таймаута",
			time.Second,
			false,
		},
		{
			"тест гораздо менее таймаута",
			time.Millisecond * 25,
			false,
		},
		{
			"тест более таймаута",
			time.Second * 3,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			err := realMain(tt.duration)
			duration := time.Since(start)

			if (err != nil) != tt.wantErr {
				t.Errorf("realMain() error = %v, wantErr %v", err, tt.wantErr)
			}

			// проверяем, что функция работала не дольше таймаута 2 секунды+100мс
			// и не дольше времени работы горутин+100мс
			// где +100мс - окно на время завершения горутин и завершения функции

			const serviceWindow = time.Millisecond * 100
			if duration > timeout+serviceWindow || duration > tt.duration+serviceWindow {
				t.Errorf("duration of realMain() = %s, expected less then %s or %s", duration, timeout+serviceWindow, tt.duration+serviceWindow)
			}
		})
	}
}
