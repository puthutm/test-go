package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"unsia.ac.id/akademic_be/pkg/validation"
)

func TestValidateExtensionFileSuccess(t *testing.T) {
	check := validation.IsValidFileExtension("test.png", []string{".png", ".jpg"})

	assert.Equal(t, true, check)
}

func TestValidateExtensionFileFailed(t *testing.T) {
	check := validation.IsValidFileExtension("test.png", []string{".jpeg", ".jpg"})

	assert.Equal(t, false, check)
}
