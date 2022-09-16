// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"golang.org/x/exp/mmap"

	"github.com/golang/protobuf/proto"
	"github.com/milvus-io/milvus/api/schemapb"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/util/tsoutil"
)

// PrintBinlogFiles call printBinlogFile in turn for the file list specified by parameter fileList.
// Return an error early if it encounters any error.
func PrintBinlogFiles(fileList []string) error {
	for _, file := range fileList {
		if err := printBinlogFile(file); err != nil {
			return err
		}
	}
	return nil
}

func printBinlogFile(filename string) error {
	fd, err := os.OpenFile(filename, os.O_RDONLY, 0400)
	if err != nil {
		return err
	}
	defer fd.Close()

	fileInfo, err := fd.Stat()
	if err != nil {
		return err
	}

	fmt.Printf("file size = %d\n", fileInfo.Size())

	at, err := mmap.Open(filename)
	if err != nil {
		return nil
	}
	defer at.Close()

	b := make([]byte, fileInfo.Size())
	at.ReadAt(b, 0)

	r, err := NewBinlogReader(b)
	if err != nil {
		return err
	}
	defer r.Close()

	fmt.Println("descriptor event header:")
	physical, _ := tsoutil.ParseTS(r.descriptorEvent.descriptorEventHeader.Timestamp)
	fmt.Printf("\tTimestamp: %v\n", physical)
	fmt.Printf("\tTypeCode: %s\n", r.descriptorEvent.descriptorEventHeader.TypeCode.String())
	fmt.Printf("\tEventLength: %d\n", r.descriptorEvent.descriptorEventHeader.EventLength)
	fmt.Printf("\tNextPosition :%d\n", r.descriptorEvent.descriptorEventHeader.NextPosition)
	fmt.Println("descriptor event data:")
	fmt.Printf("\tCollectionID: %d\n", r.descriptorEvent.descriptorEventData.CollectionID)
	fmt.Printf("\tPartitionID: %d\n", r.descriptorEvent.descriptorEventData.PartitionID)
	fmt.Printf("\tSegmentID: %d\n", r.descriptorEvent.descriptorEventData.SegmentID)
	fmt.Printf("\tFieldID: %d\n", r.descriptorEvent.descriptorEventData.FieldID)
	physical, _ = tsoutil.ParseTS(r.descriptorEvent.descriptorEventData.StartTimestamp)
	fmt.Printf("\tStartTimestamp: %v\n", physical)
	physical, _ = tsoutil.ParseTS(r.descriptorEvent.descriptorEventData.EndTimestamp)
	fmt.Printf("\tEndTimestamp: %v\n", physical)
	dataTypeName, ok := schemapb.DataType_name[int32(r.descriptorEvent.descriptorEventData.PayloadDataType)]
	if !ok {
		return fmt.Errorf("undefine data type %d", r.descriptorEvent.descriptorEventData.PayloadDataType)
	}
	fmt.Printf("\tPayloadDataType: %v\n", dataTypeName)
	fmt.Printf("\tPostHeaderLengths: %v\n", r.descriptorEvent.descriptorEventData.PostHeaderLengths)
	eventNum := 0
	for {
		event, err := r.NextEventReader()
		if err != nil {
			return err
		}
		if event == nil {
			break
		}
		fmt.Printf("event %d header:\n", eventNum)
		physical, _ = tsoutil.ParseTS(event.eventHeader.Timestamp)
		fmt.Printf("\tTimestamp: %v\n", physical)
		fmt.Printf("\tTypeCode: %s\n", event.eventHeader.TypeCode.String())
		fmt.Printf("\tEventLength: %d\n", event.eventHeader.EventLength)
		fmt.Printf("\tNextPosition: %d\n", event.eventHeader.NextPosition)
		switch event.eventHeader.TypeCode {
		case InsertEventType:
			evd, ok := event.eventData.(*insertEventData)
			if !ok {
				return errors.New("incorrect event data type")
			}
			fmt.Printf("event %d insert event:\n", eventNum)
			physical, _ = tsoutil.ParseTS(evd.StartTimestamp)
			fmt.Printf("\tStartTimestamp: %v\n", physical)
			physical, _ = tsoutil.ParseTS(evd.EndTimestamp)
			fmt.Printf("\tEndTimestamp: %v\n", physical)
			if err := printPayloadValues(r.descriptorEvent.descriptorEventData.PayloadDataType, event.PayloadReaderInterface); err != nil {
				return err
			}
		case DeleteEventType:
			evd, ok := event.eventData.(*deleteEventData)
			if !ok {
				return errors.New("incorrect event data type")
			}
			fmt.Printf("event %d delete event:\n", eventNum)
			physical, _ = tsoutil.ParseTS(evd.StartTimestamp)
			fmt.Printf("\tStartTimestamp: %v\n", physical)
			physical, _ = tsoutil.ParseTS(evd.EndTimestamp)
			fmt.Printf("\tEndTimestamp: %v\n", physical)
			if err := printPayloadValues(r.descriptorEvent.descriptorEventData.PayloadDataType, event.PayloadReaderInterface); err != nil {
				return err
			}
		case CreateCollectionEventType:
			evd, ok := event.eventData.(*createCollectionEventData)
			if !ok {
				return errors.New("incorrect event data type")
			}
			fmt.Printf("event %d create collection event:\n", eventNum)
			physical, _ = tsoutil.ParseTS(evd.StartTimestamp)
			fmt.Printf("\tStartTimestamp: %v\n", physical)
			physical, _ = tsoutil.ParseTS(evd.EndTimestamp)
			fmt.Printf("\tEndTimestamp: %v\n", physical)
			if err := printDDLPayloadValues(event.eventHeader.TypeCode, r.descriptorEvent.descriptorEventData.PayloadDataType, event.PayloadReaderInterface); err != nil {
				return err
			}
		case DropCollectionEventType:
			evd, ok := event.eventData.(*dropCollectionEventData)
			if !ok {
				return errors.New("incorrect event data type")
			}
			fmt.Printf("event %d drop collection event:\n", eventNum)
			physical, _ = tsoutil.ParseTS(evd.StartTimestamp)
			fmt.Printf("\tStartTimestamp: %v\n", physical)
			physical, _ = tsoutil.ParseTS(evd.EndTimestamp)
			fmt.Printf("\tEndTimestamp: %v\n", physical)
			if err := printDDLPayloadValues(event.eventHeader.TypeCode, r.descriptorEvent.descriptorEventData.PayloadDataType, event.PayloadReaderInterface); err != nil {
				return err
			}
		case CreatePartitionEventType:
			evd, ok := event.eventData.(*createPartitionEventData)
			if !ok {
				return errors.New("incorrect event data type")
			}
			fmt.Printf("event %d create partition event:\n", eventNum)
			physical, _ = tsoutil.ParseTS(evd.StartTimestamp)
			fmt.Printf("\tStartTimestamp: %v\n", physical)
			physical, _ = tsoutil.ParseTS(evd.EndTimestamp)
			fmt.Printf("\tEndTimestamp: %v\n", physical)
			if err := printDDLPayloadValues(event.eventHeader.TypeCode, r.descriptorEvent.descriptorEventData.PayloadDataType, event.PayloadReaderInterface); err != nil {
				return err
			}
		case DropPartitionEventType:
			evd, ok := event.eventData.(*dropPartitionEventData)
			if !ok {
				return errors.New("incorrect event data type")
			}
			fmt.Printf("event %d drop partition event:\n", eventNum)
			physical, _ = tsoutil.ParseTS(evd.StartTimestamp)
			fmt.Printf("\tStartTimestamp: %v\n", physical)
			physical, _ = tsoutil.ParseTS(evd.EndTimestamp)
			fmt.Printf("\tEndTimestamp: %v\n", physical)
			if err := printDDLPayloadValues(event.eventHeader.TypeCode, r.descriptorEvent.descriptorEventData.PayloadDataType, event.PayloadReaderInterface); err != nil {
				return err
			}
		case IndexFileEventType:
			desc := r.descriptorEvent
			extraBytes := desc.ExtraBytes
			extra := make(map[string]interface{})
			err = json.Unmarshal(extraBytes, &extra)
			if err != nil {
				return fmt.Errorf("failed to unmarshal extra: %s", err.Error())
			}
			fmt.Printf("indexBuildID: %v\n", extra["indexBuildID"])
			fmt.Printf("indexName: %v\n", extra["indexName"])
			fmt.Printf("indexID: %v\n", extra["indexID"])
			evd, ok := event.eventData.(*indexFileEventData)
			if !ok {
				return errors.New("incorrect event data type")
			}
			fmt.Printf("index file event num: %d\n", eventNum)
			physical, _ = tsoutil.ParseTS(evd.StartTimestamp)
			fmt.Printf("\tStartTimestamp: %v\n", physical)
			physical, _ = tsoutil.ParseTS(evd.EndTimestamp)
			fmt.Printf("\tEndTimestamp: %v\n", physical)
			key := fmt.Sprintf("%v", extra["key"])
			if err := printIndexFilePayloadValues(event.PayloadReaderInterface, key); err != nil {
				return err
			}
		default:
			return fmt.Errorf("undefined event typd %d", event.eventHeader.TypeCode)
		}
		eventNum++
	}

	return nil
}

func printPayloadValues(colType schemapb.DataType, reader PayloadReaderInterface) error {
	fmt.Println("\tpayload values:")
	switch colType {
	case schemapb.DataType_Bool:
		val, err := reader.GetBoolFromPayload()
		if err != nil {
			return err
		}
		for i, v := range val {
			fmt.Printf("\t\t%d : %v\n", i, v)
		}
	case schemapb.DataType_Int8:
		val, err := reader.GetInt8FromPayload()
		if err != nil {
			return err
		}
		for i, v := range val {
			fmt.Printf("\t\t%d : %d\n", i, v)
		}
	case schemapb.DataType_Int16:
		val, err := reader.GetInt16FromPayload()
		if err != nil {
			return err
		}
		for i, v := range val {
			fmt.Printf("\t\t%d : %d\n", i, v)
		}
	case schemapb.DataType_Int32:
		val, err := reader.GetInt32FromPayload()
		if err != nil {
			return err
		}
		for i, v := range val {
			fmt.Printf("\t\t%d : %d\n", i, v)
		}
	case schemapb.DataType_Int64:
		val, err := reader.GetInt64FromPayload()
		if err != nil {
			return err
		}
		for i, v := range val {
			fmt.Printf("\t\t%d : %d\n", i, v)
		}
	case schemapb.DataType_Float:
		val, err := reader.GetFloatFromPayload()
		if err != nil {
			return err
		}
		for i, v := range val {
			fmt.Printf("\t\t%d : %f\n", i, v)
		}
	case schemapb.DataType_Double:
		val, err := reader.GetDoubleFromPayload()
		if err != nil {
			return err
		}
		for i, v := range val {
			fmt.Printf("\t\t%d : %v\n", i, v)
		}
	case schemapb.DataType_String, schemapb.DataType_VarChar:
		rows, err := reader.GetPayloadLengthFromReader()
		if err != nil {
			return err
		}

		val, err := reader.GetStringFromPayload()
		if err != nil {
			return err
		}
		for i := 0; i < rows; i++ {
			fmt.Printf("\t\t%d : %s\n", i, val[i])
		}
	case schemapb.DataType_BinaryVector:
		val, dim, err := reader.GetBinaryVectorFromPayload()
		if err != nil {
			return err
		}
		dim = dim / 8
		length := len(val) / dim
		for i := 0; i < length; i++ {
			fmt.Printf("\t\t%d :", i)
			for j := 0; j < dim; j++ {
				idx := i*dim + j
				fmt.Printf(" %02x", val[idx])
			}
			fmt.Println()
		}
	case schemapb.DataType_FloatVector:
		val, dim, err := reader.GetFloatVectorFromPayload()
		if err != nil {
			return err
		}
		length := len(val) / dim
		for i := 0; i < length; i++ {
			fmt.Printf("\t\t%d :", i)
			for j := 0; j < dim; j++ {
				idx := i*dim + j
				fmt.Printf(" %f", val[idx])
			}
			fmt.Println()
		}
	default:
		return errors.New("undefined data type")
	}
	return nil
}

func printDDLPayloadValues(eventType EventTypeCode, colType schemapb.DataType, reader PayloadReaderInterface) error {
	fmt.Println("\tpayload values:")
	switch colType {
	case schemapb.DataType_Int64:
		val, err := reader.GetInt64FromPayload()
		if err != nil {
			return err
		}
		for i, v := range val {
			physical, logical := tsoutil.ParseTS(uint64(v))
			fmt.Printf("\t\t%d : physical : %v ; logical : %d\n", i, physical, logical)
		}
	case schemapb.DataType_String:
		rows, err := reader.GetPayloadLengthFromReader()
		if err != nil {
			return err
		}

		val, err := reader.GetStringFromPayload()
		if err != nil {
			return err
		}
		for i := 0; i < rows; i++ {
			valBytes := []byte(val[i])
			switch eventType {
			case CreateCollectionEventType:
				var req internalpb.CreateCollectionRequest
				if err := proto.Unmarshal(valBytes, &req); err != nil {
					return err
				}
				fmt.Printf("\t\t%d : create collection: %v\n", i, req)
			case DropCollectionEventType:
				var req internalpb.DropCollectionRequest
				if err := proto.Unmarshal(valBytes, &req); err != nil {
					return err
				}
				fmt.Printf("\t\t%d : drop collection: %v\n", i, req)
			case CreatePartitionEventType:
				var req internalpb.CreatePartitionRequest
				if err := proto.Unmarshal(valBytes, &req); err != nil {
					return err
				}
				fmt.Printf("\t\t%d : create partition: %v\n", i, req)
			case DropPartitionEventType:
				var req internalpb.DropPartitionRequest
				if err := proto.Unmarshal(valBytes, &req); err != nil {
					return err
				}
				fmt.Printf("\t\t%d : drop partition: %v\n", i, req)
			default:
				return fmt.Errorf("undefined ddl event type %d", eventType)
			}
		}
	default:
		return errors.New("undefined data type")
	}
	return nil
}

// only print slice meta and index params
func printIndexFilePayloadValues(reader PayloadReaderInterface, key string) error {
	if key == IndexParamsKey {
		content, err := reader.GetByteFromPayload()
		if err != nil {
			return err
		}
		fmt.Print("index params: \n")
		fmt.Println(content)

		return nil
	}

	if key == "SLICE_META" {
		content, err := reader.GetByteFromPayload()
		if err != nil {
			return err
		}
		// content is a json string serialized by milvus::json,
		// it's better to use milvus::json to parse the content also,
		// fortunately, the json string is readable enough.
		fmt.Print("index slice meta: \n")
		fmt.Println(content)

		return nil
	}

	return nil
}
