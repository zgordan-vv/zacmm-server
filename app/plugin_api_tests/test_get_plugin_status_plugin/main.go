// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package main

import (
	"github.com/zgordan-vv/zacmm-server/app/plugin_api_tests"
	"github.com/zgordan-vv/zacmm-server/model"
	"github.com/zgordan-vv/zacmm-server/plugin"
)

type MyPlugin struct {
	plugin.MattermostPlugin
	configuration plugin_api_tests.BasicConfig
}

func (p *MyPlugin) OnConfigurationChange() error {
	if err := p.API.LoadPluginConfiguration(&p.configuration); err != nil {
		return err
	}
	return nil
}

func (p *MyPlugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	status, err := p.API.GetPluginStatus("test_get_plugin_status_plugin")
	if err != nil {
		return nil, err.Error()
	}

	if status.State != model.PluginStateRunning {
		return nil, "State is not running"
	}

	return nil, "OK"
}

func main() {
	plugin.ClientMain(&MyPlugin{})
}
