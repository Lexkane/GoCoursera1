package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Result struct {
	StatusCode int
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("id")
	if key == "42" {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status":200, "resp":{"user":42}}`)
	}
}

func TestGetUser(t *testing.T) {
	cases := []TestCase{
		TestCase{
			ID:         "42",
			Response:   `{"status": 200, "resp" : {"user":42}}`,
			StatusCode: http.StatusOk,
		},
		TestCase{
			ID:         "500",
			Response:   `{"status": 500, "err" : "db_error"}`,
			StatusCode: http.StatusInternalServerError,
		},
	}

	for caseNum, item := range cases {
		url := "http://example.com/api/user?id=" + item.ID
		req := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()

		GetUser(w, req)
		if w.Code != item.StatusCode {
			t.Errorf("[%d] worng StatusCode: got %d, expected %d", caseNum, w.Code, item.StatusCode)
		}
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if bodyStr != item.Response {
			t.Errorf("[%d] wrong Response : got %+v, expected %+v",
				caseNum, bodyStr, item.Response)
		}
	}
}