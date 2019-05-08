package main

import (
	"regexp"
	"fmt"
	"strings"
)


func main() {
	var re = regexp.MustCompile(`\${([^}]+)}`)
	var str = `hello_${branch}_${branch}_world`

	var matches = re.FindStringSubmatch(str)
	fmt.Println(matches[0])  // ${branch}
	fmt.Println(matches[1])  // branch

	fmt.Println(strings.Replace(str,  matches[0], "xxx", 1))

	for i, match := range re.FindAllStringSubmatch(str, -1) {
		fmt.Println(match, "found at index", i)
	}
	/**
	hello_xxx_${branch}_world
	[${branch} branch] found at index 0
	[${branch} branch] found at index 1
	*/
}
