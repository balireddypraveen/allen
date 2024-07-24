package postgres

import (
	"context"
	"fmt"
	"github.com/balireddypraveen/allen/internal/common/constants"
	"github.com/balireddypraveen/allen/internal/pkg/common_utils"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	config "github.com/balireddypraveen/allen/configs"
	"github.com/balireddypraveen/allen/internal/common/configs"
	"github.com/balireddypraveen/allen/internal/pkg/newrelic_setup"
)

func init() {
	config.SetupConfig()
}

var (
	singleton  sync.Once
	postgresDB = &gorm.DB{}
)

func InitDB() *gorm.DB {
	var err error
	singleton.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			viper.GetString(configs.VKEYS_DATABASE_POSTGRES_SOURCE_HOST),
			viper.GetString(configs.VKEYS_DATABASE_POSTGRES_SOURCE_USER),
			viper.GetString(configs.VKEYS_DATABASE_POSTGRES_SOURCE_PASSWORD),
			viper.GetString(configs.VKEYS_DATABASE_POSTGRES_SOURCE_DB_NAME),
			viper.GetInt(configs.VKEYS_DATABASE_POSTGRES_SOURCE_PORT),
		)

		var migrationLogger logger.Interface
		if !common_utils.Contains(constants.ProdHosts, viper.GetString(configs.VKEYS_HOST_TYPE)) {
			migrationLogger = logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold:             time.Second, // Slow SQL threshold
					LogLevel:                  logger.Info, // Log level
					IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
					Colorful:                  false,       // Disable color
				},
			)
		} else {
			migrationLogger = nil
		}

		postgresDB, err = gorm.Open(
			postgres.New(
				postgres.Config{
					DriverName: "nrpgx",
					DSN:        dsn,
				},
			), &gorm.Config{
				Logger: migrationLogger,
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
			},
		)
		if err != nil {
			panic(fmt.Sprintf("not able to connect to the database. Error:- %s", err.Error()))
		}

		db, err := postgresDB.DB()
		if err != nil {
			panic(fmt.Sprintf("error occurred while getting db instance object. Error:- %s", err.Error()))
		}

		db.SetMaxIdleConns(viper.GetInt(configs.VKEYS_DATABASE_POSTGRES_SOURCE_MAX_IDLE_CONN))
		db.SetMaxOpenConns(viper.GetInt(configs.VKEYS_DATABASE_POSTGRES_SOURCE_MAX_OPEN_CONN))
		db.SetConnMaxLifetime(time.Duration(viper.GetInt(configs.VKEYS_DATABASE_POSTGRES_SOURCE_MAX_CONN_LIFETIME)) * time.Second)
	})
	return postgresDB
}

// GetDB gets returns the *gorm.DB,
// rCtx is required when you want to instrument the db call and attach the timing to the NewRelicTxn stored in it.
// it is nil safe.
func GetDB(ctx context.Context) *gorm.DB {
	if newrelic_setup.IsNewrelicEnable() {
		txn := newrelic.FromContext(ctx)
		if txn != nil {
			return postgresDB.WithContext(newrelic.NewContext(context.Background(), txn))
		}
	}
	return postgresDB
}

func GetDBWithoutContext() *gorm.DB {
	return postgresDB
}
