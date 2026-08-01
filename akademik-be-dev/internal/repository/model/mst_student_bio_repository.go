package repositorymodel

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
)

type MstStudentBioRepository struct {
	log *logrus.Logger
	repository.Repository[model.MstStudentBio]
	cacheRepository cached.CacheRepository
}

func NewMstStudentBioRepository(
	log *logrus.Logger,
	cacheRepository cached.CacheRepository,
) *MstStudentBioRepository {
	return &MstStudentBioRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

// Create
func (r *MstStudentBioRepository) Create(db *gorm.DB, data *model.MstStudentBio) error {
	data.CreatedAt = time.Now().UnixMilli()
	data.UpdatedAt = time.Now().UnixMilli()
	err := db.Create(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "create student_bio",
			"id":         data.ID,
		}).Error(msg.ErrCreate.Error())
		return err
	}
	return nil
}

// Update general
func (r *MstStudentBioRepository) Update(db *gorm.DB, data *model.MstStudentBio) error {
	data.UpdatedAt = time.Now().UnixMilli()
	err := db.Model(&model.MstStudentBio{}).Where("id = ?", data.ID).Updates(map[string]interface{}{
		"name":           data.Name,
		"back_degree":     data.BackDegree,
		"nik":            data.NIK,
		"no_kk":          data.NoKK,
		"birth_place_id": data.BirthPlaceID,
		"birth_date":     data.BirthDate,
		"gender":         data.Gender,
		"status_id":      data.StatusID,
		"religion_id":    data.ReligionID,
		"ethnic_id":      data.EthnicID,
		"height":         data.Height,
		"weight":         data.Weight,
		"blood_type_id":  data.BloodTypeID,
		"updated_at":     data.UpdatedAt,
	}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update student_bio",
			"id":         data.ID,
		}).Error(msg.ErrUpdate.Error())
		return err
	}
	return nil
}

// Update Information
func (r *MstStudentBioRepository) UpdateInformation(db *gorm.DB, data *model.MstStudentBio) error {
	data.UpdatedAt = time.Now().UnixMilli()
	err := db.Model(&model.MstStudentBio{}).Where("id = ?", data.ID).Updates(map[string]interface{}{
		"private_email":     data.PrivateEmail,
		"phone":             data.Phone,
		"transportation_id": data.TransportationID,
		"citizenship_id":    data.CitizenshipID,
		"almamater_size_id": data.AlmamaterSizeID,
		"job_id":            data.JobID,
		"updated_at":        data.UpdatedAt,
	}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update student_bio information",
			"id":         data.ID,
		}).Error(msg.ErrUpdate.Error())
		return err
	}
	return nil
}

// UpdateCompleteness
func (r *MstStudentBioRepository) UpdateCompleteness(db *gorm.DB, data *model.MstStudentBio) error {
	data.UpdatedAt = time.Now().UnixMilli()
	err := db.Model(&model.MstStudentBio{}).Where("id = ?", data.ID).Updates(map[string]interface{}{
		"no_passport":         data.NoPassport,
		"google_scholar":      data.GoogleScholar,
		"sinta_id":            data.SintaID,
		"scopus_id":           data.ScopusID,
		"signature_path_file": data.SignaturePathFile,
		"updated_at":          data.UpdatedAt,
	}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update student_bio completeness",
			"id":         data.ID,
		}).Error(msg.ErrUpdate.Error())
		return err
	}
	return nil
}

// Update Documents
func (r *MstStudentBioRepository) UpdateDocuments(db *gorm.DB, data *model.MstStudentBio) error {
	data.UpdatedAt = time.Now().UnixMilli()
	err := db.Model(&model.MstStudentBio{}).Where("id = ?", data.ID).Updates(map[string]interface{}{
		"npwp":                     data.NPWP,
		"npwp_filepath":            data.NPWPFilepath,
		"bpjs_healthcare":          data.BPJSHealthcare,
		"bpjs_healthcare_filepath": data.BPJSHealthcareFilepath,
		"bpjs_employment":          data.BPJSEmployment,
		"bpjs_employment_filepath": data.BPJSEmploymentFilepath,
		"updated_at":               data.UpdatedAt,
	}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update student_bio documents",
			"id":         data.ID,
		}).Error(msg.ErrUpdate.Error())
		return err
	}
	return nil
}

// Update BankAccount
func (r *MstStudentBioRepository) UpdateBankAccount(db *gorm.DB, data *model.MstStudentBio) error {
	data.UpdatedAt = time.Now().UnixMilli()
	err := db.Model(&model.MstStudentBio{}).Where("id = ?", data.ID).Updates(map[string]interface{}{
		"bank_id":          data.BankID,
		"account_number":   data.AccountNumber,
		"account_name":     data.AccountName,
		"account_filepath": data.AccountFilepath,
		"updated_at":       data.UpdatedAt,
	}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update student_bio bank account",
			"id":         data.ID,
		}).Error(msg.ErrUpdate.Error())
		return err
	}
	return nil
}

// Delete
func (r *MstStudentBioRepository) DeleteByID(db *gorm.DB, ID string) error {
	now := time.Now().UnixMilli()
	err := db.Model(&model.MstStudentBio{}).Where("id = ?", ID).Update("deleted_at", now).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "delete student_bio",
			"id":         ID,
		}).Error(msg.ErrDelete.Error())
		return err
	}
	return nil
}

// Restore
func (r *MstStudentBioRepository) RestoreByID(db *gorm.DB, ID string) error {
	err := db.Model(&model.MstStudentBio{}).Where("id = ?", ID).Update("deleted_at", nil).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "restore student_bio",
			"id":         ID,
		}).Error(msg.ErrRestore.Error())
		return err
	}
	return nil
}

// GetByID General
func (r *MstStudentBioRepository) GetByID(db *gorm.DB, ID string, data *model.MstStudentBio) error {
	err := db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get by id student_bio",
			"id":         ID,
		}).Error(msg.ErrRestore.Error())
		return err
	}
	return nil
}

// GetByID -> Information
func (r *MstStudentBioRepository) GetInformationByID(db *gorm.DB, ID string, data *model.MstStudentBio) error {
	return r.GetByID(db, ID, data)
}

// GetByID -> Completeneses
func (r *MstStudentBioRepository) GetCompletenessByID(db *gorm.DB, ID string, data *model.MstStudentBio) error {
	return r.GetByID(db, ID, data)
}

// GetByID -> Documents
func (r *MstStudentBioRepository) GetDocumentByID(db *gorm.DB, ID string, data *model.MstStudentBio) error {
	return r.GetByID(db, ID, data)
}

// GetByID -> Bank account
func (r *MstStudentBioRepository) GetBankAccountByID(db *gorm.DB, ID string, data *model.MstStudentBio) error {
	return r.GetByID(db, ID, data)
}

// CheckByUserID
func (r *MstStudentBioRepository) CheckByUserID(db *gorm.DB, UserID string, data *model.MstStudentBio) error {
	key := fmt.Sprintf("akademic-user-id-%s-check", UserID)
	v, err := r.cacheRepository.Get(key)
	if err == nil {
		if err := json.Unmarshal(v, data); err == nil {
			return nil
		}
	}
	err = db.Where("user_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", UserID).First(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "check by user_id student_bio",
			"user_id":    UserID,
		}).Error(msg.ErrRestore.Error())
		return err
	}
	v, _ = json.Marshal(data)
	r.cacheRepository.SetDefaultEx(key, v)
	return nil
}

// GetAllWithCount
func (r *MstStudentBioRepository) GetAllWithCount(
	db *gorm.DB,
	deleted bool,
	req pageable.PageableRequest,
) (T []model.MstStudentBio, count int64, err error) {
	query := db.Model(&model.MstStudentBio{})
	if deleted {
		query = query.Where("deleted_at IS NOT NULL AND deleted_at > 0")
	} else {
		query = query.Where("deleted_at IS NULL OR deleted_at = 0")
	}

	if req.Search != "" {
		query = query.Where("name ILIKE ? OR nik ILIKE ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	err = query.Count(&count).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get all student_bio count",
		}).Error(msg.ErrMultipleRead.Error())
		return T, count, err
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	err = query.Scopes(pageable.Paginate(page, limit)).Find(&T).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get all student_bio",
		}).Error(msg.ErrMultipleRead.Error())
		return T, count, err
	}

	return T, count, nil
}
