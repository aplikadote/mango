package repository

import "fmt"
import "testing"
import "github.com/aplikadote/mango/core/tuple"

func TestRepo(t *testing.T) {
	r := Repository{}
	r.Name = "oli"

	if r.Name != "oli" {
		t.Error("el nombre debe ser 'oli'")
	}

	r.Database = make(map[string][]tuple.Tuple)
	r.Database["e1"] = make([]tuple.Tuple, 0)
	r.Database["e1"] = append(r.Database["e1"], tuple.N)
	r.Database["e1"] = append(r.Database["e1"], tuple.Tuple{})

	fmt.Println(len(r.Database["e1"]))
}
