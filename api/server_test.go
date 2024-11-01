package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-feature-flag/app-api/api"
	"github.com/go-feature-flag/app-api/dao"
	"github.com/go-feature-flag/app-api/handler"
	"github.com/go-feature-flag/app-api/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setUpTest(t *testing.T) api.Server {
	// init the in-memory mock dao
	dbImpl, err := dao.NewInMemoryMockDao()
	dbImpl.SetFlags(testutils.DefaultInMemoryFlags())
	require.NoError(t, err)

	// init the API handlers
	options := &handler.FlagAPIHandlerOptions{
		Clock: testutils.ClockMock{},
	}
	flagHandlers := handler.NewFlagAPIHandler(dbImpl, options)
	healthHandlers := handler.NewHealth(dbImpl)

	// port is not important since we are not really starting the server in the tests
	apiServer := api.New(":0", flagHandlers, healthHandlers)
	require.NotNil(t, apiServer)
	return *apiServer
}

func TestHealthRouteExist(t *testing.T) {
	apiServer := setUpTest(t)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	apiServer.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message":"API is up and running","code":200}`, rec.Body.String())
}

func TestRouteExist(t *testing.T) {
	apiServer := setUpTest(t)
	tests := []struct {
		name     string
		method   string
		path     string
		body     *string
		wantCode int
		wantBody string
	}{
		{
			name:     "GET /v1/flags/:id",
			method:   http.MethodGet,
			path:     "/v1/flags/926214f3-80c1-46e6-a913-b2d40b92a932",
			body:     nil,
			wantCode: http.StatusOK,
			wantBody: `{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2024-10-25T11:50:27Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"}}`,
		},
		{
			name:     "GET /health",
			method:   http.MethodGet,
			path:     "/health",
			body:     nil,
			wantCode: http.StatusOK,
			wantBody: `{"message":"API is up and running","code":200}`,
		},
		{
			name:     "GET /v1/flags",
			method:   http.MethodGet,
			path:     "/v1/flags",
			body:     nil,
			wantCode: http.StatusOK,
			wantBody: `[{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2024-10-25T11:50:27Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"}},{"id":"926214f3-80c1-46e6-a913-b2d40b92a111","name":"flagr6w8","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2024-10-25T11:50:27Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"}},{"id":"926214f3-80c1-46e6-a913-b2d40b92a222","name":"flagr576987209","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2024-10-25T11:50:27Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"}}]`,
		},
		{
			name:     "PATCH /v1/flags/:id/status",
			method:   http.MethodPatch,
			path:     "/v1/flags/926214f3-80c1-46e6-a913-b2d40b92a932/status",
			body:     testutils.String(`{"disable":true}`),
			wantCode: http.StatusOK,
			wantBody: `{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"},"disable":true}`,
		},
		{
			name:     "POST /v1/flags",
			method:   http.MethodPost,
			path:     "/v1/flags",
			body:     testutils.String(`{"id":"926214f3-80c1-46e6-a913-b2d40b92a933","name":"flag2","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"},"disable":true}`),
			wantCode: http.StatusCreated,
			wantBody: `{"id":"926214f3-80c1-46e6-a913-b2d40b92a933","name":"flag2","createdDate":"2020-01-01T00:00:00Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"toto","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"},"disable":true}`,
		},
		{
			name:     "PATCH /v1/flags/:id",
			method:   http.MethodPut,
			path:     "/v1/flags/926214f3-80c1-46e6-a913-b2d40b92a932",
			body:     testutils.String(`{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"},"disable":true}`),
			wantCode: http.StatusOK,
			wantBody: `{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"},"disable":true}`,
		},
		{
			name:     "DELETE /v1/flags/:id",
			method:   http.MethodDelete,
			path:     "/v1/flags/926214f3-80c1-46e6-a913-b2d40b92a932",
			body:     nil,
			wantCode: http.StatusNoContent,
			wantBody: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body string = ""
			if tt.body != nil {
				body = *tt.body
			}
			req := httptest.NewRequest(tt.method, tt.path, strings.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			apiServer.ServeHTTP(rec, req)
			assert.Equal(t, tt.wantCode, rec.Code)
			fmt.Println(rec.Body.String())
			if tt.wantBody == "" {
				assert.Empty(t, rec.Body.String())
				return
			}
			assert.JSONEq(t, tt.wantBody, rec.Body.String())
		})
	}
}
