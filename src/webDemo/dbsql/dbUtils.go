package dbsql

//在 go 标准库中是不包含任何数据库驱动包的，所以连接数据库，需要使用第三方的驱动包
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//DB 是一个数据库操作句柄，代表一个具有零到多个底层连接的连接池。它可以安全的被多个 go 协程同时使用。
//sql 包会自动创建和释放连接，也会维护一个闲置连接的连接池。
//
//如果数据库具有单连接状态的概念，该状态只有在事务中被观察时才可信。一旦调用了 DB.Begin，返回的 Tx 会绑定到单个连接。当调用事务 Tx
//的 commit 或 rollback 后，该事务使用的连接会归还到 DB 的闲置连接池中。连接池大小可用 SetMaxIdleConns 方法控制。
var (
	DB  *sql.DB
	err error
)

func init() {
	//open 函数只能验证其参数，但不创建与数据库的连接。如果要检查指定的数据源是否能连接，可调用返回值的 ping 方法。
	//并且返回的 DB 可以安全的被多个 go 协程同时使用，且会维护自身的闲置连接池。如此 open 函数只需要调用一次，很少需要关闭 DB。
	//DB, err = sql.Open("mysql","username:password@tcp(localhost:3306)/test")
	DB, err = sql.Open("mysql", "username:password@/test")
	if err != nil {
		panic(any(err.Error()))
	}
}
