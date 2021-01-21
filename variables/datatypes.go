package main
import "fmt"

// Go has the following datatypes:
// 	bool

// 	string

// 	int  int8  int16  int32  int64
// 	uint uint8 uint16 uint32 uint64 uintptr

// 	byte // alias for uint8

// 	rune // alias for int32
// 			// represents a Unicode code point

// 	float32 float64

// 	complex64 complex128

var (
	toggle bool
	id int
	name string
	GamaConst complex64
)

func main() {
	toggle = false
	id = 23
	name = "akhil"
	GamaConst = -5+2i

	fmt.Printf("Type: %T, Value: %v \n", toggle, toggle)
	fmt.Printf("Type: %T, Value: %v \n", id, id)
	fmt.Printf("Type: %T, Value: %v \n", name, name)
	fmt.Printf("Type: %T, Value: %v \n", GamaConst, GamaConst)
}

