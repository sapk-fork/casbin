// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stringadapter

import (
	"bytes"
	"errors"
	"strings"

	"github.com/casbin/casbin/v3/model"
	"github.com/casbin/casbin/v3/persist"
	"github.com/casbin/casbin/v3/util"
)

// Adapter is the string adapter for Casbin.
// It can load policy from string or save policy to string.
type Adapter struct {
	lines string
}

// NewAdapter is the constructor for Adapter.
func NewAdapter(lines string) *Adapter {
	return &Adapter{lines: lines}
}

// LoadPolicy loads all policy rules from the storage.
func (a *Adapter) LoadPolicy(model model.Model) error {
	if a.lines == "" {
		return errors.New("invalid lines, lines cannot be empty")
	}

	rows := strings.Split(a.lines, "\n")
	for _, line := range rows {
		persist.LoadPolicyLine(strings.TrimSpace(line), model)
	}

	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *Adapter) SavePolicy(model model.Model) error {
	var tmp bytes.Buffer
	ptypes := model.GetPtypes("p")
	for _, ptype := range ptypes {
		for _, rule := range model.GetPolicy("p", ptype) {
			tmp.WriteString(ptype + ", ")
			tmp.WriteString(util.ArrayToString(rule))
			tmp.WriteString("\n")
		}
	}

	ptypes = model.GetPtypes("g")
	for _, ptype := range ptypes {
		for _, rule := range model.GetPolicy("g", ptype) {
			tmp.WriteString(ptype + ", ")
			tmp.WriteString(util.ArrayToString(rule))
			tmp.WriteString("\n")
		}
	}

	a.lines = strings.TrimRight(tmp.String(), "\n")

	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

// AddPolicies adds policy rules to the storage.
func (a *Adapter) AddPolicies(sec string, ptype string, rules [][]string) error {
	return errors.New("not implemented")
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

// RemovePolicies removes policy rules from the storage.
func (a *Adapter) RemovePolicies(sec string, ptype string, rules [][]string) error {
	return errors.New("not implemented")
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errors.New("not implemented")
}
