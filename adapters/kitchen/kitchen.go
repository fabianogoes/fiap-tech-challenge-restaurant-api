package kitchen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"io/ioutil"
	"log"
	"net/http"
)

type ClientAdapter struct {
}

func NewKitchenClientAdapter() *ClientAdapter {
	return &ClientAdapter{}
}

func (p *ClientAdapter) Preparation(order *entities.Order) error {
	fmt.Printf("Order Preparation %d \n", order.ID)

	requestBytes, _ := json.Marshal(toCreationRequest(order))
	fmt.Printf("Post body: %s\n", string(requestBytes))

	responseBody := bytes.NewBuffer(requestBytes)
	resp, err := http.Post("http://localhost:8020/kitchen/orders", "application/json", responseBody)
	if err != nil {
		return fmt.Errorf("error kitchen creation request: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)

	response := OrderResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://localhost:8020/kitchen/orders/%d/preparation", order.ID)
	resp, err = http.Post(url, "application/json", responseBody)
	if err != nil {
		return fmt.Errorf("error kitchen preparation request: %s", err)
	}

	return nil
}

func (p *ClientAdapter) ReadyDelivery(orderID uint) error {
	fmt.Printf("Order Ready Delivery %d \n", orderID)

	url := fmt.Sprintf("http://localhost:8020/kitchen/orders/%d/preparation", orderID)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return fmt.Errorf("error kitchen preparation request: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)

	return nil
}

func toCreationRequest(order *entities.Order) *CreationRequest {
	items := make([]*OrderItem, 0)
	for _, item := range order.Items {
		items = append(items, &OrderItem{
			Product: &Product{
				Name:     item.Product.Name,
				Category: item.Product.Category.Name,
			},
			Quantity: item.Quantity,
		})
	}

	return &CreationRequest{
		ID:    order.ID,
		Items: items,
	}
}

type CreationRequest struct {
	ID    uint         `json:"id"`
	Items []*OrderItem `json:"items"`
}

type OrderItem struct {
	ID       uint     `json:"id"`
	Product  *Product `json:"product"`
	Quantity int      `json:"quantity"`
}

type Product struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type OrderResponse struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}
