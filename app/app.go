package main

import "fmt"
import "github.com/aplikadote/mango/core/tuple"

//import "github.com/aplikadote/mango/core/tupletype"
import "github.com/vjeantet/jodaTime"

func main() {
	from, _ := jodaTime.Parse("dd/MM/YYYY", "01/01/2000")
	to, _ := jodaTime.Parse("dd/MM/YYYY", "02/01/2000")

	t := tuple.Tuple{From: from, To: to}
	fmt.Println(t.String())

}
