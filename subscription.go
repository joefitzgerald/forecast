package forecast

import "time"

type subscriptionContainer struct {
	Subscription Subscription `json:"subscription"`
}

// Subscription describes the Forecast subscription
type Subscription struct {
	ID                  int `json:"id"`
	IntervalUnitAmounts struct {
		Monthly int `json:"monthly"`
		Yearly  int `json:"yearly"`
	} `json:"interval_unit_amounts"`
	NextBillingDate          string     `json:"next_billing_date"`
	DaysUntilNextBillingDate int        `json:"days_until_next_billing_date"`
	Amount                   int        `json:"amount"`
	DefaultDeactivationAt    *time.Time `json:"default_deactivation_at"`
	ReceiptRecipient         string     `json:"receipt_recipient"`
	Status                   string     `json:"status"`
	PurchasedPeople          int        `json:"purchased_people"`
	Interval                 string     `json:"interval"`
	Discounts                struct {
		MonthlyPercentage float64 `json:"monthly_percentage"`
		YearlyPercentage  float64 `json:"yearly_percentage"`
	} `json:"discounts"`
	PlaceholderLimit   *int      `json:"placeholder_limit"`
	Invoiced           bool      `json:"invoiced"`
	DaysUntilDue       int       `json:"days_until_due"`
	Balance            int       `json:"balance"`
	PastDueBalance     int       `json:"past_due_balance"`
	SalesTaxExempt     bool      `json:"sales_tax_exempt"`
	SalesTaxPercentage float64   `json:"sales_tax_percentage"`
	ConvertedAt        time.Time `json:"converted_at"`
	Card               struct {
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
	container, err := get[subscriptionContainer](api, "billing/subscription")
	if err != nil {
		return nil, err
	}
	return &container.Subscription, nil
}
