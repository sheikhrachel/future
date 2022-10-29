package rds

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/util/errutil"
	psqlConf "github.com/sheikhrachel/future/conf/dev/api_common/aws/postgres"
)

type Postgres struct {
	Psql *sqlx.DB
}

func InitPostgres(cc call.Call) *Postgres {
	dbCredStr := fetchPostgresCreds(PsqlDev)
	db, err := sqlx.Connect("postgres", dbCredStr)
	if errutil.HandleError(err) {
		cc.InfoF(err.Error())
	}
	return &Postgres{
		Psql: db,
	}
}

type PsqlEnv string

const (
	PsqlDev PsqlEnv = "dev"
)

func fetchPostgresCreds(env PsqlEnv) string {
	var host, user, pass, name string
	var port int
	switch env {
	case PsqlDev:
		host = psqlConf.DevDbHost
		user = psqlConf.DevDbUser
		pass = psqlConf.DevDbPass
		name = psqlConf.DevDbName
		port = psqlConf.DevDbPort
	}
	connOpts := fmt.Sprintf("host=%+v user=%+v password=%+v dbname=%+v port=%+v sslmode=disable", host, user, pass, name, port)
	return connOpts
}
