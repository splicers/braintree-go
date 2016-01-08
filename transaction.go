package braintree

import (
	"time"

	"github.com/splicers/braintree-go/nullable"
)

type Transaction struct {
	Amount                *Decimal             `xml:"amount"`
	BillingAddress        *Address             `xml:"billing,omitempty"`
	CreatedAt             *time.Time           `xml:"created-at,omitempty"`
	CreditCard            *CreditCard          `xml:"credit-card,omitempty"`
	Customer              *Customer            `xml:"customer,omitempty"`
	CustomerID            string               `xml:"customer-id,omitempty"`
	DisbursementDetails   *DisbursementDetails `xml:"disbursement-details,omitempty"`
	Id                    string               `xml:"id,omitempty"`
	MerchantAccountId     string               `xml:"merchant-account-id,omitempty"`
	Options               *TransactionOptions  `xml:"options,omitempty"`
	OrderId               string               `xml:"order-id,omitempty"`
	PaymentMethodNonce    string               `xml:"payment-method-nonce,omitempty"`
	PaymentMethodToken    string               `xml:"payment-method-token,omitempty"`
	PaypalTransaction     *PaypalTransaction   `xml:"paypal,omitempty"`
	PlanId                string               `xml:"plan-id,omitempty"`
	RefundIds             *[]string            `xml:"refund-ids>item,omitempty"`
	RefundedTransactionId *string              `xml:"refunded-transaction-id,omitempty"`
	Subscription          *SubscriptionDetails `xml:"subscription,omitempty"`
	SubscriptionId        string               `xml:"subscription-id,omitempty"`
	ServiceFeeAmount      *Decimal             `xml:"service-fee-amount,attr,omitempty"`
	ShippingAddress       *Address             `xml:"shipping,omitempty"`
	Status                string               `xml:"status,omitempty"`
	TaxAmount             *Decimal             `xml:"tax-amount,omitempty"`
	Type                  string               `xml:"type,omitempty"`
	UpdatedAt             *time.Time           `xml:"updated-at,omitempty"`
	XMLName               string               `xml:"transaction"`
}

type SubscriptionDetails struct {
	BillingPeriodEndDate   string `xml:"billing-period-end-date,omitempty"`
	BillingPeriodStartDate string `xml:"billing-period-start-date,omitempty"`
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
