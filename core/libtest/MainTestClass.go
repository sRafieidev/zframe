package libtest

import "C"
import (
	"ZGOPROJ/compile/compilerule"
	"fmt"
	"runtime"
)

func StartTest() {

	Gofunctest()
	//	Testshift()
	//TestPointerStructingo()

	// Show Application Logo
	/*
		yellow := color.New(color.FgYellow).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()
		fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red("error"))
	*/
	//ZGJSON.TestJsonParse()
	//libtest.TestChannel()
	//libtest.Gofunctest()
	//mssqltest()
	//testc()
	//oracletest()
	//start web application server
	if compilerule.DEBUG {
		fmt.Println(runtime.Version())
	}

}
