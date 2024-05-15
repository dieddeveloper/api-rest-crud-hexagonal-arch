package pkg

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DBConnDto struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewConn(connData *DBConnDto) (*gorm.DB, error) {
	portInt, err := strconv.Atoi(connData.Port)
	if err != nil {
		panic(err)
	}
	var connect string
	var dialector gorm.Dialector
	if portInt == 5432 {
		connect = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			connData.Host, portInt, connData.Username, connData.Password, connData.DBName,
		)
		dialector = gormpostgres.Open(connect)
	}
	db, err := connectDB(dialector)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Errorln("NewDBClient", err)
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		logrus.Errorln("NewDBClient", err)
		return nil, err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	maxIdleConns, _ := strconv.Atoi(os.Getenv("GORM_MAX_IDLE_CONNS"))
	sqlDB.SetMaxIdleConns(maxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	maxOpenConns, _ := strconv.Atoi(os.Getenv("GORM_MAX_OPEN_CONNS"))
	sqlDB.SetMaxOpenConns(maxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	maxLifeTime, _ := strconv.Atoi(os.Getenv("GORM_MAX_LIFE_TIME"))
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(maxLifeTime))

	return db, nil
}

func connectDB(dialector gorm.Dialector) (*gorm.DB, error) {
	location, _ := time.LoadLocation(os.Getenv("TIME_ZONE"))
	db, err := gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Now().In(location)
		},
	})
	if err != nil {
		logrus.Errorln("connectDB", err)
		return nil, err
	}

	return db, nil
}
