package main

import (
  "errors"
  "fmt"
)

type Outlet struct {
  ID      string
  Name    string
  Address string
  Items   []Item // 每个门店有多个商品
}

type Item struct {
  ID          string
  Name        string
  Description string
  Price       float64
  IsAccessory bool // 是否为配件
}

type LoyaltyProgram struct {
  ID                string
  Name              string
  PointsNeeded      int  // 需要的积分数量
  FreeAccessoryItem Item // 免费的配件
}

type Subscription struct {
  ID            string
  CustomerID    string
  MonthlyFee    float64
  DiscountRate  float64 // 折扣率
  ServiceAccess bool    // 是否有服务访问权
}

type Customer struct {
  ID            string
  Name          string
  Email         string
  GadgetPoints  int
  Subscriptions []Subscription // 客户的订阅列表
}

func (c *Customer) PurchaseItem(item Item) error {
  if item.IsAccessory {
    c.GadgetPoints++
  }
  if c.GadgetPoints >= 10 {
    c.GadgetPoints -= 10
    fmt.Println("You've earned a free accessory!")
  }
  return nil
}

func (c *Customer) SubscribeService(s Subscription) error {
  for _, sub := range c.Subscriptions {
    if sub.ID == s.ID {
      return errors.New("You are already subscribed to this service")
    }
  }
  c.Subscriptions = append(c.Subscriptions, s)
  fmt.Printf("You've subscribed to the service: %s\n", s.ID)
  return nil
}

func main() {
  customer := &Customer{
    ID:           "1",
    Name:         "John Doe",
    Email:        "johndoe@example.com",
    GadgetPoints: 0,
  }

  item := Item{
    ID:          "1",
    Name:        "Gadget",
    Description: "A cool gadget",
    Price:       100.0,
    IsAccessory: true,
  }

  err := customer.PurchaseItem(item)
  if err != nil {
    fmt.Printf("Error purchasing item: %s\n", err.Error())
  }

  subscription := Subscription{
    ID:            "1",
    CustomerID:    "1",
    MonthlyFee:    20.0,
    DiscountRate:  0.1,
    ServiceAccess: true,
  }

  err = customer.SubscribeService(subscription)
  if err != nil {
    fmt.Printf("Error subscribing to service: %s\n", err.Error())
  }
}
