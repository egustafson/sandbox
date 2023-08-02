package main

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

const (
	// search_pat_tmpl represents a Regular Expression, with a golang template
	// embedded in the RE.
	search_pat_tmpl = `cn=(?P<group>[\w]+),ou=workers,dc={{.Domain}},dc=org`
)

func MakeRE(domain string) (*regexp.Regexp, error) {
	// Place the `domain` parameter into a map, keyed as the template parameter name, 'Domain'
	params := map[string]string{"Domain": domain}

	// compile the template
	tmpl, err := template.New("pattern").Parse(search_pat_tmpl)
	if err != nil {
		return nil, err
	}

	// execute the template with the parameter
	var pat strings.Builder
	err = tmpl.Execute(&pat, params)
	if err != nil {
		return nil, err
	}

	// compile the resulting pattern into a RE
	return regexp.Compile(pat.String())
}

func main() {
	fmt.Println("done.")
}
