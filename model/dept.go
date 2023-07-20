package model

import (
	"gorm.io/gorm"
	"time"
)

type DingDept struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	UserList  []DingUser `gorm:"many2many:user_dept"`
	DeptId    int        `gorm:"primaryKey" json:"dept_id"`
	Deleted   gorm.DeletedAt
	Name      string `json:"name"`
	ParentId  int    `json:"parent_id"`
	IsSendFirstPerson int    `json:"is_send_first_person"` // 0为不推送，1为推送
	RobotToken        string `json:"robot_token"`
	IsRobotAttendance int    `json:"is_robot_attendance"` //是否
	IsJianShuOrBlog   int    `json:"is_jianshu_or_blog" gorm:"column:is_jianshu_or_blog"`
}
