package postgres

import (
	"fmt"

	"github.com/galaxy-toolkit/server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// New 根据配置生成 Postgres 数据库实例
func New(conf config.Postgres) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Host,
		conf.UserName,
		conf.Password,
		conf.Database,
		conf.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GeneratorConfig 创建生成器配置
type GeneratorConfig struct {
	QueryPath     string // dao 层方法生成路径
	QueryFilename string // 数据库层生成文件名
	ModelPath     string // 模型生成路径
}

// NewGenerator 创建 Gorm 生成器
func NewGenerator(conf config.Postgres, gConf GeneratorConfig) (*gen.Generator, error) {
	db, err := New(conf)
	if err != nil {
		return nil, err
	}

	generator := gen.NewGenerator(gen.Config{
		OutPath:           gConf.QueryPath,
		OutFile:           gConf.QueryFilename,
		ModelPkgPath:      gConf.ModelPath,
		WithUnitTest:      true,
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	generator.UseDB(db)
	return generator, nil
}
