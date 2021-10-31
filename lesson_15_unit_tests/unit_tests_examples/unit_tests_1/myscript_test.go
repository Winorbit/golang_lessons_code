package unit_tests_1

import("testing"
	     "fmt")

func TestupdateUser(t *testing.T){
  expectedUser := User{Username: "testsername",  Email: "testemail"}
  newUser := User{}
  newUser.updateUser("testsername", "testemail")
  if expectedUser != newUser{
  	errorMessage := fmt.Sprintf("Expected %v, got %v ", expectedUser, newUser)
    t.Error(errorMessage)
  }
}


func TestValidatingEmail(t *testing.T){
  invalidEmail := validateEmail("testemail")
  if invalidEmail{
    errorMessage := fmt.Sprintf("Expected False with email %v, got True", invalidEmail)
    t.Error(errorMessage)
  }

  validEmail := validateEmail("testemail@gmail.com")
  if !validEmail{
    errorMessage := fmt.Sprintf("Expected True with email %v, got False", validEmail)
    t.Error(errorMessage)
  }
}