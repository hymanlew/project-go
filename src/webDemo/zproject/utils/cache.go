package utils

import "webDemo/zproject/model"

var SessionMaap map[string]*model.Session

func init() {
	SessionMaap = make(map[string]*model.Session, 1024)
}
