package user

import (
	"github.com/galaxy-toolkit/server/config"
	"github.com/galaxy-toolkit/server/db/postgres"
)

// User 用户
type User struct {
	ID          int64  `gorm:"column:id;type:BIGINT;primaryKey;autoIncrement;comment:ID"` // ID
	UserName    string `gorm:"column:user_name;type:VARCHAR(255);not null;comment:用户名"`   // 用户名
	Desc        string `gorm:"column:desc;type:TEXT;not null;comment:描述"`                 // 描述
	Avatar      string `gorm:"column:avatar;type:TEXT;not null;comment:头像"`               // 头像
	Email       string `gorm:"column:email;type:VARCHAR(255);not null;comment:邮箱"`        // 邮箱
	Phone       string `gorm:"column:phone;type:VARCHAR(255);not null;comment:手机号"`       // 手机号
	Gender      string `gorm:"column:gender;type:VARCHAR(255);not null;comment:性别"`       // 性别
	Password    string `gorm:"column:password;type:VARCHAR(255);not null;comment:密码"`     // 密码
	Point       int64  `gorm:"column:point;type:BIGINT;not null;comment:积分"`              // 积分
	Status      string `gorm:"column:status;type:VARCHAR(255);not null;comment:状态"`       // 状态
	CreatedTime int64  `gorm:"column:created_time;type:TIMESTAMP;not null;comment:创建时间"`  // 创建时间
	UpdatedTime int64  `gorm:"column:updated_time;type:TIMESTAMP;not null;comment:更新时间"`  // 更新时间
	DeletedTime int64  `gorm:"column:deleted_time;type:TIMESTAMP;not null;comment:删除时间"`  // 删除时间
}

// TableName 获取表名
func (User) TableName() string {
	return "user"
}

// Users 用户列表
type Users []User

// IDs 获取 ID 列表
func (users Users) IDs() []int64 {
	ids := make([]int64, len(users))
	for i, user := range users {
		ids[i] = user.ID
	}
	return ids
}

func migrate() {
	var conf config.Config
	if err := config.LoadAndWatch[config.Config]("config/config.yaml", &conf); err != nil {
		panic(err)
	}

	db, err := postgres.New(&conf.Database.Postgres)
	if err != nil {
		panic(err)
	}

	if err = db.Migrator().AutoMigrate(&User{}); err != nil {
		panic(err)
	}
}
