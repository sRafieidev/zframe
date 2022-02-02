package DTOTable

import "database/sql"

type Sys_System struct {
	SYS_SYSTEM_ID            int
	NAME                     string
	TITLE                    string
	SYS_DATABASE_TYPE_ID     int
	DATABASECONNECTIONSTRING string
	DATABASEUSERNAME         string
	DATABASEPASSWORD         string
	DATABASESERVERADDRESS    string
	DATABASESCHEMANAME       string
	DATABASEPORT             int32
	APPLICATIONURL           string
	JDBCCNN                  string
	LTR                      int32
	SORTORDER                int32
	IS_ACTIVE                int32
	CONNECTION_POOL_COUNT    int
	SYSTEMDB                 *sql.DB
}
