package main

import (
	"backend/internal/models"
	"backend/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

// TestHome tests the Home handler function.
func TestHome(t *testing.T) {

	// Create a request to pass to the handler.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	app := &Application{
		// Initialize fields if necessary
	}

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.Home) // app is your Application instance

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"status":"active","message":"Go Movies up and running","version":"1.0.0"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestApplication_AllMovies_Status(t *testing.T) {
	// Crear un controlador de mock y un mock de la base de datos
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mocks.NewMockDatabaseRepo(ctrl)

	// Configurar expectativas en el mock
	mockDB.EXPECT().AllMovies().Return([]*models.Movie{
		{
			ID:    1,
			Title: "Pelicula 1",
			// ... otros campos del struct Movie
		},
		{
			ID:    2,
			Title: "Pelicula 2",
			// ... otros campos del struct Movie
		},
	}, nil)

	// Crear una instancia de Application con el mock
	app := &Application{DB: mockDB}

	// Crear un request HTTP y un ResponseRecorder para capturar la respuesta
	req, err := http.NewRequest("GET", "/movies", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Llamar al handler
	handler := http.HandlerFunc(app.AllMovies)
	handler.ServeHTTP(rr, req)

	// Comprobar el código de estado y la respuesta
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	// ... puedes añadir más comprobaciones aquí, como verificar el cuerpo de la respuesta
}

func TestApplication_AllMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDatabaseRepo(ctrl)
	app := &Application{DB: mockDB}

	tests := []struct {
		name         string
		setupMock    func()
		expectedCode int
		expectedBody *[]models.Movie
	}{
		{
			name: "Caso Exitoso",
			setupMock: func() {
				mockDB.EXPECT().AllMovies().Return([]*models.Movie{{ID: 1, Title: "Ejemplo de Película 1"}}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: &[]models.Movie{{ID: 1, Title: "Ejemplo de Película 1"}},
		},
		{
			name: "Error al Recuperar Películas",
			setupMock: func() {
				mockDB.EXPECT().AllMovies().Return(nil, errors.New("error de base de datos"))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: nil, // En caso de error, no esperamos un cuerpo
		},
		// ... puedes añadir más casos aquí
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			req := httptest.NewRequest("GET", "/movies", nil)
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(app.AllMovies)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedCode)
			}

			if tc.expectedBody != nil {
				var gotMovies []models.Movie
				err := json.Unmarshal(rr.Body.Bytes(), &gotMovies)
				if err != nil {
					t.Fatal("No se pudo deserializar la respuesta:", err)
				}

				if !reflect.DeepEqual(gotMovies, *tc.expectedBody) {
					t.Errorf("handler returned unexpected body: got %v want %v", gotMovies, *tc.expectedBody)
				}
			}
		})
	}
}

// TestAuthenticate tests the authenticate function.
func TestAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockDatabaseRepo(ctrl)

	// Contraseña en texto plano y su versión hasheada para las pruebas
	plainPassword := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)

	// Configura el mock para devolver un usuario con la contraseña hasheada
	mockUser := &models.User{
		// ... otros datos del usuario ...
		Password: string(hashedPassword),
	}
	mockRepo.EXPECT().GetUserByEmail("valid@example.com").Return(mockUser, nil)

	// Configurar mocks adicionales si es necesario, por ejemplo, para la generación de JWT

	app := &Application{DB: mockRepo}

	// Crea un request con credenciales válidas
	credentials := map[string]string{"email": "valid@example.com", "password": plainPassword}
	payload, _ := json.Marshal(credentials)
	req, err := http.NewRequest("POST", "/authenticate", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.authenticate)

	handler.ServeHTTP(rr, req)

	// Realiza las aserciones necesarias
	assert.Equal(t, http.StatusAccepted, rr.Code) // Asegúrate de que el código de estado esperado sea correcto
	// Puedes añadir más aserciones para verificar el cuerpo de la respuesta y la presencia del JWT
}

func TestMovieCatalog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockDatabaseRepo(ctrl)

	// Configura el mock para devolver una lista de películas en el caso de éxito
	// Crear géneros de prueba
	genreAction := &models.Genre{ID: 1, Genre: "Action", Checked: true}
	genreComedy := &models.Genre{ID: 2, Genre: "Comedy", Checked: false}

	mockMovies := []*models.Movie{
		{
			ID:          1,
			Title:       "Movie One",
			ReleaseDate: time.Date(2021, time.January, 10, 0, 0, 0, 0, time.UTC),
			RunTime:     120,
			MPAARating:  "PG-13",
			Description: "Description of Movie One",
			Image:       "image1.jpg",
			Genres:      []*models.Genre{genreAction, genreComedy},
			GenresArray: []int{1, 2},
		},
		{
			ID:          2,
			Title:       "Movie Two",
			ReleaseDate: time.Date(2022, time.February, 20, 0, 0, 0, 0, time.UTC),
			RunTime:     90,
			MPAARating:  "R",
			Description: "Description of Movie Two",
			Image:       "image2.jpg",
			Genres:      []*models.Genre{genreComedy},
			GenresArray: []int{2},
		},
		// ... más películas si es necesario ...
	}

	mockRepo.EXPECT().AllMovies().Return(mockMovies, nil)

	app := &Application{DB: mockRepo}

	req, err := http.NewRequest("GET", "/moviecatalog", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.MovieCatalog)

	handler.ServeHTTP(rr, req)

	// Realiza las aserciones necesarias para el caso de éxito
	assert.Equal(t, http.StatusOK, rr.Code)
	// Puedes añadir más aserciones para verificar el cuerpo de la respuesta

	// Configura el mock para devolver un error en el caso de error
	mockRepo.EXPECT().AllMovies().Return(nil, errors.New("some internal error"))

	// Ejecuta de nuevo con el caso de error y realiza las aserciones necesarias
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Asegúrate de que el código de estado esperado sea correcto en caso de error
	// Cambia esto a http.StatusBadRequest si es lo que tu función está configurada para devolver
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockDatabaseRepo(ctrl)

	// Mock a successful database response
	mockMovie := &models.Movie{
		ID:          1,
		Title:       "Ejemplo de Película",
		ReleaseDate: time.Date(2021, time.January, 10, 0, 0, 0, 0, time.UTC),
		RunTime:     120,
		MPAARating:  "PG-13",
		Description: "Esta es una descripción de ejemplo para la película.",
		Image:       "imagen-ejemplo.jpg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Genres: []*models.Genre{
			{ID: 1, Genre: "Acción", Checked: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 2, Genre: "Aventura", Checked: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
		GenresArray: []int{1, 2},
	}

	mockRepo.EXPECT().OneMovie(1).Return(mockMovie, nil)

	app := &Application{DB: mockRepo}

	// Create a request for a valid movie ID
	r := chi.NewRouter()
	r.Get("/movies/{id}", app.GetMovie)
	req, err := http.NewRequest("GET", "/movies/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Assert for successful retrieval
	assert.Equal(t, http.StatusOK, rr.Code)

	// Mock a database error for an invalid movie ID
	mockRepo.EXPECT().OneMovie(999).Return(nil, errors.New("movie not found"))

	// Create a request for an invalid movie ID and test
	req, err = http.NewRequest("GET", "/movies/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Assert for error case - expecting 400 Bad Request instead of 500
	assert.Equal(t, http.StatusBadRequest, rr.Code) // Ajusta a http.StatusBadRequest

}

func TestUpdateMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockDatabaseRepo(ctrl)

	// Mock para la función OneMovie
	existingMovie := &models.Movie{ID: 1 /* otros campos */}
	mockRepo.EXPECT().OneMovie(1).Return(existingMovie, nil)

	// Mock para la función UpdateMovie
	mockRepo.EXPECT().UpdateMovie(gomock.Any()).Return(nil)

	// Mock para la función UpdateMovieGenres
	mockRepo.EXPECT().UpdateMovieGenres(1, gomock.Any()).Return(nil)

	app := &Application{DB: mockRepo}

	// Crear un payload de prueba
	updatedMovie := models.Movie{
		ID:          1,
		Title:       "Título Actualizado",
		ReleaseDate: time.Now(),
		// ... otros campos actualizados ...
	}
	payload, _ := json.Marshal(updatedMovie)
	req, err := http.NewRequest("PUT", "/updatemovie", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.UpdateMovie)

	handler.ServeHTTP(rr, req)

	// Assert para el caso de éxito
	assert.Equal(t, http.StatusAccepted, rr.Code)

	// Crear un payload mal formado
	badPayload := []byte(`{"id": "esto no es un número"}`)
	req2, err2 := http.NewRequest("PUT", "/updatemovie", bytes.NewBuffer(badPayload))
	if err2 != nil {
		t.Fatal(err2)
	}

	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(app.UpdateMovie)

	handler2.ServeHTTP(rr2, req2)

	// Assert para el caso de payload mal formado
	assert.Equal(t, http.StatusBadRequest, rr2.Code)
}
