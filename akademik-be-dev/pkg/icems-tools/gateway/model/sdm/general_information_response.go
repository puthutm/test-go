package modelsdm

import "github.com/google/uuid"

type GeneralInformationResponse struct {
	ID                      uuid.UUID  `json:"id"`
	WorkingUnitID           uuid.UUID  `json:"working_unit_id"`
	WorkingUnitName         string     `json:"working_unit_name,omitempty"`
	WorkingRelationshipID   uuid.UUID  `form:"working_relationship_id"`
	WorkingRelationshipName string     `json:"working_relationship_name,omitempty"`
	EmployeeTypeID          uuid.UUID  `json:"employee_type_id"`
	EmployeeTypeName        string     `json:"employee_type_name,omitempty"`
	WhatsappNumber          string     `json:"whatsapp_number"`
	FunctionalPositionID    uuid.UUID  `json:"functional_position_id"`
	FunctionalPositionName  string     `json:"functional_position_name,omitempty"`
	DigitalSignaturePath    string     `json:"digital_signature_path"`
	BarcodeSignaturePath    string     `json:"barcode_signature_path"`
	FingerAccountNumber     string     `json:"finger_account_number"`
	TransportationID        uuid.UUID  `json:"transportation_id"`
	TransportationName      string     `json:"transportation_name,omitempty"`
	AlmamaterSizeID         uuid.UUID  `json:"almamater_size_id"`
	AlmamaterSizeName       string     `json:"almamater_size_name,omitempty"`
	JobID                   uuid.UUID  `json:"job_id"`
	JobName                 string     `json:"job_name,omitempty"`
	EmployeeStatusID        uuid.UUID  `json:"employee_status_id"`
	EmployeeStatusName      string     `json:"employee_status_name,omitempty"`
	AcademicPositionID      uuid.UUID  `json:"academic_position_id"`
	AcademicPositionName    string     `json:"academic_position_name,omitempty"`
	StudyProgramID          *uuid.UUID `json:"study_program_id"`
	StudyProgramName        string     `json:"study_program_name,omitempty"`
	CitizenshipID           uuid.UUID  `json:"citizenship_id"`
	CitizenshipName         string     `json:"citizenship_name,omitempty"`
	UserID                  uuid.UUID  `json:"user_id"`
	CreatedAt               int64      `json:"created_at"`
	UpdatedAt               int64      `json:"updated_at"`

	PhoneNumber  string `json:"phone_number"`  // dari b.phone AS phone_number
	Email        string `json:"email"`         // dari b.email
	EmailPrivate string `json:"email_private"` // dari b.email_private
	NameOfUser   string `json:"name_of_user"`  // dari b.name AS name_of_user
}

func (p *GeneralInformationResponse) GetID() string {
	return p.ID.String()
}
