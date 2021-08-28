package main

import (
	"fmt"
	"strings"

	"github.com/fawazhussain/link_parser/link"
)

var exampleHtml = `
<html>
<body>
	<h1>Hello</h1>
	<a href="/other-page">Link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(links)

}
