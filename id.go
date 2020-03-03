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

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var NodeId uint32

func GenNodeId() uint32 {
	return uint32(rand.Int31n(1000000))
}

func InitNodeId(file string) (uint32, error) {
	if !IsFileExisted(file) {
		NodeId = GenNodeId()
		err := SaveJsonObject(file, NodeId)
		if err != nil {
			return 0, err
		}
		return NodeId, nil
	}
	err := GetJsonObject(file, &NodeId)
	if err != nil {
		return 0, err
	}
	return NodeId, nil
}
