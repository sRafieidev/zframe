package DTOTable

type Sys_Channel_Element struct {
	SYS_CHANNEL_ELEMENT_ID int32
	SYS_CHANNEL_ID         int32
	SYS_ELEMENT_TYPE_ID    int32
	SYS_PROTOCOL_ID        int32
	ELEMENT_NAME           string
	ACTIVE                 string
	CONFIG                 string
	PACKETSEPARATECHAR     string
	PACKETHEADER           string
	PACKETFOOTER           string
	PACKETDATALEN          int32
	ELEMENTSEPARATECHAR    string
	AUTHENTICATION         string
}
