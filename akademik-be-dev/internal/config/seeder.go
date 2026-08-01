package config

import (
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"unsia.ac.id/akademic_be/internal/model"
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

	// 1. AutoMigrate Database Schemas
	err := db.AutoMigrate(
		&UserSeederModel{},
		&model.MstSubject{},
		&model.MstClass{},
		&model.MstClassLecturer{},
		&model.MstClassParticipant{},
		&model.MstStudentBio{},
		&model.MstSKSLimit{},
		&model.MstValueScale{},
		&model.MstValueComposition{},
	)
	if err != nil && log != nil {
		log.Warnf("AutoMigrate warning: %v", err)
	}

	// 2. Seed Demo Users
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

	// 3. Seed Demo Sample Subjects if empty
	var subjectCount int64
	db.Model(&model.MstSubject{}).Count(&subjectCount)
	if subjectCount == 0 {
		now := time.Now().Unix()
		currYearID := uuid.New()
		courseTypeID := uuid.New()
		courseGroupID := uuid.New()
		fieldStudyID := uuid.New()

		demoSubjects := []model.MstSubject{
			{
				ID:               uuid.New(),
				CurriculumYearID: currYearID,
				StudyProgramID:   "prodi-if-01",
				Code:             "IF101",
				NameID:           "Algoritma dan Pemrograman",
				NameEN:           "Algorithm and Programming",
				CourseTypeID:     courseTypeID,
				CourseGroupID:    courseGroupID,
				FaceToFaceSKS:    3,
				TotalSKS:         3,
				FieldOfStudiesID: fieldStudyID,
				CreatedAt:        now,
			},
			{
				ID:               uuid.New(),
				CurriculumYearID: currYearID,
				StudyProgramID:   "prodi-if-01",
				Code:             "IF102",
				NameID:           "Basis Data Lanjut",
				NameEN:           "Advanced Database",
				CourseTypeID:     courseTypeID,
				CourseGroupID:    courseGroupID,
				FaceToFaceSKS:    3,
				TotalSKS:         3,
				FieldOfStudiesID: fieldStudyID,
				CreatedAt:        now,
			},
			{
				ID:               uuid.New(),
				CurriculumYearID: currYearID,
				StudyProgramID:   "prodi-if-01",
				Code:             "IF103",
				NameID:           "Pemrograman Web Modern",
				NameEN:           "Modern Web Programming",
				CourseTypeID:     courseTypeID,
				CourseGroupID:    courseGroupID,
				FaceToFaceSKS:    4,
				TotalSKS:         4,
				FieldOfStudiesID: fieldStudyID,
				CreatedAt:        now,
			},
		}

		for _, sbj := range demoSubjects {
			if err := db.Create(&sbj).Error; err != nil {
				if log != nil {
					log.Errorf("⚠️ Gagal seeding sample subject %s: %v", sbj.Code, err)
				}
			} else {
				if log != nil {
					log.Infof("✨ [AutoSeeder] Berhasil memasukkan mata kuliah demo: %s - %s", sbj.Code, sbj.NameID)
				}
			}
		}
	}
}
