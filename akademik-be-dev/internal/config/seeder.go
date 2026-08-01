package config

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserSeederModel struct {
	ID       string `gorm:"primaryKey;column:id"`
	Username string `gorm:"column:username"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Role     string `gorm:"column:role_name"`
}

func (UserSeederModel) TableName() string {
	return "mst_users"
}

func RunAutoSeeder(db *gorm.DB, log *logrus.Logger) {
	if db == nil {
		return
	}

	_ = db.AutoMigrate(&UserSeederModel{})

	demoUsers := []UserSeederModel{
		{
			ID:       "user-mahasiswa-01",
			Username: "200101001",
			Name:     "Budi Santoso (Mahasiswa)",
			Email:    "budi.santoso@student.unsia.ac.id",
			Role:     "mahasiswa",
		},
		{
			ID:       "user-dosen-01",
			Username: "0401018501",
			Name:     "Dr. Ahmad Fauzi, M.Kom (Dosen)",
			Email:    "ahmad.fauzi@lecturer.unsia.ac.id",
			Role:     "dosen",
		},
		{
			ID:       "user-kaprodi-01",
			Username: "0415088202",
			Name:     "Siti Rahmawati, S.Kom., M.T. (Kaprodi)",
			Email:    "siti.rahmawati@unsia.ac.id",
			Role:     "kaprodi",
		},
		{
			ID:       "user-akademik-01",
			Username: "adminakademik",
			Name:     "Staf Akademik Pusat (Admin)",
			Email:    "akademik@unsia.ac.id",
			Role:     "akademik",
		},
	}

	for _, u := range demoUsers {
		var count int64
		db.Model(&UserSeederModel{}).Where("username = ?", u.Username).Count(&count)
		if count == 0 {
			if err := db.Create(&u).Error; err != nil {
				if log != nil {
					log.Errorf("⚠️ Gagal seeding akun demo %s: %v", u.Username, err)
				}
			} else {
				if log != nil {
					log.Infof("✨ [AutoSeeder] Berhasil memasukkan akun demo [%s]: %s (%s)", u.Role, u.Name, u.Username)
				}
			}
		}
	}
}
