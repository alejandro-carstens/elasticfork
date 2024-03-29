// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elasticfork

import (
	"net/url"
	"testing"
)

func TestClusterStatsURLs(t *testing.T) {
	fFlag := false
	tFlag := true

	tests := []struct {
		Service        *ClusterStatsService
		ExpectedPath   string
		ExpectedParams url.Values
	}{
		{
			Service: &ClusterStatsService{
				nodeId: []string{},
			},
			ExpectedPath: "/_cluster/stats",
		},
		{
			Service: &ClusterStatsService{
				nodeId: []string{"node1"},
			},
			ExpectedPath: "/_cluster/stats/nodes/node1",
		},
		{
			Service: &ClusterStatsService{
				nodeId: []string{"node1", "node2"},
			},
			ExpectedPath: "/_cluster/stats/nodes/node1%2Cnode2",
		},
		{
			Service: &ClusterStatsService{
				nodeId:       []string{},
				flatSettings: &tFlag,
			},
			ExpectedPath:   "/_cluster/stats",
			ExpectedParams: url.Values{"flat_settings": []string{"true"}},
		},
		{
			Service: &ClusterStatsService{
				nodeId:       []string{"node1"},
				flatSettings: &fFlag,
			},
			ExpectedPath:   "/_cluster/stats/nodes/node1",
			ExpectedParams: url.Values{"flat_settings": []string{"false"}},
		},
	}

	for _, test := range tests {
		gotPath, gotParams, err := test.Service.buildURL()
		if err != nil {
			t.Fatalf("expected no error; got: %v", err)
		}
		if gotPath != test.ExpectedPath {
			t.Errorf("expected URL path = %q; got: %q", test.ExpectedPath, gotPath)
		}
		if gotParams.Encode() != test.ExpectedParams.Encode() {
			t.Errorf("expected URL params = %v; got: %v", test.ExpectedParams, gotParams)
		}
	}
}
