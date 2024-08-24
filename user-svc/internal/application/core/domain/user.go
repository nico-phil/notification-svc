package domain

type User struct {
	ID int64 `json:"id"`
	Firstname string `json:"first_name"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
	HashPassword string `json:"HashPassword"`
}

type Device struct {
	ID int64 `json:"id"`
	DeviceToken string `json:"device_token"`
	DeviceType string `json:"device_type"`
	UserID int64 `json:"user_id"`
}

func(u *User) EncriptPassword() error {

	return nil
}