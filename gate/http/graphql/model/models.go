// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Respond interface {
	IsRespond()
}

type RespondWithAddress interface {
	IsRespondWithAddress()
}

type RespondWithAnything interface {
	IsRespondWithAnything()
}

type RespondWithBase interface {
	IsRespondWithBase()
}

type RespondWithSecuritySetting interface {
	IsRespondWithSecuritySetting()
}

type RespondWithTransaction interface {
	IsRespondWithTransaction()
}

type RespondWithTransactions interface {
	IsRespondWithTransactions()
}

type RespondWithUser interface {
	IsRespondWithUser()
}

type RespondWithUsers interface {
	IsRespondWithUsers()
}

type RespondWithWallet interface {
	IsRespondWithWallet()
}

type RespondWithWallets interface {
	IsRespondWithWallets()
}

// the address type
type Address struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	User        *User      `json:"user"`
	State       string     `json:"state"`
	City        string     `json:"city"`
	HouseNumber *string    `json:"house_number,omitempty"`
	Street      string     `json:"street"`
	Zip         string     `json:"zip"`
	Verified    bool       `json:"verified"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (Address) IsEntity() {}

type Attachment struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type AttachmentInput struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Base struct {
	// Base is a type describing an arbitrary set of keys.
	Exist *bool `json:"exist,omitempty"`
}

type EmailConfirmationInput struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type EmailVerificationInput struct {
	Email string `json:"email"`
}

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Status  int    `json:"status"`
}

func (Error) IsRespond() {}

func (Error) IsRespondWithBase() {}

func (Error) IsRespondWithAddress() {}

func (Error) IsRespondWithWallet() {}

func (Error) IsRespondWithWallets() {}

func (Error) IsRespondWithSecuritySetting() {}

func (Error) IsRespondWithUser() {}

func (Error) IsRespondWithUsers() {}

func (Error) IsRespondWithAnything() {}

func (Error) IsRespondWithTransaction() {}

func (Error) IsRespondWithTransactions() {}

type History struct {
	Act string `json:"act"`
	By  string `json:"by"`
	At  string `json:"at"`
}

type Item struct {
	Key      string  `json:"key"`
	Name     string  `json:"name"`
	Amount   float64 `json:"amount"`
	Quantity int     `json:"quantity"`
	Metadata *string `json:"metadata,omitempty"`
}

type Mutation struct {
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Pages int `json:"pages"`
}

type PhoneConfirmationInput struct {
	// Phone must be in international format without + (E.g 2348012345678).
	Phone     string `json:"phone"`
	Code      string `json:"code"`
	Reference string `json:"reference"`
}

type PhoneVerificationInput struct {
	Phone string `json:"phone"`
}

type Query struct {
}

type Response struct {
	Message string `json:"message"`
}

func (Response) IsRespond() {}

type ResponseWithAddress struct {
	Message string   `json:"message"`
	Data    *Address `json:"data,omitempty"`
}

func (ResponseWithAddress) IsRespondWithAddress() {}

type ResponseWithAnything struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (ResponseWithAnything) IsRespondWithAnything() {}

type ResponseWithBase struct {
	Message string `json:"message"`
	Data    *Base  `json:"data,omitempty"`
}

func (ResponseWithBase) IsRespondWithBase() {}

type ResponseWithSecuritySetting struct {
	Message string           `json:"message"`
	Data    *SecuritySetting `json:"data,omitempty"`
}

func (ResponseWithSecuritySetting) IsRespondWithSecuritySetting() {}

type ResponseWithTransaction struct {
	Message string       `json:"message"`
	Data    *Transaction `json:"data,omitempty"`
}

func (ResponseWithTransaction) IsRespondWithTransaction() {}

type ResponseWithTransactions struct {
	Message string         `json:"message"`
	Data    []*Transaction `json:"data,omitempty"`
	// only returned when paginate is true
	Pagination *Pagination `json:"pagination,omitempty"`
}

func (ResponseWithTransactions) IsRespondWithTransactions() {}

type ResponseWithUser struct {
	Message string `json:"message"`
	Data    *User  `json:"data,omitempty"`
}

func (ResponseWithUser) IsRespondWithUser() {}

type ResponseWithUsers struct {
	Message string  `json:"message"`
	Data    []*User `json:"data,omitempty"`
	// only returned when paginate is true
	Pagination *Pagination `json:"pagination,omitempty"`
}

func (ResponseWithUsers) IsRespondWithUsers() {}

type ResponseWithWallet struct {
	Message string  `json:"message"`
	Data    *Wallet `json:"data,omitempty"`
}

func (ResponseWithWallet) IsRespondWithWallet() {}

type ResponseWithWallets struct {
	Message string    `json:"message"`
	Data    []*Wallet `json:"data,omitempty"`
}

func (ResponseWithWallets) IsRespondWithWallets() {}

// the security_setting type
type SecuritySetting struct {
	ID               string     `json:"id"`
	UserID           string     `json:"user_id"`
	User             *User      `json:"user"`
	BiometricEnabled bool       `json:"biometric_enabled"`
	PrivateMode      bool       `json:"private_mode"`
	AllowScreenshot  bool       `json:"allow_screenshot"`
	ContextMenu      bool       `json:"context_menu"`
	SmsAlert         bool       `json:"sms_alert"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}

func (SecuritySetting) IsEntity() {}

type SecuritySettingInput struct {
	BiometricEnabled bool `json:"biometric_enabled"`
	PrivateMode      bool `json:"private_mode"`
	AllowScreenshot  bool `json:"allow_screenshot"`
	ContextMenu      bool `json:"context_menu"`
	SmsAlert         bool `json:"sms_alert"`
}

type SetPasscodeInput struct {
	Passcode string  `json:"passcode"`
	Pastcode *string `json:"pastcode,omitempty"`
}

type SigninInput struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type StartSessionInput struct {
	Email    string `json:"email"`
	Passcode string `json:"passcode"`
}

type StepOneInput struct {
	Email        string  `json:"email"`
	Country      string  `json:"country"`
	ReferralCode *string `json:"referral_code,omitempty"`
}

type StepThreeInput struct {
	State       string  `json:"state"`
	City        string  `json:"city"`
	HouseNumber *string `json:"house_number,omitempty"`
	Street      string  `json:"street"`
	Zip         string  `json:"zip"`
}

type StepTwoInput struct {
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Middlename *string   `json:"middlename,omitempty"`
	Gender     EGender   `json:"gender"`
	Dob        time.Time `json:"dob"`
}

// the tx type
type Transaction struct {
	ID         string `json:"id"`
	SavingsID  string `json:"savings_id"`
	ProductID  string `json:"product_id"`
	CustomerID int    `json:"customer_id"`
	// initiator
	UserID        int                    `json:"user_id"`
	Type          ETransactionType       `json:"type"`
	Reference     string                 `json:"reference"`
	Currency      ECurrency              `json:"currency"`
	Method        EPaymentMethod         `json:"method"`
	Paid          bool                   `json:"paid"`
	PaidAt        *time.Time             `json:"paid_at,omitempty"`
	Failed        bool                   `json:"failed"`
	FailedAt      *time.Time             `json:"failed_at,omitempty"`
	Cancelled     bool                   `json:"cancelled"`
	CancelledAt   *time.Time             `json:"cancelled_at,omitempty"`
	PaymentLink   *string                `json:"payment_link,omitempty"`
	InitialConfig map[string]interface{} `json:"initial_config,omitempty"`
	FinalConfig   map[string]interface{} `json:"final_config,omitempty"`
	History       []*History             `json:"history,omitempty"`
	Invoice       []*Item                `json:"invoice,omitempty"`
	SavingsAmount float64                `json:"savings_amount"`
	Amount        float64                `json:"amount"`
	PreBalance    float64                `json:"pre_balance"`
	PostBalance   float64                `json:"post_balance"`
	Remark        *string                `json:"remark,omitempty"`
	CreatedAt     *time.Time             `json:"created_at,omitempty"`
	UpdatedAt     *time.Time             `json:"updated_at,omitempty"`
}

type TransactionFilterInput struct {
	TransactionID *string `json:"transaction_id,omitempty"`
	SavingsID     *string `json:"savings_id,omitempty"`
	UserID        *int    `json:"user_id,omitempty"`
	ProductID     *string `json:"product_id,omitempty"`
	// organization id
	CustomerID *int              `json:"customer_id,omitempty"`
	Type       *ETransactionType `json:"type,omitempty"`
	Reference  *string           `json:"reference,omitempty"`
	Currency   *ECurrency        `json:"currency,omitempty"`
	Method     *EPaymentMethod   `json:"method,omitempty"`
	Paid       *bool             `json:"paid,omitempty"`
	Failed     *bool             `json:"failed,omitempty"`
	Cancelled  *bool             `json:"cancelled,omitempty"`
	// searches by remark, reference, transaction_id
	Search *string `json:"search,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
	Page   *int    `json:"page,omitempty"`
	Sort   *ESort  `json:"sort,omitempty"`
	// when true, returns page count and related info (try not to set to true always)
	Paginate *bool `json:"paginate,omitempty"`
}

type TransactionInput struct {
	AccountNumber string           `json:"account_number"`
	UserID        int              `json:"user_id"`
	Date          time.Time        `json:"date"`
	Type          ETransactionType `json:"type"`
	Amount        float64          `json:"amount"`
	Remark        *string          `json:"remark,omitempty"`
}

// the user type
type User struct {
	ID         string     `json:"id"`
	Firstname  *string    `json:"firstname,omitempty"`
	Middlename *string    `json:"middlename,omitempty"`
	Lastname   *string    `json:"lastname,omitempty"`
	Email      string     `json:"email"`
	Phone      *string    `json:"phone,omitempty"`
	PhoneCode  *string    `json:"phone_code,omitempty"`
	Gender     *string    `json:"gender,omitempty"`
	Country    *string    `json:"country,omitempty"`
	Username   *string    `json:"username,omitempty"`
	Language   *ELanguage `json:"language,omitempty"`
	Tier       *ETier     `json:"tier,omitempty"`
	Dob        *time.Time `json:"dob,omitempty"`
	Passcode   *string    `json:"passcode,omitempty"`
	Referral   *string    `json:"referral,omitempty"`
	Avatar     *string    `json:"avatar,omitempty"`
	// defines whether or not the user is an internal team member
	Internal        *bool            `json:"internal,omitempty"`
	Role            *ERole           `json:"role,omitempty"`
	Status          *EUserStatus     `json:"status,omitempty"`
	Address         *Address         `json:"address,omitempty"`
	SecuritySetting *SecuritySetting `json:"security_setting,omitempty"`
	CreatedAt       *time.Time       `json:"created_at,omitempty"`
	UpdatedAt       *time.Time       `json:"updated_at,omitempty"`
}

func (User) IsEntity() {}

// the wallet type
type Wallet struct {
	ID        string           `json:"id"`
	UserID    string           `json:"user_id"`
	User      *User            `json:"user"`
	Type      EWalletType      `json:"type"`
	Name      string           `json:"name"`
	Balance   float64          `json:"balance"`
	Currency  ECurrency        `json:"currency"`
	Address   []*WalletAddress `json:"address,omitempty"`
	CreatedAt *time.Time       `json:"created_at,omitempty"`
	UpdatedAt *time.Time       `json:"updated_at,omitempty"`
}

func (Wallet) IsEntity() {}

type WalletAddress struct {
	Address string         `json:"address"`
	Network ECryptoNetwork `json:"network"`
}

type WalletFilter struct {
	Type *EWalletType `json:"type,omitempty"`
}

type WalletInput struct {
	Currency ECurrency `json:"currency"`
	// this is not required and really only added for internal use
	Code *string `json:"code,omitempty"`
}

type WebhookInput struct {
	Tx        any    `json:"tx"`
	Reference string `json:"reference"`
}

type ECryptoNetwork string

const (
	ECryptoNetworkEthereum ECryptoNetwork = "ETHEREUM"
	ECryptoNetworkBitcoin  ECryptoNetwork = "BITCOIN"
	ECryptoNetworkTron     ECryptoNetwork = "TRON"
	ECryptoNetworkSolana   ECryptoNetwork = "SOLANA"
)

var AllECryptoNetwork = []ECryptoNetwork{
	ECryptoNetworkEthereum,
	ECryptoNetworkBitcoin,
	ECryptoNetworkTron,
	ECryptoNetworkSolana,
}

func (e ECryptoNetwork) IsValid() bool {
	switch e {
	case ECryptoNetworkEthereum, ECryptoNetworkBitcoin, ECryptoNetworkTron, ECryptoNetworkSolana:
		return true
	}
	return false
}

func (e ECryptoNetwork) String() string {
	return string(e)
}

func (e *ECryptoNetwork) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ECryptoNetwork(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ECryptoNetwork", str)
	}
	return nil
}

func (e ECryptoNetwork) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ECurrency string

const (
	ECurrencyNgn  ECurrency = "NGN"
	ECurrencyUsd  ECurrency = "USD"
	ECurrencyUsdt ECurrency = "USDT"
	ECurrencyBtc  ECurrency = "BTC"
)

var AllECurrency = []ECurrency{
	ECurrencyNgn,
	ECurrencyUsd,
	ECurrencyUsdt,
	ECurrencyBtc,
}

func (e ECurrency) IsValid() bool {
	switch e {
	case ECurrencyNgn, ECurrencyUsd, ECurrencyUsdt, ECurrencyBtc:
		return true
	}
	return false
}

func (e ECurrency) String() string {
	return string(e)
}

func (e *ECurrency) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ECurrency(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ECurrency", str)
	}
	return nil
}

func (e ECurrency) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EFrequency string

const (
	EFrequencyDaily     EFrequency = "DAILY"
	EFrequencyWeekly    EFrequency = "WEEKLY"
	EFrequencyMonthly   EFrequency = "MONTHLY"
	EFrequencyQuarterly EFrequency = "QUARTERLY"
	EFrequencyAnnually  EFrequency = "ANNUALLY"
	EFrequencyOneTime   EFrequency = "ONE_TIME"
	EFrequencyProRata   EFrequency = "PRO_RATA"
)

var AllEFrequency = []EFrequency{
	EFrequencyDaily,
	EFrequencyWeekly,
	EFrequencyMonthly,
	EFrequencyQuarterly,
	EFrequencyAnnually,
	EFrequencyOneTime,
	EFrequencyProRata,
}

func (e EFrequency) IsValid() bool {
	switch e {
	case EFrequencyDaily, EFrequencyWeekly, EFrequencyMonthly, EFrequencyQuarterly, EFrequencyAnnually, EFrequencyOneTime, EFrequencyProRata:
		return true
	}
	return false
}

func (e EFrequency) String() string {
	return string(e)
}

func (e *EFrequency) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EFrequency(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EFrequency", str)
	}
	return nil
}

func (e EFrequency) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EGender string

const (
	EGenderMale   EGender = "MALE"
	EGenderFemale EGender = "FEMALE"
)

var AllEGender = []EGender{
	EGenderMale,
	EGenderFemale,
}

func (e EGender) IsValid() bool {
	switch e {
	case EGenderMale, EGenderFemale:
		return true
	}
	return false
}

func (e EGender) String() string {
	return string(e)
}

func (e *EGender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EGender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EGender", str)
	}
	return nil
}

func (e EGender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ELanguage string

const (
	ELanguageEn ELanguage = "EN"
	ELanguageFr ELanguage = "FR"
)

var AllELanguage = []ELanguage{
	ELanguageEn,
	ELanguageFr,
}

func (e ELanguage) IsValid() bool {
	switch e {
	case ELanguageEn, ELanguageFr:
		return true
	}
	return false
}

func (e ELanguage) String() string {
	return string(e)
}

func (e *ELanguage) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ELanguage(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ELanguage", str)
	}
	return nil
}

func (e ELanguage) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EPaymentMethod string

const (
	EPaymentMethodCard   EPaymentMethod = "CARD"
	EPaymentMethodWallet EPaymentMethod = "WALLET"
)

var AllEPaymentMethod = []EPaymentMethod{
	EPaymentMethodCard,
	EPaymentMethodWallet,
}

func (e EPaymentMethod) IsValid() bool {
	switch e {
	case EPaymentMethodCard, EPaymentMethodWallet:
		return true
	}
	return false
}

func (e EPaymentMethod) String() string {
	return string(e)
}

func (e *EPaymentMethod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EPaymentMethod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EPaymentMethod", str)
	}
	return nil
}

func (e EPaymentMethod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ERole string

const (
	ERoleOrg  ERole = "ORG"
	ERoleUser ERole = "USER"
)

var AllERole = []ERole{
	ERoleOrg,
	ERoleUser,
}

func (e ERole) IsValid() bool {
	switch e {
	case ERoleOrg, ERoleUser:
		return true
	}
	return false
}

func (e ERole) String() string {
	return string(e)
}

func (e *ERole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ERole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ERole", str)
	}
	return nil
}

func (e ERole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ESort string

const (
	ESortDesc ESort = "desc"
	ESortAsc  ESort = "asc"
)

var AllESort = []ESort{
	ESortDesc,
	ESortAsc,
}

func (e ESort) IsValid() bool {
	switch e {
	case ESortDesc, ESortAsc:
		return true
	}
	return false
}

func (e ESort) String() string {
	return string(e)
}

func (e *ESort) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ESort(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ESort", str)
	}
	return nil
}

func (e ESort) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EStatus string

const (
	EStatusActive    EStatus = "ACTIVE"
	EStatusInactive  EStatus = "INACTIVE"
	EStatusDeleted   EStatus = "DELETED"
	EStatusPending   EStatus = "PENDING"
	EStatusCompleted EStatus = "COMPLETED"
)

var AllEStatus = []EStatus{
	EStatusActive,
	EStatusInactive,
	EStatusDeleted,
	EStatusPending,
	EStatusCompleted,
}

func (e EStatus) IsValid() bool {
	switch e {
	case EStatusActive, EStatusInactive, EStatusDeleted, EStatusPending, EStatusCompleted:
		return true
	}
	return false
}

func (e EStatus) String() string {
	return string(e)
}

func (e *EStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EStatus", str)
	}
	return nil
}

func (e EStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ETier string

const (
	ETierZero ETier = "zero"
	ETierOne  ETier = "one"
	ETierTwo  ETier = "two"
)

var AllETier = []ETier{
	ETierZero,
	ETierOne,
	ETierTwo,
}

func (e ETier) IsValid() bool {
	switch e {
	case ETierZero, ETierOne, ETierTwo:
		return true
	}
	return false
}

func (e ETier) String() string {
	return string(e)
}

func (e *ETier) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ETier(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ETier", str)
	}
	return nil
}

func (e ETier) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ETransactionType string

const (
	ETransactionTypeFunding    ETransactionType = "FUNDING"
	ETransactionTypeTransfer   ETransactionType = "TRANSFER"
	ETransactionTypeWithdrawal ETransactionType = "WITHDRAWAL"
)

var AllETransactionType = []ETransactionType{
	ETransactionTypeFunding,
	ETransactionTypeTransfer,
	ETransactionTypeWithdrawal,
}

func (e ETransactionType) IsValid() bool {
	switch e {
	case ETransactionTypeFunding, ETransactionTypeTransfer, ETransactionTypeWithdrawal:
		return true
	}
	return false
}

func (e ETransactionType) String() string {
	return string(e)
}

func (e *ETransactionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ETransactionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ETransactionType", str)
	}
	return nil
}

func (e ETransactionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EUserStatus string

const (
	EUserStatusUserLocked  EUserStatus = "USER_LOCKED"
	EUserStatusUserDeleted EUserStatus = "USER_DELETED"
	EUserStatusUserActive  EUserStatus = "USER_ACTIVE"
)

var AllEUserStatus = []EUserStatus{
	EUserStatusUserLocked,
	EUserStatusUserDeleted,
	EUserStatusUserActive,
}

func (e EUserStatus) IsValid() bool {
	switch e {
	case EUserStatusUserLocked, EUserStatusUserDeleted, EUserStatusUserActive:
		return true
	}
	return false
}

func (e EUserStatus) String() string {
	return string(e)
}

func (e *EUserStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EUserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EUserStatus", str)
	}
	return nil
}

func (e EUserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EWalletType string

const (
	EWalletTypeFiat   EWalletType = "FIAT"
	EWalletTypeCrypto EWalletType = "CRYPTO"
)

var AllEWalletType = []EWalletType{
	EWalletTypeFiat,
	EWalletTypeCrypto,
}

func (e EWalletType) IsValid() bool {
	switch e {
	case EWalletTypeFiat, EWalletTypeCrypto:
		return true
	}
	return false
}

func (e EWalletType) String() string {
	return string(e)
}

func (e *EWalletType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EWalletType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EWalletType", str)
	}
	return nil
}

func (e EWalletType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}