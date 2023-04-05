package main

import(
"github.com/gin-gonic/gin")

//we registry service by register a customer
type Customer struct{
	ID  string `json:"id"`
	Name string `json:"name"`
	Adress string `json:"adress"`
	Email string `json:"email"`

}


type Product struct{
	ID  string `json:"id"`
	Category string `json:"category"`
	Description string `json:"description"`
	Quantity string `json:"quantity"`
	Price string `json:"price"`

}

type ProductCategory struct{
	ID  string `json:"id"`
	Product_type string`json:"product_type"`
	Name string `json:"name"`


}

type Order struct {
	ID  string `json:"id"`
	CutomerID string `json:"customer_id"`
	Products string `json:"products"`
	TotalAmount float64 `json:"total_amount"`
}

type OrderProduct struct{
	ProductID  string `json:"product+_id"`
	Qantity int `json:"quantity"`
}

type PaymentMethod int

const(
	BankPayment PaymentMethod=iota
	CashPayment 
	MobilePayment
)


type Payment struct{
ID string `json:"id"`
OrderID  string `json:"order_id"`
Amount  float64 `json:"amount"`
Method PaymentMethod `json:"method"`
Processed bool `json:"processed"`


}

type Inventory interface{
	CreayeProduct(id string)(*Product, error)
	UpdateProduct(id string,quantity int)error
}

type CustomerService interface{
	Registry(customer*Customer)error
	GetByID(id string)(*Customer,error)
}

type ProductService interface {
GetAll()([]Product,error)
GetByID(id string)(*Product,error)
Create(customer*Customer)error
Update(id string,quantity int)error
}


type PaymentService  interface{
	Pcocess(payment*Payment)error
}

type OrderService interface{
	Create(order*Order)error
}

type ShoppersAPI struct{
	customerService CustomerService
	productService ProductService
	orderService OrderService
	paymentService PaymentService
inventory Inventory
}


func NewShoppersAPI(customerService CustomerService,productService ProductService,paymentService PaymentService,orderService OrderService,inventory Inventory)*ShoppersAPI{
return &ShoppersAPI{
	customerService: customerService,
	productService:  productService,
	orderService:    orderService,
	paymentService:  paymentService,
	inventory:       inventory,
}
}


func(api*ShoppersAPI)RegistryCustomer(c*gin.context)








func main(){
	//Initializing the database


	//create a router

	// r:=gin.Deafault


	//initialize the endpoints
	//r.GET("/customers",registryService.GetCustomers)
	//r.POST("/products",productService.CreateProduct)
	//r.PUT("/customers",customerService.UpdateCustomers)
	//r.Get("/orders",orderService.GetOrder)
	//r.Post("/orders",orderService.UpddateOrder)
	//r.GET("/paymeent",paymentSErvice.GetPayment)
//r.POST("/payement",paymentService.CreatePayment)
	//
	//
	//Start the server aand connect the service

}