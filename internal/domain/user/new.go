package user

func New(email Email, name string, nameConfirmed bool) (User, error) {
	u := User{email, name, nameConfirmed}
	return u, u.Validate()
}
