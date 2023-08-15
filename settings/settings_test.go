package settings

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabase(t *testing.T) {
	testCases := []struct {
		Name             string
		DbSettings       Database
		ExpectedDatabase Database
		ExpectedError    error
	}{
		{
			Name: "Case success: without password",
			DbSettings: Database{
				User:     "test",
				Password: "",
				Host:     "test.com",
				Port:     8080,
				Name:     "test_db",
				Driver:   "test_driver",
			},
			ExpectedDatabase: Database{
				User:     "test",
				Password: "",
				Host:     "test.com",
				Port:     8080,
				Name:     "test_db",
				Driver:   "test_driver",
			},
			ExpectedError: nil,
		},
		{
			Name: "Case success: with password",
			DbSettings: Database{
				User:     "test",
				Password: "test_password",
				Host:     "test.com",
				Port:     8080,
				Name:     "test_db",
				Driver:   "test_driver",
			},
			ExpectedDatabase: Database{
				User:     "test",
				Password: "test_password",
				Host:     "test.com",
				Port:     8080,
				Name:     "test_db",
				Driver:   "test_driver",
			},
			ExpectedError: nil,
		},
		{
			Name: "Case failed: the DB_USER env var is missing",
			DbSettings: Database{
				User:     "",
				Password: "some arbitrary text",
				Host:     "some arbitrary text",
				Port:     8080,
				Name:     "some arbitrary text",
				Driver:   "some arbitrary text",
			},
			ExpectedDatabase: Database{},
			ExpectedError:    ErrEnvVarNotFound,
		},
		{
			Name: "Case failed: the DB_HOST env var is missing",
			DbSettings: Database{
				User:     "some arbitrary text",
				Password: "some arbitrary text",
				Host:     "",
				Port:     8080,
				Name:     "some arbitrary text",
				Driver:   "some arbitrary text",
			},
			ExpectedDatabase: Database{},
			ExpectedError:    ErrEnvVarNotFound,
		},
		{
			Name: "Case failed: the DB_PORT env var is missing",
			DbSettings: Database{
				User:     "some arbitrary text",
				Password: "some arbitrary text",
				Host:     "some arbitrary text",
				Port:     0,
				Name:     "some arbitrary text",
				Driver:   "some arbitrary text",
			},
			ExpectedDatabase: Database{},
			ExpectedError:    ErrEnvVarNotFound,
		},
		{
			Name: "Case failed: the DB_NAME env var is missing",
			DbSettings: Database{
				User:     "some arbitrary text",
				Password: "some arbitrary text",
				Host:     "some arbitrary text",
				Port:     8080,
				Name:     "",
				Driver:   "some arbitrary text",
			},
			ExpectedDatabase: Database{},
			ExpectedError:    ErrEnvVarNotFound,
		},
		{
			Name: "Case failed: the DB_DRIVER env var is missing",
			DbSettings: Database{
				User:     "some arbitrary text",
				Password: "some arbitrary text",
				Host:     "some arbitrary text",
				Port:     8080,
				Name:     "some arbitrary text",
				Driver:   "",
			},
			ExpectedDatabase: Database{},
			ExpectedError:    ErrEnvVarNotFound,
		},
	}

	for _, testCase := range testCases {
		os.Unsetenv("DB_USER")
		os.Unsetenv("DB_PASSWORD")
		os.Unsetenv("DB_HOST")
		os.Unsetenv("DB_PORT")
		os.Unsetenv("DB_NAME")
		os.Unsetenv("DB_DRIVER")

		if testCase.DbSettings.User != "" {
			os.Setenv("DB_USER", testCase.DbSettings.User)
		}
		if testCase.DbSettings.Password != "" {
			os.Setenv("DB_PASSWORD", testCase.DbSettings.Password)
		}
		if testCase.DbSettings.Host != "" {
			os.Setenv("DB_HOST", testCase.DbSettings.Host)
		}
		if testCase.DbSettings.Port != 0 {
			os.Setenv("DB_PORT", strconv.Itoa(int(testCase.DbSettings.Port)))
		}
		if testCase.DbSettings.Name != "" {
			os.Setenv("DB_NAME", testCase.DbSettings.Name)
		}
		if testCase.DbSettings.Driver != "" {
			os.Setenv("DB_DRIVER", testCase.DbSettings.Driver)
		}

		t.Run(testCase.Name, func(t *testing.T) {
			dbSettings, err := NewDatabase()

			assert.Equal(t, testCase.ExpectedDatabase, dbSettings)
			assert.Equal(t, testCase.ExpectedError, err)
		})
	}
}

func TestNewServer(t *testing.T) {
	testCases := []struct {
		Name           string
		ServerSettings Server
		ExpectedServer Server
		ExpectedError  error
	}{
		{
			Name: "Case success: both SERVER_PORT and SERVER_FRAMEWORK env vars are present",
			ServerSettings: Server{
				Port:      8080,
				Framework: "test",
			},
			ExpectedServer: Server{
				Port:      8080,
				Framework: "test",
			},
			ExpectedError: nil,
		},
		{
			Name: "Case failed: the SERVER_PORT env var is missing",
			ServerSettings: Server{
				Port:      0,
				Framework: "test",
			},
			ExpectedServer: Server{},
			ExpectedError:  ErrEnvVarNotFound,
		},
		{
			Name: "Case failed: the SERVER_FRAMEWORK env var is missing",
			ServerSettings: Server{
				Port:      8080,
				Framework: "",
			},
			ExpectedServer: Server{},
			ExpectedError:  ErrEnvVarNotFound,
		},
	}

	for _, testCase := range testCases {
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("SERVER_FRAMEWORK")

		if testCase.ServerSettings.Port != 0 {
			os.Setenv("SERVER_PORT", strconv.Itoa(int(testCase.ServerSettings.Port)))
		}
		if testCase.ServerSettings.Framework != "" {
			os.Setenv("SERVER_FRAMEWORK", testCase.ServerSettings.Framework)
		}

		t.Run(testCase.Name, func(t *testing.T) {
			serverSettings, err := NewServer()
			assert.Equal(t, testCase.ExpectedServer, serverSettings)
			assert.Equal(t, testCase.ExpectedError, err)
		})
	}
}
