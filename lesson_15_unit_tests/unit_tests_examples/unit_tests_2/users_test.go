package unit_tests_2

import("testing"
	     "fmt")

func TestUpdateUser(t *testing.T){
  email := "me@mail.com"
  username := "me"
  expectedUser := User{Username: username,  Email: email}

  newUser := User{}
  err, user := newUser.updateUser(username, email)

  if err != nil{
    errorMessage := fmt.Sprintf("Expected err as nil, got %v", err)
    t.Error(errorMessage)
  }

  if expectedUser != newUser{
  	errorMessage := fmt.Sprintf("Expected %v, got %v", expectedUser, user)
    t.Error(errorMessage)
  }
}


func TestUpdateUserInvalidEmail(t *testing.T){
  invalidEmail := "megmail.com"
  username := "me"
  newUser := User{}

  err, user := newUser.updateUser(username, invalidEmail)
  if err == nil{
    errorMessage := fmt.Sprintf("Expected error with invalid email, got user: %v  and err: %v ", user, err)
    t.Error(errorMessage)
  }
}


func TestValidatingEmail(t *testing.T){
  invalidEmail :="testemail"
  validationInvalidEmail := validateEmail(invalidEmail)
  if validationInvalidEmail{
    errorMessage := fmt.Sprintf("Expected False with email %v, got True", invalidEmail)
    t.Error(errorMessage)
  }

  validEmail := "testemail@gmail.com"
  validationValidEmail := validateEmail(validEmail)
  if !validationValidEmail{
    errorMessage := fmt.Sprintf("Expected True with email %v, got False", validEmail)
    t.Error(errorMessage)
  }
}