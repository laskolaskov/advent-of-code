package aoc2023

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Condition struct {
	Value     int
	Operation string
	Symbol    string
	Target    string
}

type WF struct {
	Name       string
	Conditions []Condition
}

type Part struct {
	X, M, A, S int
}

func Day19part1() int {
	flows, parts := scanWorkflowsAndParts("./aoc2023/input-day19.txt")
	total := 0
	for _, part := range parts {
		isAccepted := runFlow(part, flows["in"], flows)
		if isAccepted {
			total += part.X + part.M + part.A + part.S
		}
	}
	return total
}

// TODO part 2
func Day19part2() int {
	flows, _ := scanWorkflowsAndParts("./aoc2023/input-day19test.txt")
	for _, f := range flows {
		fmt.Println(f)
	}
	return 0
}

func runFlow(p Part, f WF, flows map[string]WF) bool {
	//accepted
	if f.Name == "A" {
		return true
	}
	//rejected
	if f.Name == "R" {
		return false
	}
	//check flow conditions and find next flow
	for _, c := range f.Conditions {
		//check if condition is the last (has only target)
		if c.Value == 0 {
			//run the next workflow
			next := flows[c.Target]
			return runFlow(p, next, flows)
		}
		//check condtiton
		//TODO: looks ugly(and maybe it is also slow - може да си грозен, но за сметка на това си тъп, Гумени Глави), can I optimize/refactor here?
		switch c.Symbol {
		//determine the operation and make the check
		//if condition is passed, run the next workflow
		//otherwise do nothing and pass to next condition
		case "x":
			if c.Operation == ">" {
				if p.X > c.Value {
					//run the next workflow
					next := flows[c.Target]
					return runFlow(p, next, flows)
				}
			}
			if c.Operation == "<" {
				if p.X < c.Value {
					//run the next workflow
					next := flows[c.Target]
					return runFlow(p, next, flows)
				}
			}
		case "m":
			if c.Operation == ">" {
				if p.M > c.Value {
					//run the next workflow
					next := flows[c.Target]
					return runFlow(p, next, flows)
				}
			}
			if c.Operation == "<" {
				if p.M < c.Value {
					//run the next workflow
					next := flows[c.Target]
					return runFlow(p, next, flows)
				}
			}
		case "a":
			if c.Operation == ">" {
				if p.A > c.Value {
					//run the next workflow
					next := flows[c.Target]
					return runFlow(p, next, flows)
				}
			}
			if c.Operation == "<" {
				if p.A < c.Value {
					//run the next workflow
					next := flows[c.Target]
					return runFlow(p, next, flows)
				}
			}
		case "s":
			if c.Operation == ">" {
				if p.S > c.Value {
					//run the next workflow
					next := flows[c.Target]
					return runFlow(p, next, flows)
				}
			}
			if c.Operation == "<" {
				if p.S < c.Value {
					//run the next workflow
					next := flows[c.Target]
					return runFlow(p, next, flows)
				}
			}
		}
	}
	return false
}

func scanWorkflowsAndParts(fileName string) (map[string]WF, []Part) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	isWF := true

	wf := map[string]WF{"A": {Name: "A"}, "R": {Name: "R"}}
	parts := []Part{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()
		//reach the empty line - no mork worflows, next lines are parts
		if l == "" {
			isWF = false
			continue
		}
		if isWF {
			//parse workflow
			s := strings.Split(l, "{")
			wfName := s[0]
			conditions := []Condition{}
			conditionsSplit := strings.Split(s[1][:len(s[1])-1], ",")
			last := len(conditionsSplit) - 1
			for i, c := range conditionsSplit {
				if i == last {
					//the last one is always the name of the next workflow with no additional conditions
					condition := Condition{
						Target: c,
					}
					conditions = append(conditions, condition)
					continue
				}
				spl := strings.Split(c, ":")
				target := spl[1]
				symbol := string(spl[0][0])
				operation := string(spl[0][1])
				value, _ := strconv.Atoi(spl[0][2:])

				condition := Condition{
					Target:    target,
					Operation: operation,
					Value:     value,
					Symbol:    symbol,
				}
				conditions = append(conditions, condition)
			}
			wf[wfName] = WF{Name: wfName, Conditions: conditions}
		} else {
			//parse parts
			sp := strings.Split(l[1:len(l)-1], ",")
			x, _ := strconv.Atoi(sp[0][2:])
			m, _ := strconv.Atoi(sp[1][2:])
			a, _ := strconv.Atoi(sp[2][2:])
			s, _ := strconv.Atoi(sp[3][2:])
			part := Part{X: x, M: m, A: a, S: s}
			parts = append(parts, part)
		}
	}

	return wf, parts
}
