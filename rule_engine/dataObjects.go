package rule_engine

type PhoneNumber struct {
	NumberType			string	`json:"numberType"`
	Number				string	`json:"number"`
}

type EmailAddress struct {
	EmailType			string	`json:"emailType"`
	Address				string	`json:"address"`
}

type Address struct {
	AddressType			string	`json:"addressType"`
	Line1				string	`json:"line1"`
	Line2				string	`json:"line2"`
	PostCode			string	`json:"postCode"`
	City				string	`json:"city"`
	State				string	`json:"state"`
	Country				string	`json:"country"`
}

type LegalEntity struct {
	DateOfRegistration		string	`json:"line1"`
	Type					string	`json:"type"`
	Url						string	`json:"url"`
	IndustryClassification	string	`json:"industryClassification"`
	CountryOfIncorporation	string	`json:"countryOfIncorporation"`
}

type GovernmentIds struct {
	GovernmentId		string	`json:"governmentId"`
	IdType				string	`json:"IdType"`
	IssuingAuthority	string	`json:"issuingAuthority"`
	ExpirationDate		string	`json:"expirationDate"`
}

type Actor struct {
	ActorType 		string			`json:"actorType"`
	Phones			[]PhoneNumber	`json:"phoneNumber"`
    FirstName		string			`json:"firstName"`
    LastName		string			`json:"lastName"`
    Emails			[]EmailAddress	`json:"emailAddress"`
    Addresses		[]Address		`json:"address"`
    GovernmentIds	[]GovernmentIds	`json:"governmentIds"`
	Status			string			`json:"status"`
	Domicile		string			`json:"domicile"`
	Gender			string			`json:"gender"`
	LegalEntity		[]LegalEntity	`json:"legalEntity"`
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

// input data object from kafka topic
type EnrichedTransactionInput struct {
	Actor							Actor		`json:"actor"`
	Status 							string		`json:"status"`
	Currency 						string		`json:"currency"`
	Channel 						string		`json:"channel"`
	AccountNumber					string		`json:"accountNumber"`
	Category						string		`json:"category"`
	Amount 							float64		`json:"amount"`
	LocalAmount						float64		`json:"localAmount"`
	EntryDate 						string		`json:"entryDate"`
	Description 					string		`json:"description"`
	CardNumber 						string		`json:"cardNumber"`
	ChannelLocation					string		`json:"channelLocation"`
	Balance 						float64		`json:"balance"`
	CheckNumber 					string		`json:"checkNumber"`
	TransactionMethod 				string		`json:"transactionMethod"`
	InternationalTransaction 		bool		`json:"internationalTransaction"`
	MerchantCategoryCode 			string		`json:"merchantCategoryCode"`
	MerchantCountryCode 			string		`json:"merchantCountryCode"`
	Beneficiary						Beneficiary	`json:"beneficiary"`
	ActorPEPMatch					bool		`json:"actor_pep_match"`
	ActorCrimeListMatch				bool		`json:"actor_crime_list_match"`
	ActorWatchListMatch				bool		`json:"actor_watch_list_match"`
	ActorSanctionListMatch			bool		`json:"actor_sanction_list_match"`
	BeneficiaryPEPMatch				bool		`json:"beneficiary_pep_match"`
	BeneficiaryCrimeListMatch		bool		`json:"beneficiary_crime_list_match"`
	BeneficiaryWatchListMatch		bool		`json:"beneficiary_watch_list_match"`
	BeneficiarySanctionListMatch	bool		`json:"beneficiary_sanction_list_match"`
}

// rule output action object
type AlertOutput struct {
	Tag string `json:"tag"`
}

// string representation of Input Fact
func (u *EnrichedTransactionInput) DataKey() string {
	return "Transaction"
}

// string representation of Output Fact
func (u *AlertOutput) DataKey() string {
	return "Alert"
}