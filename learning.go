package main

import "fmt"
import "time"

func main() {
	fmt.Println("Hello world")

	fmt.Println("go" + "lang")

	fmt.Println("1 + 1 =", 1+1)

	var a = "Hello"
	var b int = 2
	c := 3
	const d string = "Wont change"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	for i:=0; i<4;i++{
		if i%2 == 0{
			fmt.Printf("loop %d\n", i)
		}
	}

	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("It's before noon")
		default:
			fmt.Println("It's after noon")
	}

	var arr[5] int
	arr[3] = 200

	fmt.Println("after set", arr)
	fmt.Println("Length", len(arr))

	slice := make([]string, 3)
	slice[0] = "aawa"
	slice[1] = "bbbbb"
	slice[2] = "cc"

	fmt.Println(slice)
	fmt.Println(slice[1])

	slice = append(slice, "fdfdfdf")
	fmt.Println("appended", slice)

	fmt.Println("slice operator", slice[2:4])

	 m := make(map[string]int)
	 m["k1"] = 4

	 value, present := m["k1"]
	 fmt.Println(value)
	 fmt.Println(present)
}
