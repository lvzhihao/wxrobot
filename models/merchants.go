package models

import (
	"github.com/jinzhu/gorm"
)

type MerchantInterface interface {
	Unmarshal(iter interface{}) error //解析merchant
}

/*
 * 商户信息
 */
type Merchant struct {
	gorm.Model
	MerchantType   string `gorm:"not null;type:varchar(20);unique_index:uix_merchant_type_merchant_no" json:"merchant_type"` //商户类型
	MerchantNo     string `gorm:"not null;type:varchar(80);unique_index:uix_merchant_type_merchant_no" json:"merchant_no"`   //商户ID 或 微信号
	MerchantSecret string `gorm:"type:varchar(80)" json:"merchant_secret"`                                                   //商户开发密钥 或 微信token
	MerchantName   string `gorm:"type:varchar(100)" json:"merchant_name"`                                                    //商户名称
	IsEnabled      bool   `gorm:"default:false" json:"is_enabled"`                                                           //是否可用
}

func LoadEnabledMerchants(db *gorm.DB) (ret []Merchant, err error) {
	err = db.Where("is_enabled = ?", true).Find(&ret).Error
	return
}
