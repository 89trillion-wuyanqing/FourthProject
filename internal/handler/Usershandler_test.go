package handler

import "testing"

func TestUsersHandler_RegisterUser(t *testing.T) {
	userHandler := UsersHandler{}
	user, err := userHandler.RegisterUser("100001")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(user)
}
