package main

import (
	_ "shop-api/middleware"
	_ "shop-api/routers"
	"testing"

	_ "github.com/lib/pq"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Base case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
