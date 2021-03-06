package braintree

import (
	"github.com/lionelbarrow/braintree-go/nullable"
	"time"
)

type Transaction struct {
	XMLName             string               `xml:"transaction"`
	Id                  string               `xml:"id,omitempty"`
	CustomerID          string               `xml:"customer-id,omitempty"`
	Status              string               `xml:"status,omitempty"`
	Type                string               `xml:"type,omitempty"`
	Amount              *Decimal             `xml:"amount"`
	OrderId             string               `xml:"order-id,omitempty"`
	PaymentMethodToken  string               `xml:"payment-method-token,omitempty"`
	PaymentMethodNonce  string               `xml:"payment-method-nonce,omitempty"`
	MerchantAccountId   string               `xml:"merchant-account-id,omitempty"`
	PlanId              string               `xml:"plan-id,omitempty"`
	CreditCard          *CreditCard          `xml:"credit-card,omitempty"`
	Customer            *Customer            `xml:"customer,omitempty"`
	BillingAddress      *Address             `xml:"billing,omitempty"`
	ShippingAddress     *Address             `xml:"shipping,omitempty"`
	Options             *TransactionOptions  `xml:"options,omitempty"`
	ServiceFeeAmount    *Decimal             `xml:"service-fee-amount,attr,omitempty"`
	CreatedAt           *time.Time           `xml:"created-at,omitempty"`
	UpdatedAt           *time.Time           `xml:"updated-at,omitempty"`
	DisbursementDetails *DisbursementDetails `xml:"disbursement-details,omitempty"`
}

// TODO: not all transaction fields are implemented yet, here are the missing fields (add on demand)
//
// <transaction>
//   <currency-iso-code>USD</currency-iso-code>
//   <refund-id nil="true"></refund-id>
//   <refund-ids type="array"/>
//   <refunded-transaction-id nil="true"></refunded-transaction-id>
//   <settlement-batch-id>2013-10-08_49grybq7pbtsnvsr</settlement-batch-id>
//   <custom-fields>
//   </custom-fields>
//   <avs-error-response-code nil="true"></avs-error-response-code>
//   <avs-postal-code-response-code>I</avs-postal-code-response-code>
//   <avs-street-address-response-code>I</avs-street-address-response-code>
//   <cvv-response-code>I</cvv-response-code>
//   <gateway-rejection-reason nil="true"></gateway-rejection-reason>
//   <processor-authorization-code>YCSBWR</processor-authorization-code>
//   <processor-response-code>1000</processor-response-code>
//   <processor-response-text>Approved</processor-response-text>
//   <voice-referral-number nil="true"></voice-referral-number>
//   <purchase-order-number nil="true"></purchase-order-number>
//   <tax-amount nil="true"></tax-amount>
//   <tax-exempt type="boolean">false</tax-exempt>
//   <status-history type="array">
//     <status-event>
//       <timestamp type="datetime">2013-10-07T17:26:14Z</timestamp>
//       <status>authorized</status>
//       <amount>7.00</amount>
//       <user>eaigner</user>
//       <transaction-source>Recurring</transaction-source>
//     </status-event>
//     <status-event>
//       <timestamp type="datetime">2013-10-07T17:26:14Z</timestamp>
//       <status>submitted_for_settlement</status>
//       <amount>7.00</amount>
//       <user>eaigner</user>
//       <transaction-source>Recurring</transaction-source>
//     </status-event>
//     <status-event>
//       <timestamp type="datetime">2013-10-08T07:06:38Z</timestamp>
//       <status>settled</status>
//       <amount>7.00</amount>
//       <user nil="true"></user>
//       <transaction-source></transaction-source>
//     </status-event>
//   </status-history>
//   <plan-id>bronze</plan-id>
//   <subscription-id>jqsydb</subscription-id>
//   <subscription>
//     <billing-period-end-date type="date">2013-11-06</billing-period-end-date>
//     <billing-period-start-date type="date">2013-10-07</billing-period-start-date>
//   </subscription>
//   <add-ons type="array"/>
//   <discounts type="array"/>
//   <descriptor>
//     <name nil="true"></name>
//     <phone nil="true"></phone>
//   </descriptor>
//   <recurring type="boolean">true</recurring>
//   <channel nil="true"></channel>
//   <escrow-status nil="true"></escrow-status>
// </transaction>

type Transactions struct {
	Transaction []*Transaction `xml:"transaction"`
}

type TransactionOptions struct {
	SubmitForSettlement              bool `xml:"submit-for-settlement,omitempty"`
	StoreInVault                     bool `xml:"store-in-vault,omitempty"`
	AddBillingAddressToPaymentMethod bool `xml:"add-billing-address-to-payment-method,omitempty"`
	StoreShippingAddressInVault      bool `xml:"store-shipping-address-in-vault,omitempty"`
}

type TransactionSearchResult struct {
	XMLName           string              `xml:"credit-card-transactions"`
	CurrentPageNumber *nullable.NullInt64 `xml:"current-page-number"`
	PageSize          *nullable.NullInt64 `xml:"page-size"`
	TotalItems        *nullable.NullInt64 `xml:"total-items"`
	Transactions      []*Transaction      `xml:"transaction"`
}
