package tests

import (
	"go-fiber-api/database"
	controllers "go-fiber-api/internal/controller"
	"go-fiber-api/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func SetupTestApp() *fiber.App {
	app := fiber.New()

	database.ConnectDB()

	app.Get("/users", controllers.GetUsers)
	app.Get("/user/:id", controllers.GetUserByID)
	app.Post("/user", controllers.CreateUserNew)
	app.Put("/user/:id", controllers.UpdateUser)
	app.Delete("/user/:id", controllers.DeleteUser)

	return app
}

// Test untuk mendapatkan semua user
func TestGetUsers(t *testing.T) {
	app := SetupTestApp()

	req := httptest.NewRequest("GET", "/users", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

// Test untuk mendapatkan user berdasarkan ID
func TestGetUser(t *testing.T) {
	app := SetupTestApp()
	req := httptest.NewRequest("GET", "/user/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

// Test untuk mendapatkan user yang tidak ada
func TestGetUserNotFound(t *testing.T) {
	app := SetupTestApp()

	req := httptest.NewRequest("GET", "/user/99", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 404, resp.StatusCode)
}

// Test untuk membuat user baru
func TestCreateUser(t *testing.T) {
	app := SetupTestApp()

	newUser := models.User{ID: 11, Name: "User Tester", Email: "tester@joko.com"}
	jsonBook, _ := json.Marshal(newUser)
	req := httptest.NewRequest("POST", "/user", bytes.NewReader(jsonBook))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, 201, resp.StatusCode)
}

// Test untuk update user
func TestUpdateUser(t *testing.T) {
	app := SetupTestApp()

	updatedUser := models.User{ID: 11, Name: "Updated User Tester", Email: "Updated@joko.com"}
	jsonBook, _ := json.Marshal(updatedUser)
	req := httptest.NewRequest("PUT", "/user/11", bytes.NewReader(jsonBook))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

// Test untuk menghapus user
func TestDeleteUser(t *testing.T) {
	app := SetupTestApp()

	req := httptest.NewRequest("DELETE", "/user/11", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}
