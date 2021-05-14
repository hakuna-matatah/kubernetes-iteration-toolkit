package main

import (
	"github.com/prateekgogia/kit/cli/pkg/helm"
	"github.com/prateekgogia/kit/cli/pkg/kind"
	"github.com/spf13/cobra"
)

type ApplyFlags struct {
	File        string
	Image       string
	Tidy        bool
	ClusterName string
}

var applyFlags ApplyFlags

func init() {
	Apply.Flags().StringVarP(&applyFlags.File, "file", "f", "", "Filename, directory, or URL to files containing the resources.")
	Apply.Flags().StringVarP(&applyFlags.Image, "image", "i", "TODO", "Image of the KIT controller.")
	Apply.Flags().StringVarP(&applyFlags.ClusterName, "clusterName", "c", "default", "Name of the Kind cluster used to run the KIT controller.")
	Apply.Flags().BoolVarP(&applyFlags.Tidy, "tidy", "t", false, "Tidy up kind cluster after execution.")

	// Required
	for _, flag := range []string{
		"file",
	} {
		if err := Apply.MarkFlagRequired(flag); err != nil {
			log.Fatalf("Failed to mark flag required %s", err.Error())
		}
	}
}

// Apply command
var Apply = &cobra.Command{
	Use:   "apply",
	Short: "Applies the provided kubernetes resources.",
	Run:   func(cmd *cobra.Command, args []string) { apply() },
}

func apply() {
	log.Infof("Connecting to kind cluster %s", applyFlags.ClusterName)
	cluster, err := kind.NewCluster(applyFlags.ClusterName)
	if err != nil {
		log.Fatalf("Failed to connect to cluster, %s", err.Error())
	}

	log.Infof("Installing CertManager Controller")
	if err := cluster.ApplyChart(helm.ApplyOptions{
		ChartName:   "jetstack/cert-manager",
		ReleaseName: "cert-manager",
		Namespace:   "cert-manager",
	}); err != nil {
		log.Fatalf("Failed to install CertManager Controller, %s", err.Error())
	}

	// TODO Implement
	log.Infof("Installing KIT Controller")

	log.Infof("Applying KIT Cluster from %s", applyFlags.File)
	if err := cluster.ApplyYAML(applyFlags.File); err != nil {
		log.Fatalf("Failed to apply objects, %s", err.Error())
	}
	log.Info("Success!")
}
