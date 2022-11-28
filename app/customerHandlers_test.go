package app

import (
	"github.com/akshanshgusain/Hexagonal-Architecture/dto"
	"github.com/akshanshgusain/Hexagonal-Architecture/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	// Arrange
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := service.NewMockCustomerService(controller)

	dummyCustomers := []dto.CustomerResponse{
		{"1001", "Akshansh", "New Delhi", "110011", "2000-01-01", "1"},
		{"1002", "Ranveer", "New Delhi", "110011", "2000-01-01", "1"},
	}

	mockService.EXPECT().GetAllCustomers().Return(dummyCustomers, nil)
	ch := CustomerHandlers{mockService}

	// router
	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)

	// Create a http request
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {

}
