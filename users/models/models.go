package models

type User struct {
	FullName    string
	PhoneNumber *string
}

func NewUser(fullName string, phoneNumber *string) *User {
	return &User{
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}
