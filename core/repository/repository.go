package repository

import "github.com/aplikadote/mango/core/tuple"

type Repository struct {
	Name     string
	Database map[string][]tuple.Tuple
}
