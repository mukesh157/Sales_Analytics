package model

import "time"

type Customerdetails struct {
	ID               uint       `gorm:"primaryKey"` // Correct way to handle auto-increment
	Customer_id      string     `gorm:"column:customer_id"`
	Customer_name    string     `gorm:"column:name"`
	Customer_email   string     `gorm:"column:email"`
	Customer_address string     `gorm:"column:address"`
	CreatedAt        *time.Time `gorm:"column:created_date"`
	UpdatedAt        *time.Time `gorm:"column:updated_date"`
}

type Products struct {
	ID         uint       `gorm:"primaryKey"` // Correct way to handle auto-increment
	Product_id string     `gorm:"column:product_id"`
	Name       string     `gorm:"column:name"`
	Category   string     `gorm:"column:category"`
	CreatedAt  *time.Time `gorm:"column:created_date"`
	UpdatedAt  *time.Time `gorm:"column:updated_date"`
}

type Orders struct {
	ID             uint       `gorm:"primaryKey"` // Correct way to handle auto-increment
	Order_id       string     `gorm:"column:order_id"`
	Customer_id    uint       `gorm:"column:customer_id"`
	Region         string     `gorm:"column:region"`
	Date_of_sale   string     `gorm:"column:date_of_sale"`
	Payment_method string     `gorm:"column:payment_method"`
	Shipping_cost  string     `gorm:"column:shipping_cost"`
	CreatedAt      *time.Time `gorm:"column:created_date"`
	UpdatedAt      *time.Time `gorm:"column:updated_date"`
}

type Order_items struct {
	Order_id      uint       `gorm:"column:order_id"`
	Product_id    uint       `gorm:"column:product_id"`
	Quantity_sold int        `gorm:"column:quantity_sold"`
	Unit_price    float64    `gorm:"column:unit_price"`
	Discount      float64    `gorm:"column:discount"`
	CreatedAt     *time.Time `gorm:"column:created_date"`
	UpdatedAt     *time.Time `gorm:"column:updated_date"`
}

type Response struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type GetDetails struct {
	StartDate string `json:"fromDate"`
	EndDate   string `json:"endDate"`
}

type GetRevenueDetails struct {
	Status             string            `json:"status"`
	Total_revenue      string            `json:"total_Revenue"`
	TotProdRevenue     []ProductRevenue  `json:"totProdRevenue"`
	TotalcatRevenue    []CategoryRevenue `json:"totalcatRevenue"`
	TotalRevenue_byreg []RegionRevenue   `json:"totalregionRevenue"`
	TopProduct         []TopProduct      `json:"topProduct"`
	TopCategory        []TopCategory     `json:"topcategory"`
	TopRegion          []TopRegion       `json:"topRegion"`
}

type TopProduct struct {
	ProductName   string `json:"product_name"`
	TotalQuantity int    `json:"total_quantity"`
}

type TopCategory struct {
	Category      string `json:"category"`
	TotalQuantity int    `json:"total_quantity"`
}

type TopRegion struct {
	Region        string `json:"region"`
	TotalQuantity int    `json:"total_quantity"`
}
type ProductRevenue struct {
	ProductName  string  `json:"product_name" gorm:"column:product_name"`
	TotalRevenue float64 `json:"total_revenue" gorm:"column:total_revenue"`
}

type CategoryRevenue struct {
	Category     string  `json:"category" gorm:"column:category"`
	TotalRevenue float64 `json:"total_revenue" gorm:"column:total_revenue"`
}

type RegionRevenue struct {
	Region       string  `json:"region"`
	TotalRevenue float64 `json:"total_revenue"`
}
