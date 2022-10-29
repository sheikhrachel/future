package router

import (
	"encoding/json"
	"io"

	"github.com/sheikhrachel/future/api_common/util/errutil"
)

func UnmarshallRequestBody(reqBody io.ReadCloser, req interface{}) error {
	body, err := io.ReadAll(reqBody)
	if errutil.HandleError(err) {
		return err
	}
	return json.Unmarshal(body, &req)
}
