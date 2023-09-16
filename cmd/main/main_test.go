package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	// "reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	// "gorm.io/driver/mysql"
	"net/http"
	"net/http/httptest"

	app "github.com/rachnatiwari/bookstore_management_app/pkg/config"
	controller "github.com/rachnatiwari/bookstore_management_app/pkg/controllers"
	"github.com/rachnatiwari/bookstore_management_app/pkg/models"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// type Book struct {
// 	gorm.Model
// 	Name        string `json:"name"`
// 	Author      string `json:"author"`
// 	Publication string `json:"publication"`
// }

func TestMain(m *testing.M) {

	// dropTable
	// ensureTableExists
	db = app.GetDB()
	code := m.Run()
	fmt.Println("after the run")
	clearTable()

	os.Exit(code)
}

func clearTable() {
	db.Exec("DELETE from books")
}

func dropTable() {
	db.Exec("DROP table employee")
}

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/books", nil)
	handler := http.HandlerFunc(controller.ShowAllBooks)
	response := executeRequest(req, &handler)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["Data"] != nil {
		t.Errorf("Expected an empty array. Got %s", m["Data"])
	}
}

func TestAddBook(t *testing.T) {
	var jsonStr = []byte(`{"Name":"test name","Author":"test author","Publication":"test publication"}`)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonStr))
	handler := http.HandlerFunc(controller.AddNewBook)
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req, &handler)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	expected := map[string]string{"name": "test name", "author": "test author", "publication": "test publication"}
	if m["name"] != expected["name"] || m["author"] != expected["author"] || m["publication"] != expected["publication"] {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}

}

func TestGetBook(t *testing.T) {
	req, _ := http.NewRequest("GET", "/books", nil)
	handler := http.HandlerFunc(controller.ShowAllBooks)
	response := executeRequest(req, &handler)
	fmt.Println(response.Code)
	checkResponseCode(t, http.StatusOK, response.Code)
	expected := `'{"name": "test name", "author": "test author", "publication": "test publication"}`
	if strings.Contains(response.Body.String(), expected) {
		t.Errorf("Not getting books from db")
	}
}

func TestUpdateProduct(t *testing.T) {
	var jsonStr = []byte(`{"Name":"test updated name","Author":"test author","Publication":"test publication"}`)
	book_to_update := models.GetAllBooks()
	json.Marshal(book_to_update)
	bookId := book_to_update[0].ID
	req, err := http.NewRequest("PUT", "/book/"+strconv.FormatUint(uint64(bookId), 10), bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(controller.AddNewBook)
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req, &handler)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["name"] != "test updated name" {
		t.Errorf("Expected - test updated name. Got %s", m["name"])
	}
}

func TestDeleteEmployee(t *testing.T) {
	book_to_delete := models.GetAllBooks()
	json.Marshal(book_to_delete)
	bookId := book_to_delete[0].ID
	req, err := http.NewRequest("DELETE", "/book/"+strconv.FormatUint(uint64(bookId), 10), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(controller.DeleteBook)
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req, &handler)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["name"] != "" {
		t.Errorf("Expected horse. Got %s", m["name"])
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request, r *http.HandlerFunc) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	return rr
}
