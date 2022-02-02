package ServiceManager

import (
	"ZGOPROJ/compile/compilerule"
	"ZGOPROJ/core/DatabaseManagement/QueryParsing"
	"ZGOPROJ/core/ZFrameSystem"
	"ZGOPROJ/core/datastructure/ZGJSON"
	"ZGOPROJ/core/security"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ExecuteService(function SysRestFunction, w http.ResponseWriter, r *http.Request) {

	var UserGroupListRequest string = ""
	c := function.FunctionQuery
	ua := r.Header.Get("CCTOKEN")
	var DictioneryList []QueryParsing.Dictionery = nil
	if len(ua) > 0 {
		data, err := base64.StdEncoding.DecodeString(ua)
		if err != nil {
			// token is false
			return
		}
		Val, errenc := security.DecryptAes128Ecb(data, compilerule.GetToken())
		if errenc == nil {
			var TokenData map[string]interface{}
			errunmarshal := json.Unmarshal(Val, &TokenData)
			if errunmarshal == nil {
				var Version = "0"
				var Enc = ""
				var MID = ""
				var DID = ""
				var CDATE = ""
				var MobileNumber = ""
				if x, found := TokenData["Version"]; found {
					Version = x.(string)
					DictioneryList = append(DictioneryList, QueryParsing.Dictionery{"sessionvalueJWTVERSION", Version})
				}
				if x, found := TokenData["Enc"]; found {
					Enc = x.(string)
					DictioneryList = append(DictioneryList, QueryParsing.Dictionery{"sessionvalueKEYENC", Enc})
				}
				if x, found := TokenData["MID"]; found {
					MID = x.(string)
					DictioneryList = append(DictioneryList, QueryParsing.Dictionery{"sessionvalueMEMBER_ID", MID})
				}
				if x, found := TokenData["DID"]; found {
					DID = x.(string)
					DictioneryList = append(DictioneryList, QueryParsing.Dictionery{"sessionvalueDEVICE_ID", DID})
				}
				if x, found := TokenData["CDATE"]; found {
					CDATE = x.(string)
					DictioneryList = append(DictioneryList, QueryParsing.Dictionery{"sessionvalueTOKENDATE", CDATE})
				}
				if x, found := TokenData["MO"]; found {
					MobileNumber = x.(string)
					DictioneryList = append(DictioneryList, QueryParsing.Dictionery{"sessionvalueMOBILENO", MobileNumber})
				}
				if Version == "3" {
					DictioneryList = append(DictioneryList, QueryParsing.Dictionery{"sessionvalueUSER_ID", MID})
					DictioneryList = append(DictioneryList, QueryParsing.Dictionery{"sessionvalueUSID", MID})
					DictioneryList = append(DictioneryList, QueryParsing.Dictionery{"sessionvalueENID", MID})
					DictioneryList = append(DictioneryList, QueryParsing.Dictionery{"sessionvalueCMID", MID})
					DictioneryList = append(DictioneryList, QueryParsing.Dictionery{"sessionvalueUGL", ".2."})
					UserGroupListRequest = ".2."
				}
			}
		}
	}

	var securitycheck bool = true
	// check service security
	if function.Authentication {
		if function.UserGroupAccess != nil {
			// read data from session and Token For check user have acces
			securitycheck = false
			GroupListExistInToken := strings.Split(UserGroupListRequest, ".")
			for qc := 0; qc < len(GroupListExistInToken); qc++ {
				for ac := 0; ac < len(function.UserGroupAccess); ac++ {
					if function.UserGroupAccess[ac] == GroupListExistInToken[qc] {
						securitycheck = true
						break
					}
				}
			}
		}
	}

	if !securitycheck {

		errorHandler(w, r, 403)
		return
	}

	if function.CallMethod == 1 {

		// get method
		values := r.URL.Query()
		var paramlist []QueryParsing.Dictionery = nil
		for k, v := range values {
			//fmt.Println(k, " => ", v)
			nowdic := QueryParsing.Dictionery{}
			nowdic.ParameterName = k
			nowdic.ParameterValue = v[0]
			paramlist = append(paramlist, nowdic)
		}

		////////////////////Start Excute Function In Get

		sysdb, dbtype := GetDataBaseWithSystemID(function.SysSystemId)

		if DictioneryList != nil && len(DictioneryList) > 0 {
			for q := 0; q < len(DictioneryList); q++ {
				paramlist = append(paramlist, DictioneryList[q])
			}
		}

		outtype, rt, driverrows, err := DoServiceAction(sysdb, function.FunctionQueryForExecute, function.ParameterList, paramlist, dbtype, &function)
		if outtype == -1 {
			c = "No Function Impliment for this Part"
			w.Write([]byte(c))
			return
		}
		if err != nil {
			println(err.Error())
			errorHandler(w, r, 500)
		} else {
			if outtype == 1 {
				w.Header().Set("Content-Type", "application/json;charset=UTF-8")
				jsondata := ZGJSON.JsonFromDbRows(rt)
				for i := 0; i < len(jsondata); i++ {
					w.Write([]byte(jsondata[i]))
				}
				defer rt.Close()
				return
			} else if outtype == 2 {
				w.Header().Set("Content-Type", "application/json;charset=UTF-8")
				jsondata := ZGJSON.JsonFromDriverRows(driverrows)
				for i := 0; i < len(jsondata); i++ {
					w.Write([]byte(jsondata[i]))
				}
				defer driverrows.Close()
				return
			}
		}

		// End Function Excute for get
		// get method
	} else if function.CallMethod == 2 {

		// post method
		bodydata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		} else {

			var bodyjsondata map[string]interface{}
			err := json.Unmarshal([]byte(bodydata), &bodyjsondata)
			if err != nil {
				log.Printf("Error reading body No Json Data: %v", err)
				http.Error(w, "can't read body No Json Data", http.StatusBadRequest)
			}
			//color.Blue(result["users"])
			// fmt.Println(bodyjsondata["users"])  // Test Code
			var paramlen = len(function.ParameterList)
			var paramlist []QueryParsing.Dictionery = nil
			var existerror = false
			for a := 0; a < paramlen; a++ {
				if !function.ParameterList[a].IsSession && !function.ParameterList[a].IsOutput {
					var value string
					var ok bool
					if x, found := bodyjsondata[function.ParameterList[a].ParameterName]; found {
						if value, ok = x.(string); ok {
							nowdic := QueryParsing.Dictionery{}
							nowdic.ParameterName = function.ParameterList[a].ParameterName
							nowdic.ParameterValue = value
							paramlist = append(paramlist, nowdic)
						} else {
							fmt.Println("Error C31 Cant Find Parameter Value [" + function.ParameterList[a].ParameterName + "]")
							existerror = true
						}
					} else {
						fmt.Println("Error C32 Cant Find Parameter Value [" + function.ParameterList[a].ParameterName + "]")
						existerror = true
					}
				}
			}

			if !existerror {
				// start going to execute query
				sysdb, dbtype := GetDataBaseWithSystemID(function.SysSystemId)
				if DictioneryList != nil && len(DictioneryList) > 0 {
					for q := 0; q < len(DictioneryList); q++ {
						paramlist = append(paramlist, DictioneryList[q])
					}
				}
				outtype, rt, driverrows, err := DoServiceAction(sysdb, function.FunctionQueryForExecute, function.ParameterList, paramlist, dbtype, &function)
				if outtype == -1 {
					c = "No Function Impliment for this Part"
					w.Write([]byte(c))
					return
				}
				if err != nil {
					println(err.Error())
					errorHandler(w, r, 500)
				} else {
					if outtype == 1 {
						w.Header().Set("Content-Type", "application/json;charset=UTF-8")
						jsondata := ZGJSON.JsonFromDbRows(rt)
						for i := 0; i < len(jsondata); i++ {
							w.Write([]byte(jsondata[i]))
						}
						defer rt.Close()
						return
					} else if outtype == 2 {
						w.Header().Set("Content-Type", "application/json;charset=UTF-8")
						jsondata := ZGJSON.JsonFromDriverRows(driverrows)
						for i := 0; i < len(jsondata); i++ {
							w.Write([]byte(jsondata[i]))
						}
						defer driverrows.Close()
						return
					}
				}

			} else {
				c = "Error in Process " + c
				w.Write([]byte(c))
			}
		}
	} else {
		c = "call Method handler not found (More then 2) " + c
		w.Write([]byte(c))
	}
}
func GetDataBaseWithSystemID(systemid int) (*sql.DB, int) {
	for a := 0; a < len(ZFrameSystem.SysList); a++ {
		if ZFrameSystem.SysList[a].SYS_SYSTEM_ID == systemid {
			return ZFrameSystem.SysList[a].SYSTEMDB, ZFrameSystem.SysList[a].SYS_DATABASE_TYPE_ID
		}
	}
	return nil, 0
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		if status == 500 {
			fmt.Fprint(w, " Internal Server Error  ")
		} else if status == 403 {
			fmt.Fprint(w, "  Forbidden  ")
		}
	}
}
