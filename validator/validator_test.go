package validator

import "testing"

func TestValidator(t *testing.T) {
	validate := New(
		WithLocale("zh"),
		WithTagname("label"),
		)
	req := validate.Validate(struct {
		Name string `label:"name" validate:"required,min=6"`
	}{
		Name: "raylin666",
	})

	if len(req) > 0 {
		t.Fatal(req)
	}

	t.Log("SUCCESS")
}