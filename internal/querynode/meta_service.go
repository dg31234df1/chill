package querynode

import (
	"context"
	"fmt"
	"log"
	"path"
	"reflect"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	etcdkv "github.com/zilliztech/milvus-distributed/internal/kv/etcd"
	"github.com/zilliztech/milvus-distributed/internal/proto/datapb"
	"github.com/zilliztech/milvus-distributed/internal/proto/etcdpb"
	"go.etcd.io/etcd/clientv3"
)

const (
	CollectionPrefix = "/collection/"
	SegmentPrefix    = "/segment/"
)

type metaService struct {
	ctx     context.Context
	kvBase  *etcdkv.EtcdKV
	replica collectionReplica
}

func newMetaService(ctx context.Context, replica collectionReplica) *metaService {
	ETCDAddr := Params.ETCDAddress
	MetaRootPath := Params.MetaRootPath

	cli, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{ETCDAddr},
		DialTimeout: 5 * time.Second,
	})

	return &metaService{
		ctx:     ctx,
		kvBase:  etcdkv.NewEtcdKV(cli, MetaRootPath),
		replica: replica,
	}
}

func (mService *metaService) start() {
	// init from meta
	err := mService.loadCollections()
	if err != nil {
		log.Fatal("metaService loadCollections failed")
	}
	err = mService.loadSegments()
	if err != nil {
		log.Fatal("metaService loadSegments failed")
	}
}

func GetCollectionObjID(key string) string {
	ETCDRootPath := Params.MetaRootPath

	prefix := path.Join(ETCDRootPath, CollectionPrefix) + "/"
	return strings.TrimPrefix(key, prefix)
}

func GetSegmentObjID(key string) string {
	ETCDRootPath := Params.MetaRootPath

	prefix := path.Join(ETCDRootPath, SegmentPrefix) + "/"
	return strings.TrimPrefix(key, prefix)
}

func isCollectionObj(key string) bool {
	ETCDRootPath := Params.MetaRootPath

	prefix := path.Join(ETCDRootPath, CollectionPrefix) + "/"
	prefix = strings.TrimSpace(prefix)
	index := strings.Index(key, prefix)

	return index == 0
}

func isSegmentObj(key string) bool {
	ETCDRootPath := Params.MetaRootPath

	prefix := path.Join(ETCDRootPath, SegmentPrefix) + "/"
	prefix = strings.TrimSpace(prefix)
	index := strings.Index(key, prefix)

	return index == 0
}

func printCollectionStruct(obj *etcdpb.CollectionInfo) {
	v := reflect.ValueOf(obj)
	v = reflect.Indirect(v)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if typeOfS.Field(i).Name == "GrpcMarshalString" {
			continue
		}
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}

func printSegmentStruct(obj *datapb.SegmentInfo) {
	v := reflect.ValueOf(obj)
	v = reflect.Indirect(v)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}

func (mService *metaService) processCollectionCreate(id string, value string) {
	//println(fmt.Sprintf("Create Collection:$%s$", id))

	col := mService.collectionUnmarshal(value)
	if col != nil {
		schema := col.Schema
		err := mService.replica.addCollection(col.ID, schema)
		if err != nil {
			log.Println(err)
		}
		for _, partitionID := range col.PartitionIDs {
			err = mService.replica.addPartition(col.ID, partitionID)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (mService *metaService) processSegmentCreate(id string, value string) {
	//println("Create Segment: ", id)

	seg := mService.segmentUnmarshal(value)

	// TODO: what if seg == nil? We need to notify master and return rpc request failed
	if seg != nil {
		// TODO: get partition id from segment meta
		err := mService.replica.addSegment(seg.SegmentID, seg.PartitionID, seg.CollectionID, segTypeGrowing)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (mService *metaService) processCreate(key string, msg string) {
	//println("process create", key)
	if isCollectionObj(key) {
		objID := GetCollectionObjID(key)
		mService.processCollectionCreate(objID, msg)
	} else if isSegmentObj(key) {
		objID := GetSegmentObjID(key)
		mService.processSegmentCreate(objID, msg)
	} else {
		println("can not process create msg:", key)
	}
}

func (mService *metaService) loadCollections() error {
	keys, values, err := mService.kvBase.LoadWithPrefix(CollectionPrefix)
	if err != nil {
		return err
	}

	for i := range keys {
		objID := GetCollectionObjID(keys[i])
		mService.processCollectionCreate(objID, values[i])
	}

	return nil
}

func (mService *metaService) loadSegments() error {
	keys, values, err := mService.kvBase.LoadWithPrefix(SegmentPrefix)
	if err != nil {
		return err
	}

	for i := range keys {
		objID := GetSegmentObjID(keys[i])
		mService.processSegmentCreate(objID, values[i])
	}

	return nil
}

//----------------------------------------------------------------------- Unmarshal and Marshal
func (mService *metaService) collectionUnmarshal(value string) *etcdpb.CollectionInfo {
	col := etcdpb.CollectionInfo{}
	err := proto.UnmarshalText(value, &col)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &col
}

func (mService *metaService) collectionMarshal(col *etcdpb.CollectionInfo) string {
	value := proto.MarshalTextString(col)
	if value == "" {
		log.Println("marshal collection failed")
		return ""
	}
	return value
}

func (mService *metaService) segmentUnmarshal(value string) *datapb.SegmentInfo {
	seg := datapb.SegmentInfo{}
	err := proto.UnmarshalText(value, &seg)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &seg
}

func (mService *metaService) segmentMarshal(seg *etcdpb.SegmentMeta) string {
	value := proto.MarshalTextString(seg)
	if value == "" {
		log.Println("marshal segment failed")
		return ""
	}
	return value
}
