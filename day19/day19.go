package day19

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"strconv"
	"strings"
)

type Operation string

const (
	OLower  Operation = "lower"
	OBigger Operation = "bigger"
	OFinal  Operation = "final"
)

type Condition struct {
	field     string
	operation Operation
	oValue    int
	goToRule  string
}

type Rule struct {
	conds []*Condition
}

type Part struct {
	x, m, a, s int
	status     rune
}

type RuleMap = map[string]*Rule

func Solve(input string) {
	splitInput := utils.SplitByEmptyRow(input)

	ruleStrings := utils.SplitRows(splitInput[0])
	partStrings := utils.SplitRows(splitInput[1])

	rules := make(RuleMap)
	parts := make([]*Part, len(partStrings))

	for i, ps := range partStrings {
		var part Part
		commaSplit := strings.Split(ps[1:len(ps)-1], ",")
		for j, cs := range commaSplit {
			value, _ := strconv.Atoi(strings.Split(cs, "=")[1])
			if j == 0 {
				part.x = value
			} else if j == 1 {
				part.m = value
			} else if j == 2 {
				part.a = value
			} else if j == 3 {
				part.s = value
			}
		}
		part.status = 'n'
		parts[i] = &part
	}

	for _, rs := range ruleStrings {
		ruleSplit := strings.Split(rs[:len(rs)-1], "{")
		ruleName := ruleSplit[0]

		rules[ruleName] = &Rule{}

		conditionSplit := strings.Split(ruleSplit[1], ",")
		for j, c := range conditionSplit {
			if j == len(conditionSplit)-1 { //last rule, no condition
				rules[ruleName].conds = append(rules[ruleName].conds, &Condition{
					operation: OFinal,
					goToRule:  c, //whole "c" should be rule name (or A/R)
				})
			} else {
				ruleParts := strings.Split(c, ":")
				condParts := strings.FieldsFunc(ruleParts[0], spl)

				isLower := strings.Contains(ruleParts[0], "<")
				operation := OLower
				if !isLower {
					operation = OBigger
				}
				value, _ := strconv.Atoi(condParts[1])

				rules[ruleName].conds = append(rules[ruleName].conds, &Condition{
					field:     condParts[0],
					operation: operation,
					oValue:    value,
					goToRule:  ruleParts[1],
				})

			}
		}

	}

	fmt.Println(parts)
	fmt.Println(rules)

	total := 0
	for _, p := range parts {
		processPart(p, &rules)

		if p.status == 'A' {
			fmt.Println("Success: ", p.x, p.m, p.a, p.s, p.x+p.m+p.a+p.s)
			total += p.x + p.m + p.a + p.s
		}
	}

	fmt.Println("Part 1: ", total)

}

func processPart(p *Part, rules *RuleMap) {
	fmt.Println("processPart")
	currentRule := "in"

	for currentRule != "R" && currentRule != "A" {
		fmt.Println("CURRENT RULE: ", currentRule)
		rule := (*rules)[currentRule]
		currentRule = processRule(p, rule)
	}

	if currentRule == "R" {
		p.status = 'R'
	} else if currentRule == "A" {
		p.status = 'A'
	} else {
		panic("no valid result!")
	}
}

func processRule(p *Part, r *Rule) string {
	fmt.Println("processRule", p, r)
	for _, c := range r.conds {
		validCondition := processCondition(p, c)
		if validCondition {
			return c.goToRule
		}
	}
	panic("no valid condition!")
}

func processCondition(p *Part, c *Condition) bool {
	fmt.Println("processCondition")
	if c.operation == OFinal {
		return true
	}

	pValue := getFieldValue(p, c.field)
	if c.operation == OBigger {
		if pValue > c.oValue {
			return true
		}
		return false
	} else if c.operation == OLower {
		if pValue < c.oValue {
			return true
		}
		return false
	}
	panic("not valid operation field")
}

func getFieldValue(p *Part, field string) int {
	fmt.Println("getFieldValue")
	if field == "x" {
		return p.x
	} else if field == "m" {
		return p.m
	} else if field == "a" {
		return p.a
	} else if field == "s" {
		return p.s
	}
	panic("unknown field")
}

func spl(c rune) bool {
	return c == '<' || c == '>'
}
