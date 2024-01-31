package main

import (
	"fmt"
	"test_golang/initializers"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	fmt.Println("Hallo")
}