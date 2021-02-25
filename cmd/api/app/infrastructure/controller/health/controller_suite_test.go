package health

import (
	"testing"
)

func TestControllerHealth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller")
}
