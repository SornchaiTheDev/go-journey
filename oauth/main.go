package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	authorizeUrl = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenUrl     = "https://oauth2.googleapis.com/token"
	userInfo     = "https://www.googleapis.com/oauth2/v3/userinfo"
	baseUri      = "https://oauth.sornchaithedev.com/callback"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	app := fiber.New()

	app.Get("/login", func(c *fiber.Ctx) error {

		u, err := url.Parse(authorizeUrl)
		if err != nil {
			log.Fatal(err)
		}

		q := u.Query()

		q.Set("response_type", "code")
		q.Set("client_id", clientId)
		q.Set("redirect_uri", baseUri)
		q.Set("scope", "openid email")
		q.Set("state", "test") // change this to check if request in valid from our server

		u.RawQuery = q.Encode()

		return c.Redirect(u.String())
	})

	app.Get("/callback", func(c *fiber.Ctx) error {
		queries := c.Queries()

		if queries["code"] != "" {

			if queries["state"] != "test" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "UNAUTHORIZE",
				})
			}

			u, err := url.Parse(tokenUrl)
			if err != nil {
				log.Fatal(err)
			}

			q := u.Query()

			q.Set("grant_type", "authorization_code")
			q.Set("client_id", clientId)
			q.Set("client_secret", clientSecret)
			q.Set("redirect_uri", baseUri)
			q.Set("code", queries["code"])

			u.RawQuery = q.Encode()

			agent := fiber.Post(u.String())
			statusCode, body, errs := agent.Bytes()

			if len(errs) > 0 {
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"errors": errs,
				})
			}

			var resp fiber.Map

			err = json.Unmarshal(body, &resp)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code": "SOMETHING_WENT_WRONG",
				})
			}

			accesstoken := resp["access_token"]

			agent = fiber.Get(userInfo)

			agent.Set("Authorization", fmt.Sprintf("Bearer %s", accesstoken))

			statusCode, body, errs = agent.Bytes()
			if len(errs) > 0 {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"erros": errs,
				})
			}

			return c.Status(statusCode).Send(body)

		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "UNAUTHORIZE",
		})

	})

	app.Listen(":8080")
}
