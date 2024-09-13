package entity

// type DokumenLegalEntity struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// }

// func ToDokumenLegalEntity(dokumenlegal domain.UmkmDokumen) DokumenLegalEntity {
// 	return DokumenLegalEntity{
// 		Id: dokumenlegal.Id,
// 		Name: dokumenlegal.DokumenMaster.Name,
// 	}
// }

// func ToDokuemenLegalEntities(kategoriList []domain.UmkmDokumen) []DokumenLegalEntity {
//     var kategoriEntities []DokumenLegalEntity
//     for _, kategori := range kategoriList {
//         kategoriEntities = append(kategoriEntities, ToDokumenLegalEntity(kategori))
//     }
//     return kategoriEntities
// }