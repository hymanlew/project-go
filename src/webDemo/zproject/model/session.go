package model

import "webDemo/dbsql"

type Session struct {
	SessionId string
	User      *dbsql.User
}
