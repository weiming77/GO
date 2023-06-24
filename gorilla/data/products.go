package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	validator "github.com/go-playground/validator/v10"
)

// product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,skuval"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"updateOn,omitempty"`
	DeletedOn   string  `json:"deleteOn,omitempty"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

/* Validator:
We could construct a validator outside/global data object because
we might need to register a lot of custom stuff and keep it at one
place or for convenience just save it into validate function.
*/
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("skuval", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	// SKU format is as format xxx-xxx-xxx-xxx
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+-[a-z]+`)
	// Field to get value and casting it into a string, of cause
	// it could be vary of many types so what here going to do is
	// return me a slice of string so we can call the match(es).
	matches := re.FindAllString(fl.Field().String(), -1)
	if fl.Field().String() == "invalid" {
		return false
	}
	// matches is a string of array of multi-lines with separator in it
	// if we don't have exactly one match then the validation is fail.
	if len(matches) != 1 {
		return false
	}

	return true
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "LTT245",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "ESSPRSS199",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}
