package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingParamError(t *testing.T) {
	err := MissingParamError("param_name")
	assert.Equal(t, "ServiceError: request_param_missing", err.Error())
	assert.Equal(t, "request_param_missing", err.GetCode())
	assert.Equal(t, 400, err.GetResponseStatus())
	assert.Nil(t, err.UnWrap())
}
