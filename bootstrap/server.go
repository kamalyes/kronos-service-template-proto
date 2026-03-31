/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-10-29 21:01:49
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-10-29 21:01:49
 * @FilePath: \kronos-service-template-proto\bootstrap\server.go
 * @Description: 网关实现 | Gateway implementation
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package main

import (
	"fmt"

	goconfig "github.com/kamalyes/go-config"
	gateway "github.com/kamalyes/go-rpc-gateway"
)

// Gateway 文档网关 | Documentation gateway
type Gateway struct {
	gateway *gateway.Gateway
}

// NewGateway 创建新的网关 | Create a new gateway
func NewGateway() *Gateway {
	return &Gateway{}
}

// Start 启动网关 | Start the gateway
// 配置从 resources/ 下按 <prefix>-<env>.yaml 加载，prefix 由 WithPrefix 指定。
// [EN] Config is loaded from resources/<prefix>-<env>.yaml, prefix set by WithPrefix
func (g *Gateway) Start() error {
	gw, err := gateway.NewGateway().
		WithSearchPath("resources").
		WithEnvironment(goconfig.GetEnvironment()).
		WithPrefix("gateway-kronos").
		WithHotReload(nil).
		Build()
	if err != nil {
		return fmt.Errorf("创建网关失败 | failed to create gateway: %w", err)
	}
	g.gateway = gw

	return g.gateway.Run()
}
