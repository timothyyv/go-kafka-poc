// rule_engine/offer.go
package rule_engine

type UserOfferContext struct {
	UserOfferInput  *UserOfferInput
	UserOfferOutput *UserOfferOutput
}

func (uoc *UserOfferContext) RuleName() string {
	return "user_offers"
}

func (uoc *UserOfferContext) RuleInput() RuleInput {
	return uoc.UserOfferInput
}

func (uoc *UserOfferContext) RuleOutput() RuleOutput {
	return uoc.UserOfferOutput
}

type Users struct {
	Users []UserOfferInput `json:"users"`
}

type Aggregation struct {}

func (agg *Aggregation) Sum(input1 int64, input2 int64) int64 {
	return -1
	// if err != nil {
	// 	log.Fatal(err)
	// 	return -1
	// }

	// send kafka topic
	// {
	// 	action == "sum",
	// 	filter_fileds == "",
	// 	action_field == "transaction_amt"
	// filter_fields = [{"Name": "Tina"}]
    // date_range = ["2022-12-10", "2022-12-18"]
	// }
}


// request data attributes
type UserOfferInput struct {
	TransactionDate		string	`json:"transactionDate"`
    SettledAt			string	`json:"settledAt"`
    Amount				float64 `json:"amount"`
    Category 			string	`json:"category"`
    Currency			string	`json:"currency"`
    Description			string	`json:"description"`
    NameOfActor			string	`json:"nameOfActor"`
    Status				string	`json:"status"`
    CreatedAt			string	`json:"createdAt"`
}

type PhoneNumber struct {
	NumberType			string	`json:"numberType"`	//Enum: "HOME" "WORK" "OTHER" "MOBILE"
	Number				string	`json:"number"`
}

type EmailAddress struct {
	EmailType			string	`json:"emailType"`	//Enum: "HOME" "WORK" "OTHER"
	Address				string	`json:"address"`
}

type Address struct {
	AddressType			string	`json:"addressType"` // Enum: "HOME" "WORK" "OTHER" "PRINCIPAL_PLACE_OF_BUSINESS" "REGISTERED_ADDRESS" 
													//"BUSINESS_ADDRESS_INDIVIDUAL" "HEADQUARTERS" "INVESTMENT_ADVISOR_ADDRESS" "LOCAL_OFFICE" 
													//"MAILING_ADDRESS" "MAILING_ADDRESS_GSAM" "MAILING_ADDRESS_PWM" "MAILING_ADDRESS_TAX" 
													//"MILITARY_ADDRESS" "UNKNOWN" "NEXT_OF_KIN" "OTHER_CONTACT_PERSON" "OTHER_PHYSICAL_BUSINESS_LOC" 
													//RESIDENCE_ADDRESS"
	Line1				string	`json:"line1"`
	Line2				string	`json:"line2"`
	PostCode			string	`json:"postCode"`
	City				string	`json:"city"`
	State				string	`json:"state"`
	Country				string	`json:"country"`
}

type LegalEntity struct {
	DateOfRegistration		string	`json:"line1"`
	Type					string	`json:"type"` //Enum: "SOLE_PROPRIETORSHIP" "PARTNERSHIP_GENERAL" "PARTNERSHIP_LIMITED" 
													//"COMPANY_PRIVATE" "COMPANY_PUBLIC" "HOLDING" "SUBSIDIARY" "NONPROFIT" 
													//"ORGANIZATION" "GOVERNMENT" "TRUST" "OTHER"
	Url						string	`json:"url"`
	IndustryClassification	string	`json:"industryClassification"`	//Enum: "ISAT" "MCC" "NACE" "NAICS" "SIC"
	CountryOfIncorporation	string	`json:"countryOfIncorporation"`
}

type GovernmentIds struct {
	GovernmentId		string	`json:"governmentId"`
	IdType				string	`json:"IdType"`
	IssuingAuthoriy		string	`json:"issuingAuthority"`
	ExpirationDate		string	`json:"expirationDate"`
}

type Actor struct {
	Id 							string			`json:"id"`
	ActorType 					string			`json:"actorType"`	//Enum: "INDIVIDUAL" "LEGAL_ENTITY"
	Phones						[]PhoneNumber	`json:"phoneNumber"`
    FirtsName					string			`json:"firstName"`
    LastName					string			`json:"lastName"`
    Emails						[]EmailAddress	`json:"emailAddress"`
    Addresses					[]Address		`json:"address"`
    BirthDate					string			`json:"birthDate"`
    GovernmentIds				[]GovernmentIds	`json:"governmentIds"`
	Monitored					bool			`json:"monitored"`
	Status						string			`json:"status"`
	Domicile					string			`json:"domicile"`
	Gender						string			`json:"gender"`
	LegalEntity					[]LegalEntity	`json:"legalEntity"`
	AccountIds					[]string		`json:"accountIds"`
}

type Account struct {
	Id 							string			`json:"id"`
	AccountType 				string			`json:"accountType"`	//Enum: "CURRENT" "SAVINGS" "MONEY_MARKET" "TERM_DEPOSIT" 
																		//"LOAN" "CREDIT_CARD" "TRADING" "HELD_IN_CUSTODY" "CASH_ON_DELIVERY" "BOTH_HIC_AND_COD"
	Status						string			`json:"status"`	//Enum: "ACTIVE" "CLOSED" "TERMINATED" "NO_TRADING"
    AccountNumber				string			`json:"accountNumber"`
   	Iban						string			`json:"iban"`
	ProductCode					string			`json:"productCode"`
	ProductName					string			`json:"productName"`
	ProductNumber				string			`json:"productNumber"`
   	CurrencyCode				string			`json:"currencyCode"`
    OpenhDate					string			`json:"openDate"`
	TerminationDate				string			`json:"terminationDate"`
    ClosingBalance				string			`json:"closigBalance"`
	Depositable					bool			`json:"depositable"`
	WithDrawable				bool			`json:"withDrawable"`
	Terminated					bool			`json:"terminated"`
	Internal					bool			`json:"internal"`
	BankCode					string			`json:"bankCode"`
}

type Beneficiary struct {
	Id 						string	`json:"id"`
	Type 					string	`json:"type"`
	Name 					string	`json:"name"`
	Address 				string	`json:"address"`
	BankName 				string	`json:"bankName"`
	Bic 					string	`json:"bic"`
	AccountNumber 			string	`json:"accountNumber"`
	Iban 					string	`json:"iban"`
	MerchantCategoryCode 	string	`json:"merchantCategoryCode"`
	MerchantCountryCode 	string	`json:"merchantCountry"`
}

type TransactionInput struct {
	Id							string		`json:"id"`
	Status 						string		`json:"status"`
	Currency 					string		`json:"currency"`
	Channel 					string		`json:"channel"`
	ActorId 					string		`json:"actorId"`
	AccountId 					string		`json:"accountId"`
	Category					string		`json:"category"`
	Amount 						float64		`json:"amount"`
	LocalAmount					float64		`json:"localAmount"`
	EntryDate 					string		`json:"entryDate"`
	Description 				string		`json:"description"`
	CardNumber 					string		`json:"cardNumber"`
	ChannelLocation				string		`json:"channelLocation"`
	Balance 					float64		`json:"balance"`
	CheckNumber 				string		`json:"checkNumber"`
	TransactionMethod 			string		`json:"transactionMethod"`	//Enum: "FEE" "INTEREST" "CASH" "CHECK" "PAYMENT" "DIRECT_DEBIT" "
																		//STANDING_ORDER" "TRANSFER" "REVERSAL" "CORRECTION" "OTHER" "PREPAID" "CLOSED" "CARD"
	InternationalTransaction 	bool		`json:"internationalTransaction"`
	MerchantCategoryCode 		string		`json:"merchantCategoryCode"`
	MerchantCountryCode 		string		`json:"merchanctCountryCode"`
	TransactionBeneficiary		Beneficiary	`json:"beneficiary"`
}

func (u *UserOfferInput) DataKey() string {
	return "InputData"
}

// request output object
type UserOfferOutput struct {
	IsOfferApplicable bool `json:"is_offer_applicable"`
	PrintStatement string `json:"print_statement"`
}

func (u *UserOfferOutput) DataKey() string {
	return "OutputData"
}

func NewUserOfferContext() *UserOfferContext {
	return &UserOfferContext{
		UserOfferInput:  &UserOfferInput{},
		UserOfferOutput: &UserOfferOutput{},
	}
}

// DataKey() method returns the class name identifier for the FACT object to be used in the working memory
// Context funcs creates a store that holds all object instances to be used during rule execution
// collection of rules are a knowledge set
// output object is the set of actions or values expected to be handled when there's a rule match
// input is the object to be used to evaluate a set of conditions