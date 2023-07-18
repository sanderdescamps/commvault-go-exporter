package client

func fake_vmstatus_response() string {
	return `
	{
		"totalRecords": 167,
		"pageNo": 0,
		"errorMessage": "",
		"errorCode": 0,
		"pageSize": 10000,
		"vmStatusInfoList": [
		  {
			"vmHost": "HOTFIXCOMMCELL",
			"vmGuestSpace": 152106536960,
			"bkpStartTime": 1539716353,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19698542,
			"strOSName": "Windows Server 2012 R2 Datacenter",
			"isDeleted": false,
			"vendor": 2,
			"osType": 1,
			"vmSize": 214748364800,
			"vmUsedSpace": 567279616,
			"subclientId": 206951,
			"bkpEndTime": 1539718923,
			"vmAgent": "hotfixcommcell",
			"name": "hotfx-exchange",
			"vmHardwareVer": "5.0",
			"strGUID": "35C4C1E5-1C35-4F6E-B4B4-02BFBD05AA1C",
			"subclientName": "default",
			"client": {
			  "clientId": 58426,
			  "clientName": "hotfx-exchange"
			},
			"proxyClient": {
			  "clientId": 58422,
			  "clientName": "hotfixcommcell"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 58422,
			  "clientName": "hotfixcommcell",
			  "flags": {
				"disabled": false
			  }
			},
			"vmSubClientEntity": {
			  "subclientId": 206955,
			  "subclientName": "default"
			}
		  },
		  {
			"vmHost": "HOTFIXCOMMCELL",
			"vmGuestSpace": 53125202830,
			"bkpStartTime": 1539719035,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19698542,
			"strOSName": "",
			"isDeleted": false,
			"vendor": 2,
			"osType": 2,
			"vmSize": 214748364800,
			"vmUsedSpace": 524288000,
			"subclientId": 206951,
			"bkpEndTime": 1539719412,
			"vmAgent": "hotfixcommcell",
			"name": "hotfx-oracle",
			"vmHardwareVer": "5.0",
			"strGUID": "ADBE2E45-AE49-4D3D-A3E6-8FFCED47931D",
			"subclientName": "default",
			"client": {
			  "clientId": 58428,
			  "clientName": "hotfx-oracle"
			},
			"proxyClient": {
			  "clientId": 58422,
			  "clientName": "hotfixcommcell"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 58422,
			  "clientName": "hotfixcommcell",
			  "flags": {
				"disabled": false
			  }
			},
			"vmSubClientEntity": {
			  "subclientId": 206958,
			  "subclientName": "default"
			}
		  },
		  {
			"vmHost": "HOTFIXCOMMCELL",
			"vmGuestSpace": 51171807232,
			"bkpStartTime": 1539718926,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19698542,
			"strOSName": "Windows Server 2012 R2 Datacenter",
			"isDeleted": false,
			"vendor": 2,
			"osType": 1,
			"vmSize": 214748364800,
			"vmUsedSpace": 418381824,
			"subclientId": 206951,
			"bkpEndTime": 1539719625,
			"vmAgent": "hotfixcommcell",
			"name": "hotfx-sql",
			"vmHardwareVer": "5.0",
			"strGUID": "82773AE9-A0C0-49E9-B25C-BE87A96D4408",
			"subclientName": "default",
			"client": {
			  "clientId": 58429,
			  "clientName": "hotfx-sql"
			},
			"proxyClient": {
			  "clientId": 58422,
			  "clientName": "hotfixcommcell"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 58422,
			  "clientName": "hotfixcommcell",
			  "flags": {
				"disabled": false
			  }
			},
			"vmSubClientEntity": {
			  "subclientId": 206957,
			  "subclientName": "default"
			}
		  },
		  {
			"vmHost": "HOTFIXCOMMCELL",
			"vmGuestSpace": 197568450560,
			"bkpStartTime": 1539715576,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19698543,
			"strOSName": "Windows Server 2012 R2 Datacenter",
			"isDeleted": false,
			"vendor": 2,
			"osType": 1,
			"vmSize": 429496729600,
			"vmUsedSpace": 1151336448,
			"subclientId": 209605,
			"bkpEndTime": 1539718006,
			"vmAgent": "hotfixcommcell",
			"name": "hotfx-winma",
			"vmHardwareVer": "5.0",
			"strGUID": "07ADE5CB-4242-4D57-BA8F-8164EA5FF265",
			"subclientName": "Elisa_Test",
			"client": {
			  "clientId": 58430,
			  "clientName": "hotfx-winma"
			},
			"proxyClient": {
			  "clientId": 58422,
			  "clientName": "hotfixcommcell"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 58422,
			  "clientName": "hotfixcommcell",
			  "flags": {
				"disabled": false
			  }
			},
			"vmSubClientEntity": {
			  "subclientId": 206953,
			  "subclientName": "default"
			}
		  },
		  {
			"vmHost": "HOTFIXCOMMCELL",
			"vmGuestSpace": 118888730624,
			"bkpStartTime": 1539716352,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19698542,
			"strOSName": "Windows Server 2012 R2 Datacenter",
			"isDeleted": false,
			"vendor": 2,
			"osType": 1,
			"vmSize": 429496729600,
			"vmUsedSpace": 2982150144,
			"subclientId": 206951,
			"bkpEndTime": 1539719032,
			"vmAgent": "hotfixcommcell",
			"name": "hotfxcs",
			"vmHardwareVer": "5.0",
			"strGUID": "76228FDB-C7E0-4143-A00D-FA5E9E6C436D",
			"subclientName": "default",
			"client": {
			  "clientId": 58431,
			  "clientName": "hotfxcs"
			},
			"proxyClient": {
			  "clientId": 58422,
			  "clientName": "hotfixcommcell"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 58422,
			  "clientName": "hotfixcommcell",
			  "flags": {
				"disabled": false
			  }
			},
			"vmSubClientEntity": {
			  "subclientId": 206956,
			  "subclientName": "default"
			}
		  },
		  {
			"vmHost": "MURLOCHYPERV",
			"vmGuestSpace": 3560812544,
			"bkpStartTime": 1502078881,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 12356661,
			"strOSName": "",
			"isDeleted": true,
			"vendor": 2,
			"osType": 2,
			"vmSize": 136365211648,
			"vmUsedSpace": 85017600,
			"subclientId": 206742,
			"bkpEndTime": 1502078953,
			"vmAgent": "murlochyperv",
			"name": "NoCopyVM3",
			"vmHardwareVer": "5.0",
			"strGUID": "5014CFA1-13F3-4C7F-9024-832586C4A823",
			"subclientName": "VMBackup",
			"client": {
			  "clientId": 58529,
			  "clientName": "NoCopyVM3"
			},
			"proxyClient": {
			  "clientId": 58247,
			  "clientName": "murlochyperv"
			},
			"plan": {
			  "planName": "VSAPLAN",
			  "planId": 16
			},
			"pseudoClient": {
			  "clientId": 58249,
			  "clientName": "HyperVM",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.21.250",
			"vmGuestSpace": 13378404352,
			"bkpStartTime": 1540958499,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 20087164,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 21474848784,
			"vmUsedSpace": 0,
			"subclientId": 206759,
			"bkpEndTime": 1540958544,
			"vmAgent": "matador0",
			"name": "VSA-Murloc-6",
			"vmHardwareVer": "vmx-10",
			"strGUID": "50268674-8878-af64-78ca-bd1f860ffcbd",
			"subclientName": "VMBackup1",
			"client": {
			  "clientId": 58555,
			  "clientName": "VSA-Murloc-6"
			},
			"proxyClient": {
			  "clientId": 25156,
			  "clientName": "matador0"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 53801,
			  "clientName": "VMware",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.21.250",
			"vmGuestSpace": 13378404352,
			"bkpStartTime": 1540958515,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 20087164,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 21474848784,
			"vmUsedSpace": 0,
			"subclientId": 206759,
			"bkpEndTime": 1540958549,
			"vmAgent": "matador0",
			"name": "VSA-Murloc-7",
			"vmHardwareVer": "vmx-10",
			"strGUID": "502699f1-55fd-9dba-9262-80655efc9b9f",
			"subclientName": "VMBackup1",
			"client": {
			  "clientId": 58556,
			  "clientName": "VSA-Murloc-7"
			},
			"proxyClient": {
			  "clientId": 25156,
			  "clientName": "matador0"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 53801,
			  "clientName": "VMware",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.21.250",
			"vmGuestSpace": 13378404352,
			"bkpStartTime": 1540958516,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 20087167,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 21474848744,
			"vmUsedSpace": 0,
			"subclientId": 206760,
			"bkpEndTime": 1540958551,
			"vmAgent": "matador0",
			"name": "VSA-Murloc-2",
			"vmHardwareVer": "vmx-10",
			"strGUID": "5026bef2-e814-2718-16ba-0f903a846125",
			"subclientName": "VMbackup2",
			"client": {
			  "clientId": 58558,
			  "clientName": "VSA-Murloc-2"
			},
			"proxyClient": {
			  "clientId": 25156,
			  "clientName": "matador0"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 53801,
			  "clientName": "VMware",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.21.250",
			"vmGuestSpace": 13378404352,
			"bkpStartTime": 1540958499,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 20087167,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 21474848784,
			"vmUsedSpace": 0,
			"subclientId": 206760,
			"bkpEndTime": 1540958548,
			"vmAgent": "matador0",
			"name": "VSA-Murloc-3",
			"vmHardwareVer": "vmx-10",
			"strGUID": "5026b2b1-b90d-64b1-9cc5-61217687a705",
			"subclientName": "VMbackup2",
			"client": {
			  "clientId": 58559,
			  "clientName": "VSA-Murloc-3"
			},
			"proxyClient": {
			  "clientId": 25156,
			  "clientName": "matador0"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 53801,
			  "clientName": "VMware",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.21.250",
			"vmGuestSpace": 13378404352,
			"bkpStartTime": 1540958554,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 20087167,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 21474848744,
			"vmUsedSpace": 0,
			"subclientId": 206760,
			"bkpEndTime": 1540958580,
			"vmAgent": "matador0",
			"name": "VSA-Murloc-4",
			"vmHardwareVer": "vmx-10",
			"strGUID": "5026d1e1-8bba-02c9-5334-f0c183a741d9",
			"subclientName": "VMbackup2",
			"client": {
			  "clientId": 58560,
			  "clientName": "VSA-Murloc-4"
			},
			"proxyClient": {
			  "clientId": 25156,
			  "clientName": "matador0"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 53801,
			  "clientName": "VMware",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 29799124992,
			"bkpStartTime": 1538776369,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 53687367043,
			"vmUsedSpace": 25207767040,
			"subclientId": 211627,
			"bkpEndTime": 1538784413,
			"vmAgent": "VSAPROXY",
			"name": "2008R2",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502d63a4-d594-6032-cf59-b0fcc5b42ace",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60209,
			  "clientName": "2008R2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 12461498368,
			"bkpStartTime": 1538784466,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 48322230146,
			"vmUsedSpace": 8204058624,
			"subclientId": 211627,
			"bkpEndTime": 1538787714,
			"vmAgent": "VSAPROXY",
			"name": "HMVM-WIN2K8R2",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502f289e-010a-a8d7-60a0-74888e5cab87",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60210,
			  "clientName": "HMVM-WIN2K8R2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 11579482112,
			"bkpStartTime": 1538781517,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": true,
			"vendor": 1,
			"osType": 1,
			"vmSize": 21476743405,
			"vmUsedSpace": 7554990080,
			"subclientId": 211627,
			"bkpEndTime": 1538784187,
			"vmAgent": "VSAPROXY",
			"name": "TESTVM5",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502dbb40-eca1-9e9c-7c0a-de17639b466b",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60211,
			  "clientName": "TESTVM5"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 21325156352,
			"bkpStartTime": 1538778111,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 107375311286,
			"vmUsedSpace": 14794358784,
			"subclientId": 211627,
			"bkpEndTime": 1538782578,
			"vmAgent": "VSAPROXY",
			"name": "ViewVM",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502da527-9633-b2f9-39f8-44da4ca56625",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60212,
			  "clientName": "ViewVM"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 8461950976,
			"bkpStartTime": 1538794871,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": true,
			"vendor": 1,
			"osType": 1,
			"vmSize": 12885946226,
			"vmUsedSpace": 7572815872,
			"subclientId": 211627,
			"bkpEndTime": 1538796863,
			"vmAgent": "VSAPROXY",
			"name": "shiv-test1",
			"vmHardwareVer": "vmx-07",
			"strGUID": "50396f2b-0353-4605-787a-c48cce75558a",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60213,
			  "clientName": "shiv-test1"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 80199409664,
			"bkpStartTime": 1538783520,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 171801573126,
			"vmUsedSpace": 71755104256,
			"subclientId": 211627,
			"bkpEndTime": 1538807300,
			"vmAgent": "VSAPROXY",
			"name": "vCACServer",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502dff7a-d80f-1e2d-83c1-072c37ca9cdc",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60214,
			  "clientName": "vCACServer"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 2132289536,
			"bkpStartTime": 1538782493,
			"type": 9,
			"vmStatus": 4,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2003 (64-bit)",
			"isDeleted": true,
			"vendor": 1,
			"osType": 1,
			"vmSize": 2148126241,
			"vmUsedSpace": 2135949312,
			"subclientId": 211627,
			"bkpEndTime": 1538783048,
			"vmAgent": "VSAPROXY",
			"name": "wf_Powreredon",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502dc5ae-cb92-0ed3-0a88-6867119b7a27",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60216,
			  "clientName": "wf_Powreredon"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 11571830784,
			"bkpStartTime": 1538783111,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 21476469844,
			"vmUsedSpace": 7554990080,
			"subclientId": 211627,
			"bkpEndTime": 1538785901,
			"vmAgent": "VSAPROXY",
			"name": "ViewVM_2",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502df113-7b09-c10b-ac17-cac4b4dc096f",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60217,
			  "clientName": "ViewVM_2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 14934724608,
			"bkpStartTime": 1538789745,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 85905928283,
			"vmUsedSpace": 20222836736,
			"subclientId": 211627,
			"bkpEndTime": 1538794857,
			"vmAgent": "VSAPROXY",
			"name": "HMVM-WIN2K12R2",
			"vmHardwareVer": "vmx-08",
			"strGUID": "502fa341-243f-3012-a7a1-a72343fd5b93",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60218,
			  "clientName": "HMVM-WIN2K12R2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 2894131200,
			"bkpStartTime": 1538772086,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Other Linux (32-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 8591268434,
			"vmUsedSpace": 7121928192,
			"subclientId": 211627,
			"bkpEndTime": 1538773053,
			"vmAgent": "VSAPROXY",
			"name": "UbuntuVM",
			"vmHardwareVer": "vmx-08",
			"strGUID": "502d5fdb-6fc5-0271-6340-ced3e5b7a0d5",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60219,
			  "clientName": "UbuntuVM"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 0,
			"bkpStartTime": 1538781479,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 4666640,
			"vmUsedSpace": 0,
			"subclientId": 211627,
			"bkpEndTime": 1538781510,
			"vmAgent": "VSAPROXY",
			"name": "wf_NoIOActivity",
			"vmHardwareVer": "vmx-10",
			"strGUID": "502daf2f-7d15-8fd3-dc3e-11031f01d008",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60221,
			  "clientName": "wf_NoIOActivity"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 11578208256,
			"bkpStartTime": 1538772356,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 21476369557,
			"vmUsedSpace": 7557087232,
			"subclientId": 211627,
			"bkpEndTime": 1538774646,
			"vmAgent": "VSAPROXY",
			"name": "ViewVM_1",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502d6eee-0807-2621-4e0e-be47192e34c2",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60222,
			  "clientName": "ViewVM_1"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 82772987904,
			"bkpStartTime": 1538781806,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 214780799631,
			"vmUsedSpace": 94665441280,
			"subclientId": 211627,
			"bkpEndTime": 1538804056,
			"vmAgent": "VSAPROXY",
			"name": "PhoneApps",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502dbef2-b8c9-a24d-487d-8969466538d9",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60223,
			  "clientName": "PhoneApps"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 7321157632,
			"bkpStartTime": 1538786588,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 37583406258,
			"vmUsedSpace": 7321157632,
			"subclientId": 211627,
			"bkpEndTime": 1538788323,
			"vmAgent": "VSAPROXY",
			"name": "Solidfire-RHEL7",
			"vmHardwareVer": "vmx-10",
			"strGUID": "502f6eb4-a4b1-2568-38f8-d908f17edff8",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60224,
			  "clientName": "Solidfire-RHEL7"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 2894266368,
			"bkpStartTime": 1538784422,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Other Linux (32-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 8591243163,
			"vmUsedSpace": 7121928192,
			"subclientId": 211627,
			"bkpEndTime": 1538785988,
			"vmAgent": "VSAPROXY",
			"name": "UbuntuVM-solidfire",
			"vmHardwareVer": "vmx-08",
			"strGUID": "502f27cd-7a95-d684-4316-364c5343f95e",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60225,
			  "clientName": "UbuntuVM-solidfire"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 18282119168,
			"bkpStartTime": 1538776369,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 107381295123,
			"vmUsedSpace": 19054723072,
			"subclientId": 211627,
			"bkpEndTime": 1538781444,
			"vmAgent": "VSAPROXY",
			"name": "vRAAgent",
			"vmHardwareVer": "vmx-11",
			"strGUID": "502d7730-0f85-7222-2f46-89b35b4d3a9e",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60226,
			  "clientName": "vRAAgent"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 12113285120,
			"bkpStartTime": 1538776369,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 21476692737,
			"vmUsedSpace": 7886340096,
			"subclientId": 211627,
			"bkpEndTime": 1538778095,
			"vmAgent": "VSAPROXY",
			"name": "TESTVM4_2K8R2",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502d833a-1e21-7e29-13e0-4eb90ae3f4cc",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60227,
			  "clientName": "TESTVM4_2K8R2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 21296357376,
			"bkpStartTime": 1538782852,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 21476305742,
			"vmUsedSpace": 18983419904,
			"subclientId": 211627,
			"bkpEndTime": 1538788828,
			"vmAgent": "VSAPROXY",
			"name": "002_11",
			"vmHardwareVer": "vmx-08",
			"strGUID": "502de268-849b-bdf6-1726-90c570fdf140",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60229,
			  "clientName": "002_11"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 32926969856,
			"bkpStartTime": 1538776377,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 42951614706,
			"vmUsedSpace": 25700597760,
			"subclientId": 211627,
			"bkpEndTime": 1538782840,
			"vmAgent": "VSAPROXY",
			"name": "gkvsavm1",
			"vmHardwareVer": "vmx-08",
			"strGUID": "502d99fd-06d7-154f-27f1-540b18be48c6",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60231,
			  "clientName": "gkvsavm1"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 9252634624,
			"bkpStartTime": 1538783397,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 60135042876,
			"vmUsedSpace": 9252634624,
			"subclientId": 211627,
			"bkpEndTime": 1538786576,
			"vmAgent": "VSAPROXY",
			"name": "FBR_EXT4DEMO",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502dfb8d-e178-76ad-8455-e774ce406d08",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60232,
			  "clientName": "FBR_EXT4DEMO"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 57019465728,
			"bkpStartTime": 1538776369,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 77310407268,
			"vmUsedSpace": 57019465728,
			"subclientId": 211627,
			"bkpEndTime": 1538793635,
			"vmAgent": "VSAPROXY",
			"name": "Phoneapps_FBR",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502d2c97-0456-9fe0-bcb4-33cf9f066b90",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60233,
			  "clientName": "Phoneapps_FBR"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 15767437312,
			"bkpStartTime": 1538790697,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 85899842609,
			"vmUsedSpace": 15767437312,
			"subclientId": 211627,
			"bkpEndTime": 1538794684,
			"vmAgent": "VSAPROXY",
			"name": "FBRPristine",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502fc161-fc76-b596-e4d3-dd0a58410ff2",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60234,
			  "clientName": "FBRPristine"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 13630599168,
			"bkpStartTime": 1538771811,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 107374195386,
			"vmUsedSpace": 10786701312,
			"subclientId": 211627,
			"bkpEndTime": 1538773587,
			"vmAgent": "VSAPROXY",
			"name": "ThinAppVM",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502d259c-68c7-a384-20bf-b662bf257fae",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60235,
			  "clientName": "ThinAppVM"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 15767437312,
			"bkpStartTime": 1538785912,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 85899629896,
			"vmUsedSpace": 15767437312,
			"subclientId": 211627,
			"bkpEndTime": 1538789988,
			"vmAgent": "VSAPROXY",
			"name": "FBROvf.ova",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502f5034-ff69-758b-6e52-4e1653ed28e8",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60236,
			  "clientName": "FBROvf.ova"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 211668877312,
			"bkpStartTime": 1538776366,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 536871611429,
			"vmUsedSpace": 226438938624,
			"subclientId": 211627,
			"bkpEndTime": 1538814012,
			"vmAgent": "VSAPROXY",
			"name": "vPodVM",
			"vmHardwareVer": "vmx-09",
			"strGUID": "50115d47-8475-bafa-1656-974ff27562bf",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60237,
			  "clientName": "vPodVM"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 16207839232,
			"bkpStartTime": 1538793697,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 85900681860,
			"vmUsedSpace": 16207839232,
			"subclientId": 211627,
			"bkpEndTime": 1538803723,
			"vmAgent": "VSAPROXY",
			"name": "FBRnew.ova",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502ff281-564a-1c9f-3891-22ce426c3735",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60238,
			  "clientName": "FBRnew.ova"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 9457250304,
			"bkpStartTime": 1538782527,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 64428049756,
			"vmUsedSpace": 7559184384,
			"subclientId": 211627,
			"bkpEndTime": 1538785274,
			"vmAgent": "VSAPROXY",
			"name": "TESTVM2",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502dceec-2387-dc4d-6b59-ec4013d3052f",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60239,
			  "clientName": "TESTVM2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 2870022144,
			"bkpStartTime": 1538784201,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Other Linux (32-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 8591430324,
			"vmUsedSpace": 7121928192,
			"subclientId": 211627,
			"bkpEndTime": 1538785629,
			"vmAgent": "VSAPROXY",
			"name": "UbuntuVM-solidfire-rest",
			"vmHardwareVer": "vmx-08",
			"strGUID": "502f19b2-18cd-aa38-8732-19e6be5bc6ed",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60240,
			  "clientName": "UbuntuVM-solidfire-rest"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 26470612992,
			"bkpStartTime": 1538776368,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 107376504581,
			"vmUsedSpace": 30568087552,
			"subclientId": 211627,
			"bkpEndTime": 1538797760,
			"vmAgent": "VSAPROXY",
			"name": "VSP44",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502d4977-8b90-ec9e-03a7-b59cedb46d7a",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60241,
			  "clientName": "VSP44"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 11928395776,
			"bkpStartTime": 1538787583,
			"type": 9,
			"vmStatus": 4,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 107374505041,
			"vmUsedSpace": 10818158592,
			"subclientId": 211627,
			"bkpEndTime": 1538790924,
			"vmAgent": "VSAPROXY",
			"name": "TEGILE-win2012r2",
			"vmHardwareVer": "vmx-08",
			"strGUID": "502f7402-0909-a4ba-5155-e89a1bdd0030",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60242,
			  "clientName": "TEGILE-win2012r2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 132120576,
			"bkpStartTime": 1538771811,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Other 2.4.x Linux (32-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 137286128,
			"vmUsedSpace": 132120576,
			"subclientId": 211627,
			"bkpEndTime": 1538772041,
			"vmAgent": "VSAPROXY",
			"name": "small_demo-src",
			"vmHardwareVer": "vmx-08",
			"strGUID": "502d07be-89a2-a66f-fcc8-44c5da2dbe3b",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60243,
			  "clientName": "small_demo-src"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 0,
			"bkpStartTime": 1538790254,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 128849021773,
			"vmUsedSpace": 0,
			"subclientId": 211627,
			"bkpEndTime": 1538790298,
			"vmAgent": "VSAPROXY",
			"name": "OvaSource",
			"vmHardwareVer": "vmx-11",
			"strGUID": "502fac20-a49a-7e44-e9b8-93f03dd5896c",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60244,
			  "clientName": "OvaSource"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 11901071360,
			"bkpStartTime": 1538780004,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": true,
			"vendor": 1,
			"osType": 1,
			"vmSize": 107375247312,
			"vmUsedSpace": 10734272512,
			"subclientId": 211627,
			"bkpEndTime": 1538783518,
			"vmAgent": "VSAPROXY",
			"name": "vsa_domain",
			"vmHardwareVer": "vmx-08",
			"strGUID": "502dabef-8800-a6f6-c9ce-60d98e7f5e08",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60245,
			  "clientName": "vsa_domain"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 7319060480,
			"bkpStartTime": 1538786850,
			"type": 9,
			"vmStatus": 4,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 37582036137,
			"vmUsedSpace": 7319060480,
			"subclientId": 211627,
			"bkpEndTime": 1538789731,
			"vmAgent": "VSAPROXY",
			"name": "TEGILE-RHEL7",
			"vmHardwareVer": "vmx-10",
			"strGUID": "502f7003-c3cf-8808-4be3-3cb3f53c90b6",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60246,
			  "clientName": "TEGILE-RHEL7"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 2132867072,
			"bkpStartTime": 1538771811,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2003 (64-bit)",
			"isDeleted": true,
			"vendor": 1,
			"osType": 1,
			"vmSize": 2148747861,
			"vmUsedSpace": 2135949312,
			"subclientId": 211627,
			"bkpEndTime": 1538772282,
			"vmAgent": "VSAPROXY",
			"name": "wf_Poweredoff",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502d4e18-1c0e-bbc3-a045-bc62573757ee",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60247,
			  "clientName": "wf_Poweredoff"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 2130034688,
			"bkpStartTime": 1538782598,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2003 (64-bit)",
			"isDeleted": true,
			"vendor": 1,
			"osType": 1,
			"vmSize": 2147675883,
			"vmUsedSpace": 2135949312,
			"subclientId": 211627,
			"bkpEndTime": 1538783392,
			"vmAgent": "VSAPROXY",
			"name": "wf_NoNetwork",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502dd42e-52a3-2e18-8ddd-f3f21f5ac05f",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60248,
			  "clientName": "wf_NoNetwork"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 11389214720,
			"bkpStartTime": 1538776369,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "SUSE Linux Enterprise 11 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 115967242777,
			"vmUsedSpace": 13205766144,
			"subclientId": 211627,
			"bkpEndTime": 1538782490,
			"vmAgent": "VSAPROXY",
			"name": "vRA7",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502d7f5a-3ebf-e850-df0e-84a50075bb74",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60249,
			  "clientName": "vRA7"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 8348446720,
			"bkpStartTime": 1538794690,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": true,
			"vendor": 1,
			"osType": 1,
			"vmSize": 12885776304,
			"vmUsedSpace": 7472152576,
			"subclientId": 211627,
			"bkpEndTime": 1538796399,
			"vmAgent": "VSAPROXY",
			"name": "shiv-test2",
			"vmHardwareVer": "vmx-07",
			"strGUID": "50396189-8e6e-6ed7-27a2-0800dcdf81e1",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60250,
			  "clientName": "shiv-test2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.131",
			"vmGuestSpace": 11589636096,
			"bkpStartTime": 1538782584,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 21475876292,
			"vmUsedSpace": 7554990080,
			"subclientId": 211627,
			"bkpEndTime": 1538785312,
			"vmAgent": "VSAPROXY",
			"name": "ViewVM_3",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502dd254-6ad5-3b3f-74d5-d4ccbba37212",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60251,
			  "clientName": "ViewVM_3"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.65",
			"vmGuestSpace": 114489819136,
			"bkpStartTime": 1538776369,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "SUSE Linux Enterprise 11 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 285615651829,
			"vmUsedSpace": 114489819136,
			"subclientId": 211627,
			"bkpEndTime": 1538801085,
			"vmAgent": "VSAPROXY",
			"name": "vRealize Operations Manager Appliance",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502d5aeb-fcbf-7b19-7302-3c100a0965f1",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60252,
			  "clientName": "vRealize Operations Manager Appliance"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.65",
			"vmGuestSpace": 28602961920,
			"bkpStartTime": 1538776369,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 107375430990,
			"vmUsedSpace": 27730640896,
			"subclientId": 211627,
			"bkpEndTime": 1538779999,
			"vmAgent": "VSAPROXY",
			"name": "VSAQA-VC2",
			"vmHardwareVer": "vmx-08",
			"strGUID": "502d159e-e285-6c16-6bd6-832ef4bb2335",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60253,
			  "clientName": "VSAQA-VC2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.65",
			"vmGuestSpace": 53883670528,
			"bkpStartTime": 1538776368,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 171799654832,
			"vmUsedSpace": 147338559488,
			"subclientId": 211627,
			"bkpEndTime": 1538802677,
			"vmAgent": "VSAPROXY",
			"name": "VSAQA-RemoteVC",
			"vmHardwareVer": "vmx-09",
			"strGUID": "5011e9bb-1cca-02a3-7c71-faa76325ca3d",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60254,
			  "clientName": "VSAQA-RemoteVC"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.19.108.65",
			"vmGuestSpace": 102280056832,
			"bkpStartTime": 1538795140,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2008 R2 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 107375184557,
			"vmUsedSpace": 93143957504,
			"subclientId": 211627,
			"bkpEndTime": 1538805451,
			"vmAgent": "VSAPROXY",
			"name": "VSAQA-VC",
			"vmHardwareVer": "vmx-09",
			"strGUID": "503acbc7-e417-0b8e-ae72-7d9793db32c3",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60255,
			  "clientName": "VSAQA-VC"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 590348288,
			"bkpStartTime": 1538786431,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 85903291650,
			"vmUsedSpace": 590348288,
			"subclientId": 211627,
			"bkpEndTime": 1538786844,
			"vmAgent": "VSAPROXY",
			"name": "CentOS-3par-2",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502f5e34-5d8e-11e9-bce8-b448305339b1",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60256,
			  "clientName": "CentOS-3par-2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 22740467712,
			"bkpStartTime": 1538776368,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 92350172679,
			"vmUsedSpace": 22740467712,
			"subclientId": 211627,
			"bkpEndTime": 1538783956,
			"vmAgent": "VSAPROXY",
			"name": "CENTOS-netapp",
			"vmHardwareVer": "vmx-07",
			"strGUID": "501fb461-1b91-a1fb-ddda-3cc37da40378",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60257,
			  "clientName": "CENTOS-netapp"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 20371734528,
			"bkpStartTime": 1538776369,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 65696732845,
			"vmUsedSpace": 20371734528,
			"subclientId": 211627,
			"bkpEndTime": 1538784462,
			"vmAgent": "VSAPROXY",
			"name": "centos-persistent",
			"vmHardwareVer": "vmx-07",
			"strGUID": "501fde8d-9337-1f07-3336-4b159829ba74",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60258,
			  "clientName": "centos-persistent"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 36486250496,
			"bkpStartTime": 1538776370,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 65695162595,
			"vmUsedSpace": 36486250496,
			"subclientId": 211627,
			"bkpEndTime": 1538788332,
			"vmAgent": "VSAPROXY",
			"name": "centos-persistent-odd-size",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502d879c-f57b-7409-ff72-a90200f626e5",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60259,
			  "clientName": "centos-persistent-odd-size"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 48437919744,
			"bkpStartTime": 1538790249,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 136368289598,
			"vmUsedSpace": 48437919744,
			"subclientId": 211627,
			"bkpEndTime": 1538799933,
			"vmAgent": "VSAPROXY",
			"name": "Centos-vsatestcs-proxy",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502fab2d-5383-3fba-e13c-31253df3821c",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60260,
			  "clientName": "Centos-vsatestcs-proxy"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 50622103552,
			"bkpStartTime": 1538785394,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 53688901662,
			"vmUsedSpace": 50622103552,
			"subclientId": 211627,
			"bkpEndTime": 1538799448,
			"vmAgent": "VSAPROXY",
			"name": "Centos7.5",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502f3006-9d86-4d7e-c563-48f295100fb1",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60261,
			  "clientName": "Centos7.5"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 34532753408,
			"bkpStartTime": 1538785638,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 279175181030,
			"vmUsedSpace": 34532753408,
			"subclientId": 211627,
			"bkpEndTime": 1538795136,
			"vmAgent": "VSAPROXY",
			"name": "FREL-pratik",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502f3e47-d562-41e1-f0b6-15ee8d7e6d70",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60262,
			  "clientName": "FREL-pratik"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 15477657600,
			"bkpStartTime": 1538788693,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 142808170072,
			"vmUsedSpace": 22684893184,
			"subclientId": 211627,
			"bkpEndTime": 1538795779,
			"vmAgent": "VSAPROXY",
			"name": "GEN2VM",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502f9f06-6b5f-829c-7031-1ffddd723ef6",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60263,
			  "clientName": "GEN2VM"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 1924136960,
			"bkpStartTime": 1538785394,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Ubuntu Linux (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 4295375850,
			"vmUsedSpace": 1924136960,
			"subclientId": 211627,
			"bkpEndTime": 1538785957,
			"vmAgent": "VSAPROXY",
			"name": "hitachi_ubuntu",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502f2c45-fd8e-dc2f-1ba9-2fa8d1291e10",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60264,
			  "clientName": "hitachi_ubuntu"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 12283990016,
			"bkpStartTime": 1538794203,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 42950249528,
			"vmUsedSpace": 13669236736,
			"subclientId": 211627,
			"bkpEndTime": 1538798255,
			"vmAgent": "VSAPROXY",
			"name": "hitachi_winvm",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502fff82-0257-5e17-4b8d-eb6c80a9a607",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60265,
			  "clientName": "hitachi_winvm"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 34500247552,
			"bkpStartTime": 1538787614,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 92342383767,
			"vmUsedSpace": 34500247552,
			"subclientId": 211627,
			"bkpEndTime": 1538796762,
			"vmAgent": "VSAPROXY",
			"name": "HMVM-CENT73",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502f7d12-1790-f188-97cb-d60f7fa63437",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60266,
			  "clientName": "HMVM-CENT73"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 18496880640,
			"bkpStartTime": 1538788889,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2003 Standard (32-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 45150849834,
			"vmUsedSpace": 18496880640,
			"subclientId": 211627,
			"bkpEndTime": 1538799347,
			"vmAgent": "VSAPROXY",
			"name": "HMVM-RHEL5",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502fa024-ac63-9ce9-8978-55c060e17133",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60267,
			  "clientName": "HMVM-RHEL5"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 70422364160,
			"bkpStartTime": 1538790307,
			"type": 9,
			"vmStatus": 4,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 6 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 81604624690,
			"vmUsedSpace": 70422364160,
			"subclientId": 211627,
			"bkpEndTime": 1538803726,
			"vmAgent": "VSAPROXY",
			"name": "HMVM-RHEL64",
			"vmHardwareVer": "vmx-08",
			"strGUID": "502fb107-19e0-7f6f-6e8d-f58d9930af69",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60268,
			  "clientName": "HMVM-RHEL64"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 9029287936,
			"bkpStartTime": 1538793294,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 86973452349,
			"vmUsedSpace": 9029287936,
			"subclientId": 211627,
			"bkpEndTime": 1538797016,
			"vmAgent": "VSAPROXY",
			"name": "HMVM-RHEL70",
			"vmHardwareVer": "vmx-10",
			"strGUID": "502ff1c7-a8b0-3b38-c558-fb4c7f54d8db",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60269,
			  "clientName": "HMVM-RHEL70"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 10890510336,
			"bkpStartTime": 1538787200,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 85899597507,
			"vmUsedSpace": 10890510336,
			"subclientId": 211627,
			"bkpEndTime": 1538790245,
			"vmAgent": "VSAPROXY",
			"name": "HMVM-RHEL74",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502f715f-1ff2-1d91-38d1-4d2b5545750a",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60270,
			  "clientName": "HMVM-RHEL74"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 2635071488,
			"bkpStartTime": 1538787739,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "SUSE Linux Enterprise 10 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 32212268032,
			"vmUsedSpace": 2635071488,
			"subclientId": 211627,
			"bkpEndTime": 1538788687,
			"vmAgent": "VSAPROXY",
			"name": "HMVM-SUSE10",
			"vmHardwareVer": "vmx-11",
			"strGUID": "502f8cc0-a5d2-cb4c-6ef9-f5000a88be9a",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60271,
			  "clientName": "HMVM-SUSE10"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 7121928192,
			"bkpStartTime": 1538785521,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Other Linux (32-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 8589946705,
			"vmUsedSpace": 7121928192,
			"subclientId": 211627,
			"bkpEndTime": 1538787602,
			"vmAgent": "VSAPROXY",
			"name": "HMVM-UBUNTU14.04",
			"vmHardwareVer": "vmx-08",
			"strGUID": "502f3648-9d30-23a7-9063-55a211b49b09",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60272,
			  "clientName": "HMVM-UBUNTU14.04"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 1912602624,
			"bkpStartTime": 1538791288,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Ubuntu Linux (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 4295202486,
			"vmUsedSpace": 1912602624,
			"subclientId": 211627,
			"bkpEndTime": 1538791892,
			"vmAgent": "VSAPROXY",
			"name": "HMVM-UBUNTU16.04",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502fcf62-c162-648d-2876-8185d7d87d07",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60273,
			  "clientName": "HMVM-UBUNTU16.04"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 2634022912,
			"bkpStartTime": 1538786190,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "SUSE Linux Enterprise 10 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 32212823032,
			"vmUsedSpace": 2634022912,
			"subclientId": 211627,
			"bkpEndTime": 1538787191,
			"vmAgent": "VSAPROXY",
			"name": "Nimble-SUSE10",
			"vmHardwareVer": "vmx-11",
			"strGUID": "502f5d8f-8fe0-2c74-366a-6c026084773d",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60274,
			  "clientName": "Nimble-SUSE10"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 11898384384,
			"bkpStartTime": 1538789998,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 107374467492,
			"vmUsedSpace": 10729029632,
			"subclientId": 211627,
			"bkpEndTime": 1538793238,
			"vmAgent": "VSAPROXY",
			"name": "proxy-sp8",
			"vmHardwareVer": "vmx-08",
			"strGUID": "502fa44d-d388-1ab9-da7b-d66d8a11f786",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60275,
			  "clientName": "proxy-sp8"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 18496880640,
			"bkpStartTime": 1538791894,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2003 Standard (32-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 45156071390,
			"vmUsedSpace": 18496880640,
			"subclientId": 211627,
			"bkpEndTime": 1538799356,
			"vmAgent": "VSAPROXY",
			"name": "RHEL5-all-odd-size-disks-orig",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502fdbe5-b78c-bcfa-2aa8-2fe3c2eef48d",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60276,
			  "clientName": "RHEL5-all-odd-size-disks-orig"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 18929942528,
			"bkpStartTime": 1538776368,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 5 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 104156335567,
			"vmUsedSpace": 18929942528,
			"subclientId": 211627,
			"bkpEndTime": 1538783581,
			"vmAgent": "VSAPROXY",
			"name": "RHEL5-netapp-DS2",
			"vmHardwareVer": "vmx-09",
			"strGUID": "501f8be4-59bc-a389-0ae1-eb34a75248e2",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60277,
			  "clientName": "RHEL5-netapp-DS2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 21527265280,
			"bkpStartTime": 1538788335,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 5 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 27917733913,
			"vmUsedSpace": 21527265280,
			"subclientId": 211627,
			"bkpEndTime": 1538795291,
			"vmAgent": "VSAPROXY",
			"name": "RHEL5-VNX",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502f9620-1111-b140-e81e-fa90c96e7aa6",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60278,
			  "clientName": "RHEL5-VNX"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 10888413184,
			"bkpStartTime": 1538787045,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 85900257704,
			"vmUsedSpace": 10888413184,
			"subclientId": 211627,
			"bkpEndTime": 1538790351,
			"vmAgent": "VSAPROXY",
			"name": "RHEL7",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502f7028-cc13-f9f4-5ca1-ed147c94eebe",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60279,
			  "clientName": "RHEL7"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 7319060480,
			"bkpStartTime": 1538783959,
			"type": 9,
			"vmStatus": 4,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 37583470871,
			"vmUsedSpace": 7319060480,
			"subclientId": 211627,
			"bkpEndTime": 1538785888,
			"vmAgent": "VSAPROXY",
			"name": "RHEL7-3par-1",
			"vmHardwareVer": "vmx-10",
			"strGUID": "502f1727-8199-4995-7059-39fbb42b11f4",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60280,
			  "clientName": "RHEL7-3par-1"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 19626196992,
			"bkpStartTime": 1538785892,
			"type": 9,
			"vmStatus": 4,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 85899359505,
			"vmUsedSpace": 19626196992,
			"subclientId": 211627,
			"bkpEndTime": 1538791282,
			"vmAgent": "VSAPROXY",
			"name": "RHEV-proxy",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502f431f-300c-0c03-1cf3-17bfb174d471",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60282,
			  "clientName": "RHEV-proxy"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 48261758976,
			"bkpStartTime": 1538776369,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 136370177876,
			"vmUsedSpace": 48261758976,
			"subclientId": 211627,
			"bkpEndTime": 1538790249,
			"vmAgent": "VSAPROXY",
			"name": "Sap-Centos-proxy",
			"vmHardwareVer": "vmx-07",
			"strGUID": "501fecbe-948f-fcd6-cd8c-8e9a0758e320",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60283,
			  "clientName": "Sap-Centos-proxy"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 19777191936,
			"bkpStartTime": 1538776368,
			"type": 9,
			"vmStatus": 4,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 92367870631,
			"vmUsedSpace": 19777191936,
			"subclientId": 211627,
			"bkpEndTime": 1538784176,
			"vmAgent": "VSAPROXY",
			"name": "SapCentos-all-options",
			"vmHardwareVer": "vmx-07",
			"strGUID": "501f363b-6a0e-7b4d-5554-a645bc5457dc",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60284,
			  "clientName": "SapCentos-all-options"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 19771949056,
			"bkpStartTime": 1538776368,
			"type": 9,
			"vmStatus": 4,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 92351059503,
			"vmUsedSpace": 19771949056,
			"subclientId": 211627,
			"bkpEndTime": 1538783513,
			"vmAgent": "VSAPROXY",
			"name": "SapCentos-trial",
			"vmHardwareVer": "vmx-13",
			"strGUID": "501f38eb-7d31-261c-4984-303389a6e582",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60285,
			  "clientName": "SapCentos-trial"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 34500247552,
			"bkpStartTime": 1538776369,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 92365941720,
			"vmUsedSpace": 34500247552,
			"subclientId": 211627,
			"bkpEndTime": 1538786178,
			"vmAgent": "VSAPROXY",
			"name": "SapCentOS7",
			"vmHardwareVer": "vmx-07",
			"strGUID": "5026a51b-57e5-8fba-4814-0f7d5c6cf617",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60286,
			  "clientName": "SapCentOS7"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 9319096320,
			"bkpStartTime": 1538783539,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "CentOS 4/5 or later (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 92364824884,
			"vmUsedSpace": 34500247552,
			"subclientId": 211627,
			"bkpEndTime": 1538793288,
			"vmAgent": "VSAPROXY",
			"name": "SapCentOS7_new",
			"vmHardwareVer": "vmx-07",
			"strGUID": "502f0083-6fcf-9aa7-2eb8-d4f158c1b29a",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60287,
			  "clientName": "SapCentOS7_new"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 18833473536,
			"bkpStartTime": 1538776369,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 6 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 56912891775,
			"vmUsedSpace": 18833473536,
			"subclientId": 211627,
			"bkpEndTime": 1538782501,
			"vmAgent": "VSAPROXY",
			"name": "SapLinuxVM",
			"vmHardwareVer": "vmx-13",
			"strGUID": "5026de5a-d539-3f21-3402-1bcd912dc84e",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60288,
			  "clientName": "SapLinuxVM"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 190527533056,
			"bkpStartTime": 1538788333,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 1176821053244,
			"vmUsedSpace": 205380386816,
			"subclientId": 211627,
			"bkpEndTime": 1538808471,
			"vmAgent": "VSAPROXY",
			"name": "sapna-MA",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502f93b9-75e3-6c95-8a46-7f642cfa56e5",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60289,
			  "clientName": "sapna-MA"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 70422364160,
			"bkpStartTime": 1538776368,
			"type": 9,
			"vmStatus": 4,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 6 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 81606982843,
			"vmUsedSpace": 70422364160,
			"subclientId": 211627,
			"bkpEndTime": 1538795631,
			"vmAgent": "VSAPROXY",
			"name": "SapRHEL6",
			"vmHardwareVer": "vmx-08",
			"strGUID": "501fed28-4791-1e93-174d-fade09dbf57a",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60290,
			  "clientName": "SapRHEL6"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 9029287936,
			"bkpStartTime": 1538776369,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 87002398930,
			"vmUsedSpace": 9029287936,
			"subclientId": 211627,
			"bkpEndTime": 1538781796,
			"vmAgent": "VSAPROXY",
			"name": "SapRHEL7",
			"vmHardwareVer": "vmx-10",
			"strGUID": "5026e059-a0b5-78a6-cacf-3abb708b07da",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60291,
			  "clientName": "SapRHEL7"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 2635071488,
			"bkpStartTime": 1538771811,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "SUSE Linux Enterprise 10 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 32215861876,
			"vmUsedSpace": 2635071488,
			"subclientId": 211627,
			"bkpEndTime": 1538773851,
			"vmAgent": "VSAPROXY",
			"name": "SapSUSE10",
			"vmHardwareVer": "vmx-11",
			"strGUID": "5026db70-30bf-8a24-8513-6426291be0d1",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60292,
			  "clientName": "SapSUSE10"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 694157312,
			"bkpStartTime": 1538790361,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 138513569064,
			"vmUsedSpace": 694157312,
			"subclientId": 211627,
			"bkpEndTime": 1538790692,
			"vmAgent": "VSAPROXY",
			"name": "SOLIDFIREVM1",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502fbb17-574e-9bb2-f047-935d8d74a8ae",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60293,
			  "clientName": "SOLIDFIREVM1"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 80621568,
			"bkpStartTime": 1538785994,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 138514077436,
			"vmUsedSpace": 3698327552,
			"subclientId": 211627,
			"bkpEndTime": 1538787573,
			"vmAgent": "VSAPROXY",
			"name": "SOLIDFIREVM2",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502f5465-c8ef-8a63-18d5-349aaf2970d9",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60294,
			  "clientName": "SOLIDFIREVM2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 21870133248,
			"bkpStartTime": 1538792958,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 322123839260,
			"vmUsedSpace": 31070355456,
			"subclientId": 211627,
			"bkpEndTime": 1538800648,
			"vmAgent": "VSAPROXY",
			"name": "sqlappaware",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502fe55f-2177-f6bc-d784-e05d3e6088c8",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60295,
			  "clientName": "sqlappaware"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 2634022912,
			"bkpStartTime": 1538786188,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "SUSE Linux Enterprise 10 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 32212955975,
			"vmUsedSpace": 2634022912,
			"subclientId": 211627,
			"bkpEndTime": 1538787036,
			"vmAgent": "VSAPROXY",
			"name": "SUSE10-VMX",
			"vmHardwareVer": "vmx-11",
			"strGUID": "502f59c9-d25d-422b-b7b3-389074a48e48",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60296,
			  "clientName": "SUSE10-VMX"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 1928331264,
			"bkpStartTime": 1538785636,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Ubuntu Linux (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 4296401400,
			"vmUsedSpace": 1928331264,
			"subclientId": 211627,
			"bkpEndTime": 1538786178,
			"vmAgent": "VSAPROXY",
			"name": "Unix-VM",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502f3b5f-b97a-1a65-36f9-d91cb6024cfc",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60297,
			  "clientName": "Unix-VM"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 8107712512,
			"bkpStartTime": 1538784183,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 136366391771,
			"vmUsedSpace": 7881097216,
			"subclientId": 211627,
			"bkpEndTime": 1538786425,
			"vmAgent": "VSAPROXY",
			"name": "VM1",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502f1745-d892-9064-f70c-c0e2099bc7e7",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60298,
			  "clientName": "VM1"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 3667918848,
			"bkpStartTime": 1538793245,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Red Hat Enterprise Linux 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 136366759667,
			"vmUsedSpace": 3667918848,
			"subclientId": 211627,
			"bkpEndTime": 1538794199,
			"vmAgent": "VSAPROXY",
			"name": "VM2",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502fef0f-fa3c-5290-5508-65fe98006560",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60299,
			  "clientName": "VM2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 34231812096,
			"bkpStartTime": 1538776368,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "SUSE Linux Enterprise 11 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 2,
			"vmSize": 294205507206,
			"vmUsedSpace": 34231812096,
			"subclientId": 211627,
			"bkpEndTime": 1538785518,
			"vmAgent": "VSAPROXY",
			"name": "vRealize-Operations-Manager-Appliance-6.6.0.5707161_OVF10",
			"vmHardwareVer": "vmx-08",
			"strGUID": "501f2033-b3ce-540e-324a-525e10322eb5",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60300,
			  "clientName": "vRealize-Operations-Manager-Appliance-6.6.0.5707161_OVF10"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 21109743616,
			"bkpStartTime": 1538776368,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 42954958858,
			"vmUsedSpace": 16809721856,
			"subclientId": 211627,
			"bkpEndTime": 1538782467,
			"vmAgent": "VSAPROXY",
			"name": "VSAProxy3",
			"vmHardwareVer": "vmx-09",
			"strGUID": "5026c05f-c88f-99bf-fa20-cc119a07ea6c",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60301,
			  "clientName": "VSAProxy3"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 22527553536,
			"bkpStartTime": 1538791002,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 42952765448,
			"vmUsedSpace": 26563575808,
			"subclientId": 211627,
			"bkpEndTime": 1538798787,
			"vmAgent": "VSAPROXY",
			"name": "VSAProxy4",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502fc2ca-705d-e182-3d49-10c0c0c1085f",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60302,
			  "clientName": "VSAProxy4"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 11593334784,
			"bkpStartTime": 1538783593,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows 7 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 21475205737,
			"vmUsedSpace": 7545552896,
			"subclientId": 211627,
			"bkpEndTime": 1538785635,
			"vmAgent": "VSAPROXY",
			"name": "win-VM",
			"vmHardwareVer": "vmx-09",
			"strGUID": "502f06ae-43d7-b1b1-d514-9264ed0ef951",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60303,
			  "clientName": "win-VM"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 17381748736,
			"bkpStartTime": 1538785963,
			"type": 9,
			"vmStatus": 4,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 103081832087,
			"vmUsedSpace": 24118296576,
			"subclientId": 211627,
			"bkpEndTime": 1538792950,
			"vmAgent": "VSAPROXY",
			"name": "win-vmx",
			"vmHardwareVer": "vmx-13",
			"strGUID": "502f5106-d86d-5de4-023b-74c8597cec84",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60304,
			  "clientName": "win-vmx"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 0,
			"bkpStartTime": 1538771810,
			"type": 9,
			"vmStatus": 4,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 91273686664,
			"vmUsedSpace": 0,
			"subclientId": 211627,
			"bkpEndTime": 1538772851,
			"vmAgent": "VSAPROXY",
			"name": "winproxy-vmware",
			"vmHardwareVer": "vmx-08",
			"strGUID": "501f0bcf-bbdf-6acf-2a22-f8a70577f232",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60305,
			  "clientName": "winproxy-vmware"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "172.24.40.60",
			"vmGuestSpace": 37141880832,
			"bkpStartTime": 1538776333,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19419003,
			"strOSName": "Microsoft Windows Server 2012 (64-bit)",
			"isDeleted": false,
			"vendor": 1,
			"osType": 1,
			"vmSize": 188982977543,
			"vmUsedSpace": 25494028288,
			"subclientId": 211627,
			"bkpEndTime": 1538802946,
			"vmAgent": "VSAPROXY",
			"name": "WinProxyUnixVINODLR",
			"vmHardwareVer": "vmx-13",
			"strGUID": "09c01f50-2704-7dba-0a87-f923569ae3b1",
			"subclientName": "VSA-TEST",
			"client": {
			  "clientId": 60306,
			  "clientName": "WinProxyUnixVINODLR"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60208,
			  "clientName": "VSAQAVCDR",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "Sap2",
			"vmGuestSpace": 137438958202,
			"bkpStartTime": 1538773836,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 137438953472,
			"vmUsedSpace": 137438958202,
			"subclientId": 211626,
			"bkpEndTime": 1538773856,
			"vmAgent": "VSAPROXY",
			"name": "sab-unmanaged",
			"vmHardwareVer": "",
			"strGUID": "85e7d651-678b-419b-b104-edc99ea150e8",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60308,
			  "clientName": "sab-unmanaged"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "Sapna",
			"vmGuestSpace": 38654709761,
			"bkpStartTime": 1538774161,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Linux",
			"isDeleted": false,
			"vendor": 7,
			"osType": 2,
			"vmSize": 38654705664,
			"vmUsedSpace": 38654709761,
			"subclientId": 211626,
			"bkpEndTime": 1538774178,
			"vmAgent": "VSAPROXY",
			"name": "sap-centos-LS",
			"vmHardwareVer": "",
			"strGUID": "f2bec3c5-5705-480f-a6ee-5f171d37737d",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60309,
			  "clientName": "sap-centos-LS"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "Sapna",
			"vmGuestSpace": 42949676832,
			"bkpStartTime": 1538773885,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Linux",
			"isDeleted": false,
			"vendor": 7,
			"osType": 2,
			"vmSize": 42949672960,
			"vmUsedSpace": 42949676832,
			"subclientId": 211626,
			"bkpEndTime": 1538773901,
			"vmAgent": "VSAPROXY",
			"name": "sap-centos-r",
			"vmHardwareVer": "",
			"strGUID": "941fd262-cd41-49d5-8210-8df2f8ed5d71",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60310,
			  "clientName": "sap-centos-r"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "Sapna",
			"vmGuestSpace": 38654710835,
			"bkpStartTime": 1538773443,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Linux",
			"isDeleted": false,
			"vendor": 7,
			"osType": 2,
			"vmSize": 38654705664,
			"vmUsedSpace": 38654710835,
			"subclientId": 211626,
			"bkpEndTime": 1538773472,
			"vmAgent": "VSAPROXY",
			"name": "sap-centos-unmanaged",
			"vmHardwareVer": "",
			"strGUID": "2c5eb418-75b6-4298-955f-3e238d8c26af",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60311,
			  "clientName": "sap-centos-unmanaged"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "Sapna",
			"vmGuestSpace": 44023418863,
			"bkpStartTime": 1538773706,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Linux",
			"isDeleted": false,
			"vendor": 7,
			"osType": 2,
			"vmSize": 44023414784,
			"vmUsedSpace": 44023418863,
			"subclientId": 211626,
			"bkpEndTime": 1538773723,
			"vmAgent": "VSAPROXY",
			"name": "sap-rhel-LS",
			"vmHardwareVer": "",
			"strGUID": "4c4fe92c-a35f-442b-8a5a-db80d88210e3",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60312,
			  "clientName": "sap-rhel-LS"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "Sapna",
			"vmGuestSpace": 44023418983,
			"bkpStartTime": 1538774035,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Linux",
			"isDeleted": false,
			"vendor": 7,
			"osType": 2,
			"vmSize": 44023414784,
			"vmUsedSpace": 44023418983,
			"subclientId": 211626,
			"bkpEndTime": 1538774051,
			"vmAgent": "VSAPROXY",
			"name": "sap-rhel-premium-fb",
			"vmHardwareVer": "",
			"strGUID": "c1b0eaca-fc64-4ea1-9a67-216d67348474",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60313,
			  "clientName": "sap-rhel-premium-fb"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "Sapna",
			"vmGuestSpace": 4837,
			"bkpStartTime": 1538774051,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Linux",
			"isDeleted": false,
			"vendor": 7,
			"osType": 2,
			"vmSize": 0,
			"vmUsedSpace": 4837,
			"subclientId": 211626,
			"bkpEndTime": 1538774067,
			"vmAgent": "VSAPROXY",
			"name": "sap-rhel-premium-managed",
			"vmHardwareVer": "",
			"strGUID": "c3d42e47-b66c-41ba-936d-9f43aaa482b4",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60314,
			  "clientName": "sap-rhel-premium-managed"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "Sapna",
			"vmGuestSpace": 44023418463,
			"bkpStartTime": 1538773514,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Linux",
			"isDeleted": false,
			"vendor": 7,
			"osType": 2,
			"vmSize": 44023414784,
			"vmUsedSpace": 44023418463,
			"subclientId": 211626,
			"bkpEndTime": 1538773542,
			"vmAgent": "VSAPROXY",
			"name": "sap-rhel-r",
			"vmHardwareVer": "",
			"strGUID": "34a82e76-8ee9-41e9-99eb-e047cac112f0",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60315,
			  "clientName": "sap-rhel-r"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "harshars1",
			"vmGuestSpace": 3914,
			"bkpStartTime": 1538773749,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 0,
			"vmUsedSpace": 3914,
			"subclientId": 211626,
			"bkpEndTime": 1538773776,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM1-2k12R2",
			"vmHardwareVer": "",
			"strGUID": "6b491628-70b2-487f-a45e-61148c218e51",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60316,
			  "clientName": "SmallVM1-2k12R2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214980,
			"bkpStartTime": 1538773776,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": true,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214980,
			"subclientId": 211626,
			"bkpEndTime": 1538773803,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM10",
			"vmHardwareVer": "",
			"strGUID": "6b79977a-6241-45c4-9e5e-7e232ecac4fa",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60317,
			  "clientName": "SmallVM10"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "CI-SOLR-RG",
			"vmGuestSpace": 136365215594,
			"bkpStartTime": 1538774141,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365215594,
			"subclientId": 211626,
			"bkpEndTime": 1538774145,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM12",
			"vmHardwareVer": "",
			"strGUID": "e93f421e-e116-4cb1-9335-53eb225e4a58",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60318,
			  "clientName": "SmallVM12"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365215588,
			"bkpStartTime": 1538773803,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365215588,
			"subclientId": 211626,
			"bkpEndTime": 1538773819,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM12",
			"vmHardwareVer": "",
			"strGUID": "7e6210b7-d37b-4387-b64f-451eca864c52",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60319,
			  "clientName": "SmallVM12"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214952,
			"bkpStartTime": 1538773933,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214952,
			"subclientId": 211626,
			"bkpEndTime": 1538773948,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM13",
			"vmHardwareVer": "",
			"strGUID": "a53f0787-6fcc-4f45-a805-3181327b112d",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60320,
			  "clientName": "SmallVM13"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214952,
			"bkpStartTime": 1538773476,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214952,
			"subclientId": 211626,
			"bkpEndTime": 1538773504,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM14",
			"vmHardwareVer": "",
			"strGUID": "309c38f4-76ee-47b3-99bb-315bc4d23bed",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60321,
			  "clientName": "SmallVM14"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214967,
			"bkpStartTime": 1538773820,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214967,
			"subclientId": 211626,
			"bkpEndTime": 1538773836,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM17",
			"vmHardwareVer": "",
			"strGUID": "84953227-7906-467a-bb9b-23c42af7fb6f",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60322,
			  "clientName": "SmallVM17"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 147102634041,
			"bkpStartTime": 1538774242,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 147102629888,
			"vmUsedSpace": 147102634041,
			"subclientId": 211626,
			"bkpEndTime": 1538774271,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM18",
			"vmHardwareVer": "",
			"strGUID": "fe027cc7-bb22-4d15-9172-40133705356b",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60323,
			  "clientName": "SmallVM18"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214967,
			"bkpStartTime": 1538773620,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214967,
			"subclientId": 211626,
			"bkpEndTime": 1538773635,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM19",
			"vmHardwareVer": "",
			"strGUID": "3d62f7cd-b35b-4b96-9215-9c2097c17391",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60324,
			  "clientName": "SmallVM19"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "harshars1",
			"vmGuestSpace": 136365215659,
			"bkpStartTime": 1538773587,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365215659,
			"subclientId": 211626,
			"bkpEndTime": 1538773602,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM2-2k12R2",
			"vmHardwareVer": "",
			"strGUID": "36816460-bc0c-410e-9ed3-b2aae4e25fe6",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60325,
			  "clientName": "SmallVM2-2k12R2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214966,
			"bkpStartTime": 1538773333,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214966,
			"subclientId": 211626,
			"bkpEndTime": 1538773361,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM20",
			"vmHardwareVer": "",
			"strGUID": "0a8ab521-b183-4767-b48d-b0f9e2fa520f",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60326,
			  "clientName": "SmallVM20"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 3251,
			"bkpStartTime": 1538773984,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": true,
			"vendor": 7,
			"osType": 1,
			"vmSize": 0,
			"vmUsedSpace": 3251,
			"subclientId": 211626,
			"bkpEndTime": 1538774000,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM21MD",
			"vmHardwareVer": "",
			"strGUID": "bb818335-7377-410c-a8e7-197a6606d053",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60327,
			  "clientName": "SmallVM21MD"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214967,
			"bkpStartTime": 1538774002,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214967,
			"subclientId": 211626,
			"bkpEndTime": 1538774018,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM22",
			"vmHardwareVer": "",
			"strGUID": "bbeaf16e-ea95-4351-a559-b44b0646044e",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60328,
			  "clientName": "SmallVM22"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214967,
			"bkpStartTime": 1538773966,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214967,
			"subclientId": 211626,
			"bkpEndTime": 1538773982,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM23",
			"vmHardwareVer": "",
			"strGUID": "b41a6f2b-ed7d-4d2e-b71e-1b0a8ed67f6f",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60329,
			  "clientName": "SmallVM23"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214967,
			"bkpStartTime": 1538773604,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214967,
			"subclientId": 211626,
			"bkpEndTime": 1538773619,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM24",
			"vmHardwareVer": "",
			"strGUID": "39df5da0-cdfd-464d-b284-fbe93f10e534",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60330,
			  "clientName": "SmallVM24"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214967,
			"bkpStartTime": 1538773636,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214967,
			"subclientId": 211626,
			"bkpEndTime": 1538773663,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM25",
			"vmHardwareVer": "",
			"strGUID": "3eb77867-4514-4db6-afa0-2ffec45faef9",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60331,
			  "clientName": "SmallVM25"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 3251,
			"bkpStartTime": 1538774098,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 0,
			"vmUsedSpace": 3251,
			"subclientId": 211626,
			"bkpEndTime": 1538774114,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM26MD",
			"vmHardwareVer": "",
			"strGUID": "ddc428c5-b67a-4c27-80c9-334d1fd7c7c2",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60332,
			  "clientName": "SmallVM26MD"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214967,
			"bkpStartTime": 1538773088,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214967,
			"subclientId": 211626,
			"bkpEndTime": 1538773114,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM27",
			"vmHardwareVer": "",
			"strGUID": "074ebad4-d2a4-416b-b14c-4a59aad89afe",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60333,
			  "clientName": "SmallVM27"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214967,
			"bkpStartTime": 1538773949,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214967,
			"subclientId": 211626,
			"bkpEndTime": 1538773964,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM28",
			"vmHardwareVer": "",
			"strGUID": "a65712a9-a941-4659-880b-8aaba85baa4c",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60334,
			  "clientName": "SmallVM28"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214967,
			"bkpStartTime": 1538774145,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214967,
			"subclientId": 211626,
			"bkpEndTime": 1538774161,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM29",
			"vmHardwareVer": "",
			"strGUID": "f29b3192-8fea-4828-a4bb-d0fb20c0e4ae",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60335,
			  "clientName": "SmallVM29"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "harshars1",
			"vmGuestSpace": 136365215017,
			"bkpStartTime": 1538773392,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365215017,
			"subclientId": 211626,
			"bkpEndTime": 1538773429,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM3-2k12R2",
			"vmHardwareVer": "",
			"strGUID": "10812501-973e-4c81-aaf1-318bc1d14b17",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60336,
			  "clientName": "SmallVM3-2k12R2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214967,
			"bkpStartTime": 1538774083,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214967,
			"subclientId": 211626,
			"bkpEndTime": 1538774098,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM30",
			"vmHardwareVer": "",
			"strGUID": "c5bce3c2-b0ab-4775-b367-94e45c01c1f4",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60337,
			  "clientName": "SmallVM30"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "harshars1",
			"vmGuestSpace": 136365214940,
			"bkpStartTime": 1538773856,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214940,
			"subclientId": 211626,
			"bkpEndTime": 1538773871,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM4",
			"vmHardwareVer": "",
			"strGUID": "8a2ed21b-0b55-438a-b905-869fce94397b",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60338,
			  "clientName": "SmallVM4"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "harshars1",
			"vmGuestSpace": 136365215575,
			"bkpStartTime": 1538774019,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365215575,
			"subclientId": 211626,
			"bkpEndTime": 1538774034,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM5",
			"vmHardwareVer": "",
			"strGUID": "c005fe52-cacb-4f9f-b4cd-47353b101641",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60339,
			  "clientName": "SmallVM5"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 3240,
			"bkpStartTime": 1538774114,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 0,
			"vmUsedSpace": 3240,
			"subclientId": 211626,
			"bkpEndTime": 1538774130,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM6MD",
			"vmHardwareVer": "",
			"strGUID": "df44637b-2846-472f-ad15-44688c8bed9e",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60340,
			  "clientName": "SmallVM6MD"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214955,
			"bkpStartTime": 1538773902,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214955,
			"subclientId": 211626,
			"bkpEndTime": 1538773917,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM7",
			"vmHardwareVer": "",
			"strGUID": "9923bb2e-a7b8-4fcb-aeb1-f3160751b320",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60341,
			  "clientName": "SmallVM7"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214955,
			"bkpStartTime": 1538774067,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214955,
			"subclientId": 211626,
			"bkpEndTime": 1538774082,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM8",
			"vmHardwareVer": "",
			"strGUID": "c4078794-655c-447c-a94f-3d944928c9e3",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60342,
			  "clientName": "SmallVM8"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "SolarEastUS",
			"vmGuestSpace": 136365214955,
			"bkpStartTime": 1538773918,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": false,
			"vendor": 7,
			"osType": 1,
			"vmSize": 136365211648,
			"vmUsedSpace": 136365214955,
			"subclientId": 211626,
			"bkpEndTime": 1538773933,
			"vmAgent": "VSAPROXY",
			"name": "SmallVM9",
			"vmHardwareVer": "",
			"strGUID": "a1e752bc-f103-4efe-8aed-3d2e7e538383",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60343,
			  "clientName": "SmallVM9"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "AzureRMEastUS",
			"vmGuestSpace": 34359754677,
			"bkpStartTime": 1538773724,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Linux",
			"isDeleted": true,
			"vendor": 7,
			"osType": 2,
			"vmSize": 34359738368,
			"vmUsedSpace": 34359754677,
			"subclientId": 211626,
			"bkpEndTime": 1538773749,
			"vmAgent": "VSAPROXY",
			"name": "snapVM5",
			"vmHardwareVer": "",
			"strGUID": "5876f620-b456-4da1-ab1a-e8ebd31ab1ad",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60344,
			  "clientName": "snapVM5"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "AzureRMEastUS",
			"vmGuestSpace": 4848,
			"bkpStartTime": 1538773285,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": true,
			"vendor": 7,
			"osType": 1,
			"vmSize": 0,
			"vmUsedSpace": 4848,
			"subclientId": 211626,
			"bkpEndTime": 1538773333,
			"vmAgent": "VSAPROXY",
			"name": "VM1-Win2k16MDi",
			"vmHardwareVer": "",
			"strGUID": "0a16c5b3-0dc7-477b-85a0-efdfcc03e1f9",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60345,
			  "clientName": "VM1-Win2k16MDi"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "AzureRMNorthCentralUS",
			"vmGuestSpace": 4832,
			"bkpStartTime": 1538773871,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": true,
			"vendor": 7,
			"osType": 1,
			"vmSize": 0,
			"vmUsedSpace": 4832,
			"subclientId": 211626,
			"bkpEndTime": 1538773885,
			"vmAgent": "VSAPROXY",
			"name": "VM11-W2k16MDiP",
			"vmHardwareVer": "",
			"strGUID": "8b058eaa-8e48-49bc-9af2-6988dad854f0",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60346,
			  "clientName": "VM11-W2k16MDiP"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "AzureRMNorthCentralUS",
			"vmGuestSpace": 34359743188,
			"bkpStartTime": 1538773429,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Linux",
			"isDeleted": true,
			"vendor": 7,
			"osType": 2,
			"vmSize": 34359738368,
			"vmUsedSpace": 34359743188,
			"subclientId": 211626,
			"bkpEndTime": 1538773443,
			"vmAgent": "VSAPROXY",
			"name": "VM12-RHEL74i",
			"vmHardwareVer": "",
			"strGUID": "20d991f5-fd7c-412c-a74e-572ab3fa39b7",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60347,
			  "clientName": "VM12-RHEL74i"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "AzureRMEastUS",
			"vmGuestSpace": 4328,
			"bkpStartTime": 1538774206,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Linux",
			"isDeleted": true,
			"vendor": 7,
			"osType": 2,
			"vmSize": 0,
			"vmUsedSpace": 4328,
			"subclientId": 211626,
			"bkpEndTime": 1538774222,
			"vmAgent": "VSAPROXY",
			"name": "VM3-RHEL74MDi",
			"vmHardwareVer": "",
			"strGUID": "f6f300a7-c429-4f2b-8ee6-1e4280a555d3",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60348,
			  "clientName": "VM3-RHEL74MDi"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "AzureRMEastUS",
			"vmGuestSpace": 152471343636,
			"bkpStartTime": 1538773543,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": true,
			"vendor": 7,
			"osType": 1,
			"vmSize": 152471339008,
			"vmUsedSpace": 152471343636,
			"subclientId": 211626,
			"bkpEndTime": 1538773587,
			"vmAgent": "VSAPROXY",
			"name": "VM4-Win2k16iP",
			"vmHardwareVer": "",
			"strGUID": "34e1517d-0c86-420d-af21-f1ba413b4dc5",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60349,
			  "clientName": "VM4-Win2k16iP"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "AzureRMEastUS",
			"vmGuestSpace": 34359743099,
			"bkpStartTime": 1538774228,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Linux",
			"isDeleted": true,
			"vendor": 7,
			"osType": 2,
			"vmSize": 34359738368,
			"vmUsedSpace": 34359743099,
			"subclientId": 211626,
			"bkpEndTime": 1538774241,
			"vmAgent": "VSAPROXY",
			"name": "VM5-RHEL74i",
			"vmHardwareVer": "",
			"strGUID": "fb5a1148-8f2e-49ce-a390-1b5ef46fc2cb",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60350,
			  "clientName": "VM5-RHEL74i"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "AzureRMEastUS",
			"vmGuestSpace": 4749,
			"bkpStartTime": 1538773115,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Windows",
			"isDeleted": true,
			"vendor": 7,
			"osType": 1,
			"vmSize": 0,
			"vmUsedSpace": 4749,
			"subclientId": 211626,
			"bkpEndTime": 1538773285,
			"vmAgent": "VSAPROXY",
			"name": "VM6-Win2k16MDiP",
			"vmHardwareVer": "",
			"strGUID": "095d8fbb-47e5-42c0-b43c-4fd9bf33ab0c",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60351,
			  "clientName": "VM6-Win2k16MDiP"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "AzureRMEastUS",
			"vmGuestSpace": 4695,
			"bkpStartTime": 1538773664,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Linux",
			"isDeleted": true,
			"vendor": 7,
			"osType": 2,
			"vmSize": 0,
			"vmUsedSpace": 4695,
			"subclientId": 211626,
			"bkpEndTime": 1538773706,
			"vmAgent": "VSAPROXY",
			"name": "VM7-Ub1804MDiS",
			"vmHardwareVer": "",
			"strGUID": "477361d1-f1c8-46c2-bda4-91b59e52c92e",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60352,
			  "clientName": "VM7-Ub1804MDiS"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "AzureRMEastUS",
			"vmGuestSpace": 5015,
			"bkpStartTime": 1538773361,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19418998,
			"strOSName": "Linux",
			"isDeleted": true,
			"vendor": 7,
			"osType": 2,
			"vmSize": 0,
			"vmUsedSpace": 5015,
			"subclientId": 211626,
			"bkpEndTime": 1538773387,
			"vmAgent": "VSAPROXY",
			"name": "VM8-RHEL74MDPi567890-0123456789-0123456789-0123456789-0123456789",
			"vmHardwareVer": "",
			"strGUID": "0d5e6f8a-7338-4e9d-8737-c2d680013c6c",
			"subclientName": "VSA-STREAMING",
			"client": {
			  "clientId": 60353,
			  "clientName": "VM8-RHEL74MDPi567890-0123456789-0123456789-0123456789-0123456789"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60307,
			  "clientName": "VSAAZURE",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1a",
			"vmGuestSpace": 11811160064,
			"bkpStartTime": 1539027339,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Linux_64-bit",
			"isDeleted": false,
			"vendor": 4,
			"osType": 2,
			"vmSize": 11811160064,
			"vmUsedSpace": 11811160064,
			"subclientId": 211634,
			"bkpEndTime": 1539027339,
			"vmAgent": "VSAPROXY",
			"name": "aditya-bkp-1",
			"vmHardwareVer": "",
			"strGUID": "i-0a298c190f3096ca1",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60355,
			  "clientName": "aditya-bkp-1"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1a",
			"vmGuestSpace": 9663676416,
			"bkpStartTime": 1539027338,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Linux_64-bit",
			"isDeleted": false,
			"vendor": 4,
			"osType": 2,
			"vmSize": 9663676416,
			"vmUsedSpace": 9663676416,
			"subclientId": 211634,
			"bkpEndTime": 1539027338,
			"vmAgent": "VSAPROXY",
			"name": "aditya-bkp-2",
			"vmHardwareVer": "",
			"strGUID": "i-05b7bebf9bb675d77",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60356,
			  "clientName": "aditya-bkp-2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1c",
			"vmGuestSpace": 161061273600,
			"bkpStartTime": 1539027338,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Windows Server 2016 Datacenter_64-bit",
			"isDeleted": false,
			"vendor": 4,
			"osType": 1,
			"vmSize": 161061273600,
			"vmUsedSpace": 161061273600,
			"subclientId": 211634,
			"bkpEndTime": 1539027338,
			"vmAgent": "VSAPROXY",
			"name": "Ajay_SP13_Proxy_DO_NOT_TERMINATE",
			"vmHardwareVer": "",
			"strGUID": "i-08b2139e5843bbfb2",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60357,
			  "clientName": "Ajay_SP13_Proxy_DO_NOT_TERMINATE"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1a",
			"vmGuestSpace": 32212254720,
			"bkpStartTime": 1539027342,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Windows Server 2012 R2 Standard_64-bit",
			"isDeleted": false,
			"vendor": 4,
			"osType": 1,
			"vmSize": 32212254720,
			"vmUsedSpace": 32212254720,
			"subclientId": 211634,
			"bkpEndTime": 1539027342,
			"vmAgent": "VSAPROXY",
			"name": "Anita_TestVM_DontDelete",
			"vmHardwareVer": "",
			"strGUID": "i-0c762e5bf3a75a69c",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60358,
			  "clientName": "Anita_TestVM_DontDelete"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1a",
			"vmGuestSpace": 107374182400,
			"bkpStartTime": 1539027335,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Windows Server 2016 Datacenter_64-bit",
			"isDeleted": false,
			"vendor": 4,
			"osType": 1,
			"vmSize": 107374182400,
			"vmUsedSpace": 107374182400,
			"subclientId": 211634,
			"bkpEndTime": 1539027335,
			"vmAgent": "VSAPROXY",
			"name": "arya2016proxy_Donot_Delete",
			"vmHardwareVer": "",
			"strGUID": "i-02342f01a370285f1",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60359,
			  "clientName": "arya2016proxy_Donot_Delete"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1c",
			"vmGuestSpace": 51539607552,
			"bkpStartTime": 1539027336,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Windows Server 2016 Datacenter_64-bit",
			"isDeleted": false,
			"vendor": 4,
			"osType": 1,
			"vmSize": 51539607552,
			"vmUsedSpace": 51539607552,
			"subclientId": 211634,
			"bkpEndTime": 1539027337,
			"vmAgent": "VSAPROXY",
			"name": "AryaAppwareProxy_Donotterminate",
			"vmHardwareVer": "",
			"strGUID": "i-0453900fd2b6bae3a",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60360,
			  "clientName": "AryaAppwareProxy_Donotterminate"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1a",
			"vmGuestSpace": 25769803776,
			"bkpStartTime": 1539027335,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Linux_64-bit",
			"isDeleted": false,
			"vendor": 4,
			"osType": 2,
			"vmSize": 25769803776,
			"vmUsedSpace": 25769803776,
			"subclientId": 211634,
			"bkpEndTime": 1539027335,
			"vmAgent": "VSAPROXY",
			"name": "AryaUbuntu_SP14_Sync_Arya",
			"vmHardwareVer": "",
			"strGUID": "i-022c80d19c08a082b",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60361,
			  "clientName": "AryaUbuntu_SP14_Sync_Arya"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1a",
			"vmGuestSpace": 10737418240,
			"bkpStartTime": 1539027337,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Linux_64-bit",
			"isDeleted": false,
			"vendor": 4,
			"osType": 2,
			"vmSize": 10737418240,
			"vmUsedSpace": 10737418240,
			"subclientId": 211634,
			"bkpEndTime": 1539027337,
			"vmAgent": "VSAPROXY",
			"name": "AWSTOVMW_RHEL75_donotdelete",
			"vmHardwareVer": "",
			"strGUID": "i-0463346521d030628",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60362,
			  "clientName": "AWSTOVMW_RHEL75_donotdelete"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1a",
			"vmGuestSpace": 10737418240,
			"bkpStartTime": 1539027341,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Linux_64-bit",
			"isDeleted": false,
			"vendor": 4,
			"osType": 2,
			"vmSize": 10737418240,
			"vmUsedSpace": 10737418240,
			"subclientId": 211634,
			"bkpEndTime": 1539027341,
			"vmAgent": "VSAPROXY",
			"name": "AWSTOVMW_SUSE_12_dontdelete",
			"vmHardwareVer": "",
			"strGUID": "i-0b1c1b51638f03a9d",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60363,
			  "clientName": "AWSTOVMW_SUSE_12_dontdelete"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1a",
			"vmGuestSpace": 154618822656,
			"bkpStartTime": 1539027342,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Windows Server 2012 Standard_64-bit",
			"isDeleted": false,
			"vendor": 4,
			"osType": 1,
			"vmSize": 154618822656,
			"vmUsedSpace": 154618822656,
			"subclientId": 211634,
			"bkpEndTime": 1539027342,
			"vmAgent": "VSAPROXY",
			"name": "AWSVM5_SP14_Sync_Arya",
			"vmHardwareVer": "",
			"strGUID": "i-0bdcf1ba86af6631d",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60364,
			  "clientName": "AWSVM5_SP14_Sync_Arya"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1a",
			"vmGuestSpace": 107374182400,
			"bkpStartTime": 1539027343,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Windows Server 2016 Datacenter_64-bit",
			"isDeleted": false,
			"vendor": 4,
			"osType": 1,
			"vmSize": 107374182400,
			"vmUsedSpace": 107374182400,
			"subclientId": 211634,
			"bkpEndTime": 1539027343,
			"vmAgent": "VSAPROXY",
			"name": "bhama_sp11proxy_donotDelete",
			"vmHardwareVer": "",
			"strGUID": "i-0f8e2ab327393f9bc",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60365,
			  "clientName": "bhama_sp11proxy_donotDelete"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1c",
			"vmGuestSpace": 64424509440,
			"bkpStartTime": 1539027340,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Windows Server 2016 Datacenter_64-bit",
			"isDeleted": false,
			"vendor": 4,
			"osType": 1,
			"vmSize": 64424509440,
			"vmUsedSpace": 64424509440,
			"subclientId": 211634,
			"bkpEndTime": 1539027341,
			"vmAgent": "VSAPROXY",
			"name": "BhamaAWSTestVM1",
			"vmHardwareVer": "",
			"strGUID": "i-0ae6f3e334d751484",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60366,
			  "clientName": "BhamaAWSTestVM1"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1a",
			"vmGuestSpace": 35433480192,
			"bkpStartTime": 1539027333,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Windows Server 2012 R2 Standard_64-bit",
			"isDeleted": false,
			"vendor": 4,
			"osType": 1,
			"vmSize": 35433480192,
			"vmUsedSpace": 35433480192,
			"subclientId": 211634,
			"bkpEndTime": 1539027334,
			"vmAgent": "VSAPROXY",
			"name": "BhamaAWSTestVM2",
			"vmHardwareVer": "",
			"strGUID": "i-01519733524edceea",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60367,
			  "clientName": "BhamaAWSTestVM2"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1c",
			"vmGuestSpace": 21474836480,
			"bkpStartTime": 1539027334,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Windows Server 2012 R2 Standard_64-bit",
			"isDeleted": true,
			"vendor": 4,
			"osType": 1,
			"vmSize": 21474836480,
			"vmUsedSpace": 21474836480,
			"subclientId": 211634,
			"bkpEndTime": 1539027334,
			"vmAgent": "VSAPROXY",
			"name": "BSsanityVMWAWSConvVM1Importjavagui",
			"vmHardwareVer": "",
			"strGUID": "i-01d28d93f73b7bacf",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60368,
			  "clientName": "BSsanityVMWAWSConvVM1Importjavagui"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "us-east-1c",
			"vmGuestSpace": 21474836480,
			"bkpStartTime": 1539027336,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 19499791,
			"strOSName": "Windows_64-bit",
			"isDeleted": true,
			"vendor": 4,
			"osType": 1,
			"vmSize": 21474836480,
			"vmUsedSpace": 21474836480,
			"subclientId": 211634,
			"bkpEndTime": 1539027336,
			"vmAgent": "VSAPROXY",
			"name": "BSsanityVMWAWSVConvVM2HotaddjavaGUI",
			"vmHardwareVer": "",
			"strGUID": "i-03d8483bfd3e1ab8d",
			"subclientName": "VSASNAP",
			"client": {
			  "clientId": 60369,
			  "clientName": "BSsanityVMWAWSVConvVM2HotaddjavaGUI"
			},
			"proxyClient": {
			  "clientId": 60207,
			  "clientName": "VSAPROXY"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60354,
			  "clientName": "VSAAMAZON",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "nova",
			"vmGuestSpace": 1770,
			"bkpStartTime": 1541082429,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 20127631,
			"strOSName": "",
			"isDeleted": false,
			"vendor": 12,
			"osType": 2,
			"vmSize": 118111600640,
			"vmUsedSpace": 6484393984,
			"subclientId": 211645,
			"bkpEndTime": 1541082916,
			"vmAgent": "murloc-proxy",
			"name": "unix-inst-perf",
			"vmHardwareVer": "",
			"strGUID": "91ee559d-98dc-407b-992b-6dd3fddd4990",
			"subclientName": "Openstack-CBT-backup",
			"client": {
			  "clientId": 60565,
			  "clientName": "unix-inst-perf"
			},
			"proxyClient": {
			  "clientId": 60563,
			  "clientName": "murloc-proxy"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60564,
			  "clientName": "Openstack",
			  "flags": {
				"disabled": false
			  }
			}
		  },
		  {
			"vmHost": "nova",
			"vmGuestSpace": 1769,
			"bkpStartTime": 1541082450,
			"type": 9,
			"vmStatus": 1,
			"vmBackupJob": 20127631,
			"strOSName": "",
			"isDeleted": false,
			"vendor": 12,
			"osType": 1,
			"vmSize": 123480309760,
			"vmUsedSpace": 62914560,
			"subclientId": 211645,
			"bkpEndTime": 1541082561,
			"vmAgent": "murloc-proxy",
			"name": "win-inst-perf",
			"vmHardwareVer": "",
			"strGUID": "ba8cd2cd-c7c8-4485-9c0a-1c06ca87571b",
			"subclientName": "Openstack-CBT-backup",
			"client": {
			  "clientId": 60566,
			  "clientName": "win-inst-perf"
			},
			"proxyClient": {
			  "clientId": 60563,
			  "clientName": "murloc-proxy"
			},
			"plan": {},
			"pseudoClient": {
			  "clientId": 60564,
			  "clientName": "Openstack",
			  "flags": {
				"disabled": false
			  }
			}
		  }
		]
	  }
	`
}
