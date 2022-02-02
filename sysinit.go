/*
 	ZFRAME with go  Engine
	Author By Siavash Rafiei
 	Create data 1400/10/17


*/
package main

import (
	"ZGOPROJ/compile/compilerule"
	"ZGOPROJ/core/initialize"
	_ "ZGOPROJ/core/initialize"
	"ZGOPROJ/core/libtest"
	"ZGOPROJ/core/security"
	"ZGOPROJ/webserver"
	"ZGOPROJ/zfconfig"
	_ "ZGOPROJ/zfconfig"
	"encoding/base64"
	"fmt"
	"github.com/fatih/color"
	"time"
)

func main() {

	//securitycheck()
	dt := time.Now()
	color.Green("Start Application ..... ")
	color.Green("	Start time is: %s", dt.String())

	libtest.StartTest()
	color.Cyan(zfconfig.ZFGOVersion)
	var c = zfconfig.GetStringZFrameLogo()
	color.Blue(c)
	err := zfconfig.LoadConfigure()
	if err != nil {
		color.Red("\t\t [ Error In Loading Config File ]")
		return
	}
	defer ExitSys()
	if zfconfig.ConfigisLoaded {
		color.Cyan("\t\t [ GZFrame Project Name : " + zfconfig.Project_Name + " ]")
		initialize.StartSystemActivity()
		myc := make(chan int)
		webserver.StartWebApplicationServerZFGO(myc)
	}

}

func ExitSys() {
	dt := time.Now()
	color.Red("	Close Time : %s", dt.String())
}

func securitycheck() {
	//key, err := base64.StdEncoding.DecodeString("CLUBCOREsystemTokenValue@12791279846InfomationTechnologyWithZFarmeApp")
	data, err := base64.StdEncoding.DecodeString("2uwxT8gaPDEr9ttCPtzn2jUKTYf9YoHJ//I6MkF8vfSFiredEE8DUngcHjSrfklR8Fcy8bDoIa165/WA4glwZQbHS4SLcAVr0cmdGT3iXxUhWrVbugx6ZpcQVNiB+/D2zUCbwBqQH42UPqYCLGYpMRbxeWCHO+6PVE4ssIXJEdeAfWtCA93MgVyQN5WCD8njoW5Iyrv6VC/dWOU1GKSdm83g3LxltrbeQrMKV9TR8+dvXl2r/8s1hdq4WJfHYR49")
	if err != nil {
		fmt.Print(err)
		return
	}

	//var desc []byte
	//base64.StdEncoding.Encode(desc, []byte("CLUBCOREsystemTokenValue@12791279846InfomationTechnologyWithZFarmeApp"))
	val, errc := security.DecryptAes128Ecb(data, compilerule.GetToken())
	//key, err := hex.DecodeString("CLUBCOREsystemTokenValue@12791279846InfomationTechnologyWithZFarmeApp")
	if errc != nil {
		fmt.Println(errc)
		return
	}

	var str1 string = string(val)
	var str2 string = string(data)

	fmt.Println(str1)
	fmt.Println(str2)
}
