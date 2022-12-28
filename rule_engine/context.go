package rule_engine

type CustomContext struct {
	EnrichedTransactionInput *EnrichedTransactionInput
	AlertOutput *AlertOutput
}

func (c *CustomContext) RuleName() string {
	return "monitoring_rules"
}

func (c *CustomContext) RuleInput() RuleInput {
	return c.EnrichedTransactionInput
}

func (c *CustomContext) RuleOutput() RuleOutput {
	return c.AlertOutput
}

func NewCustomContext() *CustomContext {
	return &CustomContext{
		EnrichedTransactionInput: &EnrichedTransactionInput{},
		AlertOutput: &AlertOutput{},
	}
}

// DataKey() method returns the class name identifier for the FACT object to be used in the working memory
// Context funcs creates a store that holds all object instances to be used during rule execution
// collection of rules are a knowledge set
// output object is the set of actions or values expected to be handled when there's a rule match
// input is the object to be used to evaluate a set of conditions