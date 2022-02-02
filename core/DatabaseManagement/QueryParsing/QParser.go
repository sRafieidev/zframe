package QueryParsing

import (
	"strings"
)

type Dictionery struct {
	ParameterName  string
	ParameterValue string
}

type QueryParameter struct {
	ParameterName     string
	ParameterFullName string
	ParameterType     string
	IsSession         bool
	IndexItem         int
	ParameterValue    string
	IsOutput          bool
	OutputValue       string
	InQueryStartIndex int
	InQueryEndIndex   int
	ToeknName         string
}

func StringFxIndexAt(s, sep string, n int) int {
	idx := strings.Index(s[n:], sep)
	if idx > -1 {
		idx += n
	}
	return idx
}
func GetQueryParameter(Query string, QueryType int) []QueryParameter {

	var rtvalue []QueryParameter = nil
	var lenc = len(Query)
	var d = 0
	var paramcounter = 0
	if len(strings.TrimSpace(Query)) > 5 {
		for a := 0; a < lenc; a++ {
			nowpostion := StringFxIndexAt(Query, "@", a)
			if nowpostion > 0 {
				paramcounter++
				a = nowpostion
				var paramname = ""
				var nowchar = ""
				for c := nowpostion; c < lenc; c++ {
					d = c + 1
					nowchar = Query[c:d]
					if nowchar == " " || nowchar == "\r\n" || nowchar == "\n" || nowchar == "\r" || nowchar == ")" || nowchar == "(" || ((d) == lenc) {
						if (d) == lenc {
							paramname += nowchar
						}
						c := QueryParameter{}
						c.ParameterFullName = paramname
						c.ParameterType, c.ParameterName, c.IsOutput, _ = GetParameterInformation(paramname, QueryType)
						c.InQueryStartIndex = nowpostion
						c.InQueryEndIndex = d
						c.IndexItem = paramcounter
						if strings.Contains(paramname, "sessionvalue") {
							c.IsSession = true
							c.ToeknName = "aa"
						} else {
							c.IsSession = false
							c.ToeknName = ""
						}
						rtvalue = append(rtvalue, c)
						paramname = ""
						break
					} else {
						paramname += nowchar
					}
				}
			} else {
				break
			}

		}

	}
	return rtvalue
}
func GetParameterInformation(ParamName string, Qtype int) (paramtype string, paramname string, isout bool, isin bool) {
	var var_isout = false
	var var_isin = false
	var var_paramname = ""
	var var_param_type = ""
	var nowparam = ParamName[1:]
	if Qtype == 2 { // just only in this type you need to isout or isin in default mod is in its be true
		var oibtype = nowparam[1:2]
		var_param_type = strings.ToLower(nowparam[2:5])
		var_paramname = nowparam[5 : len(nowparam)-1]
		if oibtype == "o" {
			var_isout = true
		} else if oibtype == "i" {
			var_isin = true
		} else {
			var_isout = true
			var_isin = true
		}
	} else {
		var_isin = true
		var_isout = false
		var_param_type = strings.ToLower(nowparam[0:3])
		var_paramname = nowparam[3:]
	}
	return var_param_type, var_paramname, var_isout, var_isin
}
func GetQueryType(Query string) int /*guery type like as PPSQL PLSQL SQL*/ {

	var trimquery = strings.ToLower(strings.TrimSpace(Query))
	if len(trimquery) < 5 {
		return 0
	}
	var noc = trimquery[0:5]
	if noc == "ppsql" {
		return 2
	} else if noc == "plsql" {
		return 3
	} else if strings.ToLower(trimquery[0:4]) == "rest" {
		return 4
	} else {
		return 1
	}

}
func ChangeQueryWithTokenToParameter(Query string, Qtype int) string {

	var lenc = len(Query)
	var TokenList []string = nil
	for a := 0; a < lenc; a++ {
		nowpostion := StringFxIndexAt(Query, "#", a)
		if nowpostion > 0 {
			var nowchar = ""
			var tokenname = ""
			var d = 0
			for c := nowpostion; c < lenc; c++ {
				d = c + 1
				nowchar = Query[c:d]
				if nowchar == " " || nowchar == "\r\n" || nowchar == "\n" || nowchar == "\r" || nowchar == ")" || nowchar == "(" || ((d) == lenc) {
					if (d) == lenc {
						tokenname += nowchar
					}
					TokenList = append(TokenList, tokenname)

				} else {
					tokenname += nowchar
				}

			}
		}
	}

	if TokenList != nil && len(TokenList) > 0 {

		for a := 0; a < len(TokenList); a++ {
			nowToken := TokenList[a]
			var ParamName = "@lngsessionvalue" + nowToken[1:]

			if Qtype == 2 {
				ParamName = "@ilngsessionvalue" + nowToken[1:]
			}

			Query = strings.ReplaceAll(Query, nowToken, ParamName)
		}
	}

	return Query
}
