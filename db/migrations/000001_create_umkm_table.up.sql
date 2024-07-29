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

CREATE TABLE umkm(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    no_npwp VARCHAR(255) NOT NULL,
    gambar JSONB NOT NULL,
    kategori_umkm_id JSONB NOT NULL,
    nama_penanggung_jawab VARCHAR(255) NOT NULL,
    informasi_jambuka JSONB NOT NULL,
    no_kontak VARCHAR(255) NOT NULL,
    lokasi VARCHAR(255) NOT NULL,
    maps VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_umkm_user FOREIGN KEY ( kategori_umkm_id) REFERENCES kategori_umkm(id)
);