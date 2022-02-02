package DTOTable

type Sys_Source_Code struct {
	SYS_SOURCE_CODE_ID  int32
	SYS_OBJECT_ID       int32
	CODE                string
	SOURCE_CODE_TYPE_ID int32
	RULE_ID             int32
	DESCRIPTION         string
	COMPILEERROR        string
	CREATEDATE          string
	ACTIVE              int32
	CLASSNAME           string
	PACKAGENAME         string
}
