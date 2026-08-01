package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MstAcademicPeriod struct {
	ID                     string `json:"id"`
	Code                   string `json:"code"`
	AcademicYearId         string `json:"academic_year_id"`
	AcademicYear           string `json:"academic_year"`
	SemesterId             string `json:"semester_id"`
	Semester               string `json:"semester"`
	Fullname               string `json:"fullname"`
	Shortname              string `json:"shortname"`
	StartDateOfCollege     string `json:"start_date_of_college"`
	EndDateOfCollege       string `json:"end_date_of_college"`
	StartDateOfUts         string `json:"start_date_of_uts"`
	EndDateOfUts           string `json:"end_date_of_uts"`
	StartDateOfUas         string `json:"start_date_of_uas"`
	EndDateOfUas           string `json:"end_date_of_uas"`
	NumberOfLectureMeeting string `json:"number_of_lecture_meeting"`
	IsActive               bool   `json:"is_active"`
	CreatedAt              int64  `json:"created_at"`
	UpdatedAt              int64  `json:"updated_at"`
}

type MstAcademicPeriodDetail struct {
	ID                     string `json:"id"`
	Code                   string `json:"code"`
	AcademicYearId         string `json:"academic_year_id"`
	AcademicYear           string `json:"academic_year"`
	SemesterId             string `json:"semester_id"`
	Semester               string `json:"semester"`
	Fullname               string `json:"fullname"`
	Shortname              string `json:"shortname"`
	StartDateOfCollege     string `json:"start_date_of_college"`
	EndDateOfCollege       string `json:"end_date_of_college"`
	StartDateOfUts         string `json:"start_date_of_uts"`
	EndDateOfUts           string `json:"end_date_of_uts"`
	StartDateOfUas         string `json:"start_date_of_uas"`
	EndDateOfUas           string `json:"end_date_of_uas"`
	NumberOfLectureMeeting string `json:"number_of_lecture_meeting"`
	IsActive               bool   `json:"is_active"`
	CreatedAt              int64  `json:"created_at"`
	UpdatedAt              int64  `json:"updated_at"`
}

type MstAcademicPeriodDetailWithSession struct {
	ID                     string `json:"id"`
	Code                   string `json:"code"`
	AcademicYearId         string `json:"academic_year_id"`
	AcademicYear           string `json:"academic_year"`
	SemesterId             string `json:"semester_id"`
	Semester               string `json:"semester"`
	Fullname               string `json:"fullname"`
	Shortname              string `json:"shortname"`
	StartDateOfCollege     string `json:"start_date_of_college"`
	EndDateOfCollege       string `json:"end_date_of_college"`
	StartDateOfUts         string `json:"start_date_of_uts"`
	EndDateOfUts           string `json:"end_date_of_uts"`
	StartDateOfUas         string `json:"start_date_of_uas"`
	EndDateOfUas           string `json:"end_date_of_uas"`
	NumberOfLectureMeeting string `json:"number_of_lecture_meeting"`
	IsActive               bool   `json:"is_active"`
	UtsSession             int    `json:"uts_session"`
	UasSession             int    `json:"uas_session"`
	CreatedAt              int64  `json:"created_at"`
	UpdatedAt              int64  `json:"updated_at"`
}

type MstAcademicPeriodExport struct {
	ID                     string `json:"id"`
	Code                   string `json:"code"`
	AcademicYearId         string `json:"academic_year_id"`
	AcademicYear           string `json:"academic_year"`
	SemesterId             string `json:"semester_id"`
	Semester               string `json:"semester"`
	Fullname               string `json:"fullname"`
	Shortname              string `json:"shortname"`
	StartDateOfCollege     string `json:"start_date_of_college"`
	EndDateOfCollege       string `json:"end_date_of_college"`
	StartDateOfUts         string `json:"start_date_of_uts"`
	EndDateOfUts           string `json:"end_date_of_uts"`
	StartDateOfUas         string `json:"start_date_of_uas"`
	EndDateOfUas           string `json:"end_date_of_uas"`
	NumberOfLectureMeeting string `json:"number_of_lecture_meeting"`
	IsActive               bool   `json:"is_active"`
}

type MstAcademicPeriodSearch struct {
	ID                     string `json:"id"`
	Code                   string `json:"code"`
	AcademicYearId         string `json:"academic_year_id"`
	AcademicYear           string `json:"academic_year"`
	SemesterId             string `json:"semester_id"`
	Semester               string `json:"semester"`
	Fullname               string `json:"fullname"`
	Shortname              string `json:"shortname"`
	StartDateOfCollege     string `json:"start_date_of_college"`
	EndDateOfCollege       string `json:"end_date_of_college"`
	StartDateOfUts         string `json:"start_date_of_uts"`
	EndDateOfUts           string `json:"end_date_of_uts"`
	StartDateOfUas         string `json:"start_date_of_uas"`
	EndDateOfUas           string `json:"end_date_of_uas"`
	NumberOfLectureMeeting string `json:"number_of_lecture_meeting"`
	IsActive               bool   `json:"is_active"`
}

type MstAcademicPeriodRelation struct {
	ID                     string `json:"id"`
	Code                   string `json:"code"`
	AcademicYearId         string `json:"academic_year_id"`
	AcademicYear           string `json:"academic_year"`
	SemesterId             string `json:"semester_id"`
	Semester               string `json:"semester"`
	Fullname               string `json:"fullname"`
	Shortname              string `json:"shortname"`
	StartDateOfCollege     string `json:"start_date_of_college"`
	EndDateOfCollege       string `json:"end_date_of_college"`
	StartDateOfUts         string `json:"start_date_of_uts"`
	EndDateOfUts           string `json:"end_date_of_uts"`
	StartDateOfUas         string `json:"start_date_of_uas"`
	EndDateOfUas           string `json:"end_date_of_uas"`
	NumberOfLectureMeeting string `json:"number_of_lecture_meeting"`
	IsActive               bool   `json:"is_active"`
}

/* Action */
func GetAcademicPeriods(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAcademicPeriod, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_academic_periods_get", filter, sortBy, sortDirection, page, pageSize, &MstAcademicPeriod{})
	if err != nil {
		return []MstAcademicPeriod{}, 0, err
	}

	var modelResults []MstAcademicPeriod
	for _, item := range results {
		level, ok := item.(*MstAcademicPeriod)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchAcademicPeriods(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAcademicPeriodSearch, error) {
	var results []MstAcademicPeriodSearch
	err := handlers.SPGet("sp_mst_academic_periods_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstAcademicPeriodSearch{}, err
	}
	return results, nil
}

func ExportAcademicPeriods(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstAcademicPeriodExport
	err := handlers.SPGet("sp_mst_academic_periods_get", "", "name", "asc", 1, CountAcademicPeriods(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":                     result.ID,
			"Code":                   result.Code,
			"AcademicYearId":         result.AcademicYearId,
			"SemesterId":             result.SemesterId,
			"Fullname":               result.Fullname,
			"Shortname":              result.Shortname,
			"StartDateOfCollege":     result.StartDateOfCollege,
			"EndDateOfCollege":       result.EndDateOfCollege,
			"StartDateOfUts":         result.StartDateOfUts,
			"EndDateOfUts":           result.EndDateOfUts,
			"StartDateOfUas":         result.StartDateOfUas,
			"EndDateOfUas":           result.EndDateOfUas,
			"NumberOfLectureMeeting": result.NumberOfLectureMeeting,
			"IsActive":               result.IsActive,
		}
	}

	headers := []string{
		"ID", "Code", "Academic Year Id", "Semester Id", "Full Name", "Short Name", "Start Date Of College", "End Date Of College", "Start Date Of UTS", "End Date Of UTS", "Start Date Of UAS", "End Date Of UAS", "Number Of Lecture Meeting", "Is Active",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetAcademicPeriod(id string) (MstAcademicPeriodDetail, error) {
	var result MstAcademicPeriodDetail
	err := handlers.SPGetByID("sp_mst_academic_periods_get_by_id", id, &result)
	if err != nil {
		return MstAcademicPeriodDetail{}, err
	}
	return result, nil
}

func GetAcademicPeriodDetailWithSession(id string) (MstAcademicPeriodDetailWithSession, error) {
	var result MstAcademicPeriodDetailWithSession
	err := handlers.SPGetByID("sp_mst_academic_periods_get_by_id_with_detail_session", id, &result)
	if err != nil {
		return MstAcademicPeriodDetailWithSession{}, err
	}
	return result, nil
}

func CreateAcademicPeriod(id string, code string, academic_year_id string, semester_id string, fullname string, shortname string, start_date_of_college string, end_date_of_college string, start_date_of_uts string, end_date_of_uts string, start_date_of_uas string, end_date_of_uas string, number_of_lecture_meeting string, is_active bool) error {
	return QueryInsertAcademicPeriod(id, code, academic_year_id, semester_id, fullname, shortname, start_date_of_college, end_date_of_college, start_date_of_uts, end_date_of_uts, start_date_of_uas, end_date_of_uas, number_of_lecture_meeting, is_active)
}

func ImportAcademicPeriods(fileStatus string) error {
	headers := []string{
		"id", "code", "academic_year_id", "semester_id", "fullname", "shortname", "start_date_of_college", "end_date_of_college", "start_date_of_uts", "end_date_of_uts", "start_date_of_uas", "end_date_of_uas", "number_of_lecture_meeting", "is_active",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		academic_year_id := row["academic_year_id"]
		semester_id := row["semester_id"]
		fullname := row["fullname"]
		shortname := row["shortname"]
		start_date_of_college := row["start_date_of_college"]
		end_date_of_college := row["end_date_of_college"]
		start_date_of_uts := row["start_date_of_uts"]
		end_date_of_uts := row["end_date_of_uts"]
		start_date_of_uas := row["start_date_of_uas"]
		end_date_of_uas := row["end_date_of_uas"]
		number_of_lecture_meeting := row["number_of_lecture_meeting"]
		is_active, err := strconv.ParseBool(row["is_active"])
		if err != nil {
			return err
		}

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstAcademicPeriod{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateAcademicPeriod(id, code, academic_year_id, semester_id, fullname, shortname, start_date_of_college, end_date_of_college, start_date_of_uts, end_date_of_uts, start_date_of_uas, end_date_of_uas, number_of_lecture_meeting, is_active); err != nil {
					return err
				}
			} else {
				if err := QueryInsertAcademicPeriod(id, code, academic_year_id, semester_id, fullname, shortname, start_date_of_college, end_date_of_college, start_date_of_uts, end_date_of_uts, start_date_of_uas, end_date_of_uas, number_of_lecture_meeting, is_active); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstAcademicPeriod{})
			if err != nil {
				return err
			}
			if err := QueryInsertAcademicPeriod(id, code, academic_year_id, semester_id, fullname, shortname, start_date_of_college, end_date_of_college, start_date_of_uts, end_date_of_uts, start_date_of_uas, end_date_of_uas, number_of_lecture_meeting, is_active); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateAcademicPeriod(id string, code string, academic_year_id string, semester_id string, fullname string, shortname string, start_date_of_college string, end_date_of_college string, start_date_of_uts string, end_date_of_uts string, start_date_of_uas string, end_date_of_uas string, number_of_lecture_meeting string, is_active bool) error {
	return QueryUpdateAcademicPeriod(id, code, academic_year_id, semester_id, fullname, shortname, start_date_of_college, end_date_of_college, start_date_of_uts, end_date_of_uts, start_date_of_uas, end_date_of_uas, number_of_lecture_meeting, is_active)
}

func DeleteAcademicPeriod(id string) error {
	err := handlers.SPDelete("sp_mst_academic_periods_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashAcademicPeriods(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAcademicPeriod, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_academic_periods_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstAcademicPeriod{})
	if err != nil {
		return []MstAcademicPeriod{}, 0, err
	}

	var modelResults []MstAcademicPeriod
	for _, item := range results {
		level, ok := item.(*MstAcademicPeriod)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreAcademicPeriod(id string) error {
	err := handlers.SPRestore("sp_mst_academic_periods_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountAcademicPeriods(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstAcademicPeriod{}, nullable)
}

/* Query */
func QueryInsertAcademicPeriod(id string, code string, academic_year_id string, semester_id string, fullname string, shortname string, start_date_of_college string, end_date_of_college string, start_date_of_uts string, end_date_of_uts string, start_date_of_uas string, end_date_of_uas string, number_of_lecture_meeting string, is_active bool) error {
	query := `
		EXEC sp_mst_academic_periods_insert
		@id = ?,
		@code = ?, 
		@academic_year_id = ?, 
		@semester_id = ?, 
		@fullname = ?, 
		@shortname = ?, 
		@start_date_of_college = ?, 
		@end_date_of_college = ?, 
		@start_date_of_uts = ?, 
		@end_date_of_uts = ?, 
		@start_date_of_uas = ?, 
		@end_date_of_uas = ?, 
		@number_of_lecture_meeting = ?, 
		@is_active = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, academic_year_id, semester_id, fullname, shortname, start_date_of_college, end_date_of_college, start_date_of_uts, end_date_of_uts, start_date_of_uas, end_date_of_uas, number_of_lecture_meeting, is_active)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateAcademicPeriod(id string, code string, academic_year_id string, semester_id string, fullname string, shortname string, start_date_of_college string, end_date_of_college string, start_date_of_uts string, end_date_of_uts string, start_date_of_uas string, end_date_of_uas string, number_of_lecture_meeting string, is_active bool) error {
	query := `
		EXEC sp_mst_academic_periods_update
		@id = ?,
		@code = ?, 
		@academic_year_id = ?, 
		@semester_id = ?, 
		@fullname = ?, 
		@shortname = ?, 
		@start_date_of_college = ?, 
		@end_date_of_college = ?, 
		@start_date_of_uts = ?, 
		@end_date_of_uts = ?, 
		@start_date_of_uas = ?, 
		@end_date_of_uas = ?, 
		@number_of_lecture_meeting = ?, 
		@is_active = ?,
		@updated_at = ?,
		@updated_by = ?
	`
	err := handlers.SPUpdate(query, id, code, academic_year_id, semester_id, fullname, shortname, start_date_of_college, end_date_of_college, start_date_of_uts, end_date_of_uts, start_date_of_uas, end_date_of_uas, number_of_lecture_meeting, is_active)
	if err != nil {
		return err
	}

	return nil
}
