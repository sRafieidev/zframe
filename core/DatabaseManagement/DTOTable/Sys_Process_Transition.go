package DTOTable

type Sys_Process_Transition struct {
	SYS_PROCESS_TRANSITION_ID int64
	FROM_SYS_PROCESS_STATE_ID int32
	TO_SYS_PROCESS_STATE_ID   int32
	TRANSITIONNAME            string
	DESCRIPTION               string
	TRANSITION_CONDITION      string
	TRANSITIONQUERY           string
}
