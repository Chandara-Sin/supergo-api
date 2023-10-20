package exception

import "go.uber.org/zap"

type ErrorRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func LogError(log *zap.Logger, errRes ErrorRes) {
	log.Error("error",
		zap.String("code", errRes.Code),
		zap.String("message", errRes.Message),
	)
}
