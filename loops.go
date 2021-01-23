// go only has for loop
// it has several variants which can sever as while loop too
package main
import "fmt"

func main(){
	sum := 0 // datatype is automatically infered by Go
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("Sum is: %v\n", sum)
	fmt.Printf("Sum in while: %v\n", whileLoop())
}


// when used as while loop

func whileLoop() int{
	i := 5
	sum := 0
	for i > 0{
		sum += i
		i--
	}
	return sum
}