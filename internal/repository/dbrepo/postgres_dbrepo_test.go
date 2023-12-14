package dbrepo

import (
	"backend/internal/models"
	"backend/internal/repository"
	"backend/mocks"
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	"reflect"
	"testing"
)

const connectionString = "user=postgres password=sander7l dbname=movies sslmode=disable"

func setupDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		t.Fatal("No se pudo conectar a la base de datos:", err)
	}
	return db
}

func TestPostgresDBRepo_AllGenres(t *testing.T) {

	db := setupDB(t)
	defer db.Close()

	repo := PostgresDBRepo{DB: db}

	type fields struct {
		DB *sql.DB
	}
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{
			name:    "Gender Count Correct",
			want:    13,
			wantErr: false,
		},
		{
			name:    "Gender count Error",
			want:    0,
			wantErr: true,
		},
		// Puedes añadir más casos de prueba según sea necesario
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.AllGenres()
			if err != nil {
				t.Errorf("AllGenres() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (len(got) != tt.want) && (tt.wantErr != true) {
				t.Errorf("AllGenres() got %v genres, want %v", len(got), tt.want)
			}
		})
	}
}

func TestPostgresDBRepo_AllMovies(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		genre []int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Movie
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PostgresDBRepo{
				DB: tt.fields.DB,
			}
			got, err := m.AllMovies(tt.args.genre...)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllMovies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllMovies() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresDBRepo_Connection(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	var tests []struct {
		name   string
		fields fields
		want   *sql.DB
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PostgresDBRepo{
				DB: tt.fields.DB,
			}
			if got := m.Connection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresDBRepo_DeleteMovie(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		id int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PostgresDBRepo{
				DB: tt.fields.DB,
			}
			if err := m.DeleteMovie(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteMovie() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresDBRepo_GetUserByEmail(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	type fields struct {
		DB *sql.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name:   "Caso exitoso",
			fields: fields{DB: db},
			args:   args{email: "admin@example.com"},
			want: &models.User{
				ID:        1,
				Email:     "admin@example.com",
				FirstName: "Admin"},
			wantErr: false,
		},
		{
			name:    "Correo electrónico no encontrado",
			fields:  fields{DB: db},
			args:    args{email: "noexiste@example.com"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Correo electrónico inválido",
			fields:  fields{DB: db},
			args:    args{email: "correo_invalido"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Base de datos no accesible",
			fields:  fields{DB: nil}, // Simulación de DB no disponible
			args:    args{email: "usuario@example.com"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Correo electrónico vacío",
			fields:  fields{DB: db},
			args:    args{email: ""},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PostgresDBRepo{
				DB: tt.fields.DB,
			}
			got, err := m.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && (got.ID != tt.want.ID || got.Email != tt.want.Email || got.FirstName != tt.want.FirstName) {
				t.Errorf("GetUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresDBRepo_GetUserByEmail_sample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDatabaseRepo(ctrl)

	type fields struct {
		DB repository.DatabaseRepo
	}
	type args struct {
		email string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *models.User
		wantErr   bool
		setupMock func() // Añade un campo para configurar el mock
	}{
		// ... tus casos de prueba
		{
			name:    "Base de datos no accesible",
			fields:  fields{DB: mockDB},
			args:    args{email: "admin@example.com"},
			want:    nil,
			wantErr: true,
			setupMock: func() {
				mockDB.EXPECT().GetUserByEmail("admin@example.com").Return(nil, sql.ErrConnDone)
			},
		},
		// ... otros casos de prueba
	}

	for _, tt := range tests {
		tt.setupMock() // Configura el mock según el caso de prueba
		t.Run(tt.name, func(t *testing.T) {
			m := tt.fields.DB // 'm' es ahora del tipo 'DatabaseRepo'
			got, err := m.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresDBRepo_GetUserByEmail2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDatabaseRepo(ctrl)

	type fields struct {
		DB repository.DatabaseRepo
	}
	type args struct {
		email string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *models.User
		wantErr   bool
		setupMock func(mockDB *mocks.MockDatabaseRepo)
	}{

		{
			name:    "Caso exitoso",
			fields:  fields{DB: mockDB},
			args:    args{email: "admin@example.com"},
			want:    &models.User{ID: 1, Email: "admin@example.com", FirstName: "Admin"},
			wantErr: false,
			setupMock: func(mockDB *mocks.MockDatabaseRepo) {
				mockDB.EXPECT().GetUserByEmail("admin@example.com").Return(&models.User{ID: 1, Email: "admin@example.com", FirstName: "Admin"}, nil)
			},
		},
		{
			name:    "Correo electrónico no encontrado",
			fields:  fields{DB: mockDB},
			args:    args{email: "noexiste@example.com"},
			want:    nil,
			wantErr: true,
			setupMock: func(mockDB *mocks.MockDatabaseRepo) {
				mockDB.EXPECT().GetUserByEmail("noexiste@example.com").Return(nil, errors.New("user not found"))
			},
		},
		{
			name:    "Correo electrónico inválido",
			fields:  fields{DB: mockDB},
			args:    args{email: "correo_invalido"},
			want:    nil,
			wantErr: true,
			setupMock: func(mockDB *mocks.MockDatabaseRepo) {
				mockDB.EXPECT().GetUserByEmail("correo_invalido").Return(nil, errors.New("invalid email format"))
			},
		},
		{
			name:    "Base de datos no accesible",
			fields:  fields{DB: mockDB},
			args:    args{email: "admin@example.com"},
			want:    nil,
			wantErr: true,
			setupMock: func(mockDB *mocks.MockDatabaseRepo) {
				mockDB.EXPECT().GetUserByEmail("admin@example.com").Return(nil, sql.ErrConnDone)
			},
		},
		{
			name:    "Correo electrónico vacío",
			fields:  fields{DB: mockDB},
			args:    args{email: ""},
			want:    nil,
			wantErr: true,
			setupMock: func(mockDB *mocks.MockDatabaseRepo) {
				mockDB.EXPECT().GetUserByEmail("").Return(nil, errors.New("email is empty"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock(mockDB)
			m := tt.fields.DB
			got, err := m.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresDBRepo_GetUserByID(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		id int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PostgresDBRepo{
				DB: tt.fields.DB,
			}
			got, err := m.GetUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresDBRepo_InsertMovie(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		movie models.Movie
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PostgresDBRepo{
				DB: tt.fields.DB,
			}
			got, err := m.InsertMovie(tt.args.movie)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InsertMovie() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresDBRepo_OneMovie(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		id int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *models.Movie
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PostgresDBRepo{
				DB: tt.fields.DB,
			}
			got, err := m.OneMovie(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("OneMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OneMovie() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresDBRepo_OneMovieForEdit(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		id int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *models.Movie
		want1   []*models.Genre
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PostgresDBRepo{
				DB: tt.fields.DB,
			}
			got, got1, err := m.OneMovieForEdit(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("OneMovieForEdit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OneMovieForEdit() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OneMovieForEdit() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPostgresDBRepo_UpdateMovie(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		movie models.Movie
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PostgresDBRepo{
				DB: tt.fields.DB,
			}
			if err := m.UpdateMovie(tt.args.movie); (err != nil) != tt.wantErr {
				t.Errorf("UpdateMovie() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresDBRepo_UpdateMovieGenres(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		id       int
		genreIDs []int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PostgresDBRepo{
				DB: tt.fields.DB,
			}
			if err := m.UpdateMovieGenres(tt.args.id, tt.args.genreIDs); (err != nil) != tt.wantErr {
				t.Errorf("UpdateMovieGenres() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
