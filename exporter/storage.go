package exporter

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"gitlab.global.ingenico.com/sdescamps/commvault-go-exporter/client"
)

const BYTE_UNIT_LIST string = "KMGTPEZYRQ"

type StorageDiskCollector struct {
	commvaultClient        *client.CommvaultClient
	bytesBackedupInLast24H *prometheus.Desc
	bytesBackedupInLast1H  *prometheus.Desc
	totalNumOfMountPath    *prometheus.Desc
	onlineMountPaths       *prometheus.Desc
	totalCapacityBytes     *prometheus.Desc
	freeCapacityBytes      *prometheus.Desc
	availableCapacityBytes *prometheus.Desc
	isOnline               *prometheus.Desc
	//tape
	numOfActiveMedia     *prometheus.Desc
	numOfAgedMedia       *prometheus.Desc
	numOfAppendableMedia *prometheus.Desc
	numOfAssignedMedia   *prometheus.Desc
	numOfCleaningMedia   *prometheus.Desc
	numOfDrives          *prometheus.Desc
	numOfDrivesOffline   *prometheus.Desc
	numOfFullMedia       *prometheus.Desc
	numOfIESlotOccupied  *prometheus.Desc
	numOfIESlots         *prometheus.Desc
	numOfMedia           *prometheus.Desc
	numOfMediaExporting  *prometheus.Desc
	numOfRegSlotOccupied *prometheus.Desc
	numOfRegSlots        *prometheus.Desc
	numOfSlots           *prometheus.Desc
	numOfSpareMedia      *prometheus.Desc
	lastRestoreTime      *prometheus.Desc
	lastBackupTime       *prometheus.Desc
}

// The collector which will report on Disk and Cloud libraries
func NewStorageDiskCollector(client *client.CommvaultClient) *StorageDiskCollector {
	return &StorageDiskCollector{
		commvaultClient: client,
		bytesBackedupInLast24H: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "bytes_backedup_24h"),
			"Amount of bytes backuped in the last 24 hours",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		bytesBackedupInLast1H: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "bytes_backedup_1h"),
			"Amount of bytes backuped in the last hour",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		isOnline: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "is_online"),
			"Job status description",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		// only for disks and cloud
		totalNumOfMountPath: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "total_number_mount_paths"),
			"Total number of mount paths",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		onlineMountPaths: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "online_number_mount_paths"),
			"Number of online mount paths ",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		totalCapacityBytes: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "total_capacity_bytes"),
			"Total capacity in bytes",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		freeCapacityBytes: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "free_capacity_bytes"),
			"Free capacity in bytes",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		availableCapacityBytes: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "available_capacity_bytes"),
			"Available capacity in bytes",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		//only for tapes

		numOfActiveMedia: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_active_media"),
			"Tape library total number of active  media",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfAgedMedia: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_aged_media"),
			"Tape library total number of aged media",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfAppendableMedia: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_appendable_media"),
			"Tape library total number of appendable media",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfAssignedMedia: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_assigned_media"),
			"Tape library total number of assigned media",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfCleaningMedia: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_clean_media"),
			"Tape library total number of cleaning media",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfDrives: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_drives"),
			"Tape library total number of drives.",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfDrivesOffline: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_drives"),
			"Tape library number of offline drives",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfFullMedia: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_full_media"),
			"Tape library total number of full media",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfIESlotOccupied: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_ie_slots_occupied"),
			"Tape library total number of mail slots occupied",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfIESlots: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_ie_slots"),
			"Tape library total number of mail slots",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfMedia: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_media"),
			"Tape library total number of media in the library",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfMediaExporting: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_media_exporting"),
			"Tape library total number of exported media",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfRegSlotOccupied: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_reg_slot_occupied"),
			"Tape library total number of occupied slots",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfRegSlots: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_reg_slot"),
			"Tape library total number of registered slots",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfSlots: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_slots"),
			"Tape library total number of slots",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		numOfSpareMedia: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "num_of_spare_media"),
			"Tape library total number of spare media",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		lastRestoreTime: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "last_restore_time"),
			"Tape library time of the most recent restore job",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
		lastBackupTime: prometheus.NewDesc(prometheus.BuildFQName(namespace, "storage_library", "last_backup_time"),
			"Tape library time of the most recent backup job",
			[]string{"libraryName", "type", "libraryId"}, nil,
		),
	}
}

func (collector *StorageDiskCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.bytesBackedupInLast24H
	ch <- collector.bytesBackedupInLast1H
	ch <- collector.totalNumOfMountPath
	ch <- collector.onlineMountPaths
	ch <- collector.totalCapacityBytes
	ch <- collector.freeCapacityBytes
	ch <- collector.availableCapacityBytes
	ch <- collector.isOnline
}

func (collector *StorageDiskCollector) Collect(ch chan<- prometheus.Metric) {
	if !collector.commvaultClient.GetActiveStatus() {
		return
	}

	storageDetails, err := collector.commvaultClient.Storage.GetLibrariesDetails()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] %v", err)
	}
	for _, sd := range storageDetails {
		var libTyp string
		if sd.LibraryType == 1 {
			libTyp = "tape"
		} else if sd.LibraryType == 3 && (strings.ToLower(sd.Model) == "disk" || strings.ToLower(sd.Model) == "cloud") {
			libTyp = strings.ToLower(sd.Model)
		} else if sd.LibraryType == 4 {
			libTyp = "stand alone library"
		} else {
			libTyp = "unknown"
		}

		labels := []string{sd.Library.LibraryName, libTyp, strconv.FormatUint(sd.Library.LibraryID, 10)}

		if libTyp == "tape" {
			bbil24h, _ := byteStringToFloat64(sd.TapeLibSummary.BytesBackedupInLast24H)
			ch <- prometheus.MustNewConstMetric(collector.bytesBackedupInLast24H, prometheus.GaugeValue, bbil24h, labels...)
			bbil1h, _ := byteStringToFloat64(sd.TapeLibSummary.BytesBackedupInLast1H)
			ch <- prometheus.MustNewConstMetric(collector.bytesBackedupInLast24H, prometheus.GaugeValue, bbil1h, labels...)
			var isOnline float64
			if strings.ToLower(sd.TapeLibSummary.IsOnline) == "no" {
				isOnline = -1
			} else if strings.ToLower(sd.TapeLibSummary.IsOnline) == "ready" {
				isOnline = 1
			} else {
				isOnline = 0
			}
			ch <- prometheus.MustNewConstMetric(collector.isOnline, prometheus.GaugeValue, isOnline, labels...)

			ch <- prometheus.MustNewConstMetric(collector.numOfActiveMedia, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfActiveMedia), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfAgedMedia, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfAgedMedia), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfAppendableMedia, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfAppendableMedia), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfAssignedMedia, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfAssignedMedia), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfCleaningMedia, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfCleaningMedia), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfDrives, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfDrives), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfDrivesOffline, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfDrivesOffline), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfFullMedia, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfFullMedia), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfIESlotOccupied, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfIESlotOccupied), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfIESlots, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfIESlots), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfMedia, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfMedia), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfMediaExporting, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfMediaExporting), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfRegSlotOccupied, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfRegSlotOccupied), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfRegSlots, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfRegSlots), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfSlots, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfSlots), labels...)
			ch <- prometheus.MustNewConstMetric(collector.numOfSpareMedia, prometheus.GaugeValue, float64(sd.TapeLibSummary.NumOfSpareMedia), labels...)

			lastBackup, err := time.Parse(TIME_LAYOUT, sd.TapeLibSummary.LastBackupTime)
			if err == nil {
				ch <- prometheus.MustNewConstMetric(collector.lastBackupTime, prometheus.GaugeValue, float64(lastBackup.Unix()), labels...)
			}
			lastRestore, err := time.Parse(TIME_LAYOUT, sd.TapeLibSummary.LastBackupTime)
			if err == nil {
				ch <- prometheus.MustNewConstMetric(collector.lastRestoreTime, prometheus.GaugeValue, float64(lastRestore.Unix()), labels...)
			}

		} else if libTyp == "cloud" || libTyp == "disk" {
			bbil24h, _ := byteStringToFloat64(sd.MagLibSummary.BytesBackedupInLast24H)
			ch <- prometheus.MustNewConstMetric(collector.bytesBackedupInLast24H, prometheus.GaugeValue, bbil24h, labels...)
			bbil1h, _ := byteStringToFloat64(sd.MagLibSummary.BytesBackedupInLast24H)
			ch <- prometheus.MustNewConstMetric(collector.bytesBackedupInLast1H, prometheus.GaugeValue, bbil1h, labels...)
			var isOnline float64
			if strings.ToLower(sd.MagLibSummary.IsOnline) == "no" {
				isOnline = -1 //is not online
			} else if strings.ToLower(sd.MagLibSummary.IsOnline) == "ready" {
				isOnline = 1 //is online
			} else {
				isOnline = 0 //unknown status
			}
			ch <- prometheus.MustNewConstMetric(collector.isOnline, prometheus.GaugeValue, isOnline, labels...)
			onlineMP, totalMP := parseMountPaths(sd.MagLibSummary.OnlineMountPaths)
			ch <- prometheus.MustNewConstMetric(collector.onlineMountPaths, prometheus.GaugeValue, float64(onlineMP), labels...)
			ch <- prometheus.MustNewConstMetric(collector.totalNumOfMountPath, prometheus.GaugeValue, float64(totalMP), labels...)
			totalCap, _ := byteStringToFloat64(sd.MagLibSummary.TotalCapacity)
			freeCap, _ := byteStringToFloat64(sd.MagLibSummary.TotalFreeSpace)
			availCap, _ := byteStringToFloat64(sd.MagLibSummary.TotalAvailableSpace)
			ch <- prometheus.MustNewConstMetric(collector.totalCapacityBytes, prometheus.GaugeValue, totalCap, labels...)
			ch <- prometheus.MustNewConstMetric(collector.freeCapacityBytes, prometheus.GaugeValue, freeCap, labels...)
			ch <- prometheus.MustNewConstMetric(collector.availableCapacityBytes, prometheus.GaugeValue, availCap, labels...)
		}
	}
}

func byteStringToFloat64(value string) (float64, error) {
	value = strings.Replace(value, " ", "", -1)
	unitList := strings.Split(BYTE_UNIT_LIST, "")

	re := regexp.MustCompile(`^ *(\-?[\d]*[\.\,][\d]*) *(?i)([` + BYTE_UNIT_LIST + `])?[b]?(?-i) *$`)

	if re.MatchString(value) {
		rfind := re.FindStringSubmatch(value)
		num, _ := strconv.ParseFloat(strings.Replace(rfind[1], ",", ".", -1), 64)
		idx := IndexOf(unitList, strings.ToUpper(rfind[2]))
		byts := math.Pow(1000, float64(idx)+1) * num
		return byts, nil
	}
	return 0, nil
}

func IndexOf[T comparable](collection []T, el T) int {
	for i, x := range collection {
		if x == el {
			return i
		}
	}
	return -1
}

func parseMountPaths(mountString string) (int, int) {
	re := regexp.MustCompile(`^ *(\d*) out of (\d*) *$`)

	if re.MatchString(mountString) {
		rfind := re.FindStringSubmatch(mountString)
		num, _ := strconv.Atoi(rfind[1])
		tot, _ := strconv.Atoi(rfind[2])
		return num, tot
	}

	return 0, 0
}

const (
	TIME_LAYOUT string = "Jan 1 2006  3: 04PM"
)
