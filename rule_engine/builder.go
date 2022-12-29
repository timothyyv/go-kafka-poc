package rule_engine

import (
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// build and add new ruleset to rule engine working memory
func BuildRuleEngine(jsonData []byte) {
	ruleBuilder := builder.NewRuleBuilder(&knowledgeLibrary)
	ruleset, err := pkg.ParseJSONRuleset(jsonData)
	if err != nil {
		panic(err)
	}

	bs := pkg.NewBytesResource([]byte(ruleset))
	err = ruleBuilder.BuildRuleFromResource("Rules", "0.0.1", bs)
	if err != nil {
		panic(err)
	}
}