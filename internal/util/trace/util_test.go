package trace

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/opentracing/opentracing-go"
	oplog "github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go/config"
)

func InitTracing() io.Closer {
	cfg := &config.Configuration{
		ServiceName: "test",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return closer
}

type simpleStruct struct {
	name  string
	value string
}

func TestTracing(t *testing.T) {
	//Already Init in each framework, this can be ignored in debug
	closer := InitTracing()
	defer closer.Close()

	// context normally can be propagated through func params
	ctx := context.Background()

	//start span
	//default use function name for operation name
	sp, ctx := StartSpanFromContext(ctx)
	sp.SetTag("tag1", "tag1")
	// use self-defined operation name for span
	// sp, ctx := StartSpanFromContextWithOperationName(ctx, "self-defined name")
	defer sp.Finish()

	ss := &simpleStruct{
		name:  "name",
		value: "value",
	}
	sp.LogFields(oplog.String("key", "value"), oplog.Object("key", ss))

	err := caller(ctx)

	if err != nil {
		LogError(sp, err) //LogError do something error log in trace and returns origin error.
	}

}

func caller(ctx context.Context) error {
	for i := 0; i < 2; i++ {
		// if span starts in a loop, defer is not allowed.
		// manually call span.Finish() if error occurs or one loop ends
		sp, _ := StartSpanFromContextWithOperationName(ctx, fmt.Sprintf("test:%d", i))
		sp.SetTag(fmt.Sprintf("tags:%d", i), fmt.Sprintf("tags:%d", i))

		var err error
		if i == 1 {
			err = errors.New("test")
		}

		if err != nil {
			sp.Finish()
			return LogError(sp, err)
		}

		sp.Finish()
	}
	return nil
}
