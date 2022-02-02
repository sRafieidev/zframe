package DTOTable

type Sys_Concurrency struct {
	SYS_OBJECT_ID int32
	RECORDID      int64
	SYS_USER_ID   int32
	QUERYMODEID   int32
	LOCK          bool
	LOCKDATETIME  string
	SESSIONID     string
}
