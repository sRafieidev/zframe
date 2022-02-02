package DTOTable

type Sys_Connection_Log struct {
	SYS_CONNECTION_LOG_ID int64
	SYS_USER_ID           int32
	USERNAME              string
	USERPASS              string
	CONNECTION_DATE       string
	CONNECTION_TIME       string
	IP_ADDRESS            string
	MAC_ADDRESS           string
	EXITDATE              string
	EXITTIME              string
}
