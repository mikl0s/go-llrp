package llrp

// Generate C1G2PC parameter from hexpc string.
func C1G2PC(hexpc string) []byte {
	intpc, _ := strconv.ParseInt(hexpc, 10, 32)
	var data = []interface{}{
		uint8(140),    // 1+uint7(Type=12)
		uint16(intpc), // PC bits
	}
	return Pack(data)
}

// Generate C1G2ReadOpSpecResult parameter from readData.
func C1G2ReadOpSpecResult(readData []byte) []byte {
	var data = []interface{}{
		uint16(349), // Rsvd+Type=
		uint16(11),  // Length
		uint8(0),    // Result
		uint16(9),   // OpSpecID
		uint16(1),   // ReadDataWordCount
		readData,    // ReadData
	}
	return Pack(data)
}

// Generate ConnectionAttemptEvent parameter.
func ConnectionAttemptEvent() []byte {
	var data = []interface{}{
		uint16(256), // Rsvd+Type=256
		uint16(6),   // Length
		uint16(0),   // Status(Success=0)
	}
	return Pack(data)
	return Pack(data)
}

// Generate EPCData parameter from its length and epcLength, and epc.
func EPCData(length int64, epcLengthBits int64, epc []byte) []byte {
	var data []interface{}
	if epcLengthBits == 96 {
		data = []interface{}{
			uint8(141), // 1+uint7(Type=13)
			epc,        // 96-bit EPCData string
		}
	} else {
		data = []interface{}{
			uint16(241),           // uint8(0)+uint8(Type=241)
			uint16(length),        // Length
			uint16(epcLengthBits), // EPCLengthBits
			epc, // EPCData string
		}
	}
	return Pack(data)
}

// Generate KeepaliveSpec parameter.
func KeepaliveSpec() []byte {
	var data = []interface{}{
		uint16(220),   // Rsvd+Type=220
		uint16(9),     // Length
		uint8(1),      // KeepaliveTriggerType=Periodic(1)
		uint32(10000), // TimeInterval=10000
	}
	return Pack(data)
}

// Generate LLRPStatus parameter.
func LLRPStatus() []byte {
	var data = []interface{}{
		uint16(287), // Rsvd+Type=287
		uint16(8),   // Length
		uint16(0),   // StatusCode=M_Success(0)
		uint16(0),   // ErrorDescriptionByteCount=0
	}
	return Pack(data)
}

// Generate PeakRSSI parameter.
func PeakRSSI() []byte {
	var data = []interface{}{
		uint8(134), // 1+uint7(Type=6)
		uint8(203), // PeakRSSI
	}
	return Pack(data)
}

// Generate ReaderEventNotification parameter.
func ReaderEventNotificationData() []byte {
	utcTimeStamp := UTCTimeStamp()
	connectionAttemptEvent := ConnectionAttemptEvent()
	readerEventNotificationDataLength := len(utcTimeStamp) +
		len(connectionAttemptEvent) + 4 // Rsvd+Type+length=32bits=4bytes
	var data = []interface{}{
		uint16(246),                               // Rsvd+Type=246 (ReaderEventNotificationData parameter)
		uint16(readerEventNotificationDataLength), // Length
		utcTimeStamp,
		connectionAttemptEvent,
	}
	return Pack(data)
}

// Generate TagReportData parameter from epcData, peakRSSI, airProtocolTagData, opSpecResult.
func TagReportData(epcData []byte,
	peakRSSI []byte,
	airProtocolTagData []byte,
	opSpecResult []byte) []byte {
	tagReportDataLength := len(epcData) +
		len(peakRSSI) + len(airProtocolTagData) +
		len(opSpecResult) + 4 // Rsvd+Type+length->32bits=4bytes
	var data = []interface{}{
		uint16(240),                 // Rsvd+Type=240 (TagReportData parameter)
		uint16(tagReportDataLength), // Length
		epcData,
		peakRSSI,
		airProtocolTagData,
		opSpecResult,
	}
	return Pack(data)
}

// Generate UTCTimeStamp parameter at the current time.
func UTCTimeStamp() []byte {
	currentTime := uint64(time.Now().UTC().Nanosecond() / 1000)
	var data = []interface{}{
		uint16(128), // Rsvd+Type=128
		uint16(12),  // Length
		currentTime, // Microseconds
	}
	return Pack(data)
}
