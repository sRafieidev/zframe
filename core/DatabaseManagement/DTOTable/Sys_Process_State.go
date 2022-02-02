package DTOTable

type Sys_Process_State struct {
	SYS_PROCESS_STATE_ID       int32
	SYS_PROCESS_ID             int32
	SYS_OBJECT_ID              int32
	STATENAME                  string
	STATEINDEX                 int32
	DESCRIPTION                string
	LIMITDAY                   int32
	ACTIVE                     bool
	GRIDSOURCEQUERYINTHISSTATE string
}
