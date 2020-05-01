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

func Get(c *fiber.Ctx, dataModel interface{}, params... interface{}) {
	if err := model.CheckValidAllPK(dataModel); err != nil {
		_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		return
	}

	filter := new(db.Filter)

	v := reflect.ValueOf(dataModel)
	if isImpl := v.Type().Implements(reflect.TypeOf((*model.CanSearch)(nil)).Elem()); isImpl {
		params = nil
		filter.Search.SetValid("")
	}

	results, err := db.ListData(dataModel, filter, params...)
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

	if isImpl := v.Type().Implements(reflect.TypeOf((*model.WithPostGet)(nil)).Elem()); isImpl {
		_ = dataModel.(model.WithPostGet).PostGet()
	}

	if err := c.JSON(dataModel); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
}

func New(c *fiber.Ctx, dataModel interface{}) {
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

func Update(c *fiber.Ctx, dataModel interface{}) {
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

func Delete(c *fiber.Ctx, dataModel interface{}) {
	// todo store pk and restore or valid

	if err := db.DeleteDate(dataModel); err != nil {
		_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		return
	}

	c.JSON(fiber.Map{"status": "ok"})
}

func List(c *fiber.Ctx, dataModel interface{}, params... interface{}) {
	var getString = func(b []byte) string {
		return *(*string)(unsafe.Pointer(&b))
	}

	filter := new(db.Filter)

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
		if err := schemaDecoderQuery.Decode(filter, data); err != nil {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}
	} else if c.Body() != "" {
		if err := c.BodyParser(dataModel); err != nil {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}
		if err := c.BodyParser(filter); err != nil {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}
	}

	v := reflect.ValueOf(dataModel)
	if isImpl := v.Type().Implements(reflect.TypeOf((*model.CanSearch)(nil)).Elem()); isImpl {
		params = nil
		if !filter.Search.Valid {
			filter.Search.SetValid("")
		}
	} else {
		filter.Search.UnmarshalText([]byte(""))
	}

	results, err := db.ListData(dataModel, filter, params...)
	if err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}

	if isImpl := v.Type().Implements(reflect.TypeOf((*model.WithPostGet)(nil)).Elem()); isImpl {
		for i := range results {
			_ = results[i].(model.WithPostGet).PostGet()
		}
	}

	_ = c.JSON(results)
}