package trace

import (
	"go-mall/lib/diagnostics"
	"go-mall/lib/log"
)

func init() {
	traceMail()
}

func traceMail() {
	diagnostics.RegisterFunc("trace_lib_sendmail_error", func(v interface{}, err error) {
		obj := struct {
			Value interface{}
			Error error
		}{
			Value: v,
			Error: err,
		}
		log.Errorf("[diagnostics] send mail error -- %v", obj)
	})
}
