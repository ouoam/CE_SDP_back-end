package controller

import (
	"../db"
	"../model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"
)

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

	members, err := db.ListData(&member, nil)
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

func CheckLogin(c *fiber.Ctx)  {
	auth := c.Get(fiber.HeaderAuthorization)
	// Check if header is valid
	if len(auth) > 7 && strings.ToLower(auth[:6]) == "bearer" {

		token, err := jwt.Parse(auth[7:], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var id int64
			if val, ok := claims["user_id"]; !ok {
				c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
				return
			} else {
				id = int64(val.(float64))
			}
			c.Locals("user_id", id)

			member := new(model.Member)
			member.ID.SetValid(id)
			members, err := db.ListData(member, nil)
			if err != nil {
				_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
				return
			}
			if len(members) == 0 {
				_ = c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
				return
			}
			member = members[0].(*model.Member)

			c.Locals("user_data", member)

			c.Next()
			return
		} else {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
			return
		}
	}
	c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
}

func CheckCanGuide(c *fiber.Ctx) {
	member := c.Locals("user_data").(*model.Member)
	if !member.Verify.Bool {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		return
	}
	c.Next()
	return
}

func ForgotPassword(c *fiber.Ctx) {
	input := new(model.Member)
	if err := c.BodyParser(input); err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}

	if !(input.Username.Valid || input.Email.Valid) {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Please enter username or email"})
		return
	}

	member := model.Member{
		Username: input.Username,
		Email: input.Email,
	}

	members, err := db.ListData(&member, nil)
	if err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}
	if len(members) == 0 {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Username or E-mail or Password incorrect."})
		return
	}

	member = *(members[0].(*model.Member))

	exp := time.Now().Add(time.Minute * 30).Unix()
	tm := exp % 17 + 5
	password := member.Password.String
	password = password[tm: tm + 5]

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": exp,
		"r": member.ID.Int64,
		"u": member.Username.String,
		"c": password,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		return
	}

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		os.Getenv("EMAIL_USER"),
		os.Getenv("EMAIL_PASS"),
		os.Getenv("EMAIL_SMTP"),
	)
	msg := []byte("To: " + member.Email.String + "\r\n" +
		"Subject: ไปเที่ยวกัน : รหัสสำหรับรีเซ็ตรหัสผ่าน\r\n" +
		"\r\n" +
		"รหัสข้างล่างนี้ไว้ใช้สำหรับรีเซ็ตรหัสผ่าน รหัสนี้สามารถใช้ได้ 30 นาที นับจากขอรีเซ็ตรหัสผ่าน\r\n\r\n" + tokenString + "\r\n")
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err = smtp.SendMail(
		os.Getenv("EMAIL_SMTP") + ":" + os.Getenv("EMAIL_PORT"),
		auth,
		os.Getenv("EMAIL_USER"),
		[]string{member.Email.String},
		msg,
	)
	if err != nil {
		log.Fatal(err)
	}

	_ = c.JSON(fiber.Map{"success": "Reset code send to your E-mail"})
}

type resetReq struct {
	Code		string	`json:"code"`
	Password	string	`json:"password"`
}

func Reset(c *fiber.Ctx) {
	input := new(resetReq)
	if err := c.BodyParser(input); err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}

	if input.Code == "" || input.Password == "" {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Please enter code and password"})
		return
	}

	token, err := jwt.Parse(input.Code, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var id, exp int64
		var pass, user string
		if val, ok := claims["r"]; !ok {
			c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			return
		} else {
			id = int64(val.(float64))
		}
		if val, ok := claims["u"]; !ok {
			c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			return
		} else {
			user = val.(string)
		}
		if val, ok := claims["c"]; !ok {
			c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			return
		} else {
			pass = val.(string)
		}
		if val, ok := claims["exp"]; !ok {
			c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			return
		} else {
			exp = int64(val.(float64))
		}

		member := new(model.Member)
		member.Username.SetValid(user)
		member.ID.SetValid(id)

		members, err := db.ListData(member, nil)
		if err != nil {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}
		if len(members) == 0 {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Username or E-mail or Password incorrect."})
			return
		}

		member = members[0].(*model.Member)

		tm := exp % 17 + 5
		password := member.Password.String
		password = password[tm: tm + 5]

		if pass != password {
			c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			return
		}

		member.Password.SetValid(input.Password)
		Update(c, member)
	} else {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}
}