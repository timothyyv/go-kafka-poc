package main

import (
	"fmt"
	"context"
	"encoding/json"
	"grule-demo/kafka"
	"grule-demo/rule_engine"
	kafkago "github.com/segmentio/kafka-go"
	"github.com/hyperjumptech/grule-rule-engine/logger"
)

type AlertServiceClient struct {
	ruleEngineSvc *rule_engine.RuleEngineSvc
}

func NewAlertServiceClient(ruleEngineSvc *rule_engine.RuleEngineSvc) *AlertServiceClient {
	return &AlertServiceClient{
		ruleEngineSvc: ruleEngineSvc,
	}
}

func (svc AlertServiceClient) checkTransaction(t rule_engine.EnrichedTransactionInput) string {
	// CREATE CONTEXT TO STORE FACTS
	ctx := rule_engine.NewCustomContext()

	// map request data to FACT attributes
	ctx.EnrichedTransactionInput = &rule_engine.EnrichedTransactionInput{
		Actor:							t.Actor,
		Status: 						t.Status,
		Currency: 						t.Currency,
		Channel: 						t.Channel,
		AccountNumber:					t.AccountNumber,
		Category:						t.Category,
		Amount: 						t.Amount,
		LocalAmount:					t.LocalAmount,
		EntryDate: 						t.EntryDate,
		Description: 					t.Description,
		CardNumber: 					t.CardNumber,
		ChannelLocation:				t.ChannelLocation,
		Balance: 						t.Balance,
		CheckNumber: 					t.CheckNumber,
		TransactionMethod: 				t.TransactionMethod,
		InternationalTransaction: 		t.InternationalTransaction,
		MerchantCategoryCode: 			t.MerchantCategoryCode,
		MerchantCountryCode: 			t.MerchantCountryCode,
		Beneficiary:					t.Beneficiary,
		ActorPEPMatch:					t.ActorPEPMatch,
		ActorCrimeListMatch:			t.ActorCrimeListMatch,
		ActorWatchListMatch:			t.ActorWatchListMatch,
		ActorSanctionListMatch:			t.ActorSanctionListMatch,
		BeneficiaryPEPMatch:			t.BeneficiaryPEPMatch,
		BeneficiaryCrimeListMatch:		t.BeneficiaryCrimeListMatch,
		BeneficiaryWatchListMatch:		t.BeneficiaryWatchListMatch,
		BeneficiarySanctionListMatch:	t.BeneficiarySanctionListMatch,
	}

	// pass the new context into the engine working memory
	err := svc.ruleEngineSvc.Execute(ctx)
	if err != nil {
		logger.Log.Error("CHECK TRANSACTION RULE ENGINE FAILED", err)
	}

	return ctx.AlertOutput.Tag
}

func main() {
	// instantiate a service instance to build or fetch a rule version to be executed
	ruleEngineSvc := rule_engine.NewRuleEngineSvc()

	// create new service client with a rule engine instance
	alertSvcClient := NewAlertServiceClient(ruleEngineSvc)

	reader := kafka.NewKafkaReader("enriched_transaction_request")
	ctx := context.Background()
	transData := make(chan kafkago.Message, 1000)

	go reader.FetchMessages(ctx, transData)
	go reader.CommitMessages(ctx, transData)

	data := <- transData

	// generate request data object to execute rule against
	request := rule_engine.EnrichedTransactionInput{}
	err := json.Unmarshal([]byte(string(data.Value)), &request)

	if err != nil {
		fmt.Println(err.Error()) 
		//json: Unmarshal(non-pointer main.Request)
	}

	fmt.Println("check transaction for issues: ", alertSvcClient.checkTransaction(request))
	fmt.Println("Hello")
}