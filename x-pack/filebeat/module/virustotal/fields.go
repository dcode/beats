// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package virustotal

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "virustotal", asset.ModuleFieldsPri, AssetVirustotal); err != nil {
		panic(err)
	}
}

// AssetVirustotal returns asset data.
// This is the base64 encoded gzipped contents of module/virustotal.
func AssetVirustotal() string {
	return "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"
}
