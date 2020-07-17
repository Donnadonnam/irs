// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package service

import (
	"github.com/moov-io/identity/pkg/database"
	tmw "github.com/moov-io/tumbler/pkg/middleware"
)

type GlobalConfig struct {
	IRS Config
}

// Config defines all the configuration for the app
type Config struct {
	Servers  ServerConfig
	Database database.DatabaseConfig
	Gateway  tmw.TumblerConfig
}

// ServerConfig - Groups all the http configs for the servers and ports that get opened.
type ServerConfig struct {
	Public HTTPConfig
	Admin  HTTPConfig
}

// HTTPConfig configuration for running an http server
type HTTPConfig struct {
	Bind BindAddress
}

// BindAddress specifies where the http server should bind to.
type BindAddress struct {
	Address string
}
