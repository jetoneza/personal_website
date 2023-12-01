package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/models"
	"github.com/jetoneza/personal_website/internal/schema"
	"github.com/jetoneza/personal_website/pkg/utils"
	"gorm.io/gorm"
)

func (h *Handler) GetEvent(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var event models.Event

	result := h.App.DB.First(&event, "id = ?", id)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("%v", result.Error),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   event,
	})
}

func (h *Handler) GetAllEvents(ctx *fiber.Ctx) error {
	filter := ctx.Query("filter", "")

	var events []models.Event
	var results *gorm.DB

	// TODO: Create a cleaner implementation on separating the queries
	if filter != "" {
		var start time.Time
		var end time.Time

		now := time.Now()

		if filter == "month" {
			start = time.Date(now.Year(), now.Month(), 0, 0, 0, 0, 0, now.Location())
			end = start.AddDate(0, 1, -1)
		}

		if filter == "year" {
			start = time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
			end = time.Date(now.Year(), time.December, 31, 0, 0, 0, 0, now.Location())
		}

		results = h.App.DB.Where("start BETWEEN ? AND ?", start.Format(time.RFC3339), end.Format(time.RFC3339)).Order("start desc").Find(&events)
	} else {
		page, _ := strconv.Atoi(ctx.Query("page", "1"))
		limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
		offset := (page - 1) * limit

		results = h.App.DB.Order("created_at desc").Limit(limit).Offset(offset).Find(&events)
	}

	if results.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": results.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"results": len(events),
		"data":    events,
	})
}

// TODO: Implement get events by date range

func (h *Handler) CreateEvent(ctx *fiber.Ctx) error {
	body := new(schema.CreateEventSchema)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "error": err.Message})
	}

	allDay, err := strconv.ParseBool(body.AllDay)
	if err != nil {
		allDay = false
	}

	start, err := time.Parse(time.RFC3339, body.Start)
	if err != nil {
		start = time.Now()
	}

	end, err := time.Parse(time.RFC3339, body.End)
	if err != nil {
		end = time.Now()
	}

	event := models.Event{
		Title:  body.Title,
		Type:   body.Type,
		Notes:  body.Notes,
		AllDay: allDay,
		Start:  start,
		End:    end,
	}

	result := h.App.DB.Create(&event)

	if result.Error != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "error": result.Error.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   event,
	})
}

func (h *Handler) UpdateEvent(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var event models.Event

	result := h.App.DB.First(&event, "id = ?", id)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("%v", result.Error),
		})
	}

	body := new(schema.UpdateEventSchema)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "error": err.Message})
	}

	if body.Notes != "" {
		event.Notes = body.Notes
	}

	if body.Title != "" {
		event.Title = body.Title
	}

	if body.Type != "" {
		event.Type = body.Type
	}

	allDay, err := strconv.ParseBool(body.AllDay)
	if err == nil {
		event.AllDay = allDay
	}

	start, err := time.Parse(time.RFC3339, body.Start)
	if err == nil {
		event.Start = start
	}

	end, err := time.Parse(time.RFC3339, body.End)
	if err == nil {
		event.End = end
	}

	saveResult := h.App.DB.Save(&event)

	if saveResult.Error != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "error": saveResult.Error.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   event,
	})
}

func (h *Handler) DeleteEvent(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	result := h.App.DB.Delete(&models.Event{}, "id = ?", id)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}
