package libtest

import (
	"github.com/fatih/color"
	"time"
)

func EvenNumbersTillEight(even chan int) {
	i := 2
	for i < 1000 {
		var dt = time.Now()
		even <- i
		i = i + 2
		color.Green("	Fn_Even time is: %s  value(%d)", dt.String(), i)
		//fmt.Println("Fn_Even ", i)
	}
	close(even)
}
func OddNumberTillEight(odd chan int) {
	i := 1
	for i < 1000 {
		var dt = time.Now()
		odd <- i
		i = i + 2
		color.Blue("	Fn_Odd time is: %s   value(%d)", dt.String(), i)
		//fmt.Println("Fn_Odd ", i)
	}
	close(odd)
}
func MyFunctionTest(myc chan int) {
	i := 10
	for i < 10000 {
		var dt = time.Now()
		i = i + 1
		if i%5 == 0 {
			myc <- i
		}
		color.Red("	Fn_myc time is: %s   value(%d)", dt.String(), i)
		//fmt.Println("Fn_myc ", i)
	}
	close(myc)
}
func TestChannel() {
	even := make(chan int)
	odd := make(chan int)
	myc := make(chan int)
	go EvenNumbersTillEight(even)
	go OddNumberTillEight(odd)
	go MyFunctionTest(myc)
	for {
		even, ok1 := <-even
		odd, ok2 := <-odd
		myc, ok3 := <-myc
		if ok1 == false || ok2 == false || ok3 == false {
			break
		}
		color.White("	even %d evenok %s odd %d oddok %s mycvalue %d mycok %s ", even, showboolstring(ok1), odd, showboolstring(ok2), myc, showboolstring(ok3))
	}
}

func showboolstring(val bool) string {
	if val {
		return "true"
	}
	return "false"
}
