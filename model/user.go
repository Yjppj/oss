package model

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sso/global"
)

type DingUser struct {
	UserId              string ` gorm:"primaryKey;foreignKey:UserId" json:"userid"`
	Deleted             gorm.DeletedAt
	Name                string         `json:"name"`
	Mobile              string         `json:"mobile"`
	Password            string         `json:"password"`
	DeptIdList          []int          `json:"dept_id_list" gorm:"-"` //所属部门
	DeptList            []DingDept     `json:"dept_list" gorm:"many2many:user_dept"`
	AuthorityId         uint           `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
	Authority           SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities         []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Title               string         `json:"title"` //职位
	JianShuAddr         string         `json:"jianshu_addr"`
	BlogAddr            string         `json:"blog_addr"`
	AuthToken           string         `json:"auth_token" gorm:"-"`
	IsExcellentJianBlog bool           `json:"is_excellentBlogJian"`
}

func (d *DingUser) Login() (user *DingUser, err error) {
	user = &DingUser{
		Mobile:   d.Mobile,
		Password: d.Password,
	}
	//此处的Login函数传递的是一个指针类型的数据
	opassword := user.Password //此处是用户输入的密码，不一定是对的
	err = global.GLOAB_DB.Where(&DingUser{Mobile: user.Mobile}).Preload("Authorities").Preload("Authority").Preload("DeptList").First(user).Error
	if err != nil {
		zap.L().Error("登录时查询数据库失败", zap.Error(err))
		return
	}
	//如果到了这里还没有结束的话，那就说明该用户至少是存在的，于是我们解析一下密码
	//password := encryptPassword(opassword)
	password := opassword
	//拿到解析后的密码，我们看看是否正确
	if password != user.Password {
		return nil, errors.New("密码错误")
	}
	//如果能到这里的话，那就登录成功了
	return
}
