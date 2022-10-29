package handlers

import (
	"io"
	"net/http"
	"testing"

	"github.com/smartystreets/assertions/should"

	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/util/testutil"
)

func TestHandler_GetBookings(t *testing.T) {
	r := SetUpRouter()
	cc := call.New()
	h := NewHandler(cc)
	url := "/appointments/bookings/1"
	testCases := []struct {
		name       string
		shouldBook bool
		expected   string
	}{
		{
			name:       "endpoint should correctly return the booked appointments",
			shouldBook: true,
			expected: "{" +
				"\"booked\":[" +
				"{" +
				"\"id\":0," +
				"\"trainer_id\":1," +
				"\"user_id\":1," +
				"\"started_at\":\"2019-01-24 10:00:00 -0800 PST\"," +
				"\"ended_at\":\"2019-01-24 10:30:00 -0800 PST\"" +
				"}," +
				"{" +
				"\"id\":1," +
				"\"trainer_id\":1," +
				"\"user_id\":2," +
				"\"started_at\":\"2019-01-24 11:00:00 -0800 PST\"," +
				"\"ended_at\":\"2019-01-24 11:30:00 -0800 PST\"" +
				"}," +
				"{" +
				"\"id\":2," +
				"\"trainer_id\":1," +
				"\"user_id\":3," +
				"\"started_at\":\"2019-01-24 12:30:00 -0800 PST\"," +
				"\"ended_at\":\"2019-01-24 13:00:00 -0800 PST\"" +
				"}" +
				"]," +
				"\"trainer_id\":1" +
				"}",
		},
		{
			name:       "endpoint should correctly return no appointments if none have been booked",
			shouldBook: false,
			expected:   "{\"booked\":null,\"trainer_id\":1}",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldBook {
				h.bookAppointments()
				defer h.cleanupAppointments()
			}
			w := sendTestRequest(r, http.MethodGet, url, nil)
			responseData, _ := io.ReadAll(w.Body)
			testutil.So(t, string(responseData), should.Resemble, tc.expected)
			testutil.So(t, w.Code, should.Equal, http.StatusOK)
		})
	}
}
