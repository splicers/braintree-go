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

type Add struct {
	Type      string           `xml:"type,attr"`
	Discounts []*AddedDiscount `xml:"item,omitempty"`
}

type Update struct {
	Type      string           `xml:"type,attr"`
	Discounts []*AddedDiscount `xml:"item,omitempty"`
}

type Remove struct {
	Type      string   `xml:"type,attr"`
	Discounts []string `xml:"item,omitempty"`
}

type Discounts struct {
	XMLName string  `xml:"discounts"`
	Add     *Add    `xml:"add,omitempty"`
	Update  *Update `xml:"update,omitempty"`
	Remove  *Remove `xml:"remove,omitempty"`
}
