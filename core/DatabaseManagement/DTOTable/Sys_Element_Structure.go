package DTOTable

type Sys_Element_Structure struct {
	SYS_ELEMENT_STRUCTURE_ID int64
	SYS_CHANNEL_ELEMENT_ID   int32
	SEQUENCENUMBER           int32
	NAME                     string
	SYS_DATA_TYPE_ID         int16
	STARTINDEX               int32
	ENDINDEX                 int32
	ISOUTPUT                 string
}
