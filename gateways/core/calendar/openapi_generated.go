// +build !ignore_autogenerated

/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by main. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package calendar

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/argoproj/argo-events/gateways/core/calendar/.CalSchedule": schema_gateways_core_calendar__CalSchedule(ref),
	}
}

func schema_gateways_core_calendar__CalSchedule(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CalSchedule describes a time based dependency. One of the fields (schedule, interval, or recurrence) must be passed. Schedule takes precedence over interval; interval takes precedence over recurrence",
				Properties: map[string]spec.Schema{
					"schedule": {
						SchemaProps: spec.SchemaProps{
							Description: "Schedule is a cron-like expression. For reference, see: https://en.wikipedia.org/wiki/Cron",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"interval": {
						SchemaProps: spec.SchemaProps{
							Description: "Interval is a string that describes an interval duration, e.g. 1s, 30m, 2h...",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"recurrence": {
						SchemaProps: spec.SchemaProps{
							Description: "List of RRULE, RDATE and EXDATE lines for a recurring event, as specified in RFC5545. RRULE is a recurrence rule which defines a repeating pattern for recurring events. RDATE defines the list of DATE-TIME values for recurring events. EXDATE defines the list of DATE-TIME exceptions for recurring events. the combination of these rules and dates combine to form a set of date times. NOTE: functionality currently only supports EXDATEs, but in the future could be expanded.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"schedule", "interval"},
			},
		},
		Dependencies: []string{},
	}
}