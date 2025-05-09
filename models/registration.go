package models

// RegistrationInfo represents registration information
type RegistrationInfo struct {
	PaymentType          string         `json:"payment_type"`
	NumberOfPlayers      int            `json:"number_of_players"`
	BasePrice            string         `json:"base_price"`
	IsManualPrice        bool           `json:"is_manual_price"`
	Registrations        []Registration `json:"registrations"`
	OnlinePaymentAllowed bool           `json:"online_payment_allowed"`
}

// Registration represents a registration
type Registration struct {
	ClassRegistrationID      string      `json:"class_registration_id"`
	Player                   Player      `json:"player"`
	Price                    string      `json:"price"`
	RegistrationDate         string      `json:"registration_date"`
	Payment                  Payment     `json:"payment"`
	CustomPriceConfiguration interface{} `json:"custom_price_configuration"`
	CustomPrice              string      `json:"custom_price"`
	IsManualPrice            bool        `json:"is_manual_price"`
	CourseBillID             *string     `json:"course_bill_id"`
	CategoryName             *string     `json:"category_name"`
}
