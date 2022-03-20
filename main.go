package main

import (
	"fmt"
	"strings"
	// "github.com/serranoarevalo/something"
)

//multi return
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// a,b int 형 인자 받기
// return 값은 int
func multply(a, b int) int {
	return a * b
}

//defer 메소드가 끝나면 실행 된다
// ... 여러 인자를 받을 수 있다
//repeatMe("test1","test1","test1","test1","test1")
func repeatMe(words ...string) {
	defer fmt.Println("reapeat Me end ...")
	fmt.Println(words)
}

//for each
//result := superAdd(1, 2, 3, 4, 5, 6)
func superAdd(numbers ...int) int {
	defer fmt.Println("super Add end .. ")
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//for 문
func superAddV2(numbers ...int) int {
	defer fmt.Println("super Add  V2 end .. ")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

// if -else
// else 는 따로 써주지 않아도 된다.
func canIdrink(age int) bool {
	if age < 18 {
		return false
	}
	return true

}

// if -else
// if else 안에서 variable expression 사용
func canIdrinkV2(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true

}

func main() {
	fmt.Println(canIdrink(18))
}
