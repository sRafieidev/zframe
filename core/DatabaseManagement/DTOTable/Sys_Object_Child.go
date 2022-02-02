package DTOTable

type Sys_Object_Child struct {
	SYS_OBJECT_CHILD_ID  int32
	PARENT_SYS_OBJECT_ID int32
	CHILD_SYS_OBJECT_ID  int32
	NAME                 string
	CHILD_DATA_SOURSE    string
	CAPTION              string
	ACTIVE               bool
	KEYNAME              string
}
