package main

import (
	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/gofiber/cors"
	_ "github.com/gofiber/fiber"
	_ "github.com/jinzhu/copier"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/pkg/errors"
	_ "golang.org/x/crypto/bcrypt"
	_ "gopkg.in/guregu/null.v3"
)