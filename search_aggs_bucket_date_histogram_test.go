// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elasticfork

import (
	"encoding/json"
	"testing"
)

func TestDateHistogramAggregation(t *testing.T) {
	agg := NewDateHistogramAggregation().
		Field("date").
		Interval("month").
		Format("yyyy-MM").
		TimeZone("UTC").
		Offset("+6h")
	src, err := agg.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"date_histogram":{"field":"date","format":"yyyy-MM","interval":"month","offset":"+6h","time_zone":"UTC"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestDateHistogramAggregationWithKeyedResponse(t *testing.T) {
	agg := NewDateHistogramAggregation().Field("date").Interval("year").Missing("1900").Keyed(true)
	src, err := agg.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"date_histogram":{"field":"date","interval":"year","keyed":true,"missing":"1900"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestDateHistogramAggregationWithMissing(t *testing.T) {
	agg := NewDateHistogramAggregation().Field("date").Interval("year").Missing("1900")
	src, err := agg.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"date_histogram":{"field":"date","interval":"year","missing":"1900"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
