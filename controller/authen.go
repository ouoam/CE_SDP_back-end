package controller

import (
	"../db"
	"../model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type LoginData struct {
	ID int64
}

func Login(c *fiber.Ctx) {
	input := new(model.Member)
	if err := c.BodyParser(input); err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}

	if !(input.Username.Valid || input.Email.Valid) {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Please enter username or email"})
		return
	}
	if !input.Password.Valid {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Please enter password"})
		return
	}

	member := model.Member{
		Username: input.Username,
		Email: input.Email,
	}

	members, err := db.ListData(&member)
	if err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}
	if len(members) == 0 {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Username or E-mail or Password incorrect."})
		return
	}

	member = *(members[0].(*model.Member))

	err = bcrypt.CompareHashAndPassword([]byte(member.Password.String), []byte(input.Password.String))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Username or E-mail or Password incorrect."})
		} else {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": member.ID.Int64,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		return
	}

	_ = c.JSON(fiber.Map{"token": tokenString})
}
