package entity

type UserCredential struct {
	UserID       string
	PasswordHash string
	Salt         string
}
