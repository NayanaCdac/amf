// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

//+build !debug

package util

import (
	"github.com/free5gc/path_util"
)

var (
	AmfLogPath           = path_util.Free5gcPath("free5gc/amfsslkey.log")
	AmfPemPath           = path_util.Free5gcPath("free5gc/support/TLS/amf.pem")
	AmfKeyPath           = path_util.Free5gcPath("free5gc/support/TLS/amf.key")
	DefaultAmfConfigPath = path_util.Free5gcPath("free5gc/config/amfcfg.yaml")
)
