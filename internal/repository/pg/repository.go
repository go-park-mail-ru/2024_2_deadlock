package pg

import "github.com/go-park-mail-ru/2024_2_deadlock/internal/adapters"

type CommonRepo struct {
	PG *adapters.AdapterPG
}

type MinioRepo struct {
	MinioAdapter *adapters.MinioAdapter
}
