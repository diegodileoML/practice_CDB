package test

import (
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/base/tracing"
)

type (
	FakeTracing struct{}
	fakeSegment struct{}
)

var segment fakeSegment

func (f fakeSegment) End() {
}

func (f FakeTracing) SimpleMetric(name string, value float64, tags ...tracing.Tag) {
}

func (f FakeTracing) StartExternalSegment(url string) tracing.Segment {
	return segment
}

func (f FakeTracing) StartSegment(name string) tracing.Segment {
	return segment
}

func (f FakeTracing) StartDatastoreSegment(db tracing.DataSource, summary string, operation tracing.DBOperation, query string) tracing.Segment {
	return segment
}