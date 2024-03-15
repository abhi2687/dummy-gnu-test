package dummy

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func DummyFunction() {
	log := logrus.New()
	fmt.Println("Dummy Function")
	log.Println("Dummy Function")
}