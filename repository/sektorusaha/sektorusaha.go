package sektorusaharepo

import domain "umkm/model/domain/master"

type SektorUsaha interface {
	CreateSektorUsaha(sektorusaha domain.SektorUsaha) (domain.SektorUsaha, error)
	GetSektorUsaha() ([]domain.SektorUsaha, error)
	GetStatusTempatUsaha() ([]domain.StatusTempatUsaha, error)
	GetBentukUsaha() ([]domain.BentukUsaha, error)
}