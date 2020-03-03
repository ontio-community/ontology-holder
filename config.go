/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

var DefConfig = &Config{}

type Config struct {
	MySqlAddress                    string
	MySqlUserName                   string
	MySqlPassword                   string
	MySqlDBName                     string
	MySqlMaxIdleConnSize            uint32
	MySqlMaxConnSize                uint32
	MySqlMaxOpenConnSize            uint32
	MySqlConnMaxLifetime            uint32
	MySqlHeartbeatTimeoutTime       uint32
	MySqlHeartbeatUpdateInterval    uint32
	UpdateHolderCountInterval       uint32
	UpdateSyncedBlockHeightInterval uint32
	OntologyRpcAddress              string
	BlockHeight                     uint32
	HttpServerPort                  uint32
	DBBatchSize                     uint32
	DBBatchTime                     uint32
	MaxQueryPageSize                uint32
	Contracts                       []string
}

func (this *Config) GetHeartbeatUpdateInterval() uint32 {
	if this.MySqlHeartbeatUpdateInterval == 0 {
		return DEFAULT_HAARTBEAT_UPDATE_INTERVAL
	}
	return this.MySqlHeartbeatUpdateInterval
}

func (this *Config) GetHeartbeatTimeoutTime() uint32 {
	if this.MySqlHeartbeatTimeoutTime == 0 {
		return DEFAULT_HEARTBEAT_TIMEOUT_TIME
	}
	return this.MySqlHeartbeatTimeoutTime
}

func (this *Config) GetHolderCountUpdateInterval() uint32 {
	if this.UpdateHolderCountInterval == 0 {
		return DEFAULT_UPDATE_ASSET_HOLDER_COUNT_INTERVAL
	}
	return this.UpdateHolderCountInterval
}

func (this *Config) GetSyncedBlockHeightInterval() uint32 {
	if this.UpdateSyncedBlockHeightInterval == 0 {
		return DEFAULT_UPDATE_SYNCED_BLOCK_HEIGHT_INTERVAL
	}
	return this.UpdateSyncedBlockHeightInterval
}
