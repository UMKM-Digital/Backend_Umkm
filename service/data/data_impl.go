package dataservice

import (
	"context"
	"fmt"
	datarepo "umkm/repository/data"
)

type DataServiceImpl struct {
	datarepository datarepo.DataUserRepo
}

func NewDataservice(datarepository datarepo.DataUserRepo) *DataServiceImpl {
	return &DataServiceImpl{
		datarepository: datarepository,
	}
}

func (service *DataServiceImpl) CountAtas() (map[string]interface{}, error) {
    // Memanggil fungsi dari repository untuk menghitung persentase gender
    dataResult, err := service.datarepository.CountProductByCategoryWithPercentage()
    if err != nil {
        return nil, err
    }

	dataUmkmVerify, err := service.datarepository.CountWaitingVerify()
	if err != nil {
        return nil, err
    }

	dataUmkmBina, err := service.datarepository.CountUmkmBina()
	if err != nil {
        return nil, err
    }

	dataUmkmTertolak, err := service.datarepository.CountUmkmTertolak()
	if err != nil {
        return nil, err
    }


	dataUmkm, err := service.datarepository.TotalUmkm()
	if err != nil {
        return nil, err
    }

	dataMikro, err := service.datarepository.TotalMikro()
	if err != nil {
        return nil, err
    }

	dataMenengah, err := service.datarepository.TotalMenengah()
	if err != nil {
        return nil, err
    }

	dataKecil, err := service.datarepository.TotalKecil()
	if err != nil {
        return nil, err
    }

	dataSektorjasa, err := service.datarepository.TotalSektorJasa()
	if err != nil {
        return nil, err
    }

	dataSektorPoruduksi, err := service.datarepository.TotalSektorProduksi()
	if err != nil {
        return nil, err
    }

	dataSektorPerdagangan, err := service.datarepository.TotalSektorPerdagangan()
	if err != nil {
        return nil, err
    }


	dataSektorEkonomiKreatif, err := service.datarepository.TotalEkonomiKreatif()
	if err != nil {
        return nil, err
    }

    // Menggabungkan hasil gender, study, dan age dalam satu map
    result := map[string]interface{}{
        "total_produk": dataResult,
		"total_menunggu_Verifikasi": dataUmkmVerify,
		"total_umkm_bina": dataUmkmBina,
		"total_umkm_tertolak": dataUmkmTertolak,
		"total_umkm": dataUmkm,
		"total_umkm_mikro": dataMikro,
		"total_umkm_menengah": dataMenengah,
		"total_umkm_kecil": dataKecil,
		"total_umkm_sektor_jasa": dataSektorjasa,
		"total_umkm_sektor_produksi": dataSektorPoruduksi,
		"total_umkm_sektor_perdagangan": dataSektorPerdagangan,
		"total_umkm_ekonomi_kreatif": dataSektorEkonomiKreatif,
    }

    // Menambahkan struktur respons sesuai dengan yang Anda inginkan

    return result, nil
}


func (service *DataServiceImpl) GrafikKategoriBySektor(ctx context.Context, sektorUsahaID int, kecamatan, kelurahan string, tahun int) ([]datarepo.KategoriCount, error) {
    if sektorUsahaID <= 0 {
        return nil, fmt.Errorf("invalid sektor usaha ID")
    }

    // Panggil repository untuk mengambil data kategori UMKM berdasarkan sektor, kecamatan, dan kelurahan
    result, err := service.datarepository.GrafikKategoriBySektorUsaha(ctx, sektorUsahaID, kecamatan, kelurahan, tahun)
    if err != nil {
        return nil, err
    }

    return result, nil
}


func (service *DataServiceImpl) TotalUmkmKriteriaUsahaPerBulan(tahun int) (map[string]map[string]int64, error) {
    return service.datarepository.TotalUmkmKriteriaUsahaPerBulan(tahun)
}

func(service *DataServiceImpl) TotalUmkmBinaan()(map[string]interface{}, error){
    dataUmkmBulan, err := service.datarepository.TotalUmkmBulan()
    if err != nil {
        return nil, err
    }

	dataUmkmBulalnLalu, err := service.datarepository.TotalUmkmBulanLalu()
	if err != nil {
        return nil, err
    }

	dataUmkmTahun, err := service.datarepository.TotalUmkmTahun()
	if err != nil {
        return nil, err
    }

	dataUmkmTahunLalu, err := service.datarepository.TotalUmkmTahunLalu()
	if err != nil {
        return nil, err
    }

    dataPersentasuBulan, _ := service.datarepository.PersentasiKenaikanUmkm()
    dataPersentasuTahun, _ := service.datarepository.PersentasiKenaikanUmkmTahun()

    result := map[string]interface{}{
        "total_umkm_bulan_ini": dataUmkmBulan,
		"total_umkm_bulan_lalu": dataUmkmBulalnLalu,
        "total_umkm_tahun_ini": dataUmkmTahun,
        "total_umkm_tahun_lalu": dataUmkmTahunLalu,
        "total_umkm_persentasi_bulan": dataPersentasuBulan,
        "total_umkm_persentasi_tahun": dataPersentasuTahun,
    }

    // Menambahkan struktur respons sesuai dengan yang Anda inginkan

    return result, nil
}