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
	"errors"
	"strings"

	"github.com/casbin/casbin/v3/model"
	"github.com/casbin/casbin/v3/persist"
)

// String is the string adapter for Casbin.
// It can load policy from string.
type String string

// LoadPolicy loads all policy rules from the string.
func (a String) LoadPolicy(model model.Model) error {
	rows := strings.Split(string(a), "\n")
	for _, line := range rows {
		persist.LoadPolicyLine(strings.TrimSpace(line), model)
	}

	return nil
}

// SavePolicy is not supported by this adaptater.
func (a String) SavePolicy(model model.Model) error {
	return errors.New("not implemented")
}

// AddPolicy is not supported by this adaptater.
func (a String) AddPolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

// AddPolicies is not supported by this adaptater.
func (a String) AddPolicies(sec string, ptype string, rules [][]string) error {
	return errors.New("not implemented")
}

// RemovePolicy is not supported by this adaptater.
func (a String) RemovePolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

// RemovePolicies is not supported by this adaptater.
func (a String) RemovePolicies(sec string, ptype string, rules [][]string) error {
	return errors.New("not implemented")
}

// RemoveFilteredPolicy is not supported by this adaptater.
func (a String) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errors.New("not implemented")
}
