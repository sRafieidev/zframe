package zfconfig

import (
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var ZFGOVersion = "\t     1.0.0.1.00001"
var Server_IS_Running bool
var ZFrame_Jar_Version string
var JAVASCRIPTVERSION string
var LicenceValidate bool
var LoadAddRuleClass bool
var App_Name string
var Project_Name string
var ApplicationVersionAndInformation string
var DomainAddressOnWebServer string
var WebApplicationAddressOnWeb string
var ApplicationBaseAddressOnServer string
var CheckUserAccessInSystem bool
var DeniedViewAndEditAcessToChildControl bool
var CheckControlRuleInFormForAccess bool
var ActiveDiagnostic string
var JarFileAddress string
var TraceLevel int
var ConnectionPool int
var ZFFormLoaderPoolCount int
var MaxQueryTimeOutSecond int
var PageLoadTimeOut int
var ServiceTimeOut int
var MetaDataConnectionPoolCount int
var LogConnectionPoolCount int
var LogDataBaseTypeID string
var LogProjectConnectionstring string
var LogProjectDataBaseUserName string
var LogProjectDataPassPassword string
var LogProjectDataBaseType int
var LogProjectDataDriver string
var LogDisable int
var MaxRecordForExcelCoutput int
var SetHoldabilityforConnection int
var IS_Multi_Language int
var MaxStatements int
var MaxStatementsPerConnection int
var MetaDataMaxStatements int
var MetaDataMaxStatementsPerConnection int
var LogMaxStatements int
var LogMaxStatementsPerConnection int
var ZFrameOTP int
var ZFrameOTPCheck int
var NumHelperThreads int
var ChildWindowWidth int
var IdleConnectionTestPeriod int
var ZFrameDisableFormGenerator int
var CSRF_TOKEN_NAME_SERVER string
var CSRF_TOKEN_NAME_CLENT string
var External_Form_Content string
var REST_EXECUTE_SUCCESS_OUTPUT string
var REST_EXECUTE_SUCCESS_OUTPUT_INSERT_UPDATE string
var MSSQLDeriver string
var OracleDeriver string
var IBMDB2Driver string
var MYSQLDRIVER string
var PostgreSQLDRIVER string
var DataBaseTypeID string
var ProjectConnectionstring string
var ProjectDataBaseUserName string
var ProjectDataPassPassword string
var ProjectDataBasePort = 1433
var ProjectDataBaseServer = "localhost"
var ProjectDatabaseName = "ISC"
var ProjectDataBaseType int
var ProjectDataDriver string
var System_ID string
var SetupUserName string
var SetupPassword string
var Direction int
var SystemConnectionstring string
var SystemDataBaseUserName string
var SystemDataPassPassword string
var SystemDataDriver string
var ApplicationName string
var BasePathFileDataBaseContent string
var RestHeader string
var customformHeaderTag string
var Have_RestHeader bool
var ApplicationLunchTime string
var Enable_JWT int
var JWT_Session_TimeOut int
var ZFrame_Form_JavaScript_Cashing int
var DataIsLoaded bool
var ZFLOCK string
var __lock string
var Ismsqsqlserver = true
var ConfigisLoaded = false

type fileattrib struct {
	key   string
	value string
}

func LoadConfigure() error {

	color.Blue("Start Reading Confing File ")
	data, err := ioutil.ReadFile("zf.configdata")
	if err != nil {
		log.Fatal(err)
	} else {
		content := string(data)
		var s = strings.Split(content, "\r\n")
		if len(s) > 5 {
			ConfigisLoaded = true
		}
		for j := range s {
			var nowline = s[j]
			if len(nowline) > 1 {
				if nowline[0] == '#' {

				} else {

					var nowc = strings.Split(nowline, "=")
					nowkey := fileattrib{"", ""}
					nowkey.key = strings.TrimSpace(strings.ToLower(nowc[0]))
					nowkey.value = strings.TrimSpace(nowc[1])

					if strings.EqualFold(nowkey.key, "project") {
						Project_Name = nowkey.value
					}
					if strings.EqualFold(nowkey.key, "databasetypeid") {
						DataBaseTypeID = nowkey.value
					}

					if strings.EqualFold(nowkey.key, "logdatabasetypeid") {
						LogDataBaseTypeID = nowkey.value
					}

					if strings.EqualFold(nowkey.key, "dbusername") {
						ProjectDataBaseUserName = nowkey.value

					}

					if strings.EqualFold(nowkey.key, "weburl") {
						WebApplicationAddressOnWeb = nowkey.value
					}

					if strings.EqualFold(nowkey.key, "restheader") {
						RestHeader = strings.TrimSpace(nowkey.key)
						if len(RestHeader) > 0 {
							Have_RestHeader = true
						}
					}

					if strings.EqualFold(nowkey.key, "logdbusername") {
						LogProjectDataBaseUserName = nowkey.value

					}

					if strings.EqualFold(nowkey.key, "dbpassword") {
						ProjectDataPassPassword = nowkey.value

					}

					if strings.EqualFold(nowkey.key, "logdbpassword") {
						LogProjectDataPassPassword = nowkey.value

					}

					if strings.EqualFold(nowkey.key, "system_id") {
						System_ID = nowkey.value

					}

					if strings.EqualFold(nowkey.key, "projectconnectionstring") {
						ProjectConnectionstring = nowkey.value

					}

					if strings.EqualFold(nowkey.key, "filedatabasepath") {
						BasePathFileDataBaseContent = nowkey.value

					}

					if strings.EqualFold(nowkey.key, "logprojectconnectionstring") {
						LogProjectConnectionstring = nowkey.value

					}

					if strings.EqualFold(nowkey.key, "setupusername") {
						SetupUserName = nowkey.value

					}

					if strings.EqualFold(nowkey.key, "setuppassword") {
						SetupPassword = nowkey.value

					}

					if strings.EqualFold(nowkey.key, "debug") {
						ActiveDiagnostic = nowkey.value

					}

					if strings.EqualFold(nowkey.key, "jar") {
						JarFileAddress = nowkey.value

					}

					if strings.EqualFold(nowkey.key, "tracelevel") {
						TraceLevel, _ = strconv.Atoi(nowkey.value)

					}

					if strings.EqualFold(nowkey.key, "connectionpool") {
						ConnectionPool, _ = strconv.Atoi(nowkey.value)

					}

					NowVal := 1
					if strings.EqualFold(nowkey.key, "threadpoolcount") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							ZFFormLoaderPoolCount = NowVal

						}
					}

					if strings.EqualFold(nowkey.key, "cnnholdability") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							SetHoldabilityforConnection = NowVal

						}
					}

					if strings.EqualFold(nowkey.key, "idleconnectiontestperiod") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							IdleConnectionTestPeriod = NowVal

						}
					}

					if strings.EqualFold(nowkey.key, "maxquerytimeoutsecond") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							MaxQueryTimeOutSecond = NowVal

						}
					}

					if strings.EqualFold(nowkey.key, "javascriptcache") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							ZFrame_Form_JavaScript_Cashing = NowVal

						}
					}

					if strings.EqualFold(nowkey.key, "multilanguage") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							IS_Multi_Language = NowVal

						}
					}

					if strings.EqualFold(nowkey.key, "jwtsessiontimeout") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							JWT_Session_TimeOut = NowVal

						}
					}

					if strings.EqualFold(nowkey.key, "maxrecordforexcelcoutput") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							MaxRecordForExcelCoutput = NowVal

						}
					}

					if strings.EqualFold(nowkey.key, "childwinodwwidth") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							ChildWindowWidth = NowVal

						}
					}

					if strings.EqualFold(nowkey.key, "pageloadtimeout") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							PageLoadTimeOut = NowVal

						}
					}

					if strings.EqualFold(nowkey.key, "servicetimeout") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							ServiceTimeOut = NowVal

						}
					}

					if strings.EqualFold(nowkey.key, "metadataconnectionpoolcount") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							MetaDataConnectionPoolCount = NowVal

						}
					}

					if strings.EqualFold(nowkey.key, "logconnectionpoolcount") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							LogConnectionPoolCount = NowVal

						}
					}

					if strings.EqualFold(nowkey.key, "maxstatements") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							MaxStatements = NowVal
						}
					}

					if strings.EqualFold(nowkey.key, "maxstatementsperconnection") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							MaxStatementsPerConnection = NowVal
						}
					}

					if strings.EqualFold(nowkey.key, "metadatamaxstatements") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							MetaDataMaxStatements = NowVal
						}
					}

					if strings.EqualFold(nowkey.key, "metadatamaxstatementsperconnection") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							MetaDataMaxStatementsPerConnection = NowVal
						}
					}

					if strings.EqualFold(nowkey.key, "logmaxstatements") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							LogMaxStatements = NowVal
						}
					}

					if strings.EqualFold(nowkey.key, "logmaxstatementsperconnection") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							LogMaxStatementsPerConnection = NowVal
						}
					}

					if strings.EqualFold(nowkey.key, "numhelperthreads") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							NumHelperThreads = NowVal
						}
					}

					if strings.EqualFold(nowkey.key, "logdisable") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							LogDisable = NowVal
						}
					}

					if strings.EqualFold(nowkey.key, "zframeotp") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							ZFrameOTP = NowVal
						}
					}

					if strings.EqualFold(nowkey.key, "zframeotpcheck") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							ZFrameOTPCheck = NowVal
						}
					}

					if strings.EqualFold(nowkey.key, "enablejwt") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							Enable_JWT = NowVal
						}
					}

					if strings.EqualFold(nowkey.key, "uioutoff") {
						NowVal, _ = strconv.Atoi(nowkey.value)
						if NowVal > 0 {
							ZFrameDisableFormGenerator = NowVal
						}
					}

					if strings.EqualFold(nowkey.key, "restsuccessmessage") {
						REST_EXECUTE_SUCCESS_OUTPUT = nowkey.value
					}

					if strings.EqualFold(nowkey.key, "restsuccessmessageinsertupdate") {
						REST_EXECUTE_SUCCESS_OUTPUT_INSERT_UPDATE = nowkey.value
					}

				}
			}
		}
	}

	color.Blue("End Reading Confing File ")
	return err
}
