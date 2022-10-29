package handlers

import (
	"io"
	"net/http"
	"testing"

	"github.com/smartystreets/assertions/should"

	"github.com/sheikhrachel/future/api_common/util/testutil"
)

func TestHealthCheck(t *testing.T) {
	r := SetUpRouter()
	healthExpected := "{\"message\":\"health\"}"
	testCases := []struct {
		name string
		url  string
	}{
		{name: "method should correctly return health check for root path", url: PathRoot},
		{name: "method should correctly return health check for health path", url: PathHealth},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := sendTestRequest(r, http.MethodGet, tc.url, nil)
			responseData, _ := io.ReadAll(w.Body)
			testutil.So(t, string(responseData), should.Equal, healthExpected)
			testutil.So(t, w.Code, should.Equal, http.StatusOK)
		})
	}
}
