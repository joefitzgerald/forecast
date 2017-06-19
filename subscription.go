package forecast

type subscriptionContainer struct {
	Subscription Subscription `json:"subscription"`
}

// Subscription describes the Forecast subscription
type Subscription struct {
	NextBillingDate  string `json:"next_billing_date"`
	Amount           int    `json:"amount"`
	AmountPerPerson  int    `json:"amount_per_person"`
	ReceiptRecipient string `json:"receipt_recipient"`
	Status           string `json:"status"`
	PurchasedPeople  int    `json:"purchased_people"`
	Interval         string `json:"interval"`
	Discounts        struct {
		MonthlyPercentage int `json:"monthly_percentage"`
		YearlyPercentage  int `json:"yearly_percentage"`
	} `json:"discounts"`
	Card struct {
		Brand       string `json:"brand"`
		LastFour    string `json:"last_four"`
		ExpiryMonth int    `json:"expiry_month"`
		ExpiryYear  int    `json:"expiry_year"`
	} `json:"card"`
	Address struct {
		Line1      string `json:"line_1"`
		Line2      string `json:"line_2"`
		City       string `json:"city"`
		State      string `json:"state"`
		PostalCode string `json:"postal_code"`
		Country    string `json:"country"`
	} `json:"address"`
}

// Subscription returns information about the Forecast account subscription
func (api *API) Subscription() (*Subscription, error) {
	var container subscriptionContainer
	err := api.do("billing/subscription", &container)
	if err != nil {
		return nil, err
	}
	return &container.Subscription, nil
}
