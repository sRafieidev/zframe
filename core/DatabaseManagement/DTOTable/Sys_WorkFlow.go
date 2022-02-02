package DTOTable

type Sys_WorkFlow struct {
	SYS_WORKFLOW_ID int32
	SYS_SYSTEM_ID   int32
	NAME            string
	NAVIGATEADDRESS string
	FORMNAME        string
	ACTIVE          bool
	PARENT_ID       int32
	INDEXNUMBER     int32
	KEYNUMBER       string
}
