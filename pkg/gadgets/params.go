// Copyright 2023 The Inspektor Gadget authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gadgets

import (
	"fmt"
	"strings"

	"github.com/inspektor-gadget/inspektor-gadget/pkg/params"
	"github.com/inspektor-gadget/inspektor-gadget/pkg/parser"
)

const (
	ParamInterval = "interval"
	ParamSortBy   = "sort"
	ParamMaxRows  = "max-rows"
)

const (
	LocalContainer   params.ValueHint = "local:container"
	K8SNodeName      params.ValueHint = "k8s:node"
	K8SNodeList      params.ValueHint = "k8s:node-list"
	K8SPodName       params.ValueHint = "k8s:pod"
	K8SNamespace     params.ValueHint = "k8s:namespace"
	K8SContainerName params.ValueHint = "k8s:container"
	K8SLabels        params.ValueHint = "k8s:labels"
)

// DefaultSort can be implemented in addition to the Gadget interface, to specify the default sorting columns
type DefaultSort interface {
	SortByDefault() []string
}

// GadgetParams returns params specific to the gadgets' type - for example, it returns
// parameters for 'sort' and 'max-rows' for gadgets with sortable results, and 'interval'
// for periodically called gadgets
func GadgetParams(gadget GadgetDesc, gType GadgetType, parser parser.Parser) params.ParamDescs {
	p := params.ParamDescs{}
	if gType.IsPeriodic() {
		p.Add(IntervalParams()...)
	}
	if gType.CanSort() {
		p.Add(SortableParams(gadget, parser)...)
	}
	return p
}

func IntervalParams() params.ParamDescs {
	return params.ParamDescs{
		{
			Key:          ParamInterval,
			Title:        "Interval",
			DefaultValue: "1",
			TypeHint:     params.TypeUint32,
			Description:  "Interval (in Seconds)",
		},
	}
}

func SortableParams(gadget GadgetDesc, parser parser.Parser) params.ParamDescs {
	if parser == nil {
		return nil
	}

	var defaultSort []string
	if sortInterface, ok := gadget.(DefaultSort); ok {
		defaultSort = sortInterface.SortByDefault()
	}

	return params.ParamDescs{
		{
			Key:          ParamMaxRows,
			Title:        "Max Rows",
			Alias:        "m",
			DefaultValue: "50",
			TypeHint:     params.TypeUint32,
			Description:  "Maximum number of rows to return",
		},
		{
			Key:          ParamSortBy,
			Title:        "Sort By",
			DefaultValue: strings.Join(defaultSort, ","),
			Description:  "Sort by columns. Join multiple columns with ','. Prefix a column with '-' to sort in descending order.",
		},
	}
}

// ParamsFromMap fills the given params (gadget, runtime and operator) using values from `paramMap`. It looks up
// values using prefixes (see also `ParamsToMap`) and applies verification. If verification for a field fails, an
// error will be returned.
func ParamsFromMap(
	paramMap map[string]string,
	gadgetParams *params.Params,
	runtimeParams *params.Params,
	operatorParams params.Collection,
) error {
	err := gadgetParams.CopyFromMap(paramMap, "")
	if err != nil {
		return fmt.Errorf("setting gadget parameters: %w", err)
	}
	err = runtimeParams.CopyFromMap(paramMap, "runtime.")
	if err != nil {
		return fmt.Errorf("setting runtime parameters: %w", err)
	}
	err = operatorParams.CopyFromMap(paramMap, "operator.")
	if err != nil {
		return fmt.Errorf("setting operator parameters: %w", err)
	}
	return nil
}

// ParamsToMap adds the given params (gadget, runtime and operator) to the paramMap. It uses prefixes to ensure
// the keys remain unique.
func ParamsToMap(
	paramMap map[string]string,
	gadgetParams *params.Params,
	runtimeParams *params.Params,
	operatorParams params.Collection,
) {
	gadgetParams.CopyToMap(paramMap, "")
	runtimeParams.CopyToMap(paramMap, "runtime.")
	operatorParams.CopyToMap(paramMap, "operator.")
}
