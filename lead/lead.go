package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamza-s47/crm/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `gorm:"default:'Not provided'" json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	if err := db.Find(&leads).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch leads",
		})
	}
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	if err := db.Find(&lead, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Lead not found",
		})
	}
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := db.Create(&lead).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create lead",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	if err := db.Where("ID=?", id).Delete(&lead).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete lead",
		})
	}
	return c.JSON(fiber.Map{
		"Message": "Lead deleted successfully!",
		"status":  "200",
	})
}
