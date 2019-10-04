package gcp

import (
	"reflect"
	"testing"
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
