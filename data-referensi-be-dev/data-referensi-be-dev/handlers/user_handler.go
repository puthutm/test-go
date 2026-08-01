package handlers

import (
	"data-referensi/pkg/auth"
)

func GetUserId() string {
	auth.Mu.Lock()
	user := auth.UserClaimsGlobal
	auth.Mu.Unlock()

	return user.ID
}

func GetUserEmail() string {
	auth.Mu.Lock()
	user := auth.UserClaimsGlobal
	auth.Mu.Unlock()

	return user.Email
}
