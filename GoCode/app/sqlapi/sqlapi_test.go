package sqlapi

import (
	"context"
	"errors"
	"log"
	"projectbotticket/types/apitypes"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	viper.SetConfigName("keys")
	viper.SetConfigType("json")
	viper.AddConfigPath("../../config/")

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Println(err)
	}
}

func connectDB_test(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		log.Println("sqlx.Open failed with an error: ", err.Error())
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Println("DB.Ping failed with an error: ", err.Error())
		return nil, err
	}

	return db, err
}

func TestGetTestValue(t *testing.T) {

	db, err := connectDB_test(viper.GetString("ConnectPostgres"))

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var api = NewAPI(db)

	testCases := []struct {
		name    string
		prepare func() int
		isValid int
	}{
		{
			name: "true",
			prepare: func() int {
				return api.GetTestValue(2)
			},
			isValid: 4,
		},
		{
			name: "false",
			prepare: func() int {
				return api.GetTestValue(4)
			},
			isValid: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.prepare(), tc.isValid)
		})
	}
}

func TestSetNewUser(t *testing.T) {
	ctx := context.Background()

	db, err := connectDB_test(viper.GetString("ConnectPostgres"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRow := apitypes.UserRow{
		UserId:   "testUserId",
		NameUser: "testUser",
		ChatId:   "testChatId",
	}

	testCases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "add new user",
			wantErr: false,
		},
		{
			name:    "add new user",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			api := NewAPI(db)
			if err := api.SetNewUser(ctx, userRow); (err != nil) && !tc.wantErr {
				t.Errorf("SetNewUser end with err: %v, want err: %v", err, tc.wantErr)
			}
		})
	}
}

func TestSetNewUser2(t *testing.T) {
	ctx := context.Background()
	user := apitypes.UserRow{
		UserId:   "some_user_id",
		NameUser: "some_user_name",
		ChatId:   "some_chat_id",
	}

	const expectedQuery = `INSERT INTO prj_user\(userid, nameuser, chatid\)
	VALUES (.+)
	ON CONFLICT DO NOTHING
	;`

	tests := []struct {
		name    string
		prepare func(mock sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			"1. error on add new user",
			func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(expectedQuery).
					WithArgs(user.UserId, user.NameUser, user.ChatId).
					WillReturnError(errors.New("some error"))
			},
			true,
		},
		{
			"2. success add new user",
			func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(expectedQuery).
					WithArgs(user.UserId, user.NameUser, user.ChatId).
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseDB, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			db := sqlx.NewDb(baseDB, "postgres")
			defer db.Close()

			tt.prepare(mock)

			api := NewAPI(db)
			if err := api.SetNewUser(ctx, user); (err != nil) != tt.wantErr {
				t.Errorf("SetNewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("SetNewUser() there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestGetUserById(t *testing.T) {

	columns := []string{"userid", "nameuser", "chatid"}

	const expectedQuery = `SELECT (.+) FROM prj_user WHERE userid = (.+);`

	tests := []struct {
		name    string
		prepare func(mock sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			"1. Get user",
			func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(expectedQuery).
					WithArgs("UserId").
					WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1,1,1"))
			},
			false,
		},
		{
			"2. Get user nil",
			func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(expectedQuery).
					WithArgs("UserId").
					WillReturnRows(sqlmock.NewRows(columns).FromCSVString("0,0,0"))
			},
			false,
		},
		{
			"3. error on SELECT",
			func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(expectedQuery).
					WithArgs("UserId").
					WillReturnRows(sqlmock.NewRows(columns).FromCSVString("0,1"))
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseDb, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			db := sqlx.NewDb(baseDb, "postgres")
			defer db.Close()

			tt.prepare(mock)

			api := NewAPI(db)

			_, err = api.GetUserByID("UserId")

			if (err != nil) != tt.wantErr {
				t.Errorf("SetNewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("SetNewUser() there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestGetUserByName(t *testing.T) {
	columns := []string{"userid", "nameuser", "chatid"}

	const expectedQuery = `SELECT (.+) FROM prj_user WHERE nameuser = (.+);`

	tests := []struct {
		name    string
		prepare func(mock sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			"1. Get user",
			func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(expectedQuery).
					WithArgs("UserName").
					WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1,1,1"))
			},
			false,
		},
		{
			"2. Get user nil",
			func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(expectedQuery).
					WithArgs("UserName").
					WillReturnRows(sqlmock.NewRows(columns).FromCSVString("0,0,0"))
			},
			false,
		},
		{
			"3. error on SELECT",
			func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(expectedQuery).
					WithArgs("UserName").
					WillReturnRows(sqlmock.NewRows(columns).FromCSVString("0,1"))
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseDb, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			db := sqlx.NewDb(baseDb, "postgres")
			defer db.Close()

			tt.prepare(mock)

			api := NewAPI(db)

			_, err = api.GetUserByName("UserName")

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("GetUserByName() there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestSetNewExecutor(t *testing.T) {

	ctx := context.Background()
	executor := apitypes.ExecutorRow{
		ExecutorId:      "some_executor_id",
		ExecutorName:    "some_executor_name",
		ExecutorPasword: "some_executor_pasword",
	}

	const expectedQuery = `INSERT INTO prj_executor\(executorid, executorname, executorpasword\)
	VALUES (.+)
	ON CONFLICT DO NOTHING
	;`

	tests := []struct {
		name    string
		prepare func(mock sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			"1. error on add new executor",
			func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(expectedQuery).
					WithArgs(executor.ExecutorId, executor.ExecutorName, executor.ExecutorPasword).
					WillReturnError(errors.New("some error"))
			},
			true,
		},
		{
			"2. success add new executor",
			func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(expectedQuery).
					WithArgs(executor.ExecutorId, executor.ExecutorName, executor.ExecutorPasword).
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseDb, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			db := sqlx.NewDb(baseDb, "postgres")
			defer db.Close()

			tt.prepare(mock)

			api := NewAPI(db)

			err = api.SetNewExecutor(ctx, executor)

			if (err != nil) != tt.wantErr {
				t.Errorf("SetNewExecutor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("SetNewExecutor() there were unfulfilled expectations: %s", err)
			}
		})
	}
}
