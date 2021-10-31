package unit_tests_0

import("testing"
	   "fmt")

func TestMultiple(t *testing.T){
   res := multiple(2, 2)
   if res != 4{
        t.Error("Expected 4, got ", res)
}}

func TestupdateUser(t *testing.T){
  expectedUser := User{Username: "testsername",  Email: "testemail"}
  newUser := updateUser("testsername", "testemail")
  if expectedUser != newUser{
  	errorMessage := fmt.Sprintf("Expected %v, got %v ", expectedUser, newUser)
    t.Error(errorMessage)
  }
}