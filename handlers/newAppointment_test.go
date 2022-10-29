package handlers

import (
	"io"
	"net/http"
	"testing"

	"github.com/smartystreets/assertions/should"

	"github.com/sheikhrachel/future/api_common/util/testutil"
)

func TestHandler_NewAppointment(t *testing.T) {
	r := SetUpRouter()
	testCases := []struct {
		name      string
		bodyBytes []byte
		expected  string
	}{
		{
			name: "endpoint should correctly return the added appointment",
			bodyBytes: []byte(`{
		        "ended_at": "2019-01-24T10:30:00-08:00",
		        "id": 2,
		        "user_id": 2,
		        "started_at": "2019-01-24T10:00:00-08:00",
		        "trainer_id": 1
		    }`),
			expected: "{" +
				"\"appointment_added\":{" +
				"\"id\":2," +
				"\"trainer_id\":1," +
				"\"user_id\":2," +
				"\"started_at\":\"2019-01-24T10:00:00-08:00\"," +
				"\"ended_at\":\"2019-01-24T10:30:00-08:00\"}" +
				"}",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := sendTestRequest(r, http.MethodPost, PathNewAppointment, tc.bodyBytes)
			responseData, _ := io.ReadAll(w.Body)
			testutil.So(t, string(responseData), should.Resemble, tc.expected)
			testutil.So(t, w.Code, should.Equal, http.StatusOK)
		})
	}
}
