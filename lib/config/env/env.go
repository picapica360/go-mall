package env

import (
	"os"
)

const (
	// Dev env.
	Dev = "development"
	// Test env.
	Test = "test"
	// Prod env.
	Prod = "production"
)

// internal variable
var (
	// environment. The 'PICAPICA_ENV' is read in initialzation to set the variable.
	env = Dev
	// the rooted path name for the current directory.
	root string
)

func init() {
	setEnv(os.Getenv("PICAPICA_ENV"))
}

func setEnv(e string) {
	if len(e) > 0 {
		env = e
	}

	var err error
	root, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

// Env environment.
func Env() string {
	return env
}

// IsDevelopment check is Development env.
func IsDevelopment() bool {
	return env == Dev
}

// IsProduction check is production env.
func IsProduction() bool {
	return env == Prod
}

// IsEnvironment check is the specical env.
func IsEnvironment(env string) bool {
	return env == env
}

// Root get the rooted path name for the current directory.
func Root() string {
	return root
}
