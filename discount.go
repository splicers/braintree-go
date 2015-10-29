package braintree

type DiscountList struct {
	XMLName   string     `xml:"discounts"`
	Discounts []Discount `xml:"discount"`
}

type Discount struct {
	XMLName string `xml:"discount"`
	Modification
}

type Quantity struct {
	Type  string `xml:"type,attr"`
	Value int    `xml:",chardata"`
}

type SubDiscount struct {
	InheritedFromID string    `xml:"inherited-from-id,omitempty"`
	Quantity        *Quantity `xml:"quantity,omitempty"`
}

type AddDiscounts struct {
	Type      string         `xml:"type,attr"`
	Discounts []*SubDiscount `xml:"item,omitempty"`
}

type UpdateDiscounts struct {
	Type      string         `xml:"type,attr"`
	Discounts []*SubDiscount `xml:"item,omitempty"`
}

type RemoveDiscounts struct {
	Type      string   `xml:"type,attr"`
	Discounts []string `xml:"item,omitempty"`
}

type Discounts struct {
	XMLName   string           `xml:"discounts"`
	Add       *AddDiscounts    `xml:"add,omitempty"`
	Update    *UpdateDiscounts `xml:"update,omitempty"`
	Remove    *RemoveDiscounts `xml:"remove,omitempty"`
	Discounts []*Discount      `xml:"discount,omitempty"`
}
