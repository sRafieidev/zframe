package DTOTable

type Sys_User struct {
	SYS_USER_ID   int32
	USERNAME      string
	PASSWORD      string
	CREATEDATE    string
	ENTITY_ID     int64
	SYS_OBJECT_ID int64
	ACTIVE        bool
}
