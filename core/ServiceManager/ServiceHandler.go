package ServiceManager

import (
	"ZGOPROJ/core/DatabaseManagement"
	"ZGOPROJ/core/DatabaseManagement/QueryParsing"
	"ZGOPROJ/core/ZFrameSystem"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

var ServiceList [][]string

type SysRestFunction struct {
	SysRestFunctionId             int
	SysSystemId                   int
	FunctionName                  string
	BaseUrl                       string
	Authentication                bool
	FunctionRuleId                int
	FunctionQuery                 string
	FunctionQueryWithTokenInParam string
	FunctionQueryForExecute       string
	CallMethod                    int
	FunctionDescription           string
	OttAction                     interface{}
	DefaultOutput                 string
	ErrorOutput                   string
	ProcessFunctionName           string
	QueryType                     int
	ParameterList                 []QueryParsing.QueryParameter
	UserGroupAccess               []string
}

var ServiceListStruct []SysRestFunction = nil

func LoadService() []SysRestFunction {
	if ServiceListStruct == nil {
		rows := DatabaseManagement.ExecuteDataTable(DatabaseManagement.MetaDataDb, "select * from sys_rest_function")
		for rows.Next() {
			nowfunc := SysRestFunction{}
			err := rows.Scan(&nowfunc.SysRestFunctionId, &nowfunc.SysSystemId, &nowfunc.FunctionName, &nowfunc.BaseUrl, &nowfunc.Authentication, &nowfunc.FunctionRuleId, &nowfunc.FunctionQuery, &nowfunc.CallMethod, &nowfunc.FunctionDescription, &nowfunc.OttAction, &nowfunc.DefaultOutput, &nowfunc.ErrorOutput)
			if err == nil {
				nowfunc.ProcessFunctionName = nowfunc.BaseUrl + strings.ReplaceAll(nowfunc.FunctionName, ".zjs", "")
				//color.Blue("Process Service :" + nowfunc.ProcessFunctionName)
				nowfunc.QueryType = QueryParsing.GetQueryType(nowfunc.FunctionQuery)
				nowfunc.FunctionQueryWithTokenInParam = QueryParsing.ChangeQueryWithTokenToParameter(nowfunc.FunctionQuery, nowfunc.QueryType)
				nowfunc.ParameterList = QueryParsing.GetQueryParameter(nowfunc.FunctionQueryWithTokenInParam, nowfunc.QueryType)
				GenerateQueryForExecuteFromRawQuery(&nowfunc)
				getServiceUserGroupAccess(&nowfunc)
				sort.Slice(nowfunc.ParameterList[:], func(i, j int) bool { return nowfunc.ParameterList[i].IndexItem > nowfunc.ParameterList[j].IndexItem })
				ServiceListStruct = append(ServiceListStruct, nowfunc)
			} else {
				fmt.Println(err.Error())
			}
		}
	}
	return ServiceListStruct
}

func GenerateQueryForExecuteFromRawQuery(function *SysRestFunction) {
	var Query = function.FunctionQueryWithTokenInParam

	if function.QueryType > 1 {
		if function.QueryType == 2 {
			var stin int = strings.Index(strings.ToLower(Query), "ppsql")
			Query = Query[stin+5 : len(Query)]
		} else if function.QueryType == 3 {
			var stin int = strings.Index(strings.ToLower(Query), "plsql")
			Query = Query[stin+5 : len(Query)]
		} else if function.QueryType == 4 {
			var stin int = strings.Index(strings.ToLower(Query), "rest")
			Query = Query[stin+4 : len(Query)]
		}
	}

	if function.ParameterList != nil {
		sort.Slice(function.ParameterList[:], func(i, j int) bool {
			return len(function.ParameterList[i].ParameterFullName) > len(function.ParameterList[j].ParameterFullName)
		})
		NowDataBaseTypeForParsing := GetDatabaseTypeWithSystemID(function.SysSystemId)
		if NowDataBaseTypeForParsing == 1 { // MSSQL Server
			for a := 0; a < len(function.ParameterList); a++ {
				var nowC string = "@" + function.ParameterList[a].ParameterName
				Query = strings.ReplaceAll(Query, function.ParameterList[a].ParameterFullName, nowC)
			}
		} else if NowDataBaseTypeForParsing == 2 { //Oracle
			var lastparamid int = 0
			for a := 0; a < len(function.ParameterList); a++ {
				var nowC string = ":" + strconv.Itoa(a)
				lastparamid = a
				Query = strings.ReplaceAll(Query, function.ParameterList[a].ParameterFullName, nowC)
			}
			if function.QueryType == 3 { //PLSQL FOR REPLACE LAST ? WITH :PARAMC

				lastparamid++
				var indexofquestionmark int = strings.LastIndex(Query, "?")
				var lastcursername string = ":" + strconv.Itoa(lastparamid)
				Query = strings.Replace(Query, "?", lastcursername, indexofquestionmark)
			}
		}
		function.FunctionQueryForExecute = Query
	} else {
		function.FunctionQueryForExecute = function.FunctionQuery
	}
}

func getServiceUserGroupAccess(function *SysRestFunction) {
	queryforloadusergroupaccess := "SELECT SYS_USERGROUP_ID FROM SYS_REST_ACCESS WHERE SYS_REST_FUNCTION_ID = " + strconv.Itoa(function.SysRestFunctionId)
	rows := DatabaseManagement.ExecuteDataTable(DatabaseManagement.MetaDataDb, queryforloadusergroupaccess)
	var groupacess []string = nil
	for rows.Next() {

		now_sys_group_id := ""
		err := rows.Scan(&now_sys_group_id)
		if err == nil {
			groupacess = append(groupacess, now_sys_group_id)
		}
	}

	function.UserGroupAccess = groupacess

}

func GetDatabaseTypeWithSystemID(systemid int) int {
	for a := 0; a < len(ZFrameSystem.SysList); a++ {
		if ZFrameSystem.SysList[a].SYS_SYSTEM_ID == systemid {
			return ZFrameSystem.SysList[a].SYS_DATABASE_TYPE_ID
		}
	}
	return 0
}

func DoServiceAction(db *sql.DB, Query string, parameter []QueryParsing.QueryParameter, Values []QueryParsing.Dictionery, dbtype int, function *SysRestFunction) (int, *sql.Rows, driver.Rows, error) {

	//db.QueryRow(Query)
	if function.QueryType != 4 {
		//color.Yellow("Start Query :" + Query)
		stm, inerror := db.Prepare(Query)
		if inerror != nil {
			return 0, nil, nil, inerror
		}
		args := make([]interface{}, len(parameter))
		//sort.Slice(parameter[:], func(i, j int) bool { return parameter[i].IndexItem > parameter[j].IndexItem })
		if dbtype == 1 { //MSSQL Server
			for i := range parameter {
				if parameter[i].ParameterType == "int" {
					args[i] = sql.Named(parameter[i].ParameterName, GetIntValueFromParameterName(Values, parameter[i].ParameterName))
				} else if parameter[i].ParameterType == "lng" {
					args[i] = sql.Named(parameter[i].ParameterName, GetIntValueFromParameterName(Values, parameter[i].ParameterName))
				} else if parameter[i].ParameterType == "str" {
					args[i] = sql.Named(parameter[i].ParameterName, GetStringValueFromParameterName(Values, parameter[i].ParameterName))
				}
			}
		} else if dbtype == 2 { // Oracle
			if function.QueryType < 2 {
				var counter int = -1
				for i := range parameter {
					counter++
					if parameter[i].ParameterType == "int" {
						//args[i] = sql.Named(":"+strconv.Itoa(counter), GetIntValueFromParameterName(Values, parameter[i].ParameterName))
						args[i] = GetIntValueFromParameterName(Values, parameter[i].ParameterName)
					} else if parameter[i].ParameterType == "lng" {
						//args[i] = sql.Named(":"+strconv.Itoa(counter), GetIntValueFromParameterName(Values, parameter[i].ParameterName))
						args[i] = GetIntValueFromParameterName(Values, parameter[i].ParameterName)
					} else if parameter[i].ParameterType == "str" {
						//args[i] = sql.Named(":"+strconv.Itoa(counter), GetStringValueFromParameterName(Values, parameter[i].ParameterName))
						args[i] = GetStringValueFromParameterName(Values, parameter[i].ParameterName)
					}
				}
			} else if function.QueryType == 2 { // ppsql
				// process parameter for procedure value
				var counter int = -1
				for i := range parameter {
					counter++
					if parameter[i].ParameterType == "int" {
						//args[i] = sql.Named(":"+strconv.Itoa(counter), GetIntValueFromParameterName(Values, parameter[i].ParameterName))
						if !parameter[i].IsOutput {
							args[i] = GetIntValueFromParameterName(Values, parameter[i].ParameterName)
						} else {
							var Paramint int
							args[i] = &Paramint
						}
					} else if parameter[i].ParameterType == "lng" {
						//args[i] = sql.Named(":"+strconv.Itoa(counter), GetIntValueFromParameterName(Values, parameter[i].ParameterName))
						if !parameter[i].IsOutput {
							args[i] = GetIntValueFromParameterName(Values, parameter[i].ParameterName)
						} else {
							var paramlong int64
							args[i] = &paramlong
						}
					} else if parameter[i].ParameterType == "str" {
						//args[i] = sql.Named(":"+strconv.Itoa(counter), GetStringValueFromParameterName(Values, parameter[i].ParameterName))
						if !parameter[i].IsOutput {
							args[i] = GetStringValueFromParameterName(Values, parameter[i].ParameterName)
						} else {
							var paramstr string
							args[i] = &paramstr
						}
					}
				}

			} else if function.QueryType == 3 { //plsql
				args = make([]interface{}, len(parameter)+1)
				var counter int = -1
				for i := range parameter {
					counter++
					if parameter[i].ParameterType == "int" {
						//args[i] = sql.Named(":"+strconv.Itoa(counter), GetIntValueFromParameterName(Values, parameter[i].ParameterName))
						args[i] = GetIntValueFromParameterName(Values, parameter[i].ParameterName)
					} else if parameter[i].ParameterType == "lng" {
						//args[i] = sql.Named(":"+strconv.Itoa(counter), GetIntValueFromParameterName(Values, parameter[i].ParameterName))
						args[i] = GetIntValueFromParameterName(Values, parameter[i].ParameterName)
					} else if parameter[i].ParameterType == "str" {
						//args[i] = sql.Named(":"+strconv.Itoa(counter), GetStringValueFromParameterName(Values, parameter[i].ParameterName))
						args[i] = GetStringValueFromParameterName(Values, parameter[i].ParameterName)
					}
				}
				var rset1 driver.Rows
				args[len(parameter)] = sql.Out{Dest: &rset1}
				_, errc := stm.Exec(args...)
				if errc != nil {
					log.Fatal(errc, stm)
					return 0, nil, nil, errc
				} else {
					return 2, nil, rset1, nil
				}

			}
		}
		rt, err := stm.Query(args...)
		if err == nil {
			defer stm.Close()
			return 1, rt, nil, nil
		} else {
			return 0, nil, nil, err

		}
	} else { // just only for call rest service
		return -1, nil, nil, nil
	}
}
func ServiceExecuteDataTable1(db sql.DB) (*sql.Rows, error) {
	row, err := db.Query("select * from Business_Desk.dbo.tblinfo where infoID = @ID  and @VAl = 1 ", sql.Named("ID", 1), sql.Named("VAL", 1))
	return row, err
}
func GetStringValueFromParameterName(Values []QueryParsing.Dictionery, Name string) string {
	for a := 0; a < len(Values); a++ {
		if Values[a].ParameterName == Name {
			return Values[a].ParameterValue
		}
	}
	return ""
}
func GetIntValueFromParameterName(Values []QueryParsing.Dictionery, Name string) int {
	for a := 0; a < len(Values); a++ {
		if Values[a].ParameterName == Name {
			v, err := strconv.Atoi(Values[a].ParameterValue)
			if err == nil {
				return v
			}
			return 0
		}
	}
	return 0
}
