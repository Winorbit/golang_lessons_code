package main

import "packages_3/src/some_package"

import "fmt"

func main() {
    some_package.Imported()
    some_package.Imported2()
    fmt.Println(some_package.ValueFromPackage)
    //fmt.Println(some_package.valueFromPackage)
    fmt.Println(some_package.SliceFromPackage)


}
