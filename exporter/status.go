package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.global.ingenico.com/sdescamps/commvault-go-exporter/client"
)

type StatusCollector struct {
	commvaultClient *client.CommvaultClient
	commvaultStatus *prometheus.Desc
}

func NewCommvaultStatusCollector(client *client.CommvaultClient) *StatusCollector {
	return &StatusCollector{
		commvaultClient: client,
		commvaultStatus: prometheus.NewDesc(prometheus.BuildFQName(namespace, "status", "is_active"),
			"is Commvault API active",
			[]string{}, nil,
		),
	}
}

func (collector *StatusCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.commvaultStatus
}

func (collector *StatusCollector) Collect(ch chan<- prometheus.Metric) {
	if collector.commvaultClient.Status != nil && collector.commvaultClient.Status.GetIsActive() {
		ch <- prometheus.MustNewConstMetric(collector.commvaultStatus, prometheus.GaugeValue, 1)
	} else {
		ch <- prometheus.MustNewConstMetric(collector.commvaultStatus, prometheus.GaugeValue, 0)
	}
}
