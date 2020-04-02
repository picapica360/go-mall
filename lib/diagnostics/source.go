package diagnostics

import (
	"go-mall/lib/logs"
)

func innerRegister() {
	RegisterFunc(Event_Mail_Sendmail_Failure, func(obj interface{}, err error) {
		msg := struct {
			value interface{}
			err   error
		}{
			value: obj,
			err:   err,
		}

		logs.Errorf("[diagnostics] send mail failure: message %v", msg)
	})
}
