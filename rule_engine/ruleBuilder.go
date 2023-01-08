package rule_engine

import (
	"log"
	"encoding/json"
)

type DataObj struct {
	Name		string 	`json:"name"`
	Desc		string 	`json:"desc"`
	Salience	int8 	`json:"salience"`
	When		string 	`json:"when"`
	Then		string 	`json:"then"`
}

type KafkaData struct {
	Workflow	string			`json:"workflow"`
	Version		uint8			`json:"version"`
	Data		json.RawMessage	`json:"data"`
}

func NewRuleBuilder(data []byte) error {
	request := KafkaData{}
	json.Unmarshal(data, &request)

	err := createGRLFile(request)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}