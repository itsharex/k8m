package models

import (
	"time"

	"github.com/weibaohui/k8m/internal/dao"
	"github.com/weibaohui/k8m/pkg/comm/utils"
	"gorm.io/gorm"
)

// User 用户导入User
type User struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Username   string    `gorm:"uniqueIndex;not null" json:"username,omitempty"`
	Salt       string    `gorm:"not null" json:"salt,omitempty"`
	Password   string    `gorm:"not null" json:"password,omitempty"`
	GroupNames string    `json:"group_names"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`                             // Automatically managed by GORM for update time
	CreatedBy  string    `gorm:"index:idx_created_by" json:"created_by,omitempty"` // 创建者
	// 2FA相关字段
	TwoFAEnabled     bool   `gorm:"default:false" json:"two_fa_enabled,omitempty"` // 是否启用2FA
	TwoFAType        string `gorm:"size:20" json:"two_fa_type,omitempty"`          // 2FA类型：如 'totp', 'sms', 'email'
	TwoFASecret      string `gorm:"size:100" json:"two_fa_secret,omitempty"`       // 2FA密钥
	TwoFABackupCodes string `gorm:"size:500" json:"two_fa_backup_codes,omitempty"` // 备用恢复码，逗号分隔
	TwoFAAppName     string `gorm:"size:100" json:"two_fa_app_name,omitempty"`     // 2FA应用名称，用于提醒用户使用的是哪个软件

}

const (
	// 用户角色两种，平台管理员、普通用户
	RolePlatformAdmin = "platform_admin" // 平台管理员
	RoleGuest         = "guest"          // 普通用户，只能登录，约等于游客,无任何集群权限，也看不到集群列表

	// 平台管理员拥有所有权限
	// 普通用户需要赋予集群角色，
	// 集群角色三种，集群管理员、集群只读、集群Pod内执行命令
	RoleClusterAdmin    = "cluster_admin"    // 集群管理员
	RoleClusterReadonly = "cluster_readonly" // 集群只读权限
	RoleClusterPodExec  = "cluster_pod_exec" // 集群Pod内执行命令权限
)

func (c *User) List(params *dao.Params, queryFuncs ...func(*gorm.DB) *gorm.DB) ([]*User, int64, error) {

	return dao.GenericQuery(params, c, queryFuncs...)
}

func (c *User) Save(params *dao.Params, queryFuncs ...func(*gorm.DB) *gorm.DB) error {
	return dao.GenericSave(params, c, queryFuncs...)
}

func (c *User) Delete(params *dao.Params, ids string, queryFuncs ...func(*gorm.DB) *gorm.DB) error {
	return dao.GenericDelete(params, c, utils.ToInt64Slice(ids), queryFuncs...)
}

func (c *User) GetOne(params *dao.Params, queryFuncs ...func(*gorm.DB) *gorm.DB) (*User, error) {
	return dao.GenericGetOne(params, c, queryFuncs...)
}
