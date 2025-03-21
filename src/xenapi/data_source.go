/*
 * Copyright (c) Cloud Software Group, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 *   1) Redistributions of source code must retain the above copyright
 *      notice, this list of conditions and the following disclaimer.
 *
 *   2) Redistributions in binary form must reproduce the above
 *      copyright notice, this list of conditions and the following
 *      disclaimer in the documentation and/or other materials
 *      provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
 * INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED
 * OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package xenapi

type DataSourceRecord struct {
	// a human-readable name
	NameLabel string `json:"name_label,omitempty"`
	// a notes field containing human-readable description
	NameDescription string `json:"name_description,omitempty"`
	// true if the data source is being logged
	Enabled bool `json:"enabled,omitempty"`
	// true if the data source is enabled by default. Non-default data sources cannot be disabled
	Standard bool `json:"standard,omitempty"`
	// the units of the value
	Units string `json:"units,omitempty"`
	// the minimum value of the data source
	Min float64 `json:"min,omitempty"`
	// the maximum value of the data source
	Max float64 `json:"max,omitempty"`
	// current value of the data source
	Value float64 `json:"value,omitempty"`
}

type DataSourceRef string

// Data sources for logging in RRDs
type dataSource struct{}

var DataSource dataSource
