package handlers

import (
	"math"

	"github.com/gofiber/fiber/v2"
)

func SendSuccess(c *fiber.Ctx, statusCode int, data interface{}, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error":   false,
		"data":    data,
		"message": message,
	})
}

func SendFailed(c *fiber.Ctx, statusCode int, data interface{}, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error":   true,
		"data":    nil,
		"message": message,
	})
}

func ResponseWithMetadata(data interface{}, page int, pageSize int64, subTotal int, total int64) map[string]interface{} {
	return map[string]interface{}{
		"data": data,
		"metadata": map[string]interface{}{
			"page":        page,
			"per_page":    pageSize,
			"total_pages": calculateTotalPages(int(total), int(pageSize)),
			"sub_total":   subTotal,
			"total":       total,
		},
	}
}

func calculateTotalPages(totalItems, itemsPerPage int) int {
	if itemsPerPage == 0 {
		return 0
	}
	return int(math.Ceil(float64(totalItems) / float64(itemsPerPage)))
}
