package db

import (
	"errors"
	"log"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	Connection *gorm.DB
}

// Singleton pattern
var (
	instance *Database
)

func ConnectToDatabase() error {
	options := map[string]string{
		"CONNECTION TIMEOUT": "90",
		"LANGUAGE":           "FRENCH",
		"TERRITORY":          "CANADA",
		"SSL":                "false",
	}

	// could be environment variables or repo secrets
	serverHost := ""
	serverPort := 1
	serviceName := ""
	user := ""
	password := ""

	// oracle://user:password@host:port/service
	url := oracle.BuildUrl(serverHost, serverPort, serviceName, user, password, options)
	dialector := oracle.New(oracle.Config{
		DSN:                     url,
		IgnoreCase:              false, // query conditions are not case-sensitive
		NamingCaseSensitive:     true,  // whether naming is case-sensitive
		VarcharSizeIsCharLength: true,  // whether VARCHAR type size is character length, defaulting to byte length

		// RowNumberAliasForOracle11 is the alias for ROW_NUMBER() in Oracle 11g, defaulting to ROW_NUM
		RowNumberAliasForOracle11: "ROW_NUM",
	})

	db, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true,
			IdentifierMaxLength: 30,
		},
		PrepareStmt:     false,
		CreateBatchSize: 50,
	})

	if err != nil {
		return err
	}

	// set session parameters
	if sqlDB, err := db.DB(); err == nil {
		_, _ = oracle.AddSessionParams(sqlDB, map[string]string{
			"TIME_ZONE":               "+08:00",                       // ALTER SESSION SET TIME_ZONE = '+08:00';
			"NLS_DATE_FORMAT":         "YYYY-MM-DD",                   // ALTER SESSION SET NLS_DATE_FORMAT = 'YYYY-MM-DD';
			"NLS_TIME_FORMAT":         "HH24:MI:SSXFF",                // ALTER SESSION SET NLS_TIME_FORMAT = 'HH24:MI:SS.FF3';
			"NLS_TIMESTAMP_FORMAT":    "YYYY-MM-DD HH24:MI:SSXFF",     // ALTER SESSION SET NLS_TIMESTAMP_FORMAT = 'YYYY-MM-DD HH24:MI:SS.FF3';
			"NLS_TIME_TZ_FORMAT":      "HH24:MI:SS.FF TZR",            // ALTER SESSION SET NLS_TIME_TZ_FORMAT = 'HH24:MI:SS.FF3 TZR';
			"NLS_TIMESTAMP_TZ_FORMAT": "YYYY-MM-DD HH24:MI:SSXFF TZR", // ALTER SESSION SET NLS_TIMESTAMP_TZ_FORMAT = 'YYYY-MM-DD HH24:MI:SS.FF3 TZR';
		})
	}

	log.Println("Connected to database, GetInstance() to get the database instance.")
	instance = &Database{
		Connection: db,
	}

	return nil
}

// instance.Close() to close the database connection
func (db *Database) CloseConnection() {
	if sqlDB, err := db.Connection.DB(); err == nil {
		_ = sqlDB.Close()
	}
}

func GetInstance() (*Database, error) {
	if instance == nil {
		return nil, errors.New("database instance is not initialized")
	}
	return instance, nil
}
