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

	if data, ok := dataModel.(model.WithPostGet); ok {
		if err := data.PostGet(); err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}
	}

	if err := c.JSON(dataModel); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
}

func New(c *fiber.Ctx, dataModel interface{}) {
	v := reflect.ValueOf(dataModel).Elem()
	result := reflect.New(v.Type()).Interface()
	if err := c.BodyParser(result); err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}

	stv2 := reflect.ValueOf(result).Elem()
	for i := 0; i < stv2.NumField(); i++ {
		if v.Type().Field(i).Tag.Get("key") != "p" {
			nvField := v.Field(i)
			nvField.Set(stv2.Field(i))
		}
	}

	if data, ok := dataModel.(model.WithPreChange); ok {
		if err := data.PreChange(true); err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}
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

	if data, ok := dataModel.(model.WithPostGet); ok {
		if err := data.PostGet(); err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}
	}

	if err := c.JSON(dataModel); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
}

func Update(c *fiber.Ctx, dataModel interface{}) {
	v := reflect.ValueOf(dataModel).Elem()
	result := reflect.New(v.Type()).Interface()
	if err := c.BodyParser(result); err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}

	stv2 := reflect.ValueOf(result).Elem()
	for i := 0; i < stv2.NumField(); i++ {
		if v.Type().Field(i).Tag.Get("key") != "p" {
			nvField := v.Field(i)
			nvField.Set(stv2.Field(i))
		}
	}

	if data, ok := dataModel.(model.WithPreChange); ok {
		if err := data.PreChange(false); err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}
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

	if data, ok := dataModel.(model.WithPostGet); ok {
		if err := data.PostGet(); err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}
	}

	if err := c.JSON(dataModel); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
}

func Delete(c *fiber.Ctx, dataModel interface{}) {
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

	if _, ok := dataModel.(model.WithPostGet); ok {
		for i := range results {
			if err := results[i].(model.WithPostGet).PostGet(); err != nil {
				c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
				return
			}
		}
	}

	_ = c.JSON(results)
}