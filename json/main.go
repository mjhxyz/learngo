package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID         string      `json:"id"`
	Items      []OrderItem `json:"item"`
	TotalPrice float64     `json:"total_price"`
}

func marshal() {
	o := Order{
		ID:         "1234",
		TotalPrice: 20,
		Items: []OrderItem{
			{
				ID:    "item_2",
				Name:  "mao",
				Price: 15,
			},
			{
				ID:    "item_1",
				Name:  "mao1",
				Price: 10,
			},
		},
	}

	b, _ := json.Marshal(o)
	fmt.Printf("%s\n", b)
}

func main() {
	unmarshal()
}

func unmarshal() {
	s := `{"id":"1234","item":[{"id":"item_2","name":"mao","price":15},{"id":"item_1","name":"mao1","price":10}],"total_price":20}`
	var o Order
	// json 字节, 转换后的类型
	_ = json.Unmarshal([]byte(s), &o) // 会写入 o
	fmt.Printf("结果:%+v\n", o)
}
