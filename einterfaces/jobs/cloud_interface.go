// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package jobs

import "github.com/zgordan-vv/zacmm-server/model"

type CloudJobInterface interface {
	MakeWorker() model.Worker
	MakeScheduler() model.Scheduler
}
