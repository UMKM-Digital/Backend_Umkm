package sektorusahaservice

import (
	entity "umkm/model/entity/master"
	web "umkm/model/web/master"
)

type SektorUsaha interface {
	CreateSektorUsaha(sektorusaha web.CreateSektorUsaha) (map[string]interface{}, error)
	GetSektorUsaha() ([]entity.SektorUsahaEntity, error)
	GetStatusTempatUsaha() ([]entity.StatusTempatUsahaEntity, error)
	GetBentukUsaha() ([]entity.BentukUsahaEntity, error)
}