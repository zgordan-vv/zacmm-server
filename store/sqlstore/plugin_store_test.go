// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"testing"

	"github.com/zgordan-vv/zacmm-server/store/storetest"
)

func TestPluginStore(t *testing.T) {
	StoreTestWithSqlSupplier(t, storetest.TestPluginStore)
}
