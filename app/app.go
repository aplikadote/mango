package main

import (
	"fmt"
	"os"

	"github.com/aplikadote/mango/core/subsystem"
	_ "github.com/aplikadote/mango/core/tuple"
	_ "github.com/vjeantet/jodaTime"
)

func main() {
	//	from, _ := jodaTime.Parse("dd/MM/YYYY", "01/01/2000")
	//	to, _ := jodaTime.Parse("dd/MM/YYYY", "02/01/2000")

	//	t := tuple.Tuple{From: from, To: to}
	//	fmt.Println(t.String())

	//	fmt.Println(tupletype.MC.Name())
	//	fmt.Println(tupletype.MC.Fullname())

	//	ss := subsystem.NewSubsystem("Linea", "root")
	//	e1 := subsystem.NewSubsystem("Equipo 1", "e1")
	//	e2 := subsystem.NewSubsystem("Equipo 2", "e2")
	//	e3 := subsystem.NewSubsystem("Equipo 3", "e3")
	//	e4 := subsystem.NewSubsystem("Equipo 4", "e4")

	//	ss.AddChild(e1)
	//	ss.AddChild(e2)
	//	ss.AddChild(e3)
	//	ss.AddChild(e4)

	//	fmt.Println(ss.Children())
	//	ss.DeleteChild(e3)
	//	fmt.Println(ss.Children())

	//	array := [4]int{1, 2, 3, 4}
	//	byref(array[:])

	f, _ := os.Open("/home/rgonzalez/TRANSPORTE/addons/structure.csv")
	chain := subsystem.ImportChain(f)

	printss := func(ss *subsystem.Subsystem) {
		fmt.Println(ss)
	}

	chain.RootSubsystem().PreOrden(printss)
}
