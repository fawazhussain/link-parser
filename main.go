package main

import (
	"fmt"
	"strings"

	"github.com/fawazhussain/link_parser/link"
)

var exampleHtml = `
<html>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="testing 123">
      Check me 
    </a>
    <a href="out">
      on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", links)

}
