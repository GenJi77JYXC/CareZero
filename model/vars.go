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
