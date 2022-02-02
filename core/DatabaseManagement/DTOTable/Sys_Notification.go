package DTOTable

type Sys_Notification struct {
	SYS_NOTIFICATION_ID int32
	SYS_SYSTEM_ID       int32
	SYS_OBJECT_ID       int32
	NAME                string
	LOADQUERY           string
	CAPTION             string
	ACTIVE              int32
	INTERVAL            int32
	URLADDRESS          string
	IMAGEADDRESS        string
}
