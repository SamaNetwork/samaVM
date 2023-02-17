// Copyright (C) 2022-2023, Sama , Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SamaNetwork/SamaVM/client"
)

var activityCmd = &cobra.Command{
	Use:   "activity [options]",
	Short: "View recent activity on the network",
	RunE:  activityFunc,
}

func activityFunc(_ *cobra.Command, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("expected exactly 0 arguments, got %d", len(args))
	}
	cli := client.New(uri, requestTimeout)
	activity, err := cli.RecentActivity(context.Background())
	if err != nil {
		return err
	}
	if err := client.PPActivity(activity); err != nil {
		return err
	}
	return nil
}
