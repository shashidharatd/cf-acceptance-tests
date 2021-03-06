package cf

import (
	"github.com/cloudfoundry-incubator/cf-test-helpers/runner"
	. "github.com/onsi/gomega/gexec"
)

var Cf = func(args ...string) *Session {
	return runner.Run("cf", args...)
}
