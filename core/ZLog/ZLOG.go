package ZLog

import (
	"github.com/fatih/color"
	"log"
)

var BasePriority int = 0

func ZLOG(Periprity int, Module string, Message string) {
	if Periprity > BasePriority {
		log.Println(Module + ":" + Message)
	}
}

func ZLOGDISPLAY(Periprity int, Module string, Message string, ColorDisplay int) {

	if Periprity > BasePriority {
		//log.Println(Module + ":" + Message)
		if ColorDisplay == 1 {
			// default
			color.White(Module + ":" + Message)
		} else if ColorDisplay == 2 {
			// sucess
			color.Green(Module + ":" + Message)
		} else if ColorDisplay == 3 {
			// warning
			color.Yellow(Module + ":" + Message)
		} else if ColorDisplay == 4 {
			// error
			color.RedString(Module + ":" + Message)
		}

	}
}
