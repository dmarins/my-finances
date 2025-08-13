package env

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvWithDefaultAsStringWithEnvValue(t *testing.T) {
	testCases := []struct {
		envKey     string
		envValue   string
		defaultVal string
		expected   string
	}{
		{envKey: "Env", envValue: "prod", defaultVal: "", expected: "prod"},
		{envKey: Env, envValue: "dev", defaultVal: "", expected: DefaultEnv},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("'%s' should return '%s'", tC.envValue, tC.expected), func(t *testing.T) {
			err := os.Setenv(tC.envKey, tC.envValue)
			assert.Nil(t, err)
			val := GetEnvWithDefaultAsString(tC.envKey, "")
			assert.Equal(t, tC.envValue, val)
		})
	}
}

func TestGetEnvWithDefaultAsStringWithoutEnvValue(t *testing.T) {
	testCases := []struct {
		envKey     string
		defaultVal string
		expected   string
	}{
		{envKey: "Environment", defaultVal: "prod", expected: "prod"},
		{envKey: Env, defaultVal: "dev", expected: "dev"},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("'%s' should return '%s'", tC.defaultVal, tC.expected), func(t *testing.T) {
			val := GetEnvWithDefaultAsString(tC.envKey, tC.defaultVal)
			assert.Equal(t, tC.defaultVal, val)
		})
	}
}

func TestGetEnvWithDefaultAsIntWithEnvValue(t *testing.T) {
	testCases := []struct {
		envKey     string
		envValue   string
		defaultVal interface{}
		expected   int
	}{
		{envKey: "Env", envValue: "1", defaultVal: 0, expected: 1},
		{envKey: Env, envValue: "", defaultVal: "", expected: 0},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("'%s' should return '%d'", tC.envValue, tC.expected), func(t *testing.T) {
			err := os.Setenv(tC.envKey, tC.envValue)
			assert.Nil(t, err)
			val := GetEnvWithDefaultAsInt(tC.envKey, 0)
			value, _ := strconv.Atoi(tC.envValue)
			assert.Equal(t, value, val)
		})
	}
}

func TestGetEnvWithDefaultAsIntWithoutEnvValue(t *testing.T) {
	testCases := []struct {
		envKey     string
		defaultVal string
		expected   int
	}{
		{envKey: "Env", defaultVal: "1", expected: 1},
		{envKey: Env, defaultVal: "", expected: 0},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("'%s' should return '%d'", tC.defaultVal, tC.expected), func(t *testing.T) {
			err := os.Setenv(tC.envKey, tC.defaultVal)
			assert.Nil(t, err)
			val := GetEnvWithDefaultAsInt(tC.envKey, 0)
			value, _ := strconv.Atoi(tC.defaultVal)
			assert.Equal(t, value, val)
		})
	}
}
