package order

import (
	"time"
	"yasser-backend/internal/user"
	"yasser-backend/internal/vendor-group/vendor"
	"yasser-backend/pkg/models"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusPreparing OrderStatus = "preparing"
	OrderStatusReady     OrderStatus = "ready"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	models.BaseModel
	VendorID     int64       `gorm:"type:bigint;index;not null" json:"vendorId"`
	UserID       int64       `gorm:"type:bigint;index;not null" json:"userId"`
	Status       OrderStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
	TotalAmount  float64     `gorm:"type:decimal(10,2);not null" json:"totalAmount"`
	DeliveryFee  float64     `gorm:"type:decimal(10,2);default:0" json:"deliveryFee"`
	Notes        string      `gorm:"type:text" json:"notes,omitempty"`
	DeliveryTime *time.Time  `gorm:"type:timestamp" json:"deliveryTime,omitempty"`
	
	// Relations
	Items []OrderItem `gorm:"foreignKey:OrderID" json:"items,omitempty"`

	User   user.User     `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Vendor vendor.Vendor `gorm:"foreignKey:VendorID;references:ID" json:"vendor,omitempty"`
}

type OrderItem struct {
	models.BaseModel
	OrderID      int64   `gorm:"type:bigint;index;not null" json:"orderId"`
	ItemID       int64   `gorm:"type:bigint;index;not null" json:"itemId"`
	ItemName     string  `gorm:"type:text;not null" json:"itemName"`
	Quantity     int     `gorm:"type:int;not null" json:"quantity"`
	UnitPrice    float64 `gorm:"type:decimal(10,2);not null" json:"unitPrice"`
	TotalPrice   float64 `gorm:"type:decimal(10,2);not null" json:"totalPrice"`
	SizeID       *int64  `gorm:"type:bigint" json:"sizeId,omitempty"`
	SizeName     *string `gorm:"type:text" json:"sizeName,omitempty"`
	SizePrice    *int    `gorm:"type:int" json:"sizePrice,omitempty"`
	VariantID    *int64  `gorm:"type:bigint" json:"variantId,omitempty"`
	VariantName  *string `gorm:"type:text" json:"variantName,omitempty"`
	
	// Relations
	Addons []OrderItemAddon `gorm:"foreignKey:OrderItemID" json:"addons,omitempty"`
}

type OrderItemAddon struct {
	models.BaseModel
	OrderItemID int64  `gorm:"type:bigint;index;not null" json:"orderItemId"`
	AddonID     int64  `gorm:"type:bigint;not null" json:"addonId"`
	AddonName   string `gorm:"type:text;not null" json:"addonName"`
	AddonPrice  int64  `gorm:"type:bigint;not null" json:"addonPrice"`
	IsRemoval   bool   `gorm:"type:boolean;default:false" json:"isRemoval"`
}