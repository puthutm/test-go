package main

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserSeed struct {
	ID       string `gorm:"primaryKey;column:id"`
	Username string `gorm:"column:username"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Role     string `gorm:"column:role_name"`
}

func (UserSeed) TableName() string {
	return "mst_users"
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=akademik_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("⚠️ Konek DB gagal: %v", err)
	} else {
		log.Println("✅ Terhubung ke PostgreSQL, memulai seeding data akun...")
		_ = db.AutoMigrate(&UserSeed{})
	}

	demoUsers := []UserSeed{
		{
			ID:       uuid.New().String(),
			Username: "200101001",
			Name:     "Budi Santoso (Mahasiswa)",
			Email:    "budi.santoso@student.unsia.ac.id",
			Role:     "mahasiswa",
		},
		{
			ID:       uuid.New().String(),
			Username: "0401018501",
			Name:     "Dr. Ahmad Fauzi, M.Kom (Dosen)",
			Email:    "ahmad.fauzi@lecturer.unsia.ac.id",
			Role:     "dosen",
		},
		{
			ID:       uuid.New().String(),
			Username: "0415088202",
			Name:     "Siti Rahmawati, S.Kom., M.T. (Kaprodi)",
			Email:    "siti.rahmawati@unsia.ac.id",
			Role:     "kaprodi",
		},
		{
			ID:       uuid.New().String(),
			Username: "adminakademik",
			Name:     "Staf Akademik Pusat (Admin)",
			Email:    "akademik@unsia.ac.id",
			Role:     "akademik",
		},
	}

	if db != nil {
		for _, u := range demoUsers {
			if err := db.Where("username = ?", u.Username).FirstOrCreate(&u).Error; err != nil {
				log.Printf("Gagal seeding user %s: %v", u.Username, err)
			} else {
				log.Printf("✨ Seeding berhasil untuk role [%s]: %s (User ID: %s)", u.Role, u.Name, u.Username)
			}
		}
	}

	log.Println("🎉 Seeding selesai!")
}
