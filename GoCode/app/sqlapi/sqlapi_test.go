package sqlapi

import (
	"log"
	"projectbotticket/types/apitypes"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func runViper_test() {
	viper.SetConfigName("keys")
	viper.SetConfigType("json")
	viper.AddConfigPath("../../config/")

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Fatal(err)
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
	runViper_test()

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
	runViper_test()

	db, err := connectDB_test(viper.GetString("ConnectPostgres"))

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var api = NewAPI(db)

	userRow := apitypes.UserRow{
		UserId:   "testUserId",
		NameUser: "testUser",
		ChatId:   "testChatId",
	}
	err = errors.New("the user exists")

	testCases := []struct {
		name    string
		prepare func() (*apitypes.UserRow, error)
		wantErr bool
	}{
		{
			name: "add new user",
			prepare: func() (*apitypes.UserRow, error) {

				return api.SetNewUser("testUser", "testUserId", "testChatId")
			},
			wantErr: false,
		},
		{
			name: "add new user",
			prepare: func() (*apitypes.UserRow, error) {
				return api.SetNewUser("testUser", "testUserId", "testChatId")
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {

				u, e := tc.prepare()

				assert.Error(t, e)
				assert.Nil(t, u)
				assert.Equal(t, e.Error(), err.Error())

			} else {

				u, e := tc.prepare()

				assert.NoError(t, e)
				assert.NotNil(t, u)
				assert.Equal(t, u.ChatId, userRow.ChatId)
				assert.Equal(t, u.UserId, userRow.UserId)
				assert.Equal(t, u.NameUser, userRow.NameUser)
			}

		})
	}
}
