package rule_engine

import (
	"fmt"
	"log"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// build and add new ruleset to rule engine working memory
// build a rule version from a rules file|db|string
func BuildRuleEngine(jsonData []byte) error {
	fmt.Println("and here")
	ruleBuilder := builder.NewRuleBuilder(&knowledgeLibrary)
	ruleset, err := pkg.ParseJSONRuleset(jsonData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New ruleset: %v", ruleset)

	bs := pkg.NewBytesResource([]byte(ruleset))
	err = ruleBuilder.BuildRuleFromResource("Rules", "0.0.1", bs)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}