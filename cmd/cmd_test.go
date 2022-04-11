package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	var err error

	_, err = initConfig("invalid.yml")
	assert.NotEqual(t, nil, err)

	_, err = initConfig("../tests/invalid.yml")
	assert.NotEqual(t, nil, err)

	_, err = initConfig("../tests/config.yml")
	assert.Equal(t, nil, err)
}
