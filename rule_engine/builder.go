package rule_engine

import (
	"os"
	// "fmt"
	"log"
	"strconv"
	"io/ioutil"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// build and add new ruleset to rule engine working memory
// build a rule version from a rules file|db|string
func createGRLFile(jsonData KafkaData) error {
	ruleset, err := pkg.ParseJSONRuleset(jsonData.Data)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("New ruleset: %v\n", strconv.Itoa(int(jsonData.Version)))

	wd, _ := os.Getwd()

	if _, err := os.Stat(wd + "/workflows/" + jsonData.Workflow); os.IsNotExist(err) {
		if err := os.MkdirAll(wd + "/workflows/" + jsonData.Workflow, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	data := []byte(ruleset)
	err = ioutil.WriteFile(wd + "/workflows/" + jsonData.Workflow + "/v" + strconv.Itoa(int(jsonData.Version)) + ".grl", data, 0777)

    if err != nil {
        log.Fatal(err)
    }

	return nil
}

func BuildRuleEngine(workflow string, version string) error {
	wd, _ := os.Getwd()
	file := wd + "/workflows/" + workflow + "/v" + version + ".grl"
	resource := pkg.NewFileResource(file)

	ruleBuilder := builder.NewRuleBuilder(&knowledgeLibrary)
	err := ruleBuilder.BuildRuleFromResource(workflow, version, resource)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}