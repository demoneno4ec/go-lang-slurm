package main

import (
	"context"
	"runtime"
	"testing"
	"time"
)

func Test_realMain(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		name string
		n    int
	}{
		{
			"тест запуска 5 горутин",
			5,
		},
		{
			"тест запуска 0 горутин",
			0,
		},
		{
			"тест запуска 100 горутин",
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := runtime.NumGoroutine()
			go realMain(ctx, tt.n)
			time.Sleep(time.Millisecond * 250)

			inTime := runtime.NumGoroutine()

			time.Sleep(timeout)
			after := runtime.NumGoroutine()

			// проверяем, что realMain запустила n горутин
			if inTime != before+tt.n+1 { // +1 на саму realMain
				t.Errorf("realMain start %d goroutine, want %d", inTime, before+tt.n+1)
			}
			// проверяем, что после завершения realMain все горутины тоже завершились
			if after != before {
				t.Errorf("after realMain end %d goroutine, want %d", after, before)
			}
		})
	}
}
