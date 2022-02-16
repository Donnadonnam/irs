/*
 * IRS API
 *
 * Package github.com/moov-io/irs implements a file reader and writer written in Go along with a HTTP API and CLI for creating, parsing, validating, and transforming IRS electronic Filing Information Returns Electronically (FIRE). FIRE operates on a byte(ASCII) level making it difficult to interface with JSON and CSV/TEXT file formats.  | Input      | Output     |  |------------|------------|  | JSON       | JSON       |  | ASCII FIRE | ASCII FIRE |  |            | PDF Form   |  |            | SQL        |
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	"time"
)

// BRecordWith1099Ltc struct for BRecordWith1099Ltc
type BRecordWith1099Ltc struct {
	RecordType                  string    `json:"record_type"`
	PaymentYear                 int32     `json:"payment_year"`
	CorrectedReturnIndicator    string    `json:"corrected_return_indicator,omitempty"`
	PayeesNameControl           string    `json:"payees_name_control,omitempty"`
	TypeOfTin                   string    `json:"type_of_tin,omitempty"`
	PayeesTin                   string    `json:"payees_tin"`
	PayersAccountNumberForPayee string    `json:"payers_account_number_for_payee,omitempty"`
	PayersOfficeCode            string    `json:"payers_office_code,omitempty"`
	PaymentAmount1              int32     `json:"payment_amount_1,omitempty"`
	PaymentAmount2              int32     `json:"payment_amount_2,omitempty"`
	PaymentAmount3              int32     `json:"payment_amount_3,omitempty"`
	PaymentAmount4              int32     `json:"payment_amount_4,omitempty"`
	PaymentAmount5              int32     `json:"payment_amount_5,omitempty"`
	PaymentAmount6              int32     `json:"payment_amount_6,omitempty"`
	PaymentAmount7              int32     `json:"payment_amount_7,omitempty"`
	PaymentAmount8              int32     `json:"payment_amount_8,omitempty"`
	PaymentAmount9              int32     `json:"payment_amount_9,omitempty"`
	PaymentAmountA              int32     `json:"payment_amount_A,omitempty"`
	PaymentAmountB              int32     `json:"payment_amount_B,omitempty"`
	PaymentAmountC              int32     `json:"payment_amount_C,omitempty"`
	PaymentAmountD              int32     `json:"payment_amount_D,omitempty"`
	PaymentAmountE              int32     `json:"payment_amount_E,omitempty"`
	PaymentAmountF              int32     `json:"payment_amount_F,omitempty"`
	PaymentAmountG              int32     `json:"payment_amount_G,omitempty"`
	PaymentAmountH              int32     `json:"payment_amount_H,omitempty"`
	PaymentAmountJ              int32     `json:"payment_amount_J,omitempty"`
	ForeignCountryIndicator     string    `json:"foreign_country_indicator,omitempty"`
	FirstPayeeNameLine          string    `json:"first_payee_name_line"`
	SecondPayeeNameLine         string    `json:"second_payee_name_line,omitempty"`
	PayeeMailingAddress         string    `json:"payee_mailing_address"`
	PayeeCity                   string    `json:"payee_city"`
	PayeeState                  string    `json:"payee_state"`
	PayeeZipCode                string    `json:"payee_zip_code"`
	RecordSequenceNumber        int32     `json:"record_sequence_number"`
	TypePaymentIndicator        string    `json:"type_payment_indicator,omitempty"`
	NameInsured                 string    `json:"name_insured,omitempty"`
	AddressInsured              string    `json:"address_insured,omitempty"`
	CityInsured                 string    `json:"city_insured,omitempty"`
	StateInsured                string    `json:"state_insured,omitempty"`
	ZipCodeInsured              string    `json:"zip_code_insured"`
	StatusIllnessIndicator      string    `json:"status_illness_indicator,omitempty"`
	DateCertified               time.Time `json:"date_certified,omitempty"`
	QualifiedContractIndicator  string    `json:"qualified_contract_indicator,omitempty"`
	StateIncomeTaxWithheld      int32     `json:"state_income_tax_withheld,omitempty"`
	LocalIncomeTaxWithheld      int32     `json:"local_income_tax_withheld,omitempty"`
}
