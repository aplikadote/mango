package main

import (
	"fmt"
	"os"

	"github.com/aplikadote/mango/core/subsystem"
	_ "github.com/aplikadote/mango/core/tuple"
	_ "github.com/vjeantet/jodaTime"
)

func main() {
	//	f, err := os.Open("/home/rgonzalez/TRANSPORTE/addons/structure.csv")
	f, err := os.Open("D:\\TRANSPORTE\\addons\\structure.csv")

	if err != nil {
		fmt.Println(err)
		fmt.Println(f)
		return
	}
	chain := subsystem.ImportChain(f)

	//	chain.RootSubsystem().PreOrden(func(ss *subsystem.Subsystem) {
	//		fmt.Println(ss.Parent())
	//	})

	f1, _ := os.Create("D:\\exportTest.csv")
	subsystem.ExportChain(f1, chain)
}
