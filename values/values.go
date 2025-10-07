package values

import (
	"fmt"
	"time"
)

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

func Loops() {

	// i := 0
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for j := 0; j <= 5; j++ {
	// 	fmt.Printf("Value:%v \n", j)
	// }

	// for i := range 3 {
	// 	fmt.Printf("3rd loop %v\n", i)
	// }
	// c := 0
	// for {
	// 	fmt.Printf("print %v\n", c)
	// 	if c == 4 {
	// 		break
	// 	}
	// 	c = c + 1
	// }
	for d := range 20 {
		if d == 2 {
			continue
		}
		fmt.Printf("print %v\n", d)

		if d == 10 {
			break
		}

	}

	// aa = 250

}

func IfElse() {
	if n := 6; n == 7 {
		fmt.Printf("n %v,\n", n)
	} else {
		fmt.Printf("m  else %v\n", n)
	}
}

func SwithCase() {

	const i int = 2

	switch i {
	case 1, 2, 3:
		fmt.Println("1 2 or 3")
	case 5:
		fmt.Println("Its 5")
	default:
		fmt.Println("Unknown")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Not Weekend")
	}

	switch {
	case 2 > 4:
		fmt.Println("2 is greater than 1")
	default:
		fmt.Println("No its Not")
	}

	checkType := func(i interface{}) {

		switch i.(type) {
		case string:
			fmt.Println("Its a string")
		case int:
			fmt.Println("Its a Integer")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("Unidentified")
		}
	}

	const mm = "shiva"

	checkType(mm)
	checkType(false)
	checkType(25)
	checkType(25.5)

}

func Arrays() {
	var a [5]int

	fmt.Printf("Before %v\n", a)

	a[2] = 3

	fmt.Printf("After %v\n", a)

	var b = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("b %v %v\n", b, len(b))

	c := [...]int{1, 2, 3}

	fmt.Printf("c %v %v\n", c, len(c))

	var ab [2][3]int

	for i := range 2 {
		for j := range 3 {
			ab[i][j] = j + 1
		}
	}
	fmt.Println(ab)

	var abc = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(abc)

}
