package env

// Environment is the type of environment this project is being ran in.
type Environment string

const (
	// DevelopmentEnv is a development environment.
	DevelopmentEnv Environment = "DEVELOPMENT"
	// ProductionEnv is a production environment.
	ProductionEnv Environment = "PRODUCTION"
	// TestingEnv is a testing environment.
	TestingEnv Environment = "TESTING"
)

// Development will return true if the environment is development
func (e Environment) Development() bool {
	return e == DevelopmentEnv
}

// Production will return true if the environment is production
func (e Environment) Production() bool {
	return e == ProductionEnv
}

// Testing will return true if the environment is testing
func (e Environment) Testing() bool {
	return e == TestingEnv
}
