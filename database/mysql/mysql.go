package mysql

import (
	"fmt"

	"github.com/galaxy-toolkit/server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// New 根据配置生成 MySQL 数据库实例
func New(conf config.MySQL) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// NewGenerator 创建 Gorm 生成器
func NewGenerator(conf config.MySQL, output string) (*gen.Generator, error) {
	db, err := New(conf)
	if err != nil {
		return nil, err
	}

	generator := gen.NewGenerator(gen.Config{
		OutPath:       output, // output directory, default value is ./query
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	generator.UseDB(db)
	return generator, nil
}
