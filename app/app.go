package app

import (
	"fmt"
	"os"

	"github.com/aplikadote/mango/core/subsystem"
	_ "github.com/aplikadote/mango/core/tuple"
	_ "github.com/vjeantet/jodaTime"
)

func StartApp() {
	f, err := os.Open("/home/rgonzalez/TRANSPORTE/addons/structure.csv")
	//f, err := os.Open("D:\\TRANSPORTE\\addons\\structure.csv")

	if err != nil {
		panic(err)
	}

	chain, err := subsystem.ImportChain(f)
	if err != nil {
		panic(err)
	}

	chain.RootSubsystem().PreOrden(func(ss *subsystem.Subsystem) {
		fmt.Println(ss.Name)
	})

	//	efilename := "/home/rgonzalez/exportTest.csv"
	//	f1, _ := os.Create(efilename)
	//	subsystem.ExportChain(f1, chain)
}
