package context

import (
	"context"
	"fmt"
	"github.com/balireddypraveen/allen/internal/common/configs"
	"github.com/balireddypraveen/allen/internal/common/constants"
	"runtime/debug"

	"github.com/balireddypraveen/allen/internal/pkg/newrelic_setup"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ReqCtx struct {
	Url           string
	Method        string
	ReqId         string
	SpanReqId     string
	AmazonTraceId string
	NewRelicTxn   *newrelic.Transaction
	NrTraceId     string
	Log           *logrus.Logger
	context.Context
}

func GetRequestContext(ctx context.Context) ReqCtx {
	switch ctx.(type) {
	case *gin.Context:
		ginCtx := ctx.(*gin.Context)

		log, ok := ctx.Value(constants.ContextLogger).(logrus.Logger)

		nrTxn, ok := ctx.Value(constants.ContextNrTxn).(*newrelic.Transaction)
		if !ok {
			nrTxn = newrelic.FromContext(ctx)
		}

		context_ := &ReqCtx{
			Log:         &log,
			NewRelicTxn: nrTxn,
			Context:     ginCtx.Request.Context(),
			ReqId:       ginCtx.Request.Header.Get(string(constants.HeaderTraceReqId)),
		}
		commonFields, ok := ginCtx.Value(constants.CommonFieldsKey).(map[string]string)
		if ok {
			context_.Url = commonFields[constants.RequestPath]
			context_.Method = commonFields[constants.RequestMethod]
			context_.AmazonTraceId = commonFields[string(constants.HeaderTraceAmznID)]
			context_.ReqId = commonFields[string(constants.HeaderTraceReqId)]
			context_.SpanReqId = commonFields[string(constants.HeaderTraceSpanID)]
			context_.NrTraceId = commonFields[string(constants.HeaderNrTraceId)]
		}
		return *context_
	default:
		rCtx := &ReqCtx{
			Log:         &logrus.Logger{},
			NewRelicTxn: newrelic.FromContext(ctx),
			Context:     ctx,
		}
		return *rCtx
	}
}

func GetRCtxNonWebWithNRTxn(txnName string, ctx context.Context) ReqCtx {
	rCtx := ReqCtx{
		ReqId:         uuid.New().String(),
		AmazonTraceId: uuid.New().String(),
		SpanReqId:     uuid.New().String(),
		NewRelicTxn:   GetNonWebNewRelicTxn(txnName),
		Log:           &logrus.Logger{},
		Context:       context.Background(),
	}

	if rCtx.NewRelicTxn != nil {
		rCtx.NewRelicTxn.SetName(txnName)
		rCtx.NrTraceId = rCtx.NewRelicTxn.GetTraceMetadata().TraceID
	}
	return rCtx
}

func GetRCtxNonWeb(ctx context.Context) ReqCtx {
	rCtx := ReqCtx{
		ReqId:         uuid.New().String(),
		AmazonTraceId: uuid.New().String(),
		SpanReqId:     uuid.New().String(),
		Log:           &logrus.Logger{},
		Context:       ctx,
	}
	return rCtx
}

func GetNonWebNewRelicTxn(txnName string) *newrelic.Transaction {
	app := newrelic_setup.GetNewRelicApp(viper.GetString(configs.VKEYS_NEWRELIC_APP_NAME))
	return app.StartTransaction(txnName)
}

func GoRoutinePanicHandler(ctx context.Context) {
	if r := recover(); r != nil {
		fmt.Println("goroutine paniqued, recovering with panic handler:  ", r)
		debug.PrintStack()
	}
}
