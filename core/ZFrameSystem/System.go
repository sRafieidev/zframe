package ZFrameSystem

import (
	"ZGOPROJ/core/DatabaseManagement"
	"ZGOPROJ/core/DatabaseManagement/DTOTable"
	"ZGOPROJ/core/ZLog"
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	"log"
	"strings"
)

var SysList []DTOTable.Sys_System = nil

func LoadSystemInZframeProject() {
	rowsdata := DatabaseManagement.ExecuteDataTable(DatabaseManagement.MetaDataDb, "Select * from SYS_SYSTEM")
	if rowsdata != nil {
		for rowsdata.Next() {
			nowSystem := DTOTable.Sys_System{}
			err := rowsdata.Scan(&nowSystem.SYS_SYSTEM_ID, &nowSystem.NAME, &nowSystem.TITLE, &nowSystem.SYS_DATABASE_TYPE_ID, &nowSystem.DATABASECONNECTIONSTRING, &nowSystem.DATABASEUSERNAME, &nowSystem.DATABASEPASSWORD, &nowSystem.DATABASESERVERADDRESS, &nowSystem.DATABASESCHEMANAME, &nowSystem.DATABASEPORT, &nowSystem.APPLICATIONURL, &nowSystem.JDBCCNN, &nowSystem.LTR, &nowSystem.SORTORDER, &nowSystem.IS_ACTIVE, &nowSystem.CONNECTION_POOL_COUNT)
			ZLog.ZLOGDISPLAY(100, " InitSys ", "Load System :["+nowSystem.NAME+"]", 2)
			if err == nil {
				CreateSystemDatabaseConnection(&nowSystem)
				SysList = append(SysList, nowSystem)

			} else {
				fmt.Println(err.Error())
			}
		}
	}
}

func CreateSystemDatabaseConnection(system *DTOTable.Sys_System) {
	color.Blue("Start Create Connection :" + system.NAME)
	if system.SYS_DATABASE_TYPE_ID == 1 {
		LoadingSystemWithMSSQLServerDatabase(system)
	} else {
		LoadingSystemWithOracleDatabase(system)
	}
}

func LoadingSystemWithMSSQLServerDatabase(system *DTOTable.Sys_System) {
	var err error
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", system.DATABASESERVERADDRESS, system.DATABASEUSERNAME, system.DATABASEPASSWORD, system.DATABASEPORT, system.DATABASESCHEMANAME)
	system.SYSTEMDB, err = sql.Open("sqlserver", connString)
	if err != nil {
		color.Red("***Error For Connected To System Database => %s", system.NAME)
		log.Fatal("Error creating connection pool: " + err.Error())
	} else {
		system.SYSTEMDB.SetMaxIdleConns(system.CONNECTION_POOL_COUNT / 2)
		system.SYSTEMDB.SetMaxOpenConns(system.CONNECTION_POOL_COUNT)
		color.Green("\t\t ********* Connected To System Database MSSQL => %s", system.NAME)
	}

	// Close the database connection  mmpool after program executes
	//defer MetaDataDb.Close()
}

func LoadingSystemWithOracleDatabase(system *DTOTable.Sys_System) {

	var err error
	var OrcCnnFromJdbi string = GetGoLangConnectionStringFromJDBCConnectionOracle(system.JDBCCNN)
	connString := fmt.Sprintf("user=%s password=%s connectString=%s ", system.DATABASEUSERNAME, system.DATABASEPASSWORD, OrcCnnFromJdbi)
	system.SYSTEMDB, err = sql.Open("godror", connString)
	if err != nil {
		color.Red("\t\t Error For Connected To System Database =>" + system.NAME)
		log.Fatal("Error creating connection pool: " + err.Error())

	} else {
		system.SYSTEMDB.SetMaxIdleConns(system.CONNECTION_POOL_COUNT / 2)
		system.SYSTEMDB.SetMaxOpenConns(system.CONNECTION_POOL_COUNT)
		color.Green("\t\t ********* Connected To System Database Oracle =>" + system.NAME)
	}

	// Close the database connection  mmpool after program executes
	//defer MetaDataDb.Close()
}

func GetGoLangConnectionStringFromJDBCConnectionOracle(jdbcurl string) string {
	startindex := strings.Index(jdbcurl, "//")
	if startindex > 0 {
		var rt string = jdbcurl[startindex+2 : len(jdbcurl)]
		return rt
	} else {
		startindex = strings.Index(jdbcurl, "@")
		var rt string = jdbcurl[startindex+1 : len(jdbcurl)]
		return rt
	}
	return ""
}
