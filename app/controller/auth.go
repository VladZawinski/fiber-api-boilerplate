package controller

import (
	"fiber-api-boilerplate/app/dto"
	"fiber-api-boilerplate/app/repository"
	"fiber-api-boilerplate/pkg/config"
	"fiber-api-boilerplate/platform/database"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	body := &dto.Auth{}
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	repo := repository.NewUserRepo(database.GetDB())
	user, err := repo.GetByUsername(body.Username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "user were not found",
		})
	}
	isValid := IsValidPassword([]byte(user.Password), []byte(body.Password))
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg": "password is wrong",
		})
	}
	// Generate a new Access token.
	token, err := GenerateNewAccessToken(user.ID, user.Role == "admin")
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"msg":          fmt.Sprintf("Token will be expired within %d minutes", config.AppCfg().JWTSecretExpireMinutesCount),
		"access_token": token,
	})
}

func Register(c *fiber.Ctx) error {
	body := &dto.CreateUser{}
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	repo := repository.NewUserRepo(database.GetDB())
	err := repo.Create(body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"msg": "Register successfully",
	})
}

func GenerateNewAccessToken(userID uint, isAdmin bool) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["admin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.AppCfg().JWTSecretExpireMinutesCount)).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.AppCfg().JWTSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func IsValidPassword(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		return false
	}

	return true
}
