package usertest

import "hackbar-copilot/internal/domain/user"

var ExampleUsers = []user.User{
	{
		Email: "john.doe@example.test",
		Name:  "John Doe",
	},
}

var ExampleUsersIter = IterWithNilError(ExampleUsers)
