package unit_tests_1

import "regexp"

type User struct{Username string
				         Email string}

func (u *User) updateUser(username string, email string){
  (*u).Username = username
  (*u).Email = email
}

func validateEmail(email string) bool {
  emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
  res := emailRegex.MatchString(email)
  return res
}