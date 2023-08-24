package getredis

import (
	"encoding/json"
	dtredis "golangredis/models/dtRedis"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/thedevsaddam/renderer"
)

// MockData adalah mock dari getData interface
type MockData struct {
	mock.Mock
}

// GetDataRedis adalah implementasi mock untuk method GetDataRedis
func (m *MockData) GetDataRedis(req dtredis.DataSet) (resp dtredis.RespGetData, err error) {
	args := m.Called(req)
	return args.Get(0).(dtredis.RespGetData), args.Error(1)
}

func TestDataHandler(t *testing.T) {
	render := renderer.New()

	mockData := new(MockData)
	handler := NewHandler(mockData, render)

	// Mock untuk kasus sukses
	mockResp := dtredis.RespGetData{Key: "MockValue"}
	mockData.On("GetDataRedis", mock.Anything).Return(mockResp, nil)

	req, err := http.NewRequest(http.MethodGet, "/learning/get?key=myKey", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	httpHandler := http.HandlerFunc(handler.GetDataRedis)
	httpHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var resBody dtredis.ResponseData
	if err := json.NewDecoder(rr.Body).Decode(&resBody); err != nil {
		t.Fatal(err)
	}

	expectedBody := dtredis.ResponseData{
		Message: "Success",
		Data:    mockResp,
	}

	if resBody.Message != expectedBody.Message {
		t.Errorf("handler returned unexpected message: got %s want %s",
			resBody.Message, expectedBody.Message)
	}

	if resBody.Data != expectedBody.Data {
		t.Errorf("handler returned unexpected data: got %v want %v",
			resBody.Data, expectedBody.Data)
	}

	// Pastikan mock dipanggil sesuai ekspektasi
	mockData.AssertExpectations(t)
}
