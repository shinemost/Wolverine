package domain

type SysUserInfo struct {
	UserId        int64  `json:"user_id,omitempty" db:"user_id"`       // 用户ID
	DeptId        int64  `json:"dept_id,omitempty" db:"dept_id"`       // 部门ID
	LoginName     string `json:"login_name,omitempty" db:"login_name"` // 登录账号
	UserName      string `json:"user_name,omitempty" db:"user_name"`   // 用户昵称
	UserType      string `json:"user_type,omitempty" db:"user_type"`   // 用户类型（00系统用户 01注册用户）
	Email         string `json:"email,omitempty"`                      // 用户邮箱
	Phonenumber   string `json:"phonenumber,omitempty"`                // 手机号码
	Sex           string `json:"sex,omitempty"`                        // 用户性别（0男 1女 2未知）
	Avatar        string `json:"avatar,omitempty"`                     // 头像路径
	Password      string `json:"password,omitempty" db:"password"`     // 密码
	Salt          string `json:"salt,omitempty"`                       // 盐加密
	Status        string `json:"status,omitempty"`                     // 帐号状态（0正常 1停用）
	DelFlag       string `json:"del_flag,omitempty" db:"del_flag"`     // 删除标志（0代表存在 2代表删除）
	LoginIp       string `json:"login_ip,omitempty" db:"login_ip"`     // 最后登录IP
	LoginDate     string `json:"login_date,omitempty" db:"login_date"` // 最后登录时间
	PwdUpdateDate string `json:"pwd_update_date,omitempty"`            // 密码最后更新时间
	CreateBy      string `json:"create_by,omitempty"`                  // 创建者
	CreateTime    string `json:"create_time,omitempty"`                // 创建时间
	UpdateBy      string `json:"update_by,omitempty"`                  // 更新者
	UpdateTime    string `json:"update_time,omitempty"`                // 更新时间
	Remark        string `json:"remark,omitempty"`                     // 备注
}
