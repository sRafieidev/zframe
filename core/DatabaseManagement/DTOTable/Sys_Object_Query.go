package DTOTable

type Sys_Object_Query struct {
	SYS_OBJECT_QUERY_ID int32
	SYS_OBJECT_ID       int32
	SYS_CRUD_TYPE_ID    int32
	QUERY               string
	QUERYINDEX          int32
	EXECUTEMESSAGE      string
	EXECUTEERRORMESSAGE string
	BEFOREXECUTEMESSAGE string
	CACHEKEY            string
}
