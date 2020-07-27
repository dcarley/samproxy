package sample

import (
	"testing"

	"github.com/honeycombio/samproxy/config"
	"github.com/honeycombio/samproxy/logger"
	"github.com/honeycombio/samproxy/metrics"
	"github.com/honeycombio/samproxy/types"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestDynamicEMAAddSampleRateKeyToTrace(t *testing.T) {
	const spanCount = 5

	config := &config.MockConfig{
		GetOtherConfigVal: `{"FieldList":["http.status_code"],"AddSampleRateKeyToTrace":true,"AddSampleRateKeyToTraceField":"meta.key"}`,
	}
	metrics := metrics.MockMetrics{}
	metrics.Start()
	sampler := &EMADynamicSampler{
		Config:  config,
		Logger:  &logger.NullLogger{},
		Metrics: &metrics,
	}

	trace := &types.Trace{}
	for i := 0; i < spanCount; i++ {
		trace.AddSpan(&types.Span{
			Event: types.Event{
				Data: map[string]interface{}{
					"http.status_code": "200",
				},
			},
		})
	}
	sampler.Start()
	sampler.GetSampleRate(trace)

	spans := trace.GetSpans()
	assert.Assert(t, is.Len(spans, spanCount))
	for _, span := range spans {
		assert.DeepEqual(t, span.Event.Data, map[string]interface{}{
			"http.status_code": "200",
			"meta.key":         "200•,",
		})
	}
}
