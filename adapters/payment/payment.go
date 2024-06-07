package payment

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

func NewPaymentClientAdapter() *ClientAdapter {
	return &ClientAdapter{}
}

func (p *ClientAdapter) Pay(order *entities.Order, paymentMethod string) error {
	fmt.Printf("Order %d paid by method %s\n", order.ID, paymentMethod)

	postBody, _ := json.Marshal(map[string]interface{}{
		"orderId": order.ID,
		"method":  paymentMethod,
		"value":   order.Amount(),
	})
	fmt.Printf("Post body: %s\n", string(postBody))

	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:8010/payments/", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
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

func (p *ClientAdapter) Reverse(order *entities.Order) error {
	fmt.Printf("Order %d payment reversed\n", order.ID)
	return nil
}
