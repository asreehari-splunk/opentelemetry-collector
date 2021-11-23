// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service // import "go.opentelemetry.io/collector/service"

import (
	"github.com/spf13/cobra"

	"go.opentelemetry.io/collector/config/configmapprovider"
	"go.opentelemetry.io/collector/service/featuregate"
)

// NewCommand constructs a new cobra.Command using the given Collector.
// TODO: Make this independent of the collector internals.
func NewCommand(set CollectorSettings) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:          set.BuildInfo.Command,
		Version:      set.BuildInfo.Version,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			featuregate.Apply(featuregate.GetFlags())
			if set.ConfigMapProvider == nil {
				set.ConfigMapProvider = configmapprovider.NewDefault(getConfigFlag(), getSetFlag())
			}
			col, err := New(set)
			if err != nil {
				return err
			}
			return col.Run(cmd.Context())
		},
	}

	rootCmd.Flags().AddGoFlagSet(flags())
	return rootCmd
}
