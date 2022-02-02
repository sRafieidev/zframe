package DTOTable

type Sys_Accessibility_Level1 struct {
	SYS_ACCESSIBILITY_LEVEL1_ID int32
	SYS_OBJECT_ID               int32
	SYS_USERGROUP_ID            int32
	CANVIEW                     bool
	CANEDIT                     bool
	CANDELETE                   bool
	CANINSERT                   bool
	CANREPORT                   bool
	CANEXPORT                   bool
}
