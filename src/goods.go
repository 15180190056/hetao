package src

import (
	"fmt"
	"time"
)

// 商品仓库（模拟数据库）
var products = make(map[int]*Product)
var categories = make(map[int]*Category)

// 获取下一个可用的商品ID
func getNextProductID() int {
	maxID := 0
	for id := range products {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}

// 获取下一个可用的分类ID
func getNextCategoryID() int {
	maxID := 0
	for id := range categories {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}

// 添加商品
func AddProduct(name string, price float64, stock int, description string, image string, categoryID int) {
	productID := getNextProductID()
	category, exists := categories[categoryID]
	if !exists {
		fmt.Printf("分类ID %d 不存在\n", categoryID)
		return
	}

	product := &Product{
		ID:          productID,
		Name:        name,
		Price:       price,
		Stock:       stock,
		Description: description,
		Image:       image,
		Status:      "上架",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Category:    category,
	}

	products[productID] = product
	fmt.Printf("商品 %s 已添加\n", name)
}

// 编辑商品
func EditProduct(id int, name string, price float64, stock int, description string, image string) {
	product, exists := products[id]
	if !exists {
		fmt.Printf("商品ID %d 不存在\n", id)
		return
	}

	product.Name = name
	product.Price = price
	product.Stock = stock
	product.Description = description
	product.Image = image
	product.UpdatedAt = time.Now()

	fmt.Printf("商品 %s 已编辑\n", name)
}

// 上架商品
func ShelveProduct(id int) {
	product, exists := products[id]
	if !exists {
		fmt.Printf("商品ID %d 不存在\n", id)
		return
	}

	if product.Status == "下架" {
		product.Status = "上架"
		product.UpdatedAt = time.Now()
		fmt.Printf("商品 %s 已上架\n", product.Name)
	} else {
		fmt.Printf("商品 %s 已经是上架状态\n", product.Name)
	}
}

// 下架商品
func UnshelveProduct(id int) {
	product, exists := products[id]
	if !exists {
		fmt.Printf("商品ID %d 不存在\n", id)
		return
	}

	if product.Status == "上架" {
		product.Status = "下架"
		product.UpdatedAt = time.Now()
		fmt.Printf("商品 %s 已下架\n", product.Name)
	} else {
		fmt.Printf("商品 %s 已经是下架状态\n", product.Name)
	}
}

// 添加分类
func AddCategory(name string) {
	// 检查分类名称是否为空
	if name == "" {
		fmt.Println("分类名称不能为空")
		return
	}

	// 获取下一个可用的分类ID
	categoryID := getNextCategoryID()

	// 创建新的分类结构体实例
	category := &Category{
		ID:   categoryID,
		Name: name,
	}

	// 将新分类添加到分类映射中
	categories[categoryID] = category

	// 打印成功信息
	fmt.Printf("分类 %s 已添加，ID为：%d\n", name, categoryID)
}
