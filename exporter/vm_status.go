package exporter

import (
	"fmt"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"gitlab.global.ingenico.com/sdescamps/commvault-go-exporter/client"
)

type VmStatusCollector struct {
	commvaultClient     *client.CommvaultClient
	lastBackupStartTime *prometheus.Desc
	lastBackupEndTime   *prometheus.Desc
	lastBackupJobStatus *prometheus.Desc
	slaStatus           *prometheus.Desc
	status              *prometheus.Desc
	vmSize              *prometheus.Desc
	vmUsedSpace         *prometheus.Desc
}

func NewVmStatusCollector(client *client.CommvaultClient) *VmStatusCollector {
	return &VmStatusCollector{
		commvaultClient: client,
		lastBackupStartTime: prometheus.NewDesc(prometheus.BuildFQName(namespace, "vm_backup", "last_backup_start_time"),
			"Time when backup started",
			[]string{"name", "plan", "sub_client", "strguid"}, nil,
		),
		lastBackupEndTime: prometheus.NewDesc(prometheus.BuildFQName(namespace, "vm_backup", "last_backup_end_time"),
			"Time when backup finished",
			[]string{"name", "plan", "sub_client", "strguid"}, nil,
		),
		lastBackupJobStatus: prometheus.NewDesc(prometheus.BuildFQName(namespace, "vm_backup", "last_backup_job_status"),
			"Status of last backup job",
			[]string{"name", "plan", "sub_client", "strguid"}, nil,
		),
		slaStatus: prometheus.NewDesc(prometheus.BuildFQName(namespace, "vm_backup", "sla_status"),
			"SLA status",
			[]string{"name", "plan", "sub_client", "strguid"}, nil,
		),
		status: prometheus.NewDesc(prometheus.BuildFQName(namespace, "vm_backup", "status"),
			"Job status",
			[]string{"name", "plan", "sub_client", "strguid"}, nil,
		),
		vmSize: prometheus.NewDesc(prometheus.BuildFQName(namespace, "vm_backup", "vm_size"),
			"Time when backup finished",
			[]string{"name", "plan", "sub_client", "strguid"}, nil,
		),
		vmUsedSpace: prometheus.NewDesc(prometheus.BuildFQName(namespace, "vm_backup", "vm_used_space"),
			"Time when backup finished",
			[]string{"name", "plan", "sub_client", "strguid"}, nil,
		),
	}
}

func (collector *VmStatusCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.lastBackupEndTime
	ch <- collector.lastBackupJobStatus
	ch <- collector.slaStatus
	ch <- collector.status
	ch <- collector.vmSize
	ch <- collector.vmUsedSpace
}

func (collector *VmStatusCollector) Collect(ch chan<- prometheus.Metric) {
	if collector.commvaultClient.Status != nil && !collector.commvaultClient.Status.GetIsActive() && collector.commvaultClient.GetToken() != "" {
		return
	}

	vmStatus, err := collector.commvaultClient.Vm.GetVmStatus(client.VmStatus_All)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] %v", err)
		return
	}
	for _, vm := range *vmStatus {
		labels := []string{vm.Name, vm.Plan.PlanName, vm.SubclientName, vm.StrGUID}

		ch <- prometheus.MustNewConstMetric(collector.lastBackupStartTime, prometheus.GaugeValue, float64(vm.BkpStartTime), labels...)
		ch <- prometheus.MustNewConstMetric(collector.lastBackupEndTime, prometheus.GaugeValue, float64(vm.BkpEndTime), labels...)
		ch <- prometheus.MustNewConstMetric(collector.lastBackupJobStatus, prometheus.GaugeValue, float64(vm.VmStatus), labels...)
		ch <- prometheus.MustNewConstMetric(collector.vmSize, prometheus.GaugeValue, float64(vm.VmSize), labels...)
		ch <- prometheus.MustNewConstMetric(collector.vmUsedSpace, prometheus.GaugeValue, float64(vm.VmUsedSpace), labels...)
		ch <- prometheus.MustNewConstMetric(collector.status, prometheus.GaugeValue, float64(vm.VmStatus), labels...)
	}
}
