package rule_engine

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"grule-demo/operations"
	// "os"
	// "io/ioutil"
	// "fmt"
)

var knowledgeLibrary = *ast.NewKnowledgeLibrary()

type RuleInput interface {
	DataKey() string
}

type RuleOutput interface {
	DataKey() string
}

// configs associated with each rule,
// basically the expected content of the rule context
type RuleConfig interface {
	RuleName() string
	RuleInput() RuleInput
	RuleOutput() RuleOutput
}

type RuleEngineSvc struct {
}


// create a new rule engine instance
func NewRuleEngineSvc() *RuleEngineSvc {
	// external call to fetch current rule version can be added here
	buildRuleEngine()
	return &RuleEngineSvc{}
}

// build a rule version from a rules file|db|string
func buildRuleEngine() {
	ruleBuilder := builder.NewRuleBuilder(&knowledgeLibrary)
	// jsonData, err := ioutil.ReadFile("rules.json")
	// fmt.Println("Here", jsonData)
	// if err != nil {
	// 	panic(err)
	// }
	// ruleset, err := pkg.ParseJSONRuleset(jsonData)
	
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Parsed ruleset: ")
	// fmt.Println("Hello", ruleset)
	

	resource := pkg.NewFileResource("rules.grl")
	// f, err := os.Open("rules.json")
	// if err != nil {
	// 	panic(err)
	// }
	// underlying := pkg.NewReaderResource(f)
	// resource := pkg.NewJSONResourceFromResource(underlying)
	err := ruleBuilder.BuildRuleFromResource("Rules", "0.0.2", resource)
	if err != nil {
		panic(err)
	}

}

// fetch rules version from application memory
// add all request objects and action properties in to the working memory
func (svc *RuleEngineSvc) Execute(ruleConf RuleConfig) error {
	// get rule version to execute
	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("Rules", "0.0.2")

	// Defining the data attributes on which the rules will evaluate
	dataCtx := ast.NewDataContext()

	// add FACT and its identifier string into the data context
	err := dataCtx.Add(ruleConf.RuleInput().DataKey(), ruleConf.RuleInput())
	if err != nil {
		return err
	}

	err = dataCtx.Add(ruleConf.RuleOutput().DataKey(), ruleConf.RuleOutput())
	if err != nil {
		return err
	}

	err = dataCtx.Add("CF", &operations.CustomFunction{})
	if err != nil {
		return err
	}

	ruleEngine := engine.NewGruleEngine()

	// pass FACT and rule version into the rule engine exection flow
	err = ruleEngine.Execute(dataCtx, knowledgeBase)
	if err != nil {
		return err
	}
	return nil
}