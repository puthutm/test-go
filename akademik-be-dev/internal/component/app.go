package component

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	restapidatareferensi "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/rest-api/datareferensi"
	restapisdm "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/rest-api/sdm"
	restsapisso "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/rest-api/sso"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/config"
	"unsia.ac.id/akademic_be/internal/delivery/http/controller"
	"unsia.ac.id/akademic_be/internal/delivery/http/middleware"
	"unsia.ac.id/akademic_be/internal/delivery/http/router"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/internal/service"
	servicemodel "unsia.ac.id/akademic_be/internal/service/model"
)

type Component struct {
	App         *fiber.App
	DB          *gorm.DB
	Log         *logrus.Logger
	Validate    *validator.Validate
	Config      *config.Config
	Minio       *minio.Client
	RedisClient *redis.Client
}

func Bootstrap(component *Component) {
	// TODO: Repository
	cacheRepository := cached.NewRedisCacheRepository(component.RedisClient)

	// external tools for gateway
	permissionService := restsapisso.NewPermissionCheckRest(
		component.Log, component.Config.InternalService.ApiKey, component.Config.InternalService.SsoApp.PermissionAction.PermissionActionEndpoint,
	)
	generalInformationGatewayRest := restapisdm.NewGeneralInformationRest(
		component.Log, component.Config.InternalService.ApiKey, component.Config.InternalService.SdmApp.GeneralInformation.GeneralInformationEndpoint,
	)
	academicPeriodGatewayRest := restapidatareferensi.NewAcademicPeriodDetaiWithSessionlRest(
		component.Log, component.Config.InternalService.ApiKey, component.Config.InternalService.DatareferensiApp.AcademicPeriod.AcademicPeriodEndpoint,
	)

	mstClassRepository := repositorymodel.NewMstClassRepository(
		component.Log,
		cacheRepository,
	)

	mstClassLecturerRepository := repositorymodel.NewMstClassLecturerRepository(
		component.Log,
		cacheRepository,
	)

	mstClassParticipantRepository := repositorymodel.NewMstClassParticipantRepository(
		component.Log,
		cacheRepository,
	)

	mstClassScheduleRepository := repositorymodel.NewMstClassScheduleRepository(
		component.Log,
		cacheRepository,
	)

	mstClassScheduleTemplateRepository := repositorymodel.NewMstClassScheduleTemplateRepository(
		component.Log,
		cacheRepository,
	)
	mstClassScheduleTaskRepository := repositorymodel.NewMstClassScheduleTaskRepository(
		component.Log,
		cacheRepository,
	)

	// mstClassScheduleTaskCollectRepository :

	// mstClassScheduleTeachingMaterialRepository :

	mstSKSLimitRepository := repositorymodel.NewMstSKSLimitRepository(
		component.Log,
	)

	mstStudentBioRepository := repositorymodel.NewMstStudentBioRepository(
		component.Log,
		cacheRepository,
	)

	// mstStudentDocumentRepository :

	mstStudentDomicileRepository := repositorymodel.NewMstStudentDomicileRepository(
		component.Log,
	)

	mstStudentEducationRepository := repositorymodel.NewMstStudentEducationRepository(
		component.Log,
	)

	mstStudentFamilyRepository := repositorymodel.NewMstStudentFamilyRepository(
		component.Log,
	)

	mstStudentStudyProgramRepository := repositorymodel.NewMstStudentStudyProgramRepository(
		component.Log,
		cacheRepository,
	)

	mstStudyProgramCurriculumRepository := repositorymodel.NewMstStudyProgramCurriculumRepository(
		component.Log,
		cacheRepository,
	)
	mstStudyProgramCurriculumSubjectPrerequisiteRepository := repositorymodel.NewMstStudyProgramCurriculumSubjectPrerequisiteRepository(
		component.Log,
		cacheRepository,
	)

	mstSubjectRepository := repositorymodel.NewMstSubjectRepository(
		component.Log,
		cacheRepository,
	)

	mstValueScaleRepository := repositorymodel.NewMstValueScaleRepository(
		component.Log,
	)
	mstValueCompositionRepository := repositorymodel.NewMstValueCompositionRepository(
		component.Log,
		cacheRepository,
	)
	mstStudentPresenceSettingRepository := repositorymodel.NewMstStudentPresenceSettingRepository(
		component.Log,
	)

	trxFinalProjectProposalRepository := repositorymodel.NewTrxFinalProjectProposalRepository(
		component.Log,
		cacheRepository,
	)
	trxStudentPresenceSettingRepository := repositorymodel.NewTrxStudentPresenceSettingRepository(
		component.Log,
	)
	krsRepository := repositorymodel.NewKrsRepository(
		component.Log,
	)

	classScoreRepository := repositorymodel.NewClassScoreRepository(
		component.Log,
	)

	khsRepository := repositorymodel.NewKhsRepository(
		component.Log,
	)

	portalStudentRepository := repositorymodel.NewPortalStudentRepository(
		component.Log,
	)

	// trxFinalProjectProposalMentorLecturerRepository

	// For external
	mstLecturerRepository := repositorymodel.NewMstLecturerRepository(
		component.Log,
		cacheRepository,
	)
	mstStudyProgramRepository := repositorymodel.NewMstStudyProgramRepository(
		component.Log,
		cacheRepository,
	)
	mstBiodataOfUserRepository := repositorymodel.NewMstBiodataOfUserRepository(
		component.Log,
		cacheRepository,
	)
	mstMediaLibraryRepository := repositorymodel.NewMstMediaLibraryRepository(
		component.Log, cacheRepository,
	)
	mstAssessmentWeightRepository := repositorymodel.NewMstAssessmentWeightRepository(
		component.Log, cacheRepository,
	)

	// TODO: Service
	storageService := service.NewStorageService(
		component.Config,
		component.Minio,
		component.Log,
	)

	mstClassService := servicemodel.NewMstClassService(
		component.Log,
		component.DB,
		cacheRepository,
		mstClassRepository,
		mstSubjectRepository,
		storageService,

		// external
		generalInformationGatewayRest,
	)

	mstClassLecturerService := servicemodel.NewMstClassLecturerService(
		component.Log,
		component.DB,
		cacheRepository,
		mstClassLecturerRepository,
	)

	mstClassParticipantService := servicemodel.NewMstClassParticipantService(
		component.Log,
		component.DB,
		cacheRepository,
		mstClassParticipantRepository,
		mstStudentBioRepository,
		mstClassRepository,

		// external
		generalInformationGatewayRest,
	)

	mstClassScheduleService := servicemodel.NewMstClassScheduleService(
		component.Log,
		component.DB,
		cacheRepository,
		mstClassScheduleRepository,
		storageService,
		mstClassRepository,
		mstClassScheduleTemplateRepository,

		// external
		generalInformationGatewayRest,
		academicPeriodGatewayRest,
	)

	mstClassScheduleTemplateService := servicemodel.NewMstClassScheduleTemplateService(
		component.Log,
		component.DB,
		cacheRepository,
		mstClassScheduleTemplateRepository,
		mstClassRepository,
		mstClassScheduleService,

		// external
		generalInformationGatewayRest,
	)

	mstClassScheduleTaskService := servicemodel.NewMstClassScheduleTaskService(
		component.Log,
		component.DB,
		cacheRepository,
		mstClassScheduleTaskRepository,
	)

	// mstClassScheduleTaskCollectService
	// mstClassScheduleTeachingMaterialService

	mstSKSLimitService := servicemodel.NewMstSKSLimitService(
		component.Log,
		component.DB,
		cacheRepository,
		mstSKSLimitRepository,
	)

	mstStudentBioService := servicemodel.NewMstStudentBioService(
		component.Log,
		component.DB,
		cacheRepository,
		mstStudentBioRepository,
		storageService,
	)

	// mstStudentDocumentService

	mstStudentDomicileService := servicemodel.NewMstStudentDomicileService(
		component.Log,
		component.DB,
		cacheRepository,
		mstStudentDomicileRepository,
	)

	mstStudentEducationService := servicemodel.NewMstStudentEducationService(
		component.Log,
		component.DB,
		cacheRepository,
		mstStudentEducationRepository,
	)

	mstStudentFamilyService := servicemodel.NewMstStudentFamilyService(
		component.Log,
		component.DB,
		cacheRepository,
		mstStudentFamilyRepository,
	)

	mstStudentStudyProgramService := servicemodel.NewMstStudentStudyProgramService(
		component.Log,
		component.DB,
		cacheRepository,
		mstStudentStudyProgramRepository,
	)

	mstStudyProgramCurriculumService := servicemodel.NewMstStudyProgramCurriculumService(
		component.Log,
		component.DB,
		cacheRepository,
		mstStudyProgramCurriculumRepository,
		mstClassRepository,
		mstSubjectRepository,
		mstStudyProgramCurriculumSubjectPrerequisiteRepository,

		// external
		generalInformationGatewayRest,
	)

	// mstStudyProgramCurriculumRepository

	mstSubjectService := servicemodel.NewMstSubjectService(
		component.Log,
		component.DB,
		cacheRepository,
		mstSubjectRepository,
		generalInformationGatewayRest,
	)

	mstValueCompositionService := servicemodel.NewMstValueCompositionService(
		component.Log,
		component.DB,
		cacheRepository,
		mstValueCompositionRepository,
	)

	mstValueScaleService := servicemodel.NewMstValueScaleService(
		component.Log,
		component.DB,
		cacheRepository,
		mstValueScaleRepository,
	)
	mstStudentPresenceSettingService := servicemodel.NewMstStudentPresenceSettingService(
		component.DB,
		component.Log,
		cacheRepository,
		mstStudentPresenceSettingRepository,
	)

	trxFinalProjectProposalService := servicemodel.NewTrxFinalProjectProposalService(
		component.Log,
		component.DB,
		cacheRepository,
		trxFinalProjectProposalRepository,
		mstStudentBioRepository,
		storageService,
	)
	trxStudentPresenceSettingService := servicemodel.NewTrxStudentPresenceSettingService(
		component.DB,
		component.Log,
		cacheRepository,
		trxStudentPresenceSettingRepository,
	)
	krsService := servicemodel.NewKrsService(
		component.Log,
		component.DB,
		krsRepository,
	)

	classScoreService := servicemodel.NewClassScoreService(
		component.Log,
		component.DB,
		classScoreRepository,
	)

	khsService := servicemodel.NewKhsService(
		component.Log,
		component.DB,
		khsRepository,
	)

	portalStudentService := servicemodel.NewPortalStudentService(
		component.Log,
		component.DB,
		portalStudentRepository,
	)

	// trxFinalProjectProposalMentorLecturerService

	// For External
	mstLecturerService := servicemodel.NewMstLecturerService(
		component.Log,
		component.DB,
		cacheRepository,
		mstLecturerRepository,
	)
	mstStudyProgramService := servicemodel.NewMstStudyProgramService(
		component.Log,
		component.DB,
		cacheRepository,
		mstStudyProgramRepository,
	)
	mstBiodataOfUserService := servicemodel.NewMstBiodataOfUserService(
		component.Log,
		component.DB,
		cacheRepository,
		mstBiodataOfUserRepository,
	)
	mstMediaLibraryService := servicemodel.NewMstMediaLibraryService(
		component.Log,
		component.DB,
		cacheRepository,
		mstMediaLibraryRepository,
	)

	mstAssessmentWeightService := servicemodel.NewMstAssessmentWeightService(
		component.Log,
		component.DB,
		cacheRepository,
		mstAssessmentWeightRepository,
	)

	// TODO: Controller
	mstClassController := controller.NewMstClassController(
		component.Log,
		mstClassService,
		component.Validate,
	)

	mstClassLecturerController := controller.NewMstClassLecturerController(
		component.Log,
		mstClassLecturerService,
		component.Validate,
	)

	// mstClassLecturerController

	mstClassParticipantController := controller.NewMstClassParticipantController(
		component.Log,
		mstClassParticipantService,
		component.Validate,
	)

	mstClassScheduleController := controller.NewMstClassScheduleController(
		component.Log,
		mstClassScheduleService,
		component.Validate,
	)
	mstClassScheduleTemplateController := controller.NewMstClassScheduleTemplateController(
		component.Log,
		mstClassScheduleTemplateService,
		component.Validate,
	)

	mstClassScheduleTaskController := controller.NewMstClassScheduleTaskController(
		component.Log,
		mstClassScheduleTaskService,
		component.Validate,
	)

	// mstClassScheduleTaskCollectController
	// mstClassScheduleTeachingMaterialController

	mstSKSLimitController := controller.NewMstSKSLimitController(
		component.Log,
		mstSKSLimitService,
		component.Validate,
	)

	mstStudentBioController := controller.NewMstStudentBioController(
		component.Log,
		mstStudentBioService,
		component.Validate,
	)

	// mstStudentDocumentController

	mstStudentDomicileController := controller.NewMstStudentDomicileController(
		component.Log,
		mstStudentDomicileService,
		mstStudentBioService,
		component.Validate,
	)

	mstStudentEducationController := controller.NewMstStudentEducationController(
		component.Log,
		mstStudentEducationService,
		mstStudentBioService,
		component.Validate,
		storageService,
		component.Config,
	)

	mstStudentFamilyController := controller.NewMstStudentFamilyController(
		component.Log,
		mstStudentFamilyService,
		mstStudentBioService,
		component.Validate,
	)

	mstStudentStudyProgramController := controller.NewMstStudentStudyProgramController(
		component.Log,
		mstStudentStudyProgramService,
	)

	mstStudyProgramCurriculumController := controller.NewMstStudyProgramCurriculumController(
		component.Log,
		mstStudyProgramCurriculumService,
		component.Validate,
	)

	mstSubjectController := controller.NewMstSubjectController(
		component.Log,
		mstSubjectService,
		component.Validate,
	)

	mstValueCompositionController := controller.NewMstValueCompositionController(
		component.Log,
		mstValueCompositionService,
		component.Validate,
	)

	mstValueScaleController := controller.NewMstValueScaleController(
		component.Log,
		mstValueScaleService,
		component.Validate,
	)
	mstStudentPresenceSettingController := controller.NewMstStudentPresenceSettingController(
		component.Log,
		component.Validate,
		mstStudentPresenceSettingService,
	)

	trxFinalProjectProposalController := controller.NewTrxFinalProjectProposalController(
		component.Log,
		trxFinalProjectProposalService,
		component.Validate,
	)
	trxStudentPresenceSettingController := controller.NewTrxStudentPresenceSettingController(
		component.Log,
		component.Validate,
		trxStudentPresenceSettingService,
	)
	krsController := controller.NewKrsController(
		component.Log,
		krsService,
		component.Validate,
	)

	classScoreController := controller.NewClassScoreController(
		component.Log,
		classScoreService,
		component.Validate,
	)

	khsController := controller.NewKhsController(
		component.Log,
		khsService,
	)

	portalStudentController := controller.NewPortalStudentController(
		component.Log,
		portalStudentService,
		component.Validate,
	)

	// trxFinalProjectProposalMentorLecturerController

	// For External
	mstLecturerController := controller.NewMstLecturerController(
		component.Log,
		mstLecturerService,
	)
	mstStudyProgramController := controller.NewMstStudyProgramController(
		component.Log,
		mstStudyProgramService,
	)
	mstBiodataOfUserController := controller.NewMstBiodataOfUserController(
		component.Log,
		mstBiodataOfUserService,
	)
	mstMediaLibraryController := controller.NewMstMediaLibraryController(
		component.Log,
		mstMediaLibraryService,
	)

	mstAssessmentWeightController := controller.NewMstAssessmentWeightController(
		component.Log,
		mstAssessmentWeightService,
		component.Validate,
	)

	//

	// TODO: Middleware
	codeError := middleware.ErrorCode()
	authenticationMiddleware := middleware.Authentication(component.Config)
	permissionMiddleware := middleware.NewMiddlewarePermissions(
		component.Log,
		component.Config,
		cacheRepository,
		permissionService,
	)

	// TODO: Config route
	route := router.RouterConfig{
		Config: component.Config,
		App:    component.App,

		CodeError:                codeError,
		AuthenticationMiddleware: authenticationMiddleware,
		PermissionsMiddleware:    permissionMiddleware,

		/* Controller */
		MstClassController:                 mstClassController,
		MstCLassLecturerController:         mstClassLecturerController,
		MstCLassParticipantController:      mstClassParticipantController,
		MstClassScheduleController:         mstClassScheduleController,
		MstClassScheduleTemplateController: mstClassScheduleTemplateController,
		MstClassScheduleTaskController:     mstClassScheduleTaskController,
		// mstClassScheduleTaskCollectController
		// mstClassScheduleTeachingMaterialController
		MstSKSLimitController:   mstSKSLimitController,
		MstStudentBioController: mstStudentBioController,
		// mstStudentDocumentController
		MstStudentDomicileController:        mstStudentDomicileController,
		MstStudentEducationController:       mstStudentEducationController,
		MstStudentFamilyController:          mstStudentFamilyController,
		MstStudentStudyProgramController:    mstStudentStudyProgramController,
		MstStudyProgramCurriculumController: mstStudyProgramCurriculumController,
		MstSubjectController:                mstSubjectController,
		MstValueCompositionController:       mstValueCompositionController,
		MstValueScaleController:             mstValueScaleController,
		TrxFinalProjectProposalController:   trxFinalProjectProposalController,
		TrxStudentPresenceSettingController: trxStudentPresenceSettingController,
		KrsController:                       krsController,
		ClassScoreController:                classScoreController,
		KhsController:                       khsController,
			PortalStudentController:             portalStudentController,
		// trxFinalProjectProposalMentorLecturerController
		MstStudentPresenceSettingController: mstStudentPresenceSettingController,

		// For external
		MstLecturerController:         mstLecturerController,
		MstStudyProgramController:     mstStudyProgramController,
		MstBiodataOfUserController:    mstBiodataOfUserController,
		MstMediaLibraryController:     mstMediaLibraryController,
		MstAssessmentWeightController: mstAssessmentWeightController,
	}

	route.Setup()
}
