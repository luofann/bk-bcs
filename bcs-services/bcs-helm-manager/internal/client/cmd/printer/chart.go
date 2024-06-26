/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package printer xxx
package printer

import (
	"os"

	"github.com/olekukonko/tablewriter"

	helmmanager "github.com/Tencent/bk-bcs/bcs-services/bcs-helm-manager/proto/bcs-helm-manager"
)

// PrintChartInTable print chart data in table format
func PrintChartInTable(wide bool, chart *helmmanager.ChartListData) {
	if chart == nil {
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(func() []string {
		r := []string{"NAME", "VERSION", "APP_VERSION"}
		if wide {
			r = append(r, "UPDATE_BY", "UPDATE_TIME", "DESCRIPTION")
		}
		return r
	}())
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("    ")
	table.SetNoWhiteSpace(true)

	for _, ct := range chart.Data {
		table.Append(func() []string {
			r := []string{
				ct.GetName(), ct.GetLatestVersion(), ct.GetLatestAppVersion(),
			}

			if wide {
				r = append(r, ct.GetUpdateBy(), ct.GetUpdateTime(), cut(ct.GetLatestDescription(), 50))
			}

			return r
		}())
	}
	table.Render()
}
