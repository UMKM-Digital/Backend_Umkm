package hakaksesservice

import (
	"fmt"
	"umkm/model/domain"
	"umkm/model/web"
	hakaksesrepo "umkm/repository/hakakses"

	"github.com/google/uuid"
)

type HakaksesServiceImpl struct {
	hakakses hakaksesrepo.CreateHakakses
}

func NewKHakAkesService(hakakses hakaksesrepo.CreateHakakses) *HakaksesServiceImpl {
	return &HakaksesServiceImpl{
		hakakses: hakakses,
	}
}

func (service *HakaksesServiceImpl) UpdateBulkHakAkses(request web.HakAksesUpdate) ([]map[string]interface{}, error) {
    var umkmUUIDs []uuid.UUID
    for _, umkmIdStr := range request.UmkmIds {
        umkmid, err := uuid.Parse(umkmIdStr)
        if err != nil {
            return nil, fmt.Errorf("invalid UUID format for umkm_id: %s", umkmIdStr)
        }
        umkmUUIDs = append(umkmUUIDs, umkmid)
    }

    hakAksesUpdate := domain.HakAkses{
        Status:  domain.StatusEnum(request.Status),
        Pesan:   request.Pesan,
        Pembina: request.Pembina,
    }

    err := service.hakakses.AcceptBulkStatus(umkmUUIDs, hakAksesUpdate)
    if err != nil {
        return nil, err
    }

    // Buat response untuk dikembalikan
    response := []map[string]interface{}{}
    for _, umkmid := range umkmUUIDs {
        response = append(response, map[string]interface{}{
            "umkm_id": umkmid.String(),
            "status":  string(hakAksesUpdate.Status),
            "pesan":   hakAksesUpdate.Pesan,
            "pembina": hakAksesUpdate.Pembina,
        })
    }

    return response, nil
}
