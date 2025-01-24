package model

// 主体 sub
const (
	AdminRole        = "admin_role"         // 管理员
	StoreManagerRole = "store_manager_role" // 店长
	StaffRole        = "staff_role"         // 员工
	CustomerRole     = "customer_role"      // 顾客
)

// 领域 domain
const (
	UserDomain = "user_domain" // 用户领域
)

// 对象 obj
const (
	UserObj = "user_data" // 用户数据
)

// 操作 act
const (
	Read  = "read"
	Write = "write"
)

const Salt = "suki_"

const (
	PasswordError        = iota + 40001 // "密码错误"
	UserExists                          // "用户已经存在"
	EmailError                          // “邮箱格式错误"
	ConfirmPasswordError                // "两次密码不一致"
	UserUnExists                        // "用户不存在"
	TokenInvalid                        // Token失效
)

var ErrorMsg = map[int]string{
	PasswordError:        "密码错误",
	UserExists:           "用户已经存在",
	EmailError:           "邮箱格式错误",
	ConfirmPasswordError: "两次密码不一致",
	UserUnExists:         "用户不存在",
	TokenInvalid:         "Token失效",
}
