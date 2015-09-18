package braintree

type DiscountList struct {
	XMLName   string     `xml:"discounts"`
	Discounts []Discount `xml:"discount"`
}

type Discount struct {
	XMLName string `xml:"discount"`
	Modification
}

type AddedDiscount struct {
	InheritedFromID string `xml:"inherited-from-id,omitempty"`
	Quantity        int    `xml:"quantity,omitempty"`
}

type UpdatedDiscount struct {
	InheritedFromID string `xml:"inherited-from-id,omitempty"`
	Quantity        int    `xml:"quantity,omitempty"`
}

type Discounts struct {
	XMLName string             `xml:"discounts"`
	Add     []*AddedDiscount   `xml:"add,omitempty"`
	Update  []*UpdatedDiscount `xml:"update,omitempty"`
	Remove  []string           `xml:"remove,omitempty"`
}
