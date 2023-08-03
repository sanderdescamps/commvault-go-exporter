package exporter

import (
	"fmt"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"gitlab.global.ingenico.com/sdescamps/commvault-go-exporter/client"
)

type FsBackupStatusCollector struct {
	commvaultClient *client.CommvaultClient
	applicationSize *prometheus.Desc
}

func NewFsBackupStatusCollector(client *client.CommvaultClient) *FsBackupStatusCollector {
	return &FsBackupStatusCollector{
		commvaultClient: client,
		applicationSize: prometheus.NewDesc(prometheus.BuildFQName(namespace, "fs_backup", "application_size"),
			"size of the application",
			[]string{"hostname", "organisation", "client_name", "app_name", "instance_name"}, nil,
		),
	}
}

func (collector *FsBackupStatusCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.applicationSize

}

func (collector *FsBackupStatusCollector) Collect(ch chan<- prometheus.Metric) {
	if collector.commvaultClient.Status != nil && !collector.commvaultClient.Status.GetIsActive() && collector.commvaultClient.GetToken() != "" {
		return
	}

	bclients, err := collector.commvaultClient.Bkp.GetBkpClients()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to get the clients\n%s\n", err)
		return
	}
	for _, bc := range bclients {
		binstances, err := collector.commvaultClient.Bkp.GetBkpInstancesByClientId(uint64(bc.Client.ClientEntity.ClientID))
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] Failed to get the instances for client (%s,ID=%d) \n%s\n", bc.Client.ClientEntity.ClientName, bc.Client.ClientEntity.ClientID, err)
			return
		}
		for _, bi := range binstances {
			labels := []string{bc.Client.ClientEntity.HostName, bc.Client.ClientEntity.EntityInfo.CompanyName, bc.Client.ClientEntity.ClientName, bi.Instance.AppName, bi.Instance.InstanceName}
			ch <- prometheus.MustNewConstMetric(collector.applicationSize, prometheus.GaugeValue, float64(bi.ApplicationSize), labels...)
		}

	}
}
