// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package configmanager

import (
	"context"

	"github.com/jaegertracing/jaeger/pkg/metrics"
	"github.com/jaegertracing/jaeger/proto-gen/api_v2"
	"github.com/jaegertracing/jaeger/thrift-gen/baggage"
)

// configManagerMetrics holds metrics related to ClientConfigManager
type configManagerMetrics struct {
	// Number of successful sampling rate responses from collector
	SamplingSuccess metrics.Counter `metric:"collector-proxy" tags:"result=ok,endpoint=sampling"`

	// Number of failed sampling rate responses from collector
	SamplingFailures metrics.Counter `metric:"collector-proxy" tags:"result=err,endpoint=sampling"`

	// Number of successful baggage restriction responses from collector
	BaggageSuccess metrics.Counter `metric:"collector-proxy" tags:"result=ok,endpoint=baggage"`

	// Number of failed baggage restriction responses from collector
	BaggageFailures metrics.Counter `metric:"collector-proxy" tags:"result=err,endpoint=baggage"`
}

// ManagerWithMetrics is manager with metrics integration.
type ManagerWithMetrics struct {
	wrapped ClientConfigManager
	metrics configManagerMetrics
}

// WrapWithMetrics wraps ClientConfigManager and creates metrics for its invocations.
func WrapWithMetrics(manager ClientConfigManager, mFactory metrics.Factory) *ManagerWithMetrics {
	m := configManagerMetrics{}
	metrics.Init(&m, mFactory, nil)
	return &ManagerWithMetrics{wrapped: manager, metrics: m}
}

// GetSamplingStrategy returns sampling strategy from server.
func (m *ManagerWithMetrics) GetSamplingStrategy(ctx context.Context, serviceName string) (*api_v2.SamplingStrategyResponse, error) {
	r, err := m.wrapped.GetSamplingStrategy(ctx, serviceName)
	if err != nil {
		m.metrics.SamplingFailures.Inc(1)
	} else {
		m.metrics.SamplingSuccess.Inc(1)
	}
	return r, err
}

// GetBaggageRestrictions returns baggage restrictions from server.
func (m *ManagerWithMetrics) GetBaggageRestrictions(ctx context.Context, serviceName string) ([]*baggage.BaggageRestriction, error) {
	r, err := m.wrapped.GetBaggageRestrictions(ctx, serviceName)
	if err != nil {
		m.metrics.BaggageFailures.Inc(1)
	} else {
		m.metrics.BaggageSuccess.Inc(1)
	}
	return r, err
}
