package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValid(t *testing.T) {
	// create test data
	validPlatforms := []Platform{
		{
			Name:   "common",
			Code:   "000",
			Prefix: "",
			Dir:    "",
			Modules: []Module{
				{
					Name:   "service system 服务系统",
					Prefix: "Service",
					Code:   "000",
					SpecificErrors: []SpecificError{
						{
							Suffix: "InternalError",
							Code:   "000",
							Desc:   "Internal error 通用服务内部错误",
						},
						{
							Suffix: "Unavailable",
							Code:   "001",
							Desc:   "Service unavailable 服务不可用",
						},
					},
				},
				{
					Name:   "request error 请求错误",
					Prefix: "Request",
					Code:   "001",
					SpecificErrors: []SpecificError{
						{
							Suffix: "Fail",
							Code:   "000",
							Desc:   "Common request fail 通用请求错误",
						},
						{
							Suffix: "SvcNotFound",
							Code:   "001",
							Desc:   "Service not found 通用请求错误",
						},
					},
				},
			},
		},
	}

	t.Run("Check valid platforms", func(t *testing.T) {
		err := CheckValid(validPlatforms)
		assert.NoError(t, err)
	})

}
