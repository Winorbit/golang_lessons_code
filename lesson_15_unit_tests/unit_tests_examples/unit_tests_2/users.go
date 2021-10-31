package unit_tests_2

import("regexp"
       "errors"
       "fmt")

type User struct{Username string
				         Email string}

func (u *User) updateUser(username string, email string) (error, *User) {
  if validateEmail(email){
    (*u).Username = username
    (*u).Email = email
    return nil, u
    } else {
      errorMessage := fmt.Sprintf("Invalid email %v", email)
      err := errors.New(errorMessage)
      return err, u
  }
}

func validateEmail(email string) bool {
  emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
  res := emailRegex.MatchString(email)
  return res
}