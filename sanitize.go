package main

import (
	"fmt"
	"os"

	"github.com/Northern-Lights/yara-parser/data"
	"github.com/Northern-Lights/yara-parser/grammar"
	"github.com/alecthomas/kong"
)

type _CleanCmd struct {
	Filename string `arg name:"file" type:"existingfile"`
}

func (self _CleanCmd) Run(ctx *kong.Context) error {
	yaraFile, err := os.Open(self.Filename)
	if err != nil {
		return err
	}
	defer yaraFile.Close()

	ruleset, err := grammar.Parse(yaraFile, os.Stdout)
	if err != nil {
		return err
	}

	self.sanitize(&ruleset)

	serialized, err := ruleset.Serialize()
	if err != nil {
		return err
	}

	fmt.Println(string(serialized))

	return nil
}

func (self _CleanCmd) sanitize(ruleset *data.RuleSet) {
	new_rules := make([]data.Rule, 0, len(ruleset.Rules))

	for _, rule := range ruleset.Rules {
		rule.Meta = nil
		new_rules = append(new_rules, rule)
	}
	ruleset.Rules = new_rules
}
