package user

type User struct {
	Email         Email
	Name          string
	NameConfirmed bool
}

type Email string
