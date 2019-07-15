// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package cisco

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "cisco", asset.ModuleFieldsPri, AssetCisco); err != nil {
		panic(err)
	}
}

// AssetCisco returns asset data.
// This is the base64 encoded gzipped contents of module/cisco.
func AssetCisco() string {
	return "eJy8lkFv2zwMhu/5Fbx8tzTfPYcBRYYCBda1QLdzoEmUzVUWPYpOmn8/SHEaN8hSF/WmQwAz8vs+pCTKV/CEuyVYSpZnAEoacAmr/tFhskKtEsclfJoBANyx6wKCZ4HaRBcoVvvpEFG3LE/gcEMWIXCVFjMATxhcWpaXryCaBo92eeiuxSVUwl3bR8645nFThMALN73jwSKPoc3QyiTzEjtndsFwaMrSe14/XsMNCW5NCIvB1FP/I0GDKZkK1+ReKe9RnnC3ZXn9zwUcgG81Dkh6bSCHUckTypHpDErqvKfnkRj4bJo274aEKRHH8Yz3JW5C7wfGKwr8l4HHgnInFtcUFcUbi1NU7rFowotmWVStEXzgLbAAbjDqRSyHSSmarD8t2+ej8IcAAyWdaKN9NQ0C+wJwbS2mBCuOKhzgCyUFrY1CY9TW6EBrSiPw+mXtEkoOTI2ZdfdclEpg79fXcRThcIX/GebA9D2sjWlbdOvDWWnPcJ4E3+wsKiamYBTdoXa3D2CcE0zpHSwti56hCRyrj/Jk6TEkr47qxKUZrtf76jOk+htFGpK9WSmtBY2uA24wTHM1ZT0oemUfNyZsjSD8Dz9YI2om9Z7sAu5jOQ0blN1V4O0c8s+JXMMOxSjOoaaqzv2vTM8PY9KyRrFi2U2R2arXeunIf87sJvfpw825IenSvJ9zmp8K/zRxDqj2Yj6WY0S738iTdPbvkX51w2+GkpYp18xFErJNu86mZyhSfbqdLzLcru4eyptvG1p2UxlmqcXsdwAAAP//GtXx0Q=="
}
