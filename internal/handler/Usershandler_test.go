package handler

import "testing"

func TestUsersHandler_RegisterUser(t *testing.T) {
	userHandler := UsersHandler{}
	user := userHandler.RegisterUser("100001")

	t.Log(user)
}
