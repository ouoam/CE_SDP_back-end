package controller

import (
	"../db"
	"../model"
	"github.com/gofiber/fiber"
	"github.com/gorilla/schema"
	"github.com/jinzhu/copier"
	"net/http"
	"reflect"
	"strings"
	"unsafe"
)

var schemaDecoderQuery = schema.NewDecoder()

func GetID(c *fiber.Ctx, dataModel interface{}) {
	if err := model.CheckValidAllPK(dataModel); err != nil {
		_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		return
	}

	results, err := db.ListData(dataModel)
	if err != nil {
		_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		return
	}
	if l := len(results); l > 1 {
		_ = c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "have more than 1 row"})
		return
	} else if l == 0 {
		_ = c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "don't have any row"})
		return
	}

	_ = copier.Copy(dataModel, results[0])

	v := reflect.ValueOf(dataModel)
	if isImpl := v.Type().Implements(reflect.TypeOf((*model.WithPostGet)(nil)).Elem()); isImpl {
		_ = dataModel.(model.WithPostGet).PostGet()
	}

	if err := c.JSON(dataModel); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
}

func Post(c *fiber.Ctx, dataModel interface{}) {
	// todo store pk and restore or valid

	if err := c.BodyParser(dataModel); err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}

	v := reflect.ValueOf(dataModel)
	if isImpl := v.Type().Implements(reflect.TypeOf((*model.WithPreChange)(nil)).Elem()); isImpl {
		_ = dataModel.(model.WithPreChange).PreChange(true)
	}

	if err := db.AddData(dataModel); err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			if strings.Contains(err.Error(), "username") {
				_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Other member already used this username."})
				return
			}
			if strings.Contains(err.Error(), "email") {
				_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Other member already used this E-Mail."})
				return
			}
		}
		_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		return
	}

	if isImpl := v.Type().Implements(reflect.TypeOf((*model.WithPostGet)(nil)).Elem()); isImpl {
		_ = dataModel.(model.WithPostGet).PostGet()
	}

	if err := c.JSON(dataModel); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
}

func PutID(c *fiber.Ctx, dataModel interface{}) {
	// todo store pk and restore or valid

	if err := c.BodyParser(dataModel); err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}

	v := reflect.ValueOf(dataModel)
	if isImpl := v.Type().Implements(reflect.TypeOf((*model.WithPreChange)(nil)).Elem()); isImpl {
		_ = dataModel.(model.WithPreChange).PreChange(false)
	}

	if err := db.UpdateDate(dataModel); err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			if strings.Contains(err.Error(), "username") {
				_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Other member already used this username."})
				return
			}
			if strings.Contains(err.Error(), "email") {
				_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Other member already used this E-Mail."})
				return
			}
		}
		_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		return
	}

	if isImpl := v.Type().Implements(reflect.TypeOf((*model.WithPostGet)(nil)).Elem()); isImpl {
		_ = dataModel.(model.WithPostGet).PostGet()
	}

	if err := c.JSON(dataModel); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
}

func List(c *fiber.Ctx, dataModel interface{}) {
	var getString = func(b []byte) string {
		return *(*string)(unsafe.Pointer(&b))
	}

	// query Params
	if c.Fasthttp.QueryArgs().Len() > 0 {
		data := make(map[string][]string)
		c.Fasthttp.QueryArgs().VisitAll(func(key []byte, val []byte) {
			data[getString(key)] = []string{getString(val)}
		})
		schemaDecoderQuery.IgnoreUnknownKeys(true)
		if err := schemaDecoderQuery.Decode(dataModel, data); err != nil {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}
	} else if c.Body() != "" {
		if err := c.BodyParser(dataModel); err != nil {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}
	}

	results, err := db.ListData(dataModel)
	if err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}
	v := reflect.ValueOf(dataModel)
	if isImpl := v.Type().Implements(reflect.TypeOf((*model.WithPostGet)(nil)).Elem()); isImpl {
		for i := range results {
			_ = results[i].(model.WithPostGet).PostGet()
		}
	}

	_ = c.JSON(results)
}