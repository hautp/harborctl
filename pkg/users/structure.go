package users

import (
    "time"
)

type UserInfo []struct {
    UserId          string          `json:"user_id"`
    UserName        string          `json:"username"`
    Email           string          `json:"email"`
    Password        string          `json:"password"`
    PasswordVer     string          `json:"password_version"`
    RealName        string          `json:"realname"`
    Comment         string          `json:"comment"`
    Deleted         string          `json:"deleted"`
    RoleName        string          `json:"role_name"`
    RoleId          int             `json:"role_id"`
    SysadFlag       bool            `json:"sysadmin_flag"`
    AdRoleAuth      bool            `json:"admin_role_in_auth"`
    ResetUUID       string          `json:"reset_uuid"`
    CreateTime      time.Time       `json:"creation_time"`
    UpdateTime      time.Time       `json:"update_time"`
}
