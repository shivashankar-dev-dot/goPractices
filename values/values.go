package values

import "fmt"

func CheckVals() {
	fmt.Println("shiva" + "go")
	fmt.Println(10 + 2)
	fmt.Println(7 / 2)
	fmt.Println(true || false)
	fmt.Println(!true)

}

func Vars() {
	var a = "shiva"
	fmt.Println(a)

	var c, b int = 1, 2

	fmt.Printf("c=%v, b=%v\n", c, b)

	var init int
	var initCheck bool
	fmt.Printf("init=%v, initCheck=%v\n", init, initCheck)

	check := false

	fmt.Println(check)

	var (
		narendra int     = 420
		shiva    string  = "good boy"
		hema     float32 = 45.56
	)

	fmt.Printf("narendra %v shiva %v hema %v", narendra, shiva, hema)

}

func Consts() {

	const aa int = 20050
	fmt.Printf("check=%v", aa)

	// aa = 250

	fmt.Printf("check2=%v", aa)

}
