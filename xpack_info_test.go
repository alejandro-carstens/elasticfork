// Copyright 2012-2018 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elasticfork

import (
	"testing"
)

func TestXPackInfoBuildURL(t *testing.T) {
	client := setupTestClientForXpackSecurity(t)

	tests := []struct {
		Expected  string
		ExpectErr bool
	}{
		{
			"/_xpack",
			false,
		},
	}

	for i, test := range tests {
		builder := client.XPackInfo()
		err := builder.Validate()
		if err != nil {
			if !test.ExpectErr {
				t.Errorf("case #%d: %v", i+1, err)
				continue
			}
		} else {
			// err == nil
			if test.ExpectErr {
				t.Errorf("case #%d: expected error", i+1)
				continue
			}
			path, _, _ := builder.buildURL()
			if path != test.Expected {
				t.Errorf("case #%d: expected %q; got: %q", i+1, test.Expected, path)
			}
		}
	}
}
