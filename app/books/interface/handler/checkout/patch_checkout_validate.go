package checkout

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Validate content
func validate(a Params) error {
	return validation.ValidateStruct(&a,
		// IDは空を許容せず、1から30までの長さ
		validation.Field(&a.ID, validation.Required, validation.Length(1, 30)),
	)
}
