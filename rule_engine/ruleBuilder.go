package rule_engine

import (
	"fmt"
	"log"
)

func NewRuleBuilder(data []byte) error {
	fmt.Println("Got here")
	err := BuildRuleEngine(data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}