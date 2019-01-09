package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/my3/cloud-native-go/ping"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http/httptest"
)

var _ = Describe("Ping", func() {

	var (
		c *gin.Context
	)

	BeforeEach(func() {
		resp := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, _ = gin.CreateTestContext(resp)
	})


	Context("response to ping", func() {
		It("should be pong", func() {
			ping.RespondToPing(c)
			status := c.Writer.Status()
			Expect(status).To(Equal(200))
		})
	})
})