package client

// type GetLibraryRespXML struct {
// 	XMLname xml.Name        `xml:"App_GenericEntityResp"`
// 	Data    *GetLibraryResp `xml:"response"`
// }

// type GetLibraryResp struct {
// 	XMLname                   xml.Name       `xml:"response"`
// 	Response                  *[]LibraryItem `json:"response" xml:"entityInfo"`
// 	Processinginstructioninfo map[string]any `json:"processinginstructioninfo"`
// }

// type LibraryItem struct {
// 	XMLname xml.Name `xml:"entityInfo"`
// 	Name    string   `json:"name" xml:"name,attr"`
// 	Id      uint64   `json:"id" xml:"id,attr"`
// }

type GetLibraryResp struct {
	Response                  *[]EntityInfoResponse     `json:"response"`
	Processinginstructioninfo Processinginstructioninfo `json:"processinginstructioninfo"`
}

type EntityInfoResponse struct {
	EntityInfo EntityInfo `json:"entityInfo"`
}

type EntityInfo struct {
	Name string `json:"name"`
	Id   uint64 `json:"id"`
}

//---

type Processinginstructioninfo struct {
	Attributes []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"attributes"`
}

//---

type GetLibraryDetailsResp struct {
	LibraryInfo *LibraryDetail `json:"libraryInfo"`
}

type LibraryDetail struct {
	Description              string `json:"description"`
	LibraryUsedForLogCaching bool   `json:"libraryUsedForLogCaching"`
	ExtendedAttributes       uint64 `json:"extendedAttributes"`
	Manufacturer             string `json:"manufacturer"`
	StoragePoolType          uint64 `json:"storagePoolType"`
	LibraryVendorId          uint64 `json:"libraryVendorId"`
	Model                    string `json:"model"`
	LibraryType              uint64 `json:"libraryType"`
	Status                   string `json:"status"`
	TapeLibSummary           *struct {
		NumOfCleaningMedia             uint64 `json:"numOfCleaningMedia"`
		OfflineReason                  string `json:"offlineReason"`
		NumOfIESlots                   uint64 `json:"numOfIESlots"`
		LastDayThroughput              string `json:"lastDayThroughput"`
		IsOnline                       string `json:"isOnline"`
		NumOfIESlotOccupied            uint64 `json:"numOfIESlotOccupied"`
		LastRestoreTime                string `json:"lastRestoreTime"`
		NumOfActiveMedia               uint64 `json:"numOfActiveMedia"`
		NumOfSpareMedia                uint64 `json:"numOfSpareMedia"`
		Vendor                         string `json:"vendor"`
		NumOfRegSlots                  uint64 `json:"numOfRegSlots"`
		Model                          string `json:"model"`
		NumOfMedia                     uint64 `json:"numOfMedia"`
		Attribute                      uint64 `json:"attribute"`
		NumOfDrivesOffline             uint64 `json:"numOfDrivesOffline"`
		FirmwareRevision               string `json:"firmwareRevision"`
		AssociatedMediaAgents          string `json:"associatedMediaAgents"`
		Controller                     string `json:"controller"`
		BytesBackedupInLast24H         string `json:"bytesBackedupInLast24H"`
		SerialNumber                   string `json:"serialNumber"`
		NumOfAgedMedia                 uint64 `json:"numOfAgedMedia"`
		NumOfAssignedMedia             uint64 `json:"numOfAssignedMedia"`
		NumOfAppendableMedia           uint64 `json:"numOfAppendableMedia"`
		NumOfRegSlotOccupied           uint64 `json:"numOfRegSlotOccupied"`
		BackupReservations             uint64 `json:"backupReservations"`
		NumOfSlots                     uint64 `json:"numOfSlots"`
		BytesBackedupInLast1H          string `json:"bytesBackedupInLast1H"`
		LastBackupTime                 string `json:"lastBackupTime"`
		LastHourThroughput             string `json:"lastHourThroughput"`
		NumOfDrives                    uint64 `json:"numOfDrives"`
		BarcodeReader                  string `json:"barcodeReader"`
		NumOfMediaExporting            uint64 `json:"numOfMediaExporting"`
		AuxiliaryCopyWriteReservations uint64 `json:"auxiliaryCopyWriteReservations"`
		IsEnabled                      uint64 `json:"isEnabled"`
		NumOfFullMedia                 uint64 `json:"numOfFullMedia"`
		TotalNumberOfWriteReservations uint64 `json:"totalNumberOfWriteReservations"`
		Status                         uint64 `json:"status"`
	} `json:"tapeLibSummary"`
	MountPathList *[]struct {
		MediaAgents                string `json:"mediaAgents"`
		MountPathUsedForLogCaching bool   `json:"mountPathUsedForLogCaching"`
		DataServerType             uint64 `json:"dataServerType"`
		MountPathType              uint64 `json:"mountPathType"`
		JobIds                     string `json:"jobIds"`
		PreventDataBlockReferences bool   `json:"preventDataBlockReferences"`
		DisabledForNewWrite        bool   `json:"disabledForNewWrite"`
		LibraryVendorId            uint64 `json:"libraryVendorId"`
		MountPathId                uint64 `json:"mountPathId"`
		MountPathName              string `json:"mountPathName"`
		EnabledMP                  bool   `json:"enabledMP"`
		Status                     string `json:"status"`
		RpStoreLibraryInfo         *struct {
			CopyId               uint64         `json:"copyId"`
			ArchGroupId          uint64         `json:"archGroupId"`
			Flags                uint64         `json:"flags"`
			CommCellId           uint64         `json:"commCellId"`
			MinSpacePerRPStoreGB uint64         `json:"minSpacePerRPStoreGB"`
			OrigRPStoreId        uint64         `json:"origRPStoreId"`
			RpStoreId            uint64         `json:"rpStoreId"`
			MaxSpacePerRPStoreGB uint64         `json:"maxSpacePerRPStoreGB"`
			IntervalWindow       map[string]any `json:"intervalWindow"`
		} `json:"rpStoreLibraryInfo"`
		MountPathSummary *struct {
			LibraryName                   string `json:"libraryName"`
			AvgMediaConsumedPerDay        uint64 `json:"avgMediaConsumedPerDay"`
			TotalDeduplicationAppSize     uint64 `json:"totalDeduplicationAppSize"`
			AvgDataWrittenPerDayMB        uint64 `json:"avgDataWrittenPerDayMB"`
			TotalReserveSpaceMB           uint64 `json:"totalReserveSpaceMB"`
			LastNumDays                   uint64 `json:"lastNumDays"`
			LibraryId                     uint64 `json:"libraryId"`
			TotalSpace                    uint64 `json:"totalSpace"`
			TotalValidData                uint64 `json:"totalValidData"`
			EstimatedSpaceRunoutDays      uint64 `json:"estimatedSpaceRunoutDays"`
			WarningWaterMark              uint64 `json:"warningWaterMark"`
			TotalFreeSpace                uint64 `json:"totalFreeSpace"`
			AvgMediaFreedPerDay           uint64 `json:"avgMediaFreedPerDay"`
			TotalDataWritten              uint64 `json:"totalDataWritten"`
			AvgCapacityFreedMB            uint64 `json:"avgCapacityFreedMB"`
			LowWaterMark                  uint64 `json:"lowWaterMark"`
			NumberOfWriters               uint64 `json:"numberOfWriters"`
			Attribute                     uint64 `json:"attribute"`
			TotalAppSize                  uint64 `json:"totalAppSize"`
			TotalDeduplicationDataWritten uint64 `json:"totalDeduplicationDataWritten"`
		} `json:"mountPathSummary"`
		DeviceInfo          map[string]any `json:"deviceInfo"`
		MetallicStorageInfo map[string]any `json:"metallicStorageInfo"`
		SavedCredential     map[string]any `json:"SavedCredential"`
	} `json:"MountPathList"`
	Library struct {
		LibraryName string `json:"libraryName"`
		Type        uint64 `json:"_type_"`
		LibraryID   uint64 `json:"libraryId"`
	} `json:"library"`
	DriveList *struct {
		Access           string `json:"access"`
		DriveType        string `json:"driveType"`
		DriveName        string `json:"driveName"`
		DrivePoolID      int    `json:"drivePoolId"`
		MountStatus      string `json:"mountStatus"`
		DriveTypeID      int    `json:"driveTypeId"`
		JobIds           string `json:"jobIds"`
		DrivePoolName    string `json:"drivePoolName"`
		DriveAttributes  int    `json:"driveAttributes"`
		DriveID          int    `json:"driveId"`
		ActiveMediaAgent string `json:"activeMediaAgent"`
		Barcode          string `json:"barcode"`
		Status           string `json:"status"`
		DriveSummary     struct {
			OfflineReason                   string `json:"offlineReason"`
			Controller                      string `json:"controller"`
			SerialNumber                    string `json:"serialNumber"`
			LastUseTime                     string `json:"lastUseTime"`
			BytesBackedupInLast24Hour       string `json:"bytesBackedupInLast24Hour"`
			LastDayThroughput               string `json:"lastDayThroughput"`
			IsOnline                        string `json:"isOnline"`
			DriveName                       string `json:"driveName"`
			NumberOfRunningJobs             int    `json:"numberOfRunningJobs"`
			ReadWriteErrosSinceLastClean    int    `json:"readWriteErrosSinceLastClean"`
			Library                         string `json:"library"`
			LastHourThroughput              string `json:"lastHourThroughput"`
			MountUnmountErrosSinceLastClean int    `json:"mountUnmountErrosSinceLastClean"`
			HoursUsedSinceLastClean         int    `json:"hoursUsedSinceLastClean"`
			Vendor                          string `json:"vendor"`
			IsEnabled                       int    `json:"isEnabled"`
			Model                           string `json:"model"`
			BytesBackedupInLast1Hour        string `json:"bytesBackedupInLast1Hour"`
			Attribute                       int    `json:"attribute"`
			FirmwareRevision                string `json:"firmwareRevision"`
			Status                          int    `json:"status"`
		} `json:"driveSummary"`
		Library struct {
			Type int `json:"_type_"`
		} `json:"library"`
	} `json:"DriveList"`
	MagLibSummary *struct {
		AssociatedMediaAgents          string `json:"associatedMediaAgents"`
		BytesBackedupInLast24H         string `json:"bytesBackedupInLast24H"`
		LastDayThroughput              string `json:"lastDayThroughput"`
		IsOnline                       string `json:"isOnline"`
		BackupReservations             uint64 `json:"backupReservations"`
		BytesBackedupInLast1H          string `json:"bytesBackedupInLast1H"`
		LastRestoreTime                string `json:"lastRestoreTime"`
		OnlineMountPaths               string `json:"onlineMountPaths"`
		LastBackupTime                 string `json:"lastBackupTime"`
		TotalAvailableSpace            string `json:"totalAvailableSpace"`
		LastHourThroughput             string `json:"lastHourThroughput"`
		TotalFreeSpace                 string `json:"totalFreeSpace"`
		TotalCapacity                  string `json:"totalCapacity"`
		AuxiliaryCopyWriteReservations uint64 `json:"auxiliaryCopyWriteReservations"`
		NumOfMountPath                 uint64 `json:"numOfMountPath"`
		TotalNumberOfWriteReservations uint64 `json:"totalNumberOfWriteReservations"`
		MountPathUsage                 string `json:"mountPathUsage"`
	} `json:"magLibSummary"`
}
