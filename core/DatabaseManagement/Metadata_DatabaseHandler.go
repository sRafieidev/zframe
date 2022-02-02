package DatabaseManagement

import (
	"ZGOPROJ/core/ZLog"
	"ZGOPROJ/zfconfig"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/fatih/color"
	_ "github.com/godror/godror"
	"log"
	"strconv"
)

var MetaDataDb *sql.DB
var err error

func InitDatabaseMetadata() {

	dbType, err := strconv.Atoi(zfconfig.DataBaseTypeID)
	//1	MSSQL Server
	//2	Oracle
	//3	IBM DB2
	//4	MY SQL
	//5	PostgreSQL

	if err == nil {
		switch dbType {
		case 1:
			color.Blue(" Metadata Database : MSSQL Server ")
			LoadMSSQLServerDatabase()
		case 2:
			color.Blue(" Metadata Database : Oracle")
			LoadOracleDatabase()
		case 3:
			color.Blue(" Metadata Database : IBM DB 2")
			LoadDatabaseIBMDB2()
		case 4:
			color.Blue(" Metadata Database : MySql")
			LoadMYSQLDatabase()
		case 5:
			color.Blue(" Metadata Database : PostgreSql")
			LoadPostgreSQLDatabase()
		}

	}

}

func LoadMSSQLServerDatabase() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", zfconfig.ProjectDataBaseServer, zfconfig.ProjectDataBaseUserName, zfconfig.ProjectDataPassPassword, zfconfig.ProjectDataBasePort, zfconfig.ProjectDatabaseName)
	MetaDataDb, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	} else {
		MetaDataDb.SetMaxIdleConns(zfconfig.MetaDataConnectionPoolCount / 2)
		MetaDataDb.SetMaxOpenConns(zfconfig.MetaDataConnectionPoolCount)
	}

	log.Printf("Connected!\n")
	// Close the database connection  mmpool after program executes
	//defer MetaDataDb.Close()
}

func LoadOracleDatabase() {

	connString := fmt.Sprintf("user=%s password=%s connectString=%s", zfconfig.ProjectDataBaseUserName, zfconfig.ProjectDataPassPassword, zfconfig.ProjectConnectionstring)
	MetaDataDb, err = sql.Open("godror", connString)

	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	} else {
		MetaDataDb.SetMaxIdleConns(zfconfig.MetaDataConnectionPoolCount)
		MetaDataDb.SetMaxOpenConns(zfconfig.MetaDataConnectionPoolCount)
	}
	log.Printf("Connected!\n")
	// Close the database connection  mmpool after program executes
	//defer MetaDataDb.Close()

}

func LoadDatabaseIBMDB2() {
	println("exception not implemented  IBM DB2 ")
}

func LoadMYSQLDatabase() {
	println("exception not implemented mysql ")
}

func LoadPostgreSQLDatabase() {
	println("exception not implemented PostgreSql ")
}

func ExecuteDataTable(db *sql.DB, Query string) *sql.Rows {

	//ZLog.ZLOG(100, " Start Excute Query  ", Query)
	if db == nil {
		ZLog.ZLOGDISPLAY(100, " Connection Is NULL ", Query, 4)
		return nil
	}
	rows, errc := db.Query(Query)
	if errc != nil {
		fmt.Println(errc.Error())
		log.Fatal("Scan failed:", errc.Error())
	}
	return rows
}

func ShowRowsData(rows *sql.Rows) {
	cols, cols_err := rows.Columns()

	if cols_err != nil {
		log.Fatalln(cols_err)
	}
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))
	dest := make([]interface{}, len(cols))
	for i := range cols {
		dest[i] = &rawResult[i]
	}
	for rows.Next() {
		rows.Scan(dest...)

		for i, raw := range rawResult {
			result[i] = string(raw)
		}

		for j, v := range result {
			fmt.Printf("%s", v)
			if j != len(result)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("\n")
	}
}
func GetDataTable(rows *sql.Rows) [][]string {
	var rt [][]string
	counter := 0
	cols, cols_err := rows.Columns()

	if cols_err != nil {
		log.Fatalln(cols_err)
	}

	rawResult := make([][]byte, len(cols))

	dest := make([]interface{}, len(cols))
	for i := range cols {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		rows.Scan(dest...)
		result := make([]string, len(cols))
		for i, raw := range rawResult {
			result[i] = string(raw)
		}
		rt = append(rt, result)
		counter++
	}
	return rt
}
