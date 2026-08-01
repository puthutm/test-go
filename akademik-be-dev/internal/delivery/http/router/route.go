package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"unsia.ac.id/akademic_be/internal/config"
	"unsia.ac.id/akademic_be/internal/delivery/http/controller"
	"unsia.ac.id/akademic_be/internal/delivery/http/middleware"
)

type RouterConfig struct {
	Config *config.Config
	App    *fiber.App
	// middleware
	CodeError                fiber.Handler
	AuthenticationMiddleware fiber.Handler
	PermissionsMiddleware    *middleware.MiddlewarePermissions

	// Controller
	MstClassController                 *controller.MstClassController
	MstCLassLecturerController         *controller.MstClassLecturerController
	MstCLassParticipantController      *controller.MstClassParticipantController
	MstClassScheduleController         *controller.MstClassScheduleController
	MstClassScheduleTemplateController *controller.MstClassScheduleTemplateController
	MstClassScheduleTaskController     *controller.MstClassScheduleTaskController
	// mstClassScheduleTaskCollectController
	// mstClassScheduleTeachingMaterialController
	MstSKSLimitController   *controller.MstSKSLimitController
	MstStudentBioController *controller.MstStudentBioController
	// mstStudentDocumentController
	MstStudentDomicileController        *controller.MstStudentDomicileController
	MstStudentEducationController       *controller.MstStudentEducationController
	MstStudentFamilyController          *controller.MstStudentFamilyController
	MstStudentStudyProgramController    *controller.MstStudentStudyProgramController
	MstStudyProgramCurriculumController *controller.MstStudyProgramCurriculumController
	MstSubjectController                *controller.MstSubjectController
	MstValueCompositionController       *controller.MstValueCompositionController
	MstValueScaleController             *controller.MstValueScaleController
	TrxFinalProjectProposalController   *controller.TrxFinalProjectProposalController
	TrxStudentPresenceSettingController *controller.TrxStudentPresenceSettingController
	KrsController                       *controller.KrsController
	// TrxFinalProjectProposalMentorLecturerController
	// For External
	MstLecturerController               *controller.MstLecturerController
	MstStudyProgramController           *controller.MstStudyProgramController
	MstBiodataOfUserController          *controller.MstBiodataOfUserController
	MstMediaLibraryController           *controller.MstMediaLibraryController
	MstAssessmentWeightController       *controller.MstAssessmentWeightController
	MstStudentPresenceSettingController *controller.MstStudentPresenceSettingController
	ClassScoreController                *controller.ClassScoreController
	PortalStudentController             *controller.PortalStudentController
	KhsController                       *controller.KhsController
}

func (r *RouterConfig) Setup() {
	api := r.App.Group("/api")
	r.SetupMiddleware()
	r.SetupRoutes(api)
}

func (r *RouterConfig) SetupMiddleware() {
	r.App.Use(cors.New(cors.ConfigDefault))
	r.App.Use(recover.New(recover.ConfigDefault))
	// r.App.Use(limiter.New(limiter.Config{
	// 	Max:        100,
	// 	Expiration: 60 * time.Second,
	// }))

	// prometheus
	middleware.PrometheusInit()
	r.App.Get("/metrics", middleware.ApiKeyMetrics(r.Config), adaptor.HTTPHandler(promhttp.Handler()))
	r.App.Use(middleware.TrackMetrics())

	// custom
	r.App.Use(r.CodeError)
}

func (r *RouterConfig) SetupRoutes(api fiber.Router) {
	r.App.Use(r.AuthenticationMiddleware)
	r.SetupStudentRoute(api)
	r.SetupAcademicRoute(api)
	r.SetupProgramHeadRoute(api)
	r.SetupLecturerRoute(api)

	// coba coba
	api.Get("/test/auth",
		r.PermissionsMiddleware.PermissionCheckHandler(map[string]string{
			"berita - informasi": "read", "eadada": "asdas",
		}), func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{"message": "Authentication and Permission Check passed"})
		})
}
