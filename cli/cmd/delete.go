package main

import (
	"github.com/prateekgogia/kit/cli/pkg/kind"
	"github.com/spf13/cobra"
)

type DeleteFlags struct {
	File        string
	Image       string
	Tidy        bool
	ClusterName string
}

var deleteFlags DeleteFlags

func init() {
	Delete.Flags().StringVarP(&deleteFlags.File, "file", "f", "", "Filename, directory, or URL to files containing the resources.")
	Delete.Flags().StringVarP(&deleteFlags.ClusterName, "clusterName", "c", "default", "Name of the Kind cluster used to run the KIT controller.")
	Delete.Flags().BoolVarP(&deleteFlags.Tidy, "tidy", "t", false, "Tidy up kind cluster after execution.")

	for _, flag := range []string{
		"file",
	} {
		if err := Delete.MarkFlagRequired(flag); err != nil {
			log.Fatalf("Failed to mark flag required %s", err.Error())
		}
	}
}

// Delete command
var Delete = &cobra.Command{
	Use:   "delete",
	Short: "Applies the provided kubernetes resources.",
	Run:   func(cmd *cobra.Command, args []string) { delete() },
}

func delete() {
	log.Infof("Connecting to kind cluster %s", deleteFlags.ClusterName)
	cluster, err := kind.NewCluster(deleteFlags.ClusterName)
	if err != nil {
		log.Fatalf("Failed to connect to kind cluster, %s", err.Error())
	}

	log.Infof("Deleting KIT cluster %s", deleteFlags.File)
	// TODO cluster.DeleteYAML()

	if deleteFlags.Tidy {
		log.Infof("Stopping kind cluster %s", deleteFlags.File)
		if err := cluster.Stop(); err != nil {
			log.Fatal("Failed to stop kind cluster, %s", err.Error())
		}
	}
	log.Info("Success!")
}
