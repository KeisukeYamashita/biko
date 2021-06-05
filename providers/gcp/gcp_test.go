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

package gcp

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testDataPath = "./testdata/gcloud"
)

func TestGetSDKConfig(t *testing.T) {
	defaultGoogleSDKConfigPath = testDataPath

	conf, err := getSDKConfig()
	want := &SDKConfig{
		&Core{
			Account: "biko.is.nice.tool@maybe.com",
			Project: "i-am-biko",
		},
		&Compute{
			Zone:   "asia-roppongi-c",
			Region: "asia-minatoku",
		},
		&Container{
			Cluster: "biko-cluster",
		},
	}
	if err != nil {
		t.Fatalf("getSDKConfig failed error:%v", err)
	}

	if !reflect.DeepEqual(conf, want) {
		t.Fatalf("getSDKConfig failed not deep equal got:%v, want:%v", conf, want)
	}
}

func TestConstructPageStateParam(t *testing.T) {
	tests := map[string]struct {
		namespace string
		want      string
	}{
		"single namespace": {
			namespace: "abc",
			want:      "(\"savedViews\":(\"n\":[\"abc\"]))",
		},
		"two namespaces": {
			namespace: "test1,test2",
			want:      "(\"savedViews\":(\"n\":[\"test1\",\"test2\"]))",
		},
		"three namespaces": {
			namespace: "test1,test2,test3",
			want:      "(\"savedViews\":(\"n\":[\"test1\",\"test2\",\"test3\"]))",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, constructPageStateParam(tt.namespace))
		})
	}
}
