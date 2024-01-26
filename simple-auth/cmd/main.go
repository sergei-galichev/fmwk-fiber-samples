package main

import (
	"errors"
	"fmt"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"time"
)

type (
	AuthHandler struct {
		storage *AuthStorage
	}
	AuthStorage struct {
		users map[string]User
	}
	UserHandler struct {
		storage *AuthStorage
	}
	User struct {
		Email    string
		Name     string
		password string
	}
)
type (
	RegisterRequest struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		AccessToken string `json:"access_token"`
	}

	ProfileResponse struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
)

// Defined global variables
var (
	errBadCredentials = errors.New("email or password is incorrect")
	errUserNotFound   = errors.New("user not found")

	jwtSecretKey = []byte("very-very-secret-key")
)

// Defined constants
const (
	contextKeyUser = "user"
)

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	regReq := LoginRequest{}
	if err := c.BodyParser(&regReq); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	user, exists := h.storage.users[regReq.Email]
	if !exists {
		return errBadCredentials
	}
	if user.password != regReq.Password {
		return errBadCredentials
	}

	payload := jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		logrus.WithError(err).Errorf("failed to sign jwt token")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(LoginResponse{t})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	regReq := RegisterRequest{}
	if err := c.BodyParser(&regReq); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}
	if _, exists := h.storage.users[regReq.Email]; exists {
		return fmt.Errorf("user with email %s already exists", regReq.Email)
	}
	h.storage.users[regReq.Email] = User{
		Email:    regReq.Email,
		Name:     regReq.Name,
		password: regReq.Password,
	}
	return c.SendStatus(fiber.StatusCreated)
}

func (h *UserHandler) Profile(c *fiber.Ctx) error {
	jwtPayload, ok := jwtPayloadFromRequest(c)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	userInfo, ok := h.storage.users[jwtPayload["sub"].(string)]
	if !ok {
		return errUserNotFound
	}

	return c.JSON(ProfileResponse{
		Email: userInfo.Email,
		Name:  userInfo.Name,
	})
}

// Private methods
func jwtPayloadFromRequest(c *fiber.Ctx) (jwt.MapClaims, bool) {
	contextValue := c.Context().Value(contextKeyUser)
	jwtToken, ok := contextValue.(*jwt.Token)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"jwt_token_context_value": contextValue,
		}).Error("wrong type of JWT token in context")
		return nil, false
	}
	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"jwt_token_claims": jwtToken.Claims,
		}).Error("wrong type of JWT token claims")
		return nil, false
	}
	return payload, true
}

func main() {
	webApp := fiber.New()

	authStorage := &AuthStorage{map[string]User{}}
	authHandler := &AuthHandler{authStorage}
	userHandler := &UserHandler{authStorage}

	publicGroup := webApp.Group("")
	publicGroup.Post("/register", authHandler.Register)
	publicGroup.Get("/login", authHandler.Login)

	authorizedGroup := webApp.Group("")
	authorizedGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: jwtSecretKey,
		},
		ContextKey: contextKeyUser,
	}))
	authorizedGroup.Get("/profile", userHandler.Profile)

	logrus.Fatal(webApp.Listen(":8080"))
}
