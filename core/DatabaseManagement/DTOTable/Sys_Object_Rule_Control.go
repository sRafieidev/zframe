package DTOTable

type Sys_Object_Rule_Control struct {
	SYS_OBJECT_RULE_CONTROL_ID int32
	SYS_OBJECT_ID              int32
	RNEW                       bool
	RSAVE                      bool
	RDELETE                    bool
	RSEARCH                    bool
	RPRINT                     bool
	RFIREST                    bool
	RBACK                      bool
	RNEXT                      bool
	REND                       bool
	REXCELGRID                 bool
	REXCELREPORT               bool
	RDESIGN                    bool
	RHELP                      bool
	RREFRESH                   bool
}
