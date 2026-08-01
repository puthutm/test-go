package routes

import (
	controllers "data-referensi/app/controllers/education"
	"data-referensi/app/middlewares"
	"data-referensi/app/requests"

	"github.com/gofiber/fiber/v2"
)

func EducationRoute(app fiber.Router) {
	educationGroup := app.Group("/education", middlewares.Auth)

	/* Educational Levels */
	educationalLevel := educationGroup.Group("educational-levels")
	educationalLevelTrash := educationalLevel.Group("trashs")
	educationalLevelTrash.Get("/", requests.ValidatePagination, controllers.GetTrashEducationalLevels)
	educationalLevelTrash.Put("/:id", controllers.RestoreEducationalLevel)

	educationalLevel.Get("/", requests.ValidatePagination, controllers.GetEducationalLevels)
	educationalLevel.Get("/search", controllers.SearchEducationalLevels)
	educationalLevel.Get("/export", controllers.ExportEducationalLevels)
	educationalLevel.Get("/:id", controllers.GetEducationalLevel)
	educationalLevel.Post("/", requests.ValidateEducationalLevel, controllers.CreateEducationalLevel)
	educationalLevel.Post("/import", controllers.ImportEducationalLevels)
	educationalLevel.Put("/:id", requests.ValidateEducationalLevel, controllers.UpdateEducationalLevel)
	educationalLevel.Delete("/:id", controllers.DeleteEducationalLevel)

	/* Study Programs */
	studyProgram := educationGroup.Group("study-programs")
	studyProgramTrash := studyProgram.Group("trashs")
	studyProgramTrash.Get("/", requests.ValidatePagination, controllers.GetTrashStudyPrograms)
	studyProgramTrash.Put("/:id", controllers.RestoreStudyProgram)

	studyProgram.Get("/", requests.ValidatePagination, controllers.GetStudyPrograms)
	studyProgram.Get("/search", controllers.SearchStudyPrograms)
	studyProgram.Get("/export", controllers.ExportStudyPrograms)
	studyProgram.Get("/:id", controllers.GetStudyProgram)
	studyProgram.Post("/", requests.ValidateStudyProgram, controllers.CreateStudyProgram)
	studyProgram.Post("/import", controllers.ImportStudyPrograms)
	studyProgram.Put("/:id", requests.ValidateStudyProgram, controllers.UpdateStudyProgram)
	studyProgram.Delete("/:id", controllers.DeleteStudyProgram)

	/* Unsia Study Programs */
	unsiaStudyProgram := educationGroup.Group("unsia-study-programs")
	unsiaStudyProgramTrash := unsiaStudyProgram.Group("trashs")
	unsiaStudyProgramTrash.Get("/", requests.ValidatePagination, controllers.GetTrashUnsiaStudyPrograms)
	unsiaStudyProgramTrash.Put("/:id", controllers.RestoreUnsiaStudyProgram)

	unsiaStudyProgram.Get("/", requests.ValidatePagination, controllers.GetUnsiaStudyPrograms)
	unsiaStudyProgram.Get("/search", controllers.SearchUnsiaStudyPrograms)
	unsiaStudyProgram.Get("/export", controllers.ExportUnsiaStudyPrograms)
	unsiaStudyProgram.Get("/:id", controllers.GetUnsiaStudyProgram)
	unsiaStudyProgram.Post("/", requests.ValidateUnsiaStudyProgram, controllers.CreateUnsiaStudyProgram)
	unsiaStudyProgram.Post("/import", controllers.ImportUnsiaStudyPrograms)
	unsiaStudyProgram.Put("/:id", requests.ValidateUnsiaStudyProgram, controllers.UpdateUnsiaStudyProgram)
	unsiaStudyProgram.Delete("/:id", controllers.DeleteUnsiaStudyProgram)

	/* Educations */
	education := educationGroup.Group("educations")
	educationTrash := education.Group("trashs")
	educationTrash.Get("/", requests.ValidatePagination, controllers.GetTrashEducations)
	educationTrash.Put("/:id", controllers.RestoreEducation)

	education.Get("/", requests.ValidatePagination, controllers.GetEducations)
	education.Get("/search", controllers.SearchEducations)
	education.Get("/export", controllers.ExportEducations)
	education.Get("/by-educational-level/:educational_level_id", controllers.GetEducationByEducationalLevelId)
	education.Get("/:id", controllers.GetEducation)
	education.Post("/", requests.ValidateEducation, controllers.CreateEducation)
	education.Post("/import", controllers.ImportEducations)
	education.Put("/:id", requests.ValidateEducation, controllers.UpdateEducation)
	education.Delete("/:id", controllers.DeleteEducation)

	/* Schools */
	school := educationGroup.Group("schools")
	schoolTrash := school.Group("trashs")
	schoolTrash.Get("/", requests.ValidatePagination, controllers.GetTrashSchools)
	schoolTrash.Put("/:id", controllers.RestoreSchool)

	school.Get("/", requests.ValidatePagination, controllers.GetSchools)
	school.Get("/search", controllers.SearchSchools)
	school.Get("/export", controllers.ExportSchools)
	school.Get("/:id", controllers.GetSchool)
	school.Post("/", requests.ValidateSchool, controllers.CreateSchool)
	school.Post("/import", controllers.ImportSchools)
	school.Put("/:id", requests.ValidateSchool, controllers.UpdateSchool)
	school.Delete("/:id", controllers.DeleteSchool)

	/* Colleges */
	college := educationGroup.Group("colleges")
	collegeTrash := college.Group("trashs")
	collegeTrash.Get("/", requests.ValidatePagination, controllers.GetTrashColleges)
	collegeTrash.Put("/:id", controllers.RestoreCollege)

	college.Get("/", requests.ValidatePagination, controllers.GetColleges)
	college.Get("/search", controllers.SearchColleges)
	college.Get("/export", controllers.ExportColleges)
	college.Get("/:id", controllers.GetCollege)
	college.Post("/", requests.ValidateCollege, controllers.CreateCollege)
	college.Post("/import", controllers.ImportColleges)
	college.Put("/:id", requests.ValidateCollege, controllers.UpdateCollege)
	college.Delete("/:id", controllers.DeleteCollege)

	/* Field Of Studies */
	fieldOfStudy := educationGroup.Group("field-of-studies")
	fieldOfStudyTrash := fieldOfStudy.Group("trashs")
	fieldOfStudyTrash.Get("/", requests.ValidatePagination, controllers.GetTrashFieldOfStudies)
	fieldOfStudyTrash.Put("/:id", controllers.RestoreFieldOfStudy)

	fieldOfStudy.Get("/", requests.ValidatePagination, controllers.GetFieldOfStudies)
	fieldOfStudy.Get("/search", controllers.SearchFieldOfStudies)
	fieldOfStudy.Get("/export", controllers.ExportFieldOfStudies)
	fieldOfStudy.Get("/:id", controllers.GetFieldOfStudy)
	fieldOfStudy.Post("/", requests.ValidateFieldOfStudy, controllers.CreateFieldOfStudy)
	fieldOfStudy.Post("/import", controllers.ImportFieldOfStudies)
	fieldOfStudy.Put("/:id", requests.ValidateFieldOfStudy, controllers.UpdateFieldOfStudy)
	fieldOfStudy.Delete("/:id", controllers.DeleteFieldOfStudy)

	/* Sciences */
	science := educationGroup.Group("sciences")
	scienceTrash := science.Group("trashs")
	scienceTrash.Get("/", requests.ValidatePagination, controllers.GetTrashSciences)
	scienceTrash.Put("/:id", controllers.RestoreScience)

	science.Get("/", requests.ValidatePagination, controllers.GetSciences)
	science.Get("/search", controllers.SearchSciences)
	science.Get("/export", controllers.ExportSciences)
	science.Get("/:id", controllers.GetScience)
	science.Post("/", requests.ValidateScience, controllers.CreateScience)
	science.Post("/import", controllers.ImportSciences)
	science.Put("/:id", requests.ValidateScience, controllers.UpdateScience)
	science.Delete("/:id", controllers.DeleteScience)
}
