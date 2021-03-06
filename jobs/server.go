// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package jobs

import (
	ejobs "github.com/zgordan-vv/zacmm-server/einterfaces/jobs"
	tjobs "github.com/zgordan-vv/zacmm-server/jobs/interfaces"
	"github.com/zgordan-vv/zacmm-server/model"
	"github.com/zgordan-vv/zacmm-server/services/configservice"
	"github.com/zgordan-vv/zacmm-server/store"
)

type JobServer struct {
	ConfigService configservice.ConfigService
	Store         store.Store
	Workers       *Workers
	Schedulers    *Schedulers

	DataRetentionJob        ejobs.DataRetentionJobInterface
	MessageExportJob        ejobs.MessageExportJobInterface
	ElasticsearchAggregator ejobs.ElasticsearchAggregatorInterface
	ElasticsearchIndexer    tjobs.IndexerJobInterface
	LdapSync                ejobs.LdapSyncInterface
	Migrations              tjobs.MigrationsJobInterface
	Plugins                 tjobs.PluginsJobInterface
	BleveIndexer            tjobs.IndexerJobInterface
	ExpiryNotify            tjobs.ExpiryNotifyJobInterface
	ProductNotices          tjobs.ProductNoticesJobInterface
	ActiveUsers             tjobs.ActiveUsersJobInterface
	Cloud                   ejobs.CloudJobInterface
}

func NewJobServer(configService configservice.ConfigService, store store.Store) *JobServer {
	return &JobServer{
		ConfigService: configService,
		Store:         store,
	}
}

func (srv *JobServer) Config() *model.Config {
	return srv.ConfigService.Config()
}

func (srv *JobServer) StartWorkers() {
	srv.Workers = srv.Workers.Start()
}

func (srv *JobServer) StartSchedulers() {
	srv.Schedulers = srv.Schedulers.Start()
}

func (srv *JobServer) StopWorkers() {
	if srv.Workers != nil {
		srv.Workers.Stop()
	}
}

func (srv *JobServer) StopSchedulers() {
	if srv.Schedulers != nil {
		srv.Schedulers.Stop()
	}
}
