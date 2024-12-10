package main

import (
	"crypto/md5"
	"encoding/hex"
//	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/redis/rueidis"
	"log"
	"os"
	"os/exec"
)

func main() {
	e := echo.New()
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"localhost:6379"},
	})
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	e.GET("/:repository_address", func(c echo.Context) error {
		repoURL := c.Param("repository_address")
		repoHash := getMD5(repoURL)
		redisKey := fmt.Sprintf("repo:%s", repoHash)

		result, err := client.Do(c.Request().Context(), client.B().Get().Key(redisKey).Build()).ToString()
		if err == nil && result != "" {
			// Cached result found
			return c.Redirect(302, fmt.Sprintf("/result/%s", repoHash))
		}

		if err := cloneRepository(repoURL); err != nil {
			return c.String(500, fmt.Sprintf("Failed to clone repository: %v", err))
		}
		defer os.RemoveAll("repo") // Clean up cloned repo

		tokeiResult, err := runTokei()
		if err != nil {
			return c.String(500, fmt.Sprintf("Failed to run Tokei: %v", err))
		}

		// Store result in Redis
		if err := client.Do(c.Request().Context(), client.B().Set().Key(redisKey).Value(tokeiResult).Build()).Error(); err != nil {
			return c.String(500, fmt.Sprintf("Failed to store result in Redis: %v", err))
		}

		return c.Redirect(302, fmt.Sprintf("/result/%s", repoHash))
	})

	e.GET("/result/:md5_of_repository_address", func(c echo.Context) error {
		repoHash := c.Param("md5_of_repository_address")
		redisKey := fmt.Sprintf("repo:%s", repoHash)

		result, err := client.Do(c.Request().Context(), client.B().Get().Key(redisKey).Build()).ToString()
		if err != nil || result == "" {
			return c.String(404, "Not found")
		}

		return c.JSONBlob(200, []byte(result))
	})

	e.Logger.Fatal(e.Start(":5001"))
}

func getMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func cloneRepository(repoURL string) error {
	cmd := exec.Command("git", "clone", repoURL, "repo")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runTokei() (string, error) {
	cmd := exec.Command("tokei", ".", "-o", "json")
	cmd.Dir = "./repo"
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
