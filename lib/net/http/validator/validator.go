package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

var validators = make(map[string]validator.FuncCtx)

// Register regist custom validator.
func Register() {
	if va, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for k, v := range validators {
			va.RegisterValidationCtx(k, v)
		}
	}
}

// List all custom validators.
func List() []string {
	var keys = make([]string, 0, len(validators))
	for k := range validators {
		keys = append(keys, k)
	}
	return keys
}

func register(tag string, fn validator.FuncCtx) {
	validators[tag] = fn
}
