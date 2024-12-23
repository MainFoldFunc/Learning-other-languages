package main
package custpack

import (
	"fmt"
	"github.com/MainFoldFunc/Learning-other-languages/Go/Projects/custpack"
)


func main(){
	var str string
	fmt.Println("What is the string: ")
	fmt.Scan(&str)

	fmt.Printf("The string reversed is: %v", custpack.Reverse(str))


}
