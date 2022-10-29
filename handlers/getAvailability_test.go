package handlers

import (
	"io"
	"net/http"
	"testing"

	"github.com/smartystreets/assertions/should"

	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/util/testutil"
)

func TestHandler_GetAvailability(t *testing.T) {
	r := SetUpRouter()
	cc := call.New()
	h := NewHandler(cc)
	testCases := []struct {
		name       string
		bodyBytes  []byte
		shouldBook bool
		expected   string
	}{
		{
			name: "endpoint should correctly return the open availability",
			bodyBytes: []byte(`{
		        "trainer_id": 1,
				"starts_at": "2019-01-24T09:30:00-08:00",
				"ends_at": "2019-01-24T13:00:00-08:00"
		    }`),
			shouldBook: false,
			expected: "{" +
				"\"trainer_id\":1," +
				"\"upcoming_availability\":[" +
				"\"2019-01-24T09:30:00-08:00\"," +
				"\"2019-01-24T10:00:00-08:00\"," +
				"\"2019-01-24T10:30:00-08:00\"," +
				"\"2019-01-24T11:00:00-08:00\"," +
				"\"2019-01-24T11:30:00-08:00\"," +
				"\"2019-01-24T12:00:00-08:00\"," +
				"\"2019-01-24T12:30:00-08:00\"" +
				"]" +
				"}",
		},
		{
			name: "endpoint should correctly account for bookings arising",
			bodyBytes: []byte(`{
		        "trainer_id": 1,
				"starts_at": "2019-01-24T09:30:00-08:00",
				"ends_at": "2019-01-24T13:00:00-08:00"
		    }`),
			shouldBook: true,
			expected: "{" +
				"\"trainer_id\":1," +
				"\"upcoming_availability\":[" +
				"\"2019-01-24T09:30:00-08:00\"," +
				"\"2019-01-24T10:30:00-08:00\"," +
				"\"2019-01-24T11:30:00-08:00\"," +
				"\"2019-01-24T12:00:00-08:00\"" +
				"]" +
				"}",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldBook {
				h.bookAppointments()
				defer h.cleanupAppointments()
			}
			w := sendTestRequest(r, http.MethodGet, PathGetAvailability, tc.bodyBytes)
			responseData, _ := io.ReadAll(w.Body)
			testutil.So(t, string(responseData), should.Resemble, tc.expected)
			testutil.So(t, w.Code, should.Equal, http.StatusOK)
		})
	}
}
