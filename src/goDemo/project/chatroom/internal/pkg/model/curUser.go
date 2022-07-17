package model

import (
	"net"
)

// CurUser 因为在客户端，很多地方会使用到 curUser, 因此将其作为一个全局
type CurUser struct {
	Conn net.Conn
	User
}
