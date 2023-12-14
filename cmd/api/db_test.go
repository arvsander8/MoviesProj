package main

import (
	"backend/internal/repository"
	"database/sql"
	"reflect"
	"testing"
)

func Test_application_ConnectToDB(t *testing.T) {
	type fields struct {
		DSN          string
		Domain       string
		DB           repository.DatabaseRepo
		auth         Auth
		JWTSecret    string
		JWTIssuer    string
		JWTAudience  string
		CookieDomain string
		APIKey       string
	}
	var tests []struct {
		name    string
		fields  fields
		want    *sql.DB
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				DSN:          tt.fields.DSN,
				Domain:       tt.fields.Domain,
				DB:           tt.fields.DB,
				auth:         tt.fields.auth,
				JWTSecret:    tt.fields.JWTSecret,
				JWTIssuer:    tt.fields.JWTIssuer,
				JWTAudience:  tt.fields.JWTAudience,
				CookieDomain: tt.fields.CookieDomain,
				APIKey:       tt.fields.APIKey,
			}
			got, err := app.ConnectToDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectToDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectToDB() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_openDB(t *testing.T) {
	type args struct {
		dsn string
	}
	var tests []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := openDB(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("openDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("openDB() got = %v, want %v", got, tt.want)
			}
		})
	}
}
