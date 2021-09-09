package main

import ("os"
	//"io/ioutil"
	"fmt"
	"bufio"
	)
	
func main(){
/*
os.Create("hello.txt")
fmt.Println(os.Open("hello.txt"))
*/

// Возвращает два объекта - сам файл и ошибку. Если ошибки нет, то это будет nil

/*
file, err := os.Open("hello.txt")
fmt.Println(file, err)
defer file.Close()
*/

// Запись в файл


//os.Create("testfile.txt")
// а вот так не работает: file, err := os.Open("testfile.txt"). надо именно так, с разрешением 

/*
file, err := os.OpenFile("testfile.txt", os.O_RDWR, 0644)
if err != nil {
        fmt.Println(err.Error())
    }
defer file.Close()
file.WriteString("Text example")
*/


// Чтение из файла

file, err := os.Open("testfile.txt")
if err != nil {
        fmt.Println(err.Error())
    }
defer file.Close()
/*
content,_ := ioutil.ReadFile("testfile.txt")

fmt.Println(content)
fmt.Println(string(content))
*/


scanner := bufio.NewScanner(file)
fmt.Println(scanner.Text())

// Читаем строчка за строчкой

for scanner.Scan() {
    fmt.Println(scanner.Text())
}


}

