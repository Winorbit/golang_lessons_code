package main

import ("os"
		"io/ioutil"
		"fmt")

func write_text_into_file(text string, filepath string){
    file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
    if err != nil {
        fmt.Println(err.Error())
    } else {
	file.WriteString(text)}}


func readFromFile(filepath string) string{
    content, err := ioutil.ReadFile(filepath);
    if err != nil {
	return string(err.Error())
	} else {
	    return string(content)}}


func appendToFile(text string, filepath string){
    file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
	fmt.Println(err)
    } else{
    	defer file.Close()
	file.WriteString(text)}}


func main(){
	//write_text_into_file("Hello new!", "testfile.txt")

	//readFromFile("testfile.txt")
	//readFromFile("unexisted")
	fmt.Println(readFromFile("testf"))

	//appendToFile("Newtext into file","testfile.txt")
	//appendToFile("Newtext into file \n","testfile.txt")
	//appendToFile("Some lorem iosum","unexisted")
}

