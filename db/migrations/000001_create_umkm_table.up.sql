    CREATE TABLE users(
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        role VARCHAR(255) NOT NULL,
        no_phone VARCHAR(255) NOT NULL,
        picture VARCHAR(255) DEFAULT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE kategori_umkm(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- CREATE TABLE umkm(
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     name VARCHAR(255) NOT NULL,
--     no_npwp VARCHAR(255) NOT NULL,
--     gambar JSONB NOT NULL,
--     kategori_umkm_id JSONB NOT NULL,
--     nama_penanggung_jawab VARCHAR(255) NOT NULL,
--     informasi_jambuka JSONB NOT NULL,
--     no_kontak VARCHAR(255) NOT NULL,
--     lokasi VARCHAR(255) NOT NULL,
--     maps VARCHAR(255) NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
-- );

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
    tanggal DATE NOT NULL,
    name_client VARCHAR(255) NOT NULL,
    no_hp VARCHAR(255) NOT NULL
    id_kategori_produk JSONB NOT NULL,
    total_jml NUMERIC(15,2) NOT NULL,
    keterangan text NOT NULL,
    status INT NOT NULL,
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