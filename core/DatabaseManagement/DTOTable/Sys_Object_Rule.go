package DTOTable

type Sys_Object_Rule struct {
	SYS_OBJECT_RULE_ID    int32
	SYS_OBJECT_ID         int32
	RULE_NAME             string
	COLUMNS_NAME          string
	OPERATOR              string
	VALUE                 string
	SYS_OBJECT_CONTROL_ID int32
	PROPERTY_NAME         string
	SETVALUE              string
	SORT_ORDER            int32
	ACTIVE                bool
	PARENT_ID             int32
}
