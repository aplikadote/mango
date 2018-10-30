package main

import "fmt"
import _ "github.com/aplikadote/mango/core/tuple"
import _ "github.com/aplikadote/mango/core/tupletype"
import "github.com/aplikadote/mango/core/subsystem"
import _ "github.com/vjeantet/jodaTime"

func main() {
	//	from, _ := jodaTime.Parse("dd/MM/YYYY", "01/01/2000")
	//	to, _ := jodaTime.Parse("dd/MM/YYYY", "02/01/2000")

	//	t := tuple.Tuple{From: from, To: to}
	//	fmt.Println(t.String())

	//	fmt.Println(tupletype.MC.Name())
	//	fmt.Println(tupletype.MC.Fullname())

	ss := subsystem.Init("Linea", "root")
	e1 := subsystem.Init("Equipo 1", "e1")
	e2 := subsystem.Init("Equipo 2", "e2")
	e3 := subsystem.Init("Equipo 3", "e3")
	e4 := subsystem.Init("Equipo 4", "e4")

	ss.AddChild(e1)
	ss.AddChild(e2)
	ss.AddChild(e3)
	ss.AddChild(e4)

	fmt.Println(ss.Children())
	ss.DeleteChild(e3)
	fmt.Println(ss.Children())

	array := [4]int{1, 2, 3, 4}
	byref(array[:])
}

func byref(x []int) {
	x[3] = 8
}
