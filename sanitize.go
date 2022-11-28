package main

import (
	"fmt"
	"os"

	"github.com/Northern-Lights/yara-parser/data"
	"github.com/Northern-Lights/yara-parser/grammar"

	yara "github.com/Velocidex/go-yara"
	"github.com/alecthomas/kong"
)

type _CleanCmd struct {
	Filename string `arg name:"file" type:"existingfile"`
	Verify   bool   `flag name:"verify" type:"bool" doc:"Also verify the rules"`
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

	err = self.sanitize(&ruleset)
	if err != nil {
		return err
	}

	serialized, err := ruleset.Serialize()
	if err != nil {
		return err
	}

	fmt.Println(string(serialized))

	return nil
}

func (self _CleanCmd) sanitize(ruleset *data.RuleSet) error {
	new_rules := make([]data.Rule, 0, len(ruleset.Rules))

	for _, rule := range ruleset.Rules {
		rule.Meta = nil

		if self.Verify {
			tmp_rules := &data.RuleSet{
				Imports: ruleset.Imports,
			}
			tmp_rules.Rules = append(tmp_rules.Rules, rule)
			serialized, err := tmp_rules.Serialize()
			if err != nil {
				continue
			}

			compiler, err := yara.NewCompiler()
			if err != nil {
				return err
			}

			err = compiler.AddString(serialized, "")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Rule %v rejected: %v\n", rule, err)
				continue
			}

		}

		new_rules = append(new_rules, rule)
	}
	ruleset.Rules = new_rules

	return nil
}
