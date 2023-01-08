package main

import (
	"fmt"
	"strconv"
	"context"
	"encoding/json"
	"grule-demo/kafka"
	"grule-demo/rule_engine"
	"github.com/hyperjumptech/grule-rule-engine/logger"
	kafkago "github.com/segmentio/kafka-go"
)

type AlertServiceClient struct {
	ruleEngineSvc *rule_engine.RuleEngineSvc
}

func NewAlertServiceClient(ruleEngineSvc *rule_engine.RuleEngineSvc) *AlertServiceClient {
	return &AlertServiceClient{
		ruleEngineSvc: ruleEngineSvc,
	}
}

type EnrichedData struct {
	Workflow	string									`json:"workflow"`
	Version		uint8									`json:"version"`
	Data		rule_engine.EnrichedTransactionInput	`json:"data"`
}

func (svc AlertServiceClient) checkTransaction(t EnrichedData) string {
	version := strconv.Itoa(int(t.Version))

	// build rule from resource based on version adn workflow name
	rule_engine.BuildRuleEngine(t.Workflow, version)

	// CREATE CONTEXT TO STORE FACTS
	ctx := rule_engine.NewCustomContext()
	// fmt.Printf("SENT DATA %v\n", t)

	// map request data to FACT attributes
	ctx.EnrichedTransactionInput = &rule_engine.EnrichedTransactionInput{
		Id:								t.Data.Id,
		Actor:							t.Data.Actor,
		Status: 						t.Data.Status,
		Currency: 						t.Data.Currency,
		Channel: 						t.Data.Channel,
		AccountNumber:					t.Data.AccountNumber,
		Category:						t.Data.Category,
		Amount: 						t.Data.Amount,
		LocalAmount:					t.Data.LocalAmount,
		EntryDate: 						t.Data.EntryDate,
		Description: 					t.Data.Description,
		CardNumber: 					t.Data.CardNumber,
		ChannelLocation:				t.Data.ChannelLocation,
		Balance: 						t.Data.Balance,
		CheckNumber: 					t.Data.CheckNumber,
		TransactionMethod: 				t.Data.TransactionMethod,
		InternationalTransaction: 		t.Data.InternationalTransaction,
		MerchantCategoryCode: 			t.Data.MerchantCategoryCode,
		MerchantCountryCode: 			t.Data.MerchantCountryCode,
		Beneficiary:					t.Data.Beneficiary,
		ActorPEPMatch:					t.Data.ActorPEPMatch,
		ActorCrimeListMatch:			t.Data.ActorCrimeListMatch,
		ActorWatchListMatch:			t.Data.ActorWatchListMatch,
		ActorSanctionListMatch:			t.Data.ActorSanctionListMatch,
		BeneficiaryPEPMatch:			t.Data.BeneficiaryPEPMatch,
		BeneficiaryCrimeListMatch:		t.Data.BeneficiaryCrimeListMatch,
		BeneficiaryWatchListMatch:		t.Data.BeneficiaryWatchListMatch,
		BeneficiarySanctionListMatch:	t.Data.BeneficiarySanctionListMatch,
		CreatedAt: 						t.Data.CreatedAt,
	}

	// pass the new context into the engine working memory
	err := svc.ruleEngineSvc.Execute(ctx, t.Workflow, version)
	if err != nil {
		logger.Log.Error("CHECK TRANSACTION RULE ENGINE FAILED", err)
	}

	return ctx.AlertOutput.Tag
}

func main() {
	// instantiate a service instance to build or fetch a rule version to be executed
	ruleEngineSvc := rule_engine.NewRuleEngineSvc()

	// // create new service client with a rule engine instance
	alertSvcClient := NewAlertServiceClient(ruleEngineSvc)

	topics := []string{"enriched_transaction_request", "rule-topic"}

	reader := kafka.NewKafkaReader(topics)
	ctx := context.Background()
	transData := make(chan kafkago.Message, 1000)

	go reader.FetchMessages(ctx, transData)
	data := <- transData

	if string(data.Topic) == "rule-topic" {
		err := rule_engine.NewRuleBuilder(data.Value)
		if err != nil {
			logger.Log.Fatal("Unable to build", err)
		}
	}

	if string(data.Topic) == "enriched_transaction_request" {

		// generate request data object to execute rule against
		request := EnrichedData{}
		err := json.Unmarshal([]byte(string(data.Value)), &request)
		// fmt.Printf("NEW REQUEST %v\n", request.Data.Actor.Emails)
	
		if err != nil {
			fmt.Println(err.Error()) 
			//json: Unmarshal(non-pointer main.Request)
		}

		fmt.Println("check transaction for issues: ", alertSvcClient.checkTransaction(request))
	}

	// fmt.Println("Hello")
	reader.CommitMessages(ctx, transData)
}