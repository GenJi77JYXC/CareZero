package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string   `gorm:"size:255;unique;not null"` // 商品名称
	Description string   `gorm:"type:text"`                // 商品描述
	Picture     []string `gorm:"type:text"`                // 商品图片（JSON 数组）
	Price       float32  `gorm:"type:decimal(10,2)"`       // 商品价格
	Category    string   `gorm:"size:255"`                 // 商品分类
	Stock       int      `gorm:"default:0"`                // 商品库存
	SKU         string   `gorm:"size:100;unique;not null"` // 商品唯一编码
	IsActive    bool     `gorm:"default:true"`             // 商品是否上架
	CreatedByID uint     `gorm:"not null"`                 // 创建人 ID
	UpdatedByID uint     `gorm:"not null"`                 // 更新人 ID
	CreatedBy   User     `gorm:"foreignKey:CreatedByID"`   // 关联创建人
	UpdatedBy   User     `gorm:"foreignKey:UpdatedByID"`   // 关联更新人
	Tags        []string `gorm:"type:text"`                // 商品标签（JSON 数组）
}

func CreateProduct(db *gorm.DB, name string, createdByID uint, updatedByID uint) error {
	product := Product{
		Name:        name,
		CreatedByID: createdByID,
		UpdatedByID: updatedByID,
		// 设置其他字段...
	}
	return db.Create(&product).Error
}

func GetProductWithUsers(db *gorm.DB, productID uint) (*Product, error) {
	var product Product
	err := db.Preload("CreatedBy").Preload("UpdatedBy").First(&product, productID).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func UpdateProduct(db *gorm.DB, productID uint, updatedByID uint, updates map[string]interface{}) error {
	updates["updated_by_id"] = updatedByID
	return db.Model(&Product{}).Where("id = ?", productID).Updates(updates).Error
}
