package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func MstStudyProgramCurriculumSubjectPrerequisiteModelToResponse(model model.MstStudyProgramCurriculumSubjectPrerequisite) dto.MstStudyProgramCurriculumSubjectPrerequisiteResponse {
	return dto.MstStudyProgramCurriculumSubjectPrerequisiteResponse{
		ID:                              model.ID,
		StudyProgramCurriculumID:        model.StudyProgramCurriculumID,
		StudyProgramCurriculumSubjectID: model.StudyProgramCurriculumSubjectID,
		SubjectNameID:                   model.SubjectNameID,
		SubjectNameEN:                   model.SubjectNameEN,
		SubjectCode:                     model.SubjectCode,
	}
}
