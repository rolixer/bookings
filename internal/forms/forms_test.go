package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	form := New(url.Values{})

	isValid := form.Valid()
	if !isValid {
		t.Error("Could not validate valid form")
	}
}

func TestForm_Required(t *testing.T) {

	form := New(url.Values{})

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form is validated not having required fields")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	form = New(postedData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Form is not validated having required fields")
	}
}

func TestForm_MinLength(t *testing.T) {

	postedData := url.Values{}
	postedData.Add("short", "123")
	postedData.Add("equal", "1234")
	postedData.Add("longer", "1235")

	form := New(postedData)

	form.MinLength("short", 4)

	if form.Valid() {
		t.Error("Form got validated with to short field")
	}

	isError := form.Errors.Get("short")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	form = New(postedData)

	form.MinLength("longer", 4)

	if !form.Valid() {
		t.Error("Form not validated with longer field")
	}

	isError = form.Errors.Get("longer")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}

}

func TestFrom_IsEmail(t *testing.T) {

	postedData := url.Values{}
	postedData.Add("valid", "test@gmail.com")
	postedData.Add("invalid", "asdfasdfsdf")

	form := New(postedData)

	form.IsEmail("valid")

	if !form.Valid() {
		t.Error("Valid email not validated")
	}

	form.IsEmail("invalid")

	if form.Valid() {
		t.Error("Inalid email validated")
	}
}

func TestForm_Has(t *testing.T) {

	postedData := url.Values{}
	postedData.Add("data", "data")

	form := New(postedData)

	form.Has("data")

	if !form.Valid() {
		t.Error("Form not validated having field")
	}

	form.Has("diffrent_data")

	if !form.Valid() {
		t.Error("Form validated not having field")
	}

}
