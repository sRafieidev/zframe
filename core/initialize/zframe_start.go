package initialize

import (
	"ZGOPROJ/core/DatabaseManagement"
	"ZGOPROJ/core/ServiceManager"
	"ZGOPROJ/core/ZFrameSystem"
	color2 "github.com/fatih/color"
)

func StartSystemActivity() {

	// Before This Step System Read Config File And Config Is Loaded
	DatabaseManagement.InitDatabaseMetadata()
	ZFrameSystem.LoadSystemInZframeProject()
	servicelist := ServiceManager.LoadService()
	for i := 0; i < len(servicelist); i++ {
		color2.Blue("Engine Load Function :( " + servicelist[i].ProcessFunctionName + " )")
	}

}
