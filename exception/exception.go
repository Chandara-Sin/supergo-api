package exception

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

type ErrorRes struct {
	Code    string `json:"code"`
	TraceId string `json:"trace_id"`
	Message string `json:"message"`
}

func GetTraceId() string {
	bi, _ := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(10)))),
	)
	return "SUP-" + fmt.Sprintf("%010d", bi)
}
