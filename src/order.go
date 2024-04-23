package src

import (
	"fmt"
	"time"
)

// 订单状态枚举
type OrderStatus string

const (
	Pending   OrderStatus = "待付款"
	Paid      OrderStatus = "已付款"
	Shipped   OrderStatus = "已发货"
	Delivered OrderStatus = "已送达"
	Canceled  OrderStatus = "已取消"
)

// 订单仓库（模拟数据库）
var orders = make(map[int]*Order)

// 获取下一个可用的订单ID
func getNextOrderID() int {
	maxID := 0
	for id := range orders {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}

// 添加订单
func AddOrder(userID, productID, quantity int, amount float64) {
	orderID := getNextOrderID()
	order := &Order{
		ID:        orderID,
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
		Amount:    amount,
		Status:    Pending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	orders[orderID] = order
	fmt.Printf("订单 %d 已添加\n", orderID)
}

// 查询订单
func GetOrder(orderID int) (*Order, bool) {
	order, exists := orders[orderID]
	return order, exists
}

// 修改订单状态
func UpdateOrderStatus(orderID int, status OrderStatus) {
	order, exists := orders[orderID]
	if !exists {
		fmt.Printf("订单ID %d 不存在\n", orderID)
		return
	}
	order.Status = status
	order.UpdatedAt = time.Now()
	fmt.Printf("订单 %d 的状态已更新为 %s\n", orderID, status)
}

// 删除订单
func DeleteOrder(orderID int) {
	delete(orders, orderID)
	fmt.Printf("订单 %d 已删除\n", orderID)
}

// 打印所有订单信息
func PrintAllOrders() {
	for _, order := range orders {
		fmt.Printf("订单ID: %d, 用户ID: %d, 商品ID: %d, 数量: %d, 金额: %.2f, 状态: %s, 创建时间: %s, 更新时间: %s\n",
			order.ID, order.UserID, order.ProductID, order.Quantity, order.Amount, order.Status, order.CreatedAt, order.UpdatedAt)
	}
}

//func main() {
//	// 添加订单示例
//	AddOrder(1, 1, 2, 100.00)
//	AddOrder(2, 2, 1, 50.00)
//
//	// 打印所有订单
//	PrintAllOrders()
//
//	// 修改订单状态示例
//	UpdateOrderStatus(1, Paid)
//
//	// 再次打印所有订单以验证状态更新
//	PrintAllOrders()
//
//	// 删除订单示例
//	DeleteOrder(2)
//
//	// 再次打印所有订单以验证订单已被删除
//	PrintAllOrders()
//}
