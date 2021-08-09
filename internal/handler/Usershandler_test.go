package handler

import (
	"fmt"
	"testing"
)

func TestUsersHandler_RegisterUser(t *testing.T) {

	tests := []struct {
		name string
		args string
		want int
	}{
		{"test1", "1003", 24},
		{"test2", "1003", 24},
		{"test3", "1004", 24},
	}
	userHandler := UsersHandler{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := userHandler.RegisterUser(tt.args)
			if op, ok := got.Data.(string); ok && len(op) != tt.want {
				t.Errorf("RegisterUser() = %v, want %v", got, tt.want)
			}
			fmt.Println(got.Data)
		})
	}

}
