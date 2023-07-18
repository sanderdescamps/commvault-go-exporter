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

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
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
		fmt.Printf("Starting token auto renew...\n")
		client.StartTokenAutoRenew(cmd.Context())
		fmt.Printf("Token autorenew enabled\n")

		vmStatusCollector := exporter.NewVmStatusCollector(client)
		storageCollector := exporter.NewStorageDiskCollector(client)
		fsBackupCollector := exporter.NewFsBackupStatusCollector(client)
		dbStatusCollector := exporter.NewDbStatusCollector(client)
		prometheus.MustRegister(vmStatusCollector, storageCollector, fsBackupCollector, dbStatusCollector)

		http.Handle("/metrics", promhttp.Handler())
		fmt.Printf("Starting server on 0.0.0.0:%d....\n", tcpport)

		server := &http.Server{
			Addr:              fmt.Sprintf(":%d", tcpport),
			ReadHeaderTimeout: 3 * time.Second,
		}
		err = server.ListenAndServe()
		if err != nil {
			fmt.Printf("[error] %+v", err)
		}
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is just a test")
	},
}

var test2Cmd = &cobra.Command{
	Use:   "test2",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {

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

		bclients, err := client.Bkp.GetBkpClients()
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "%s\n", err)
			os.Exit(1)
		}
		for _, bc := range bclients {
			fmt.Printf("CLIENT clientId=%d hostname: %s company=%s\n", bc.Client.ClientEntity.ClientID, bc.Client.ClientEntity.HostName, bc.Client.ClientEntity.EntityInfo.CompanyName)

			bagents, err := client.Bkp.GetBkpAgents(uint64(bc.Client.ClientEntity.ClientID))
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "%s\n", err)
				os.Exit(1)
			}
			for _, ba := range bagents {
				fmt.Printf("AGENT clientId=%d clientName='%s' appName='%s' applicationId=%d\n", ba.IdaEntity.ClientID, ba.IdaEntity.ClientName, ba.IdaEntity.AppName, ba.IdaEntity.ApplicationID)
			}

			binstances, err := client.Bkp.GetBkpInstancesByClientId(uint64(bc.Client.ClientEntity.ClientID))
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "%s\n", err)
				os.Exit(1)
			}
			for _, bi := range binstances {
				fmt.Printf("INSTANCE appName='%s' applicationSize=%d applicationId=%d\n", bi.Instance.AppName, bi.ApplicationSize, bi.Instance.ApplicationID)
			}

		}

	},
}

var test3Cmd = &cobra.Command{
	Use:   "test3",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {

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

		fmt.Printf("------------------\n")
		dbs2, err := client.Database.GetDatabaseInstances()
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "%s\n", err)
			os.Exit(1)
		}

		for _, db := range dbs2 {
			fmt.Printf("%s\n", db.Instance.ClientName)
		}
	},
}

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.PersistentFlags().Int16P("port", "l", 2112, "Port the metricserver will listen on")
	rootCmd.PersistentFlags().StringP("endpoint", "e", "", "Commvault endpoint")
	rootCmd.PersistentFlags().StringP("user", "u", "", "Username to connect to Commvault")
	rootCmd.PersistentFlags().StringP("password", "p", "", "Password to connect to Commvault")
	rootCmd.PersistentFlags().StringP("token", "t", "", "Token to connect to Commvault")
	rootCmd.PersistentFlags().Bool("insecure", false, "Skip SSL certificate verification")
	rootCmd.MarkFlagsRequiredTogether("user", "password")
	rootCmd.MarkFlagsMutuallyExclusive("user", "token")
	rootCmd.MarkFlagsMutuallyExclusive("password", "token")
	// #nosec G104
	rootCmd.MarkFlagRequired("endpoint")
	rootCmd.AddCommand(runCmd, testCmd, test2Cmd, test3Cmd)
	// #nosec G104
	rootCmd.Execute()
}
