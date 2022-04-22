package lead

import (
	"crm-simple/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return err
	}
	db.Create(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var lead Lead
	db.Delete(&lead, id)
	return c.SendString("Lead Deleted")
}
