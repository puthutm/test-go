-- Seeders Script untuk Akun Demo Per Role (SIA AKADEMIK UNSIA)
-- PostgreSQL Database: akademik_db

CREATE TABLE IF NOT EXISTS mst_users (
    id VARCHAR(36) PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    role_name VARCHAR(50) NOT NULL,
    created_at BIGINT DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT
);

-- 1. Role: Mahasiswa
INSERT INTO mst_users (id, username, name, email, role_name)
VALUES ('user-mahasiswa-01', '200101001', 'Budi Santoso (Mahasiswa)', 'budi.santoso@student.unsia.ac.id', 'mahasiswa')
ON CONFLICT (username) DO NOTHING;

-- 2. Role: Dosen
INSERT INTO mst_users (id, username, name, email, role_name)
VALUES ('user-dosen-01', '0401018501', 'Dr. Ahmad Fauzi, M.Kom (Dosen)', 'ahmad.fauzi@lecturer.unsia.ac.id', 'dosen')
ON CONFLICT (username) DO NOTHING;

-- 3. Role: Kaprodi (Ketua Program Studi)
INSERT INTO mst_users (id, username, name, email, role_name)
VALUES ('user-kaprodi-01', '0415088202', 'Siti Rahmawati, S.Kom., M.T. (Kaprodi)', 'siti.rahmawati@unsia.ac.id', 'kaprodi')
ON CONFLICT (username) DO NOTHING;

-- 4. Role: Admin Akademik
INSERT INTO mst_users (id, username, name, email, role_name)
VALUES ('user-akademik-01', 'adminakademik', 'Staf Akademik Pusat (Admin)', 'akademik@unsia.ac.id', 'akademik')
ON CONFLICT (username) DO NOTHING;
