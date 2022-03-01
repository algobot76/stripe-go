//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72/form"
)

// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
type SubscriptionSchedulePhaseBillingCycleAnchor string

// List of values that SubscriptionSchedulePhaseBillingCycleAnchor can take
const (
	SubscriptionSchedulePhaseBillingCycleAnchorAutomatic  SubscriptionSchedulePhaseBillingCycleAnchor = "automatic"
	SubscriptionSchedulePhaseBillingCycleAnchorPhaseStart SubscriptionSchedulePhaseBillingCycleAnchor = "phase_start"
)

// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` and `cancel`.
type SubscriptionScheduleEndBehavior string

// List of values that SubscriptionScheduleEndBehavior can take
const (
	SubscriptionScheduleEndBehaviorCancel  SubscriptionScheduleEndBehavior = "cancel"
	SubscriptionScheduleEndBehaviorNone    SubscriptionScheduleEndBehavior = "none"
	SubscriptionScheduleEndBehaviorRelease SubscriptionScheduleEndBehavior = "release"
	SubscriptionScheduleEndBehaviorRenew   SubscriptionScheduleEndBehavior = "renew"
)

// If the subscription schedule will prorate when transitioning to this phase. Possible values are `create_prorations` and `none`.
type SubscriptionSchedulePhaseProrationBehavior string

// List of values that SubscriptionSchedulePhaseProrationBehavior can take
const (
	SubscriptionSchedulePhaseProrationBehaviorAlwaysInvoice    SubscriptionSchedulePhaseProrationBehavior = "always_invoice"
	SubscriptionSchedulePhaseProrationBehaviorCreateProrations SubscriptionSchedulePhaseProrationBehavior = "create_prorations"
	SubscriptionSchedulePhaseProrationBehaviorNone             SubscriptionSchedulePhaseProrationBehavior = "none"
)

// The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
type SubscriptionScheduleStatus string

// List of values that SubscriptionScheduleStatus can take
const (
	SubscriptionScheduleStatusActive     SubscriptionScheduleStatus = "active"
	SubscriptionScheduleStatusCanceled   SubscriptionScheduleStatus = "canceled"
	SubscriptionScheduleStatusCompleted  SubscriptionScheduleStatus = "completed"
	SubscriptionScheduleStatusNotStarted SubscriptionScheduleStatus = "not_started"
	SubscriptionScheduleStatusReleased   SubscriptionScheduleStatus = "released"
)

// Retrieves the list of your subscription schedules.
type SubscriptionScheduleListParams struct {
	ListParams `form:"*"`
	// Only return subscription schedules that were created canceled the given date interval.
	CanceledAt int64 `form:"canceled_at"`
	// Only return subscription schedules that were created canceled the given date interval.
	CanceledAtRange *RangeQueryParams `form:"canceled_at"`
	// Only return subscription schedules that completed during the given date interval.
	CompletedAt int64 `form:"completed_at"`
	// Only return subscription schedules that completed during the given date interval.
	CompletedAtRange *RangeQueryParams `form:"completed_at"`
	// Only return subscription schedules that were created during the given date interval.
	Created int64 `form:"created"`
	// Only return subscription schedules that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return subscription schedules for the given customer.
	Customer string `form:"customer"`
	// Only return subscription schedules that were released during the given date interval.
	ReleasedAt int64 `form:"released_at"`
	// Only return subscription schedules that were released during the given date interval.
	ReleasedAtRange *RangeQueryParams `form:"released_at"`
	// Only return subscription schedules that have not started yet.
	Scheduled *bool `form:"scheduled"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleInvoiceSettingsParams struct {
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due"`
}

// Object representing the subscription schedule's default settings.
type SubscriptionScheduleDefaultSettingsParams struct {
	Params `form:"*"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent,high_precision"`
	// Default settings for automatic tax computation.
	AutomaticTax *SubscriptionAutomaticTaxParams `form:"automatic_tax"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionBillingThresholdsParams `form:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionScheduleInvoiceSettingsParams `form:"invoice_settings"`
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams `form:"transfer_data"`
}

// SubscriptionSchedulePhaseAddInvoiceItemPriceDataRecurringParams is a structure representing the
// parameters to create an inline recurring price for a given invoice item.
type SubscriptionSchedulePhaseAddInvoiceItemPriceDataRecurringParams struct {
	AggregateUsage  *string `form:"aggregate_usage"`
	Interval        *string `form:"interval"`
	IntervalCount   *int64  `form:"interval_count"`
	TrialPeriodDays *int64  `form:"trial_period_days"`
	UsageType       *string `form:"usage_type"`
}

// SubscriptionSchedulePhaseAddInvoiceItemPriceDataParams is a structure representing the parameters
// to create an inline price for a given invoice item.
type SubscriptionSchedulePhaseAddInvoiceItemPriceDataParams struct {
	Currency          *string                                                          `form:"currency"`
	Product           *string                                                          `form:"product"`
	Recurring         *SubscriptionSchedulePhaseAddInvoiceItemPriceDataRecurringParams `form:"recurring"`
	TaxBehavior       *string                                                          `form:"tax_behavior"`
	UnitAmount        *int64                                                           `form:"unit_amount"`
	UnitAmountDecimal *float64                                                         `form:"unit_amount_decimal,high_precision"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice. You may pass up to 20 items.
type SubscriptionSchedulePhaseAddInvoiceItemParams struct {
	// The ID of the price object.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
	PriceData *InvoiceItemPriceDataParams `form:"price_data"`
	// Quantity for this item. Defaults to 1.
	Quantity *int64 `form:"quantity"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates []*string `form:"tax_rates"`
}

// Automatic tax settings for this phase.
type SubscriptionSchedulePhaseAutomaticTaxParams struct {
	// Enabled automatic tax calculation which will automatically compute tax rates on all invoices generated by the subscription.
	Enabled *bool `form:"enabled"`
}

// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
type SubscriptionSchedulePhaseItemParams struct {
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. When updating, pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionItemBillingThresholdsParams `form:"billing_thresholds"`
	// The plan ID to subscribe to. You may specify the same ID in `plan` and `price`.
	Plan *string `form:"plan"`
	// The ID of the price object.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
	PriceData *SubscriptionItemPriceDataParams `form:"price_data"`
	// Quantity for the given price. Can be set only if the price's `usage_type` is `licensed` and not `metered`.
	Quantity *int64 `form:"quantity"`
	// A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates"`
}

// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase.
type SubscriptionSchedulePhaseParams struct {
	// A list of prices and quantities that will generate invoice items appended to the next invoice. You may pass up to 20 items.
	AddInvoiceItems []*SubscriptionSchedulePhaseAddInvoiceItemParams `form:"add_invoice_items"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent"`
	// Automatic tax settings for this phase.
	AutomaticTax *SubscriptionSchedulePhaseAutomaticTaxParams `form:"automatic_tax"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionBillingThresholdsParams `form:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method"`
	// The identifier of the coupon to apply to this phase of the subscription schedule.
	Coupon *string `form:"coupon"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method"`
	// A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will set the Subscription's [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates), which means they will be the Invoice's [`default_tax_rates`](https://stripe.com/docs/api/invoices/create#create_invoice-default_tax_rates) for any Invoices issued by the Subscription during this Phase.
	DefaultTaxRates []*string `form:"default_tax_rates"`
	// The date at which this phase of the subscription schedule ends. If set, `iterations` must not be set.
	EndDate    *int64 `form:"end_date"`
	EndDateNow *bool  `form:"-"` // See custom AppendTo
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionScheduleInvoiceSettingsParams `form:"invoice_settings"`
	// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
	Items []*SubscriptionSchedulePhaseItemParams `form:"items"`
	// Integer representing the multiplier applied to the price interval. For example, `iterations=2` applied to a price with `interval=month` and `interval_count=3` results in a phase of duration `2 * 3 months = 6 months`. If set, `end_date` must not be set.
	Iterations *int64 `form:"iterations"`
	// If a subscription schedule will create prorations when transitioning to this phase. Possible values are `create_prorations` or `none`, and the default value is `create_prorations`. See [Prorations](https://stripe.com/docs/billing/subscriptions/prorations).
	ProrationBehavior *string `form:"proration_behavior"`
	// The date at which this phase of the subscription schedule starts or `now`. Must be set on the first phase.
	StartDate    *int64 `form:"start_date"`
	StartDateNow *bool  `form:"-"` // See custom AppendTo
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams `form:"transfer_data"`
	// If set to true the entire phase is counted as a trial and the customer will not be charged for any fees.
	Trial *bool `form:"trial"`
	// Sets the phase to trialing from the start date to this date. Must be before the phase end date, can not be combined with `trial`
	TrialEnd    *int64 `form:"trial_end"`
	TrialEndNow *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for SubscriptionSchedulePhaseParams.
func (s *SubscriptionSchedulePhaseParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(s.EndDateNow) {
		body.Add(form.FormatKey(append(keyParts, "end_date")), "now")
	}
	if BoolValue(s.TrialEndNow) {
		body.Add(form.FormatKey(append(keyParts, "trial_end")), "now")
	}
	if BoolValue(s.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
}

// Creates a new subscription schedule object. Each customer can have up to 500 active or scheduled subscriptions.
type SubscriptionScheduleParams struct {
	Params `form:"*"`
	// The identifier of the customer to create the subscription schedule for.
	Customer *string `form:"customer"`
	// Object representing the subscription schedule's default settings.
	DefaultSettings *SubscriptionScheduleDefaultSettingsParams `form:"default_settings"`
	// Configures how the subscription schedule behaves when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running.`cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior *string `form:"end_behavior"`
	// Migrate an existing subscription to be managed by a subscription schedule. If this parameter is set, a subscription schedule will be created using the subscription's item(s), set to auto-renew using the subscription's interval. When using this parameter, other parameters (such as phase values) cannot be set. To create a subscription schedule with other modifications, we recommend making two separate API calls.
	FromSubscription *string `form:"from_subscription"`
	// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase. Note that past phases can be omitted.
	Phases []*SubscriptionSchedulePhaseParams `form:"phases"`
	// If the update changes the current phase, indicates if the changes should be prorated. Possible values are `create_prorations` or `none`, and the default value is `create_prorations`.
	ProrationBehavior *string `form:"proration_behavior"`
	// When the subscription schedule starts. We recommend using `now` so that it starts the subscription immediately. You can also use a Unix timestamp to backdate the subscription so that it starts on a past date, or set a future date for the subscription to start on.
	StartDate    *int64 `form:"start_date"`
	StartDateNow *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for SubscriptionScheduleParams.
func (s *SubscriptionScheduleParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(s.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
}

// Cancels a subscription schedule and its associated subscription immediately (if the subscription schedule has an active subscription). A subscription schedule can only be canceled if its status is not_started or active.
type SubscriptionScheduleCancelParams struct {
	Params `form:"*"`
	// If the subscription schedule is `active`, indicates if a final invoice will be generated that contains any un-invoiced metered usage and new/pending proration invoice items. Defaults to `true`.
	InvoiceNow *bool `form:"invoice_now"`
	// If the subscription schedule is `active`, indicates if the cancellation should be prorated. Defaults to `true`.
	Prorate *bool `form:"prorate"`
}

// Releases the subscription schedule immediately, which will stop scheduling of its phases, but leave any existing subscription in place. A schedule can only be released if its status is not_started or active. If the subscription schedule is currently associated with a subscription, releasing it will remove its subscription property and set the subscription's ID to the released_subscription property.
type SubscriptionScheduleReleaseParams struct {
	Params `form:"*"`
	// Keep any cancellation on the subscription that the schedule has set
	PreserveCancelDate *bool `form:"preserve_cancel_date"`
}

// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
type SubscriptionScheduleCurrentPhase struct {
	// The end of this phase of the subscription schedule.
	EndDate int64 `json:"end_date"`
	// The start of this phase of the subscription schedule.
	StartDate int64 `json:"start_date"`
}

// The subscription schedule's default invoice settings.
type SubscriptionScheduleInvoiceSettings struct {
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int64 `json:"days_until_due"`
}
type SubscriptionScheduleDefaultSettings struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner's Stripe account during this phase of the schedule.
	ApplicationFeePercent float64                   `json:"application_fee_percent,string"`
	AutomaticTax          *SubscriptionAutomaticTax `json:"automatic_tax"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor SubscriptionSchedulePhaseBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
	BillingThresholds *SubscriptionBillingThresholds `json:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions.
	CollectionMethod SubscriptionCollectionMethod `json:"collection_method"`
	// ID of the default payment method for the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// The subscription schedule's default invoice settings.
	InvoiceSettings *SubscriptionScheduleInvoiceSettings `json:"invoice_settings"`
	// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
	TransferData *SubscriptionTransferData `json:"transfer_data"`
}

// A list of prices and quantities that will generate invoice items appended to the first invoice for this phase.
type SubscriptionSchedulePhaseAddInvoiceItem struct {
	// ID of the price used to generate the invoice item.
	Price *Price `json:"price"`
	// The quantity of the invoice item.
	Quantity int64 `json:"quantity"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates []*TaxRate `json:"tax_rates"`
}

// Subscription items to configure the subscription to during this phase of the subscription schedule.
type SubscriptionSchedulePhaseItem struct {
	// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
	BillingThresholds *SubscriptionItemBillingThresholds `json:"billing_thresholds"`
	// ID of the plan to which the customer should be subscribed.
	Plan *Plan `json:"plan"`
	// ID of the price to which the customer should be subscribed.
	Price *Price `json:"price"`
	// Quantity of the plan to which the customer should be subscribed.
	Quantity int64 `json:"quantity"`
	// The tax rates which apply to this `phase_item`. When set, the `default_tax_rates` on the phase do not apply to this `phase_item`.
	TaxRates []*TaxRate `json:"tax_rates"`
}

// Configuration for the subscription schedule's phases.
type SubscriptionSchedulePhase struct {
	// A list of prices and quantities that will generate invoice items appended to the first invoice for this phase.
	AddInvoiceItems []*SubscriptionSchedulePhaseAddInvoiceItem `json:"add_invoice_items"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner's Stripe account during this phase of the schedule.
	ApplicationFeePercent float64                   `json:"application_fee_percent"`
	AutomaticTax          *SubscriptionAutomaticTax `json:"automatic_tax"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor SubscriptionSchedulePhaseBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
	BillingThresholds *SubscriptionBillingThresholds `json:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions.
	CollectionMethod SubscriptionCollectionMethod `json:"collection_method"`
	// ID of the coupon to use during this phase of the subscription schedule.
	Coupon *Coupon `json:"coupon"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// The default tax rates to apply to the subscription during this phase of the subscription schedule.
	DefaultTaxRates []*TaxRate `json:"default_tax_rates"`
	// The end of this phase of the subscription schedule.
	EndDate int64 `json:"end_date"`
	// The invoice settings applicable during this phase.
	InvoiceSettings *SubscriptionScheduleInvoiceSettings `json:"invoice_settings"`
	// Subscription items to configure the subscription to during this phase of the subscription schedule.
	Items []*SubscriptionSchedulePhaseItem `json:"items"`
	// If the subscription schedule will prorate when transitioning to this phase. Possible values are `create_prorations` and `none`.
	ProrationBehavior SubscriptionSchedulePhaseProrationBehavior `json:"proration_behavior"`
	// The start of this phase of the subscription schedule.
	StartDate int64 `json:"start_date"`
	// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
	TransferData *SubscriptionTransferData `json:"transfer_data"`
	// When the trial ends within the phase.
	TrialEnd int64 `json:"trial_end"`
}

// A subscription schedule allows you to create and manage the lifecycle of a subscription by predefining expected changes.
//
// Related guide: [Subscription Schedules](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
type SubscriptionSchedule struct {
	APIResource
	// Time at which the subscription schedule was canceled. Measured in seconds since the Unix epoch.
	CanceledAt int64 `json:"canceled_at"`
	// Time at which the subscription schedule was completed. Measured in seconds since the Unix epoch.
	CompletedAt int64 `json:"completed_at"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
	CurrentPhase *SubscriptionScheduleCurrentPhase `json:"current_phase"`
	// ID of the customer who owns the subscription schedule.
	Customer        *Customer                            `json:"customer"`
	DefaultSettings *SubscriptionScheduleDefaultSettings `json:"default_settings"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` and `cancel`.
	EndBehavior SubscriptionScheduleEndBehavior `json:"end_behavior"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Configuration for the subscription schedule's phases.
	Phases []*SubscriptionSchedulePhase `json:"phases"`
	// Time at which the subscription schedule was released. Measured in seconds since the Unix epoch.
	ReleasedAt int64 `json:"released_at"`
	// ID of the subscription once managed by the subscription schedule (if it is released).
	ReleasedSubscription *Subscription `json:"released_subscription"`
	// The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
	Status SubscriptionScheduleStatus `json:"status"`
	// ID of the subscription managed by the subscription schedule.
	Subscription *Subscription `json:"subscription"`
	// ID of the test clock this subscription schedule belongs to.
	TestClock *TestHelpersTestClock `json:"test_clock"`
}

// SubscriptionScheduleList is a list of SubscriptionSchedules as retrieved from a list endpoint.
type SubscriptionScheduleList struct {
	APIResource
	ListMeta
	Data []*SubscriptionSchedule `json:"data"`
}

// UnmarshalJSON handles deserialization of a SubscriptionSchedule.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SubscriptionSchedule) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type subscriptionSchedule SubscriptionSchedule
	var v subscriptionSchedule
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SubscriptionSchedule(v)
	return nil
}
