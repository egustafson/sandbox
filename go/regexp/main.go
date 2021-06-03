package main

import (
	"fmt"
	"regexp"
)

func main() {
	// no-op: see test code for execution
	fmt.Println("done.")
}

// Example input:
// cn=group1,ou=Services,ou=division2,ou=Environments,dc=elfwerks,dc=org
// uid=user3,ou=People,dc=elfwerks,dc=org

const (
	group_pat  = `^cn=(?P<gr>[\w-]+),ou=Services,ou=(?P<st>[\w-]+),ou=Environments,dc=elfwerks,dc=org$`
	person_pat = `^uid=(?P<uid>[\w-]+),ou=People,dc=elfwerks,dc=org$`
)

var (
	group_re  = regexp.MustCompile(group_pat)
	person_re = regexp.MustCompile(person_pat)
)

func Parse(input string) (id string, div string) {
	m := group_re.FindStringSubmatch(input)
	if m != nil {
		return m[1], m[2]
	}
	m = person_re.FindStringSubmatch(input)
	if m != nil {
		return m[1], ""
	}
	return "", ""
}
