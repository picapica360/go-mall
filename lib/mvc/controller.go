package mvc

import (
	"fmt"
	"strconv"

	"go-mall/lib/errcode"
)

// Controller mvc controller.
type Controller struct{}

// OK success output data.
func (*Controller) OK(data interface{}) map[string]interface{} {
	return output(true, "0", nil, data)
}

// OKNull success output data, but the data is null.
func (*Controller) OKNull() map[string]interface{} {
	return output(true, "0", nil, nil)
}

// Bad bad output data.
func (*Controller) Bad(code string, err interface{}) map[string]interface{} {
	return output(false, code, err, nil)
}

// BadCode bad output data, use the errcode.
func (*Controller) BadCode(code errcode.ErrorCode, err interface{}) map[string]interface{} {
	return output(false, strconv.Itoa(int(code)), err, nil)
}

func output(success bool, code string, err, data interface{}) map[string]interface{} {
	var errmsg string
	if err != nil {
		errmsg = fmt.Sprint(err)
	}

	// TODO: think how keep the sort by map key ...
	return map[string]interface{}{
		"success": success,
		"code":    code,
		"errmsg":  errmsg,
		"data":    data,
	}
}
