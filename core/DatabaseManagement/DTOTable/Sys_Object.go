package DTOTable

type Sys_Object struct {
	SYS_OBJECT_ID             int32
	SYS_SYSTEM_ID             int32
	SYS_OBJECT_NAME           string
	SYS_OBJECT_CAPTION        string
	RECORD_SOURCE             string
	GRID_SOURCE               string
	SYS_DESIGN_PATTERN_ID     int32
	VALIDATOR                 bool
	ALTERNATIVE_SYS_OBJECT_ID int32
	DATABASEOBJECTNAME        string
	RULEID                    int32
	REPORTRULEID              int32
	SHOWMODE                  int32
	AUTHENTICATION            bool
	UPDATEMODE                int32
	LTR                       int32
	GRIDSELECTTEXT            string
	CONCURRENCYMODE           int32
	CACHEKEY                  string
	GRIDSESSIONNAME           string
	CSSCLASS                  string
	HEADER_TAG                string
	GRID_LOAD_MODE            int32
	GRID_SELECT_COUNT         string
	OBJECT_GUID               string
	SYS_PACKAGE_ID            int32
}
