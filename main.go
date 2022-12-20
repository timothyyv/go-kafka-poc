package main

import (
	"fmt"
	"grule-demo/rule_engine"
	"github.com/go-faker/faker/v4"
	"github.com/hyperjumptech/grule-rule-engine/logger"
)

type User struct {
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

type OfferService interface {
	CheckOfferForUser(user User) bool
}

type OfferServiceClient struct {
	ruleEngineSvc *rule_engine.RuleEngineSvc
}

func NewOfferService(ruleEngineSvc *rule_engine.RuleEngineSvc) *OfferServiceClient {
	return &OfferServiceClient{
		ruleEngineSvc: ruleEngineSvc,
	}
}

type dataOut struct {
	isApplicable bool
	msg string
}

func (svc OfferServiceClient) CheckOfferForUser(user User) dataOut {
	// create a new datacontext instance to be used during this cycle
	// this is essential to set the FACTs and actions required during the rule execution
	offerCard := rule_engine.NewUserOfferContext()

	// store the request data in the working memory of the engine
	offerCard.UserOfferInput = &rule_engine.UserOfferInput{
		TransactionDate:	user.TransactionDate,
		SettledAt:			user.SettledAt,
		Amount:				user.Amount,
		Category:			user.Category,
		Currency:			user.Currency,
		Description:		user.Description,
		NameOfActor:		user.NameOfActor,
		Status:				user.Status,
		CreatedAt:			user.CreatedAt,
	}

	err := svc.ruleEngineSvc.Execute(offerCard)
	if err != nil {
		logger.Log.Error("GET USER OFFER RULE ENGINE FAILED", err)
	}

	return dataOut{
		offerCard.UserOfferOutput.IsOfferApplicable,
		offerCard.UserOfferOutput.PrintStatement,
	}
}

func main() {
	// instantiate a service instance to build or fetch a rule version to be executed
	// ruleEngineSvc := rule_engine.NewRuleEngineSvc()

	// // create a service client with the rule service instance
	// // the client instance runs
	// offerSvc := NewOfferService(ruleEngineSvc)

	// userA := User{
	// 	Name:              "Mohit Khare",
	// 	Username:          "mkfeuhrer",
	// 	Email:             "me@mohitkhare.com",
	// 	Gender:            "Male",
	// 	Age:               45,
	// 	TotalOrders:       50,
	// 	AverageOrderValue: 225,
	// }

	// fmt.Println("offer validity for user A: ", offerSvc.CheckOfferForUser(userA))

	// userB := User{
	// 	Name:              "Pranjal Sharma",
	// 	Username:          "pj",
	// 	Email:             "pj@abc.com",
	// 	Gender:            "Male",
	// 	Age:               25,
	// 	TotalOrders:       10,
	// 	AverageOrderValue: 80,
	// }

	// fmt.Println("offer validity for user B: ", offerSvc.CheckOfferForUser(userB))
	// fmt.Println("Hello")
	a := rule_engine.TransactionInput{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
}