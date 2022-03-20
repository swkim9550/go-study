package main

import (
	"fmt"
	"strings"
	// "github.com/serranoarevalo/something"
)

func lenAndUpper(name string)(int,string){
	return len(name), strings.ToUpper(name)
}

func multply(a ,b int) int{
	return a * b
}

func repeatMe(words ...string){
	defer fmt.Println("reapeat Me end ...") //메소드가 끝나면 실행된다.
	fmt.Println(words)
}

//for each
func superAdd(numbers ...int) int {
	defer fmt.Println("super Add end .. ")
	total := 0
	for _ ,number := range numbers {
		total += number		
	}
	return total;
}

//for 문
func superAddV2(numbers ...int) int {
	defer fmt.Println("super Add  V2 end .. ")
	for i:=0 ; i< len(numbers); i++{
		fmt.Println(numbers[i])
	}
	return 1;
}


func main() {
	result := superAdd(1,2,3,4,5,6)
	fmt.Println(result)

	//superAddV2(1,2,3,4,5)


	//totalLenght, upperName := lenAndUpper("seongwoo")
	//fmt.Println(totalLenght,upperName)

	//totalLenght2, _ := lenAndUpper("seongwoo")
	//fmt.Println(totalLenght2)

	//repeatMe("test1","test1","test1","test1","test1")


	// fmt.Println(multply(2,2))
	// name := "test"
	// //var name string = "seongwoo"
	// name = "lynn"
	// fmt.Println(name)
}


