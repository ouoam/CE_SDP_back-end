package controller

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
)

func Upload(c *fiber.Ctx) {
	file, err := c.FormFile("pic")

	if err == nil {
		f, err := file.Open()
		src, err := imaging.Decode(f)
		_ = f.Close()
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error" : err.Error()})
			return
		}

		srcW := src.Bounds().Dx()
		srcH := src.Bounds().Dy()
		if srcH * srcW > 4000000 {
			// a * b = 4000000
			// a = 4000000 / b ---> 1
			// a / b = ratio = w / h
			// a = w / h * b ---> 2
			// 4000000 / b = w / h * b
			// b * b = 4000000 * h / w
			// b = sqrt(4000000 * h / w)
			b := math.Sqrt(4000000 * float64(srcH) / float64(srcW))
			src = imaging.Resize(src, 0, int(math.Floor(b)), imaging.Lanczos)
		}

		// create temp file
		tmpfile, err := ioutil.TempFile("", "upload-pic.*.jpg")
		if err != nil {
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		// save file to temp
		if err = imaging.Encode(tmpfile, src, imaging.JPEG); err != nil {
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
			return
		}

		// calc sha256
		if _, err = tmpfile.Seek(0,0); err != nil {
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
			return
		}
		h := sha256.New()
		if _, err := io.Copy(h, tmpfile); err != nil {
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
			return
		}

		// convent sha256 to base 64
		b := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(h.Sum(nil))

		// move temp file to pic store
		if _, err = tmpfile.Seek(0,0); err != nil {
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
			return
		}
		outputFile, err := os.Create(fmt.Sprintf("./pic/%s", b))
		if err != nil {
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
			return
		}
		_, err = io.Copy(outputFile, tmpfile)
		_ = tmpfile.Close()
		_ = os.Remove(tmpfile.Name())
		_ = outputFile.Close()

		c.JSON(fiber.Map{"file" : b})
	}
}
