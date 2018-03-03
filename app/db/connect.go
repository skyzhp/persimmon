package db

import (
	"github.com/revel/revel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"time"
)

var MasterDB *xorm.Engine

func InitDB() {
	driver := revel.Config.StringDefault("db.driver", "mysql")
	connectString := revel.Config.StringDefault("db.connect", "root:root@tcp(127.0.0.1:3306)/test")
	var err error
	MasterDB, err = xorm.NewEngine(driver, connectString)
	if err != nil {
		revel.INFO.Printf("DB Connected Error: %s", err)
	}
	MasterDB.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "pit_")
	MasterDB.SetTableMapper(tbMapper)
	MasterDB.ShowSQL(true)
}
