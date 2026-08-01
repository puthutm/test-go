package routes

import (
	biodataControllers "data-referensi/app/controllers/biodata"
	educationControllers "data-referensi/app/controllers/education"
	enrollmentControllers "data-referensi/app/controllers/enrollment"
	pmbControllers "data-referensi/app/controllers/pmb"
	regionControllers "data-referensi/app/controllers/region"

	"github.com/gofiber/fiber/v2"
)

func PublicRoute(app fiber.Router) {
	public := app.Group("/public")

	public.Get("/study-programs", educationControllers.SearchStudyPrograms)
	public.Get("/registration-paths", enrollmentControllers.SearchRegistrationPaths)
	public.Get("/countries", regionControllers.SearchCountries)
	public.Get("/provinces", regionControllers.SearchProvinces)
	public.Get("/provinces/by-country/:country_id", regionControllers.GetProvinceByCountryId)
	public.Get("/cities", regionControllers.SearchCities)
	public.Get("/cities/by-province/:province_id", regionControllers.GetCityByProvinceId)
	public.Get("/districts", regionControllers.SearchDistricts)
	public.Get("/districts/by-city/:city_id", regionControllers.GetDistrictByCityId)
	public.Get("/villages", regionControllers.SearchVillages)
	public.Get("/villages/by-district/:district_id", regionControllers.GetVillageByDistrictId)
	public.Get("/school-types", pmbControllers.SearchSchoolTypes)
	public.Get("/religions", biodataControllers.SearchReligions)
	public.Get("/ethnics", biodataControllers.SearchEthnics)
	public.Get("/almamater-sizes", biodataControllers.SearchAlmamaterSizes)
	public.Get("/education-levels", educationControllers.SearchEducationalLevels)
	public.Get("/school-types", pmbControllers.SearchSchoolTypes)
	public.Get("/school-types", pmbControllers.SearchSchoolTypes)
	public.Get("/academic-periods", pmbControllers.SearchAcademicPeriods)
	public.Get("/enrollment-batchs", enrollmentControllers.SearchEnrollmentBatchs)
}
