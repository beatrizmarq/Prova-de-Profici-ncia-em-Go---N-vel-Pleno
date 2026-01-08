package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Product struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validate() error {
	if p.ID == "" {
		return errors.New("product: id é obrigatório")
	}
	if p.Name == "" {
		return errors.New("product: nome é obrigatório")
	}
	if p.Price <= 0 {
		return errors.New("product: preço deve ser maior que zero")
	}
	if p.Stock < 0 {
		return errors.New("product: estoque não pode ser negativo")
	}
	return nil
}

// -------------------
// Order
// -------------------

type Order struct {
	ID          string    `json:"id"`
	CustomerID  string    `json:"customer_id"`
	Products    []Product `json:"products,omitempty"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func (o *Order) CalculateTotal() {
	var total float64
	for _, p := range o.Products {
		total += p.Price
	}
	o.TotalAmount = total
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("order: id é obrigatório")
	}
	if o.CustomerID == "" {
		return errors.New("order: customer_id é obrigatório")
	}
	if len(o.Products) == 0 {
		return errors.New("order: pedido deve conter ao menos um produto")
	}
	if o.Status == "" {
		return errors.New("order: status é obrigatório")
	}
	return nil
}

// -------------------
// Exemplo de uso
// -------------------

func main() {
	product := Product{
		ID:        "p1",
		Name:      "Notebook",
		Price:     3500.00,
		Stock:     10,
		CreatedAt: time.Now(),
	}

	if err := product.Validate(); err != nil {
		fmt.Println("Erro no produto:", err)
		return
	}

	order := Order{
		ID:         "o1",
		CustomerID: "c123",
		Products:   []Product{product},
		Status:     "created",
		CreatedAt:  time.Now(),
	}

	order.CalculateTotal()

	if err := order.Validate(); err != nil {
		fmt.Println("Erro no pedido:", err)
		return
	}

	// Marshal (struct -> JSON)
	orderJSON, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		fmt.Println("Erro ao serializar JSON:", err)
		return
	}

	fmt.Println("JSON gerado:")
	fmt.Println(string(orderJSON))

	// Unmarshal (JSON -> struct)
	var decodedOrder Order
	if err := json.Unmarshal(orderJSON, &decodedOrder); err != nil {
		fmt.Println("Erro ao desserializar JSON:", err)
		return
	}

	fmt.Println("Pedido decodificado:", decodedOrder)
}
