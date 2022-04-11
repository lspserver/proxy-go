package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	var err error

	_, err = initConfig("invalid.yml")
	assert.NotEqual(t, nil, err)

	_, err = initConfig("../test/invalid.yml")
	assert.NotEqual(t, nil, err)

	_, err = initConfig("../test/config.yml")
	assert.Equal(t, nil, err)
}
