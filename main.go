package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"

	"gitlab.global.ingenico.com/sdescamps/commvault-go-exporter/client"
	"gitlab.global.ingenico.com/sdescamps/commvault-go-exporter/exporter"
)

func rootRun(cmd *cobra.Command, args []string) {
	tcpport, err := cmd.Flags().GetInt16("port")
	if err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "Invalid tcp port\n")
		os.Exit(1)
	}

	commvaultTarget, _ := cmd.Flags().GetString("endpoint")
	username, _ := cmd.Flags().GetString("user")
	password, _ := cmd.Flags().GetString("password")
	insecure, _ := cmd.Flags().GetBool("insecure")

	// var client client.CommvaultClient
	client, err := client.NewClient(commvaultTarget, username, password, insecure)
	if err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "%s\n", err)
		os.Exit(1)
	}
	if client.Status.GetIsActive() {
		fmt.Printf("[info] Commvault endpoint status: Active\n")
	} else {
		fmt.Printf("[info] Commvault endpoint status: Standby\n")
	}

	fmt.Printf("[info] Starting automatic Commvault status check...\n")
	client.Status.Start(cmd.Context())
	fmt.Printf("[info] Commvault status check started\n")

	fmt.Printf("[info] Starting token auto-renew...\n")
	client.StartTokenAutoRenew(cmd.Context())
	fmt.Printf("[info] Token auto-renew enabled\n")

	vmStatusCollector := exporter.NewVmStatusCollector(client)
	storageCollector := exporter.NewStorageDiskCollector(client)
	fsBackupCollector := exporter.NewFsBackupStatusCollector(client)
	dbStatusCollector := exporter.NewDbStatusCollector(client)
	statusCollector := exporter.NewCommvaultStatusCollector(client)
	prometheus.MustRegister(statusCollector, vmStatusCollector, storageCollector, fsBackupCollector, dbStatusCollector)

	http.Handle("/metrics", promhttp.Handler())
	fmt.Printf("[info] Starting server on 0.0.0.0:%d....\n", tcpport)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", tcpport),
		ReadHeaderTimeout: 3 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("[error] %+v", err)
	}
}

func main() {
	var rootCmd = &cobra.Command{Use: "commvault-go-exporter", Run: rootRun}
	rootCmd.PersistentFlags().Int16P("port", "l", 2112, "Port the metricserver will listen on")
	rootCmd.PersistentFlags().StringP("endpoint", "e", "", "Commvault endpoint")
	rootCmd.PersistentFlags().StringP("user", "u", "", "Username to connect to Commvault")
	rootCmd.PersistentFlags().StringP("password", "p", "", "Password to connect to Commvault")
	rootCmd.PersistentFlags().Bool("insecure", false, "Skip SSL certificate verification")
	rootCmd.MarkFlagsRequiredTogether("user", "password")
	// #nosec G104
	rootCmd.MarkFlagRequired("endpoint")
	// #nosec G104
	rootCmd.MarkFlagRequired("user")
	// #nosec G104
	rootCmd.MarkFlagRequired("password")
	// #nosec G104
	rootCmd.Execute()
}
