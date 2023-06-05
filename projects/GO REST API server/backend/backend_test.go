// Integration test:
// go mod tidy
// go test
package backend

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a App

const (
	tableProductCreationQuery = `CREATE TABLE IF NOT EXISTS products
	(
		id INT NOT NULL PRIMARY KEY AUTOINCREMENT,
		productCode VARCHAR(25) NOT NULL,
		name VARCHAR(256) NOT NULL,
		inventory DOUBLE NOT NULL,
		price DOUBLE NOT NULL,
		status VARCHAR(64) NOT NULL 
	)`

	tableOrderCreationQuery = `CREATE TABLE IF NOT EXISTS orders
	(
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		customerName VARCHAR(256) NOT NULL,
		total DOUBLE NOT NULL,
		status VARCHAR(64) NOT NULL
	)`

	tableOrderItemsCreationQuery = `CREATE Table IF NOT EXISTS order_items 
	(
		order_id INT,
		product_id INT,
		quantity DOUBLE NOT NULL,
		FOREIGN KEY (order_id) REFERENCES orders (id),
		FOREIGN KEY (product_id) REFERENCES products (id),
		PRIMARY KEY (order_id, product_id)
	)`
)

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize()
	ensureTableExists()
	code := m.Run()

	clearProductTable()
	clearOrderItemsTable()
	clearOrderTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableProductCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec(tableOrderCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec(tableOrderItemsCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearProductTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("DELECT FROM sqlite_sequence WHERE name = 'products'")
}

func clearOrderTable() {
	a.DB.Exec("DELETE FROM orders")
	a.DB.Exec("DELETE FROM sqlite_sequence WHERE name = 'orders'")
}

func clearOrderItemsTable() {
	a.DB.Exec("DELETE FROM order_items")
}

func TestGetNonExistenProduct(t *testing.T) {
	clearProductTable()

	req, _ := http.NewRequest(http.MethodGet, "/product/111", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusInternalServerError, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "sql: no rows in result set" {
		t.Errorf("Expected the 'error' key of the response to be set 'sql:no rows in result set'. Got '%s'", m["error"])
	}
}

func TestGetProduct(t *testing.T) {
	clearProductTable()
	addProducts(1)

	req, _ := http.NewRequest(http.MethodGet, "/product/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addProducts(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO products(productCode, name, inventory, price, status) Values (?, ?, ?, ?, ?)", "IT"+strconv.Itoa(i), "Product X"+strconv.Itoa(i), i, i, "test")
	}
}

func TestNonExistentOrder(t *testing.T) {
	clearOrderItemsTable()
	clearOrderTable()

	req, _ := http.NewRequest(http.MethodGet, "/order/111", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusInternalServerError, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "sql: no rows in result set" {
		t.Errorf("Expected the 'error' key of the response to be set 'sql:no rows in result set'. Got '%s'", m["error"])
	}
}

func TestCreateOrder(t *testing.T) {
	clearOrderItemsTable()
	clearOrderTable()

	payload := []byte(`{"customerName":"Customer TEST", "total":1, "status":"testing", "items": []}`)

	req, _ := http.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["customerName"] != "Customer TEST" {
		t.Errorf("Expected customerName to be 'Customer Test'. Actual '%v'", m["customerName"])
	}
	if m["total"] != 1.0 {
		t.Errorf("Expected total to be '1.0'. Actual '%s'", m["total"])
	}
	if m["status"] != "testing" {
		t.Errorf("Expected status to be 'testing'. Actual '%s'", m["status"])
	}
	if m["id"] != 1.0 {
		t.Errorf("Expected id to be '1'. Actual '%s'", m["id"])
	}
}

func TestGetOrder(t *testing.T) {
	clearOrderItemsTable()
	clearOrderTable()
	addOrders(1)

	req, _ := http.NewRequest(http.MethodGet, "/order/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addOrders(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO orders(customerName, total, status) Values (?, ?, ?)", "Customer "+strconv.Itoa(i), 1, "testing")
	}
}

func TestCreateOrderItem(t *testing.T) {
	clearOrderItemsTable()
	clearOrderTable()

	addProducts(1)
	addOrders(1)

	payload := []byte(`[{"order_id":0, "product_id":1, "quantity":1}]`)

	req, _ := http.NewRequest(http.MethodPost, "/orderitems", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m [](map[string]interface{})
	json.Unmarshal(response.Body.Bytes(), &m)

	if m[0]["order_id"] != 0.0 {
		t.Errorf("Expected order_id to be '0'. Actual '%v'", m[0]["order_id"])
	}
	if m[0]["product_id"] != 1.0 {
		t.Errorf("Expected product_id to be '1.0'. Actual '%s'", m[0]["product_id"])
	}
	if m[0]["quantity"] != 1.0 {
		t.Errorf("Expected quantity to be '1'. Actual '%s'", m[0]["quantity"])
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Actual %d\n", expected, actual)
	}
}

func TestCreateProduct(t *testing.T) {
	clearProductTable()

	payload := []byte(`{"productCode":"TEST12345", "name":"ProductTest", "inventory":1, "price":1.11, "status":"testing"}`)

	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["productCode"] != "TEST12345" {
		t.Errorf("Expected productCode to be 'TEST12345'. Actual '%v'", m["productCode"])
	}
	if m["name"] != "ProductTest" {
		t.Errorf("Expected name to be 'ProductTest'. Actual '%s'", m["name"])
	}
	if m["inventory"] != 1.0 {
		t.Errorf("Expected inventory to be '1.0'. Actual '%s'", m["inventory"])
	}
	if m["price"] != 1.11 {
		t.Errorf("Expected price to be '1.11'. Actual '%s'", m["price"])
	}
	if m["status"] != "testing" {
		t.Errorf("Expected status to be 'testing'. Actual '%s'", m["status"])
	}
	if m["id"] != 1.0 {
		t.Errorf("Expected id to be '1'. Actual '%s'", m["id"])
	}
}
