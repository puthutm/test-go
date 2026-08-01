package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"unsia.ac.id/akademic_be/pkg/utils"
)

func BenchmarkGeneratorUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid := utils.GenerateUUID()
		assert.NotNil(b, uuid, "generate uuid succes")
	}
}

func BenchmarkGeneratorRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomString := utils.GeneratorRandomString(15)
		assert.NotNil(b, randomString, "yeay succes")
	}
}

func BenchmarkGeneratorRandomInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomNumber := utils.GeneratorRandomNumber()
		assert.NotNil(b, randomNumber, "yeay succes")
	}
}
