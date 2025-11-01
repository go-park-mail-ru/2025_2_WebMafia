package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONResponses(t *testing.T) {
	t.Run("JSON success", func(t *testing.T) {
		rr := httptest.NewRecorder()
		data := map[string]string{"status": "ok"}
		JSON(rr, http.StatusOK, data)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		var respData map[string]string
		err := json.Unmarshal(rr.Body.Bytes(), &respData)
		require.NoError(t, err)
		assert.Equal(t, data, respData)
	})

	testCases := []struct {
		name       string
		handler    func(w http.ResponseWriter)
		expected   ErrorResponse
		statusCode int
	}{
		{"BadRequestJSON", BadRequestJSON, ErrBadRequest, http.StatusBadRequest},
		{"UnauthorizedJSON", UnauthorizedJSON, ErrUnauthorized, http.StatusUnauthorized},
		{"ForbiddenJSON", ForbiddenJSON, ErrForbidden, http.StatusForbidden},
		{"NotFoundJSON", NotFoundJSON, ErrNotFound, http.StatusNotFound},
		{"ConflictJSON", ConflictJSON, ErrConflict, http.StatusConflict},
		{"InternalErrorJSON", InternalErrorJSON, ErrInternalServer, http.StatusInternalServerError},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			tc.handler(rr)

			assert.Equal(t, tc.statusCode, rr.Code)
			var respBody ErrorResponse
			err := json.Unmarshal(rr.Body.Bytes(), &respBody)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, respBody)
		})
	}
}
