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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

const (
	ONT_ADDRESS = iota
	ONG_ADDRESS
	OEP4_ADDRESS
	UNKNOW_ADDRESS
)

type OneThreadExecLock struct {
	isWorking bool
	lock      sync.Mutex
}

func NewOneThreadExecLock() *OneThreadExecLock {
	return &OneThreadExecLock{}
}

func (this *OneThreadExecLock) TryLock() bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.isWorking {
		return false
	}
	this.isWorking = true
	return true
}

func (this *OneThreadExecLock) Release() {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.isWorking = false
}

func GetJsonObject(filePath string, jsonObject interface{}) error {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll %s error %s", filePath, err)
	}
	err = json.Unmarshal(data, jsonObject)
	if err != nil {
		return fmt.Errorf("json.Unmarshal %s error %s", data, err)
	}
	return nil
}

func SaveJsonObject(filePath string, jsonObject interface{}) error {
	data, err := json.Marshal(jsonObject)
	if err != nil {
		return fmt.Errorf("json.Marshal error:%s", err)
	}
	return ioutil.WriteFile(filePath, data, 0666)
}

func IsMonitorContract(contract string) bool {
	for _, item := range DefConfig.Contracts {
		if contract == item {
			return true
		}
	}
	return false
}

func TypeOfContract(contract string) uint32 {
	if contract == "0100000000000000000000000000000000000000" {
		return ONT_ADDRESS
	} else if contract == "0200000000000000000000000000000000000000" {
		return ONG_ADDRESS
	}

	if IsMonitorContract(contract) {
		return OEP4_ADDRESS
	} else {
		return UNKNOW_ADDRESS
	}
}

func IsFileExisted(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
