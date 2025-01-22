//go:build unit

package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHealth(t *testing.T) {
	tc := []struct {
		desc   string
		expect int
	}{
		{
			desc:   "on success should return 200",
			expect: http.StatusOK,
		},
	}

	for i, c := range tc {
		h := handleHealth()
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		h.ServeHTTP(w, r)

		if got := w.Code; got != c.expect {
			t.Errorf("[%d] %s: expected %d; got %d", i, c.desc, c.expect, got)
		}
	}
}
