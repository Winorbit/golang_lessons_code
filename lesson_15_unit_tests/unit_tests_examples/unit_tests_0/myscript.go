package unit_tests_0

func multiple(a int, b int) int {
  res := a*b
  return res
}

type User struct{Username string
				 Email string}

func updateUser(username string, email string) User {
  newUser := User{Username: username,
				  Email: email}
  return newUser
}