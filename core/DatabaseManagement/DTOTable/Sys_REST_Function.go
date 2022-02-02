package DTOTable

type Sys_REST_Function struct {
	SYS_REST_FUNCTION_ID int32
	SYS_SYSTEM_ID        int32
	FUNCTIONNAME         string
	BASEURL              string
	AUTHENTICATION       bool
	FUNCTION_RULEID      int32
	FUNCTION_QUERY       string
	CALL_METHOD          int32
	FUNCTION_DESCRIPTION string
	OTTACTION            bool
	DEFAULT_OUTPUT       string
}
