// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/repository.go

// Package mocks is for a generated GoMock package.
package mocks

import (
	models "backend/internal/models"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDatabaseRepo is a mock of DatabaseRepo interface.
type MockDatabaseRepo struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseRepoMockRecorder
}

// MockDatabaseRepoMockRecorder is the mock recorder for MockDatabaseRepo.
type MockDatabaseRepoMockRecorder struct {
	mock *MockDatabaseRepo
}

// NewMockDatabaseRepo creates a new mock instance.
func NewMockDatabaseRepo(ctrl *gomock.Controller) *MockDatabaseRepo {
	mock := &MockDatabaseRepo{ctrl: ctrl}
	mock.recorder = &MockDatabaseRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseRepo) EXPECT() *MockDatabaseRepoMockRecorder {
	return m.recorder
}

// AllGenres mocks base method.
func (m *MockDatabaseRepo) AllGenres() ([]*models.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllGenres")
	ret0, _ := ret[0].([]*models.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllGenres indicates an expected call of AllGenres.
func (mr *MockDatabaseRepoMockRecorder) AllGenres() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllGenres", reflect.TypeOf((*MockDatabaseRepo)(nil).AllGenres))
}

// AllMovies mocks base method.
func (m *MockDatabaseRepo) AllMovies(genre ...int) ([]*models.Movie, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range genre {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AllMovies", varargs...)
	ret0, _ := ret[0].([]*models.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllMovies indicates an expected call of AllMovies.
func (mr *MockDatabaseRepoMockRecorder) AllMovies(genre ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMovies", reflect.TypeOf((*MockDatabaseRepo)(nil).AllMovies), genre...)
}

// Connection mocks base method.
func (m *MockDatabaseRepo) Connection() *sql.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connection")
	ret0, _ := ret[0].(*sql.DB)
	return ret0
}

// Connection indicates an expected call of Connection.
func (mr *MockDatabaseRepoMockRecorder) Connection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connection", reflect.TypeOf((*MockDatabaseRepo)(nil).Connection))
}

// DeleteMovie mocks base method.
func (m *MockDatabaseRepo) DeleteMovie(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMovie", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMovie indicates an expected call of DeleteMovie.
func (mr *MockDatabaseRepoMockRecorder) DeleteMovie(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMovie", reflect.TypeOf((*MockDatabaseRepo)(nil).DeleteMovie), id)
}

// GetUserByEmail mocks base method.
func (m *MockDatabaseRepo) GetUserByEmail(email string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockDatabaseRepoMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockDatabaseRepo)(nil).GetUserByEmail), email)
}

// GetUserByID mocks base method.
func (m *MockDatabaseRepo) GetUserByID(id int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockDatabaseRepoMockRecorder) GetUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockDatabaseRepo)(nil).GetUserByID), id)
}

// InsertMovie mocks base method.
func (m *MockDatabaseRepo) InsertMovie(movie models.Movie) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMovie", movie)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMovie indicates an expected call of InsertMovie.
func (mr *MockDatabaseRepoMockRecorder) InsertMovie(movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMovie", reflect.TypeOf((*MockDatabaseRepo)(nil).InsertMovie), movie)
}

// OneMovie mocks base method.
func (m *MockDatabaseRepo) OneMovie(id int) (*models.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OneMovie", id)
	ret0, _ := ret[0].(*models.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OneMovie indicates an expected call of OneMovie.
func (mr *MockDatabaseRepoMockRecorder) OneMovie(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OneMovie", reflect.TypeOf((*MockDatabaseRepo)(nil).OneMovie), id)
}

// OneMovieForEdit mocks base method.
func (m *MockDatabaseRepo) OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OneMovieForEdit", id)
	ret0, _ := ret[0].(*models.Movie)
	ret1, _ := ret[1].([]*models.Genre)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// OneMovieForEdit indicates an expected call of OneMovieForEdit.
func (mr *MockDatabaseRepoMockRecorder) OneMovieForEdit(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OneMovieForEdit", reflect.TypeOf((*MockDatabaseRepo)(nil).OneMovieForEdit), id)
}

// UpdateMovie mocks base method.
func (m *MockDatabaseRepo) UpdateMovie(movie models.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMovie", movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMovie indicates an expected call of UpdateMovie.
func (mr *MockDatabaseRepoMockRecorder) UpdateMovie(movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMovie", reflect.TypeOf((*MockDatabaseRepo)(nil).UpdateMovie), movie)
}

// UpdateMovieGenres mocks base method.
func (m *MockDatabaseRepo) UpdateMovieGenres(id int, genreIDs []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMovieGenres", id, genreIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMovieGenres indicates an expected call of UpdateMovieGenres.
func (mr *MockDatabaseRepoMockRecorder) UpdateMovieGenres(id, genreIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMovieGenres", reflect.TypeOf((*MockDatabaseRepo)(nil).UpdateMovieGenres), id, genreIDs)
}
