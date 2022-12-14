// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/honyshyota/weather-api/internal/models"
)

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepo) Create(user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepoMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepo)(nil).Create), user)
}

// FindByName mocks base method.
func (m *MockUserRepo) FindByName(name string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", name)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockUserRepoMockRecorder) FindByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockUserRepo)(nil).FindByName), name)
}

// UpdateFavCity mocks base method.
func (m *MockUserRepo) UpdateFavCity(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateFavCity", arg0, arg1)
}

// UpdateFavCity indicates an expected call of UpdateFavCity.
func (mr *MockUserRepoMockRecorder) UpdateFavCity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFavCity", reflect.TypeOf((*MockUserRepo)(nil).UpdateFavCity), arg0, arg1)
}

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUsecase) CreateUser(arg0 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUsecaseMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUsecase)(nil).CreateUser), arg0)
}

// FindUser mocks base method.
func (m *MockUsecase) FindUser(arg0 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockUsecaseMockRecorder) FindUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockUsecase)(nil).FindUser), arg0)
}

// GetAllCities mocks base method.
func (m *MockUsecase) GetAllCities() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCities")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCities indicates an expected call of GetAllCities.
func (mr *MockUsecaseMockRecorder) GetAllCities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCities", reflect.TypeOf((*MockUsecase)(nil).GetAllCities))
}

// GetForecastForTime mocks base method.
func (m *MockUsecase) GetForecastForTime(arg0, arg1 string) (*models.FullForecast, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForecastForTime", arg0, arg1)
	ret0, _ := ret[0].(*models.FullForecast)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForecastForTime indicates an expected call of GetForecastForTime.
func (mr *MockUsecaseMockRecorder) GetForecastForTime(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForecastForTime", reflect.TypeOf((*MockUsecase)(nil).GetForecastForTime), arg0, arg1)
}

// GetShortForecast mocks base method.
func (m *MockUsecase) GetShortForecast(arg0 string) (*models.ShortForecastResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShortForecast", arg0)
	ret0, _ := ret[0].(*models.ShortForecastResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShortForecast indicates an expected call of GetShortForecast.
func (mr *MockUsecaseMockRecorder) GetShortForecast(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShortForecast", reflect.TypeOf((*MockUsecase)(nil).GetShortForecast), arg0)
}

// UpdateFavCity mocks base method.
func (m *MockUsecase) UpdateFavCity(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFavCity", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFavCity indicates an expected call of UpdateFavCity.
func (mr *MockUsecaseMockRecorder) UpdateFavCity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFavCity", reflect.TypeOf((*MockUsecase)(nil).UpdateFavCity), arg0, arg1)
}

// MockCityRepo is a mock of CityRepo interface.
type MockCityRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCityRepoMockRecorder
}

// MockCityRepoMockRecorder is the mock recorder for MockCityRepo.
type MockCityRepoMockRecorder struct {
	mock *MockCityRepo
}

// NewMockCityRepo creates a new mock instance.
func NewMockCityRepo(ctrl *gomock.Controller) *MockCityRepo {
	mock := &MockCityRepo{ctrl: ctrl}
	mock.recorder = &MockCityRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCityRepo) EXPECT() *MockCityRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCityRepo) Create(cities models.CityArray) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", cities)
}

// Create indicates an expected call of Create.
func (mr *MockCityRepoMockRecorder) Create(cities interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCityRepo)(nil).Create), cities)
}

// GetAll mocks base method.
func (m *MockCityRepo) GetAll() ([]*models.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*models.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCityRepoMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCityRepo)(nil).GetAll))
}

// GetByName mocks base method.
func (m *MockCityRepo) GetByName(arg0 string) (*models.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", arg0)
	ret0, _ := ret[0].(*models.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockCityRepoMockRecorder) GetByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockCityRepo)(nil).GetByName), arg0)
}

// MockWeatherRepo is a mock of WeatherRepo interface.
type MockWeatherRepo struct {
	ctrl     *gomock.Controller
	recorder *MockWeatherRepoMockRecorder
}

// MockWeatherRepoMockRecorder is the mock recorder for MockWeatherRepo.
type MockWeatherRepoMockRecorder struct {
	mock *MockWeatherRepo
}

// NewMockWeatherRepo creates a new mock instance.
func NewMockWeatherRepo(ctrl *gomock.Controller) *MockWeatherRepo {
	mock := &MockWeatherRepo{ctrl: ctrl}
	mock.recorder = &MockWeatherRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeatherRepo) EXPECT() *MockWeatherRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWeatherRepo) Create(weathers []*models.CompleteWeather) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", weathers)
}

// Create indicates an expected call of Create.
func (mr *MockWeatherRepoMockRecorder) Create(weathers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWeatherRepo)(nil).Create), weathers)
}

// GetAll mocks base method.
func (m *MockWeatherRepo) GetAll() ([]*models.CompleteWeather, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*models.CompleteWeather)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockWeatherRepoMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockWeatherRepo)(nil).GetAll))
}

// GetByName mocks base method.
func (m *MockWeatherRepo) GetByName(arg0 string) (*models.CompleteWeather, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", arg0)
	ret0, _ := ret[0].(*models.CompleteWeather)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockWeatherRepoMockRecorder) GetByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockWeatherRepo)(nil).GetByName), arg0)
}

// Update mocks base method.
func (m *MockWeatherRepo) Update(weathers []*models.CompleteWeather) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", weathers)
}

// Update indicates an expected call of Update.
func (mr *MockWeatherRepoMockRecorder) Update(weathers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWeatherRepo)(nil).Update), weathers)
}

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetCities mocks base method.
func (m *MockClient) GetCities(arg0 []string) (models.CityArray, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCities", arg0)
	ret0, _ := ret[0].(models.CityArray)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCities indicates an expected call of GetCities.
func (mr *MockClientMockRecorder) GetCities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCities", reflect.TypeOf((*MockClient)(nil).GetCities), arg0)
}

// GetForecast mocks base method.
func (m *MockClient) GetForecast(arg0 []*models.City) ([]*models.CompleteWeather, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForecast", arg0)
	ret0, _ := ret[0].([]*models.CompleteWeather)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForecast indicates an expected call of GetForecast.
func (mr *MockClientMockRecorder) GetForecast(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForecast", reflect.TypeOf((*MockClient)(nil).GetForecast), arg0)
}
