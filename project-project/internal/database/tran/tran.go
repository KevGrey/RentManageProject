package tran

import "pomfret.cn/project-project/internal/database"

// Transaction 事物的操作 一定要跟数据库有关 注入数据库的连接 gorm.db
type Transaction interface {
	Action(func(conn database.DbConn) error) error
}
