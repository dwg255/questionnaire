package common

import (
	"database/sql"
)

var Source *Conf
type Conf struct {

	Db *sql.DB		//db

	EmailStmpHost string//email
	EmailStmpPort int
	EmailFrom string
	EmailPassword string
	EmailToers string
	EmailCcers string
	EmailLimit int
	EmailExpire int
}