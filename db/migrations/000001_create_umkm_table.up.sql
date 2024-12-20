    CREATE TABLE users(
        id SERIAL PRIMARY KEY,
        username VARCHAR(255),
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        role VARCHAR(255) NOT NULL,
        no_phone VARCHAR(255) UNIQUE NOT NULL,
        picture VARCHAR(255) ,
        google_id VARCHAR(255),
        fullnam VARCHAR(255),
        alamat TEXT,
        no_kk 
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE kategori_umkm(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
        CREATE TABLE umkm(
            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
            name VARCHAR(255),
            no_npwp VARCHAR(255),
            gambar JSONB,
            kategori_umkm_id JSONB,
            nama_penanggung_jawab VARCHAR(255),
            informasi_jambuka JSONB,
            no_kontak VARCHAR(255),
            lokasi VARCHAR(255),
            maps JSONB,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );


CREATE TABLE hak_akses(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    umkm_id UUID NOT NULL,
    status INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_hak_akses_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_hak_akses_umkm FOREIGN KEY (umkm_id) REFERENCES umkm(id)
);

CREATE TABLE  save_otps(
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    otp_code VARCHAR(6) NOT NULL,
    status BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL
);

CREATE TABLE transaksi(
    id SERIAL PRIMARY KEY,
    umkm_id UUID NOT NULL,
    no_invoice VARCHAR(255) NOT NULL,
    tanggal TIMESTAMPTZ NOT NULL,
    name_client VARCHAR(255) NOT NULL,
    no_hp VARCHAR(255) NOT NULL,
    id_kategori_produk JSONB NOT NULL,
    total_jml NUMERIC(15,2) NOT NULL,
    keterangan text NOT NULL,
    status INT,
    alasan_perubahan INT,
    tiket_validasi VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_transaksi FOREIGN KEY (umkm_id) REFERENCES umkm(id)
);

CREATE TABLE kategori_produk (
    id SERIAL PRIMARY KEY,
    umkm_id UUID NOT NULL,
    nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_kategori_produk FOREIGN KEY (umkm_id) REFERENCES umkm(id)
);

CREATE TABLE produk(
    id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    umkm_id  UUID NOT NULL,
    nama VARCHAR(255) NOT NULL,
    gambar_id JSONB NOT NULL,
    harga INT NOT NULL,
    satuan INT NOT NULL,
    min_pesanan INT NOT NULL,
    kategori_produk_id JSONB NOT NULL,
    deskripsi text NOT NULL,
    active INT
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_produk_umkm FOREIGN KEY (umkm_id) REFERENCES umkm(id)
);   

CREATE TABLE homepage.sleder(
id SERIAL PRIMARY KEY,
slide_desc VARCHAR(255) NOT NULL,
slide_title VARCHAR(255) NOT NULL,
gambar VARCHAR(255) NOT NULL,
active INT NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);


CREATE TABLE homepage.testimonial(
id SERIAL PRIMARY KEY,
quote Text NOT NULL,
name VARCHAR(255) NOT NULL,
active int,
gambar_testi VARCHAR(255) NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);


CREATE TABLE homepage.brandlogo(
id SERIAL PRIMARY  KEY,
brand_name VARCHAR(255) NOT NULL,
brand_logo VARCHAR(255) NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE homepage.about(
id SERIAL PRIMARY  KEY,
image VARCHAR(255) NOT NULL,
description TEXT NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE master_dokumen_legal(
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    is_wajib INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE umkm_dokumen_legal (
    id SERIAL PRIMARY KEY,
    umkm_id UUID NOT NULL,
    dokumen_id INT NOT NULL,
    dok_upload JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (umkm_id) REFERENCES umkm(id) ON DELETE CASCADE
    FOREIGN KEY (dokumen_id) REFERENCES master_dokumen_legal(id) ON DELETE CASCADE
);

CREATE TABLE homepage.berita (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author) REFERENCES users(id) ON DELETE CASCADE
);


CREATE TABLE master.sektor_usaha(
    id SERIAL PRIMARY KEY,
    urai VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE master.status_tempat_usaha(
    id SERIAL PRIMARY KEY,
    urai VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE master.bentuk_usaha(
    id SERIAL PRIMARY KEY,
    urai VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE omset(
    id SERIAL PRIMARY KEY,
    umkm_id UUID NOT NULL,
    jumlah_omset  NUMERIC(15,2) NOT NULL,
    bulan VARCHAR(20) NOT NULL,
    tahun VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_hak_akses_umkm FOREIGN KEY (umkm_id) REFERENCES umkm(id)
);