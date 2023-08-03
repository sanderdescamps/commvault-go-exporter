package exporter

import (
	"fmt"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"gitlab.global.ingenico.com/sdescamps/commvault-go-exporter/client"
)

type DbStatusCollector struct {
	commvaultClient     *client.CommvaultClient
	lastBackupStartTime *prometheus.Desc
	lastBackupEndTime   *prometheus.Desc
	lastBackupJobStatus *prometheus.Desc
	slaStatus           *prometheus.Desc
	applicationSize     *prometheus.Desc
	backupSize          *prometheus.Desc
}

func NewDbStatusCollector(client *client.CommvaultClient) *DbStatusCollector {
	return &DbStatusCollector{
		commvaultClient: client,
		lastBackupStartTime: prometheus.NewDesc(prometheus.BuildFQName(namespace, "db_backup", "last_backup_start_time"),
			"Time when backup started",
			[]string{"client_name", "app_name", "instance_name", "company"}, nil,
		),
		lastBackupEndTime: prometheus.NewDesc(prometheus.BuildFQName(namespace, "db_backup", "last_backup_end_time"),
			"Time when backup finished",
			[]string{"client_name", "app_name", "instance_name", "company"}, nil,
		),
		lastBackupJobStatus: prometheus.NewDesc(prometheus.BuildFQName(namespace, "db_backup", "last_backup_job_status"),
			"Status of last backup job",
			[]string{"client_name", "app_name", "instance_name", "company"}, nil,
		),
		slaStatus: prometheus.NewDesc(prometheus.BuildFQName(namespace, "db_backup", "sla_status"),
			"SLA status",
			[]string{"client_name", "app_name", "instance_name", "company"}, nil,
		),
		applicationSize: prometheus.NewDesc(prometheus.BuildFQName(namespace, "db_backup", "application_size"),
			"Application size",
			[]string{"client_name", "app_name", "instance_name", "company"}, nil,
		),
		backupSize: prometheus.NewDesc(prometheus.BuildFQName(namespace, "db_backup", "backup_size"),
			"Backup size",
			[]string{"client_name", "app_name", "instance_name", "company"}, nil,
		),
	}
}

func (collector *DbStatusCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.lastBackupEndTime
	ch <- collector.lastBackupStartTime
	ch <- collector.lastBackupJobStatus
	ch <- collector.slaStatus
	ch <- collector.applicationSize
	ch <- collector.backupSize
}

func (collector *DbStatusCollector) Collect(ch chan<- prometheus.Metric) {
	if collector.commvaultClient.Status != nil && !collector.commvaultClient.Status.GetIsActive() && collector.commvaultClient.GetToken() != "" {
		return
	}

	dbInstance, err := collector.commvaultClient.Database.GetDatabaseInstances()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] %v", err)
		return
	}
	for _, dbi := range dbInstance {
		labels := []string{dbi.Instance.ClientName, dbi.Instance.AppName, dbi.Instance.InstanceName, dbi.Instance.EntityInfo.CompanyName}
		ch <- prometheus.MustNewConstMetric(collector.lastBackupEndTime, prometheus.GaugeValue, float64(dbi.LastBackupJobInfo.EndTime.Time), labels...)
		ch <- prometheus.MustNewConstMetric(collector.lastBackupStartTime, prometheus.GaugeValue, float64(dbi.LastBackupJobInfo.StartTime.Time), labels...)
		ch <- prometheus.MustNewConstMetric(collector.lastBackupJobStatus, prometheus.GaugeValue, float64(dbi.LastBackupJobInfo.Status), labels...)
		ch <- prometheus.MustNewConstMetric(collector.slaStatus, prometheus.GaugeValue, float64(dbi.SLAStatus), labels...)
		ch <- prometheus.MustNewConstMetric(collector.applicationSize, prometheus.GaugeValue, float64(dbi.ApplicationSize), labels...)
		ch <- prometheus.MustNewConstMetric(collector.backupSize, prometheus.GaugeValue, float64(dbi.BackupSize), labels...)
	}
}
