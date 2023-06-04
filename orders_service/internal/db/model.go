package db

type Order struct {
	OrderUid          string   `json:"order_uid" db:"orderuid"`
	TrackNumber       string   `json:"track_number" db:"tracknumber"`
	Entry             string   `json:"entry" db:"entry"`
	DeliveryId        int      `json:"-" db:"deliveryid"`
	Delivery          Delivery `json:"delivery" db:"-"`
	PaymentId         string   `json:"-" db:"paymentid"`
	Payment           Payment  `json:"payment" db:"-"`
	Items             []Item   `json:"items" db:"-"`
	Locale            string   `json:"locale" db:"locale"`
	InternalSignature string   `json:"internal_signature" db:"internalsignature"`
	CustomerId        string   `json:"customer_id" db:"customerid"`
	DeliveryService   string   `json:"delivery_service" db:"deliveryservice"`
	Shardkey          string   `json:"shardkey" db:"shardkey"`
	SmId              int      `json:"sm_id" db:"smid"`
	DateCreated       string   `json:"date_created" db:"datecreated"`
	OofShard          string   `json:"oof_shard" db:"oofshard"`
}

type Payment struct {
	Transaction  string `json:"transaction" db:"transaction"`
	RequestId    string `json:"request_id" db:"requestid"`
	Currency     string `json:"currency" db:"currency"`
	Provider     string `json:"provider" db:"provider"`
	Amount       int    `json:"amount" db:"amount"`
	PaymentDt    int    `json:"payment_dt" db:"paymentdt"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost int    `json:"delivery_cost" db:"deliverycost"`
	GoodsTotal   int    `json:"goods_total" db:"goodstotal"`
	CustomFee    int    `json:"custom_fee" db:"customfee"`
}

type Item struct {
	ChrtId      int    `json:"chrt_id" db:"chrtid"`
	TrackNumber string `json:"track_number" db:"tracknumber"`
	Price       int    `json:"price" db:"price"`
	Rid         string `json:"rid" db:"rid"`
	Name        string `json:"name" db:"name"`
	Sale        int    `json:"sale" db:"sale"`
	Size        string `json:"size" db:"size"`
	TotalPrice  int    `json:"total_price" db:"totalprice"`
	NmId        int    `json:"nm_id" db:"nmid"`
	Brand       string `json:"brand" db:"brand"`
	Status      int    `json:"status" db:"status"`
	OrderUid    string `json:"-" db:"orderuid"`
}

type Delivery struct {
	Id      int    `json:"-" db:"id"`
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	Zip     string `json:"zip" db:"zip"`
	City    string `json:"city" db:"city"`
	Address string `json:"address" db:"address"`
	Region  string `json:"region" db:"region"`
	Email   string `json:"email" db:"email"`
}
