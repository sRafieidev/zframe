package DTOTable

type Sys_Application_Log struct {
	SYS_APPLICATION_LOG_ID int64
	SYS_OBJECT_ID          int32
	DATEVALUE              string
	TIMEVALUE              string
	SYS_USER_ID            int32
	ACTIONID               int32
	RECORDID               int64
	SYS_CONNECTION_LOG_ID  int64
	DATA                   string
}
