package DTOTable

type Sys_Process_Transition_L struct {
	SYS_PROCESS_TRANSITION_L_ID int64
	SYS_PROCESS_TRANSITION_ID   int64
	DATE                        string
	TIME                        string
	FROM_SYS_OBJECT_ID          int32
	FORM_SYS_OBJECT_KEYID       int64
	TO_SYS_OBJECT_ID            int32
	TO_SYS_OBJECT_KEYID         int64
	SYS_APPLICATION_LOG_ID      int64
	ACTIVE                      bool
	TRACECODE                   int64
	CHANGECOUNT                 int32
}
