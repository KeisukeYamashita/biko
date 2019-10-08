// Copyright 2019 The Biko Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alias

import (
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	// This should be same in the testdata/config.toml
	wantConf := &TomlConfig{
		GCP: map[string]interface{}{
			"alias": map[string]interface{}{
				"myalias": "biko",
			},
		},
	}

	testCases := map[string]struct {
		path       string
		wantConf   *TomlConfig
		wantResult bool
	}{
		"ok": {
			path:       "./testdata/config.toml",
			wantConf:   wantConf,
			wantResult: true,
		},
		"bad path": {
			path:       "bad/path",
			wantConf:   nil,
			wantResult: false,
		},
	}

	for n, tc := range testCases {
		defaultConfigPath = tc.path
		conf, err := GetConfig()
		if err != nil {
			if tc.wantResult {
				t.Fatalf("GetConfig fail testCase => %s, error:%v", n, err)
				continue
			}
			continue
		}

		if conf == nil {
			t.Fatalf("GetConfig fail with nil conf testCase => %s", n)
		}

		if reflect.DeepEqual(conf, wantConf) {
			t.Fatalf("GetConfig fail not get config, want:%v, got:%v", wantConf, conf)
		}
	}
}
