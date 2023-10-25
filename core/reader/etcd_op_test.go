package reader

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/pkg/common"
	"github.com/milvus-io/milvus/pkg/util/retry"
	"github.com/stretchr/testify/assert"
	"github.com/zilliztech/milvus-cdc/core/api"
	"github.com/zilliztech/milvus-cdc/core/pb"
	"github.com/zilliztech/milvus-cdc/core/util"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// HINT: Before running these cases, you should start the etcd server
const TestCasePrefix = "test_case/cdc"

func TestEtcdOp(t *testing.T) {
	etcdOp, err := NewEtcdOp(nil, "", "", "")
	assert.NoError(t, err)
	realOp := etcdOp.(*EtcdOp)

	realOp.rootPath = TestCasePrefix
	realOp.retryOptions = []retry.Option{
		retry.Attempts(1),
	}
	defer func() {
		realOp.etcdClient.Delete(context.Background(), realOp.rootPath, clientv3.WithPrefix())
	}()

	assert.Equal(t, fmt.Sprintf("%s/%s/%s", TestCasePrefix, realOp.metaSubPath, collectionPrefix), realOp.collectionPrefix())
	assert.Equal(t, fmt.Sprintf("%s/%s/%s", TestCasePrefix, realOp.metaSubPath, partitionPrefix), realOp.partitionPrefix())
	assert.Equal(t, fmt.Sprintf("%s/%s/%s", TestCasePrefix, realOp.metaSubPath, fieldPrefix), realOp.fieldPrefix())

	// watch collection
	{
		var success util.Value[bool]
		success.Store(false)
		etcdOp.WatchCollection(context.Background(), func(ci *pb.CollectionInfo) bool {
			if strings.Contains(ci.Schema.Name, "test") {
				return true
			}
			return false
		})
		etcdOp.SubscribeCollectionEvent("empty_event", func(ci *pb.CollectionInfo) bool {
			return false
		})
		etcdOp.SubscribeCollectionEvent("collection_event_123", func(ci *pb.CollectionInfo) bool {
			success.Store(true)
			return true
		})

		// "not created collection"
		{
			collectionInfo := &pb.CollectionInfo{
				ID:    100008,
				State: pb.CollectionState_CollectionCreating,
			}
			collectionBytes, _ := proto.Marshal(collectionInfo)
			realOp.etcdClient.Put(context.Background(), realOp.collectionPrefix()+"/1/100008", string(collectionBytes))
		}

		// invalid data
		{
			realOp.etcdClient.Put(context.Background(), realOp.collectionPrefix()+"/1/100009", "hoo")
		}

		// filted collection
		{
			collectionInfo := &pb.CollectionInfo{
				State: pb.CollectionState_CollectionCreated,
				Schema: &schemapb.CollectionSchema{
					Name: "test123",
				},
			}
			collectionBytes, _ := proto.Marshal(collectionInfo)
			realOp.etcdClient.Put(context.Background(), realOp.collectionPrefix()+"/1/100000", string(collectionBytes))
		}

		// "no filed info"
		{
			collectionInfo := &pb.CollectionInfo{
				ID:    100003,
				State: pb.CollectionState_CollectionCreated,
				Schema: &schemapb.CollectionSchema{
					Name: "hoo",
				},
			}
			collectionBytes, _ := proto.Marshal(collectionInfo)
			realOp.etcdClient.Put(context.Background(), realOp.collectionPrefix()+"/1/100003", string(collectionBytes))
		}

		// success
		{
			field1 := &schemapb.FieldSchema{
				FieldID: 1,
				Name:    "ID",
			}
			field2 := &schemapb.FieldSchema{
				FieldID: 2,
				Name:    common.MetaFieldName,
			}
			field3 := &schemapb.FieldSchema{
				FieldID: 100,
				Name:    "age",
			}
			realOp.etcdClient.Put(context.Background(), realOp.fieldPrefix()+"/100002/1", getStringForMessage(field1))
			realOp.etcdClient.Put(context.Background(), realOp.fieldPrefix()+"/100002/2", getStringForMessage(field2))
			realOp.etcdClient.Put(context.Background(), realOp.fieldPrefix()+"/100002/100", getStringForMessage(field3))

			collectionInfo := &pb.CollectionInfo{
				ID:    100002,
				State: pb.CollectionState_CollectionCreated,
				Schema: &schemapb.CollectionSchema{
					Name: "foo",
				},
			}
			collectionBytes, _ := proto.Marshal(collectionInfo)
			realOp.etcdClient.Put(context.Background(), realOp.collectionPrefix()+"/1/100002", string(collectionBytes))
			time.Sleep(500 * time.Millisecond)
			assert.Eventually(t, func() bool {
				return success.Load()
			}, time.Second, time.Millisecond*100)
		}

		{
			etcdOp.UnsubscribeEvent("empty_event", api.CollectionEventType)
			etcdOp.UnsubscribeEvent("collection_event_123", api.CollectionEventType)
			etcdOp.UnsubscribeEvent("collection_event_123", api.WatchEventType(0))
		}
	}

	// get all collection
	{
		collections, err := etcdOp.GetAllCollection(context.Background(), func(ci *pb.CollectionInfo) bool {
			if strings.Contains(ci.Schema.Name, "test") {
				return true
			}
			return false
		})
		assert.NoError(t, err)
		assert.Len(t, collections, 1)
		assert.EqualValues(t, 100002, collections[0].ID)
	}

	// get collection name by id
	{
		collectionName := etcdOp.GetCollectionNameByID(context.Background(), 100000)
		assert.NoError(t, err)
		assert.Equal(t, "test123", collectionName)
	}

	// watch partition
	{
		var success util.Value[bool]
		success.Store(false)

		etcdOp.WatchPartition(context.Background(), func(pi *pb.PartitionInfo) bool {
			if strings.Contains(pi.PartitionName, "test") {
				return true
			}
			return false
		})
		etcdOp.SubscribePartitionEvent("empty_event", func(pi *pb.PartitionInfo) bool {
			return false
		})
		etcdOp.SubscribePartitionEvent("partition_event_123", func(pi *pb.PartitionInfo) bool {
			success.Store(true)
			return true
		})
		// no created state
		{
			info := &pb.PartitionInfo{
				State: pb.PartitionState_PartitionCreating,
			}
			realOp.etcdClient.Put(context.Background(), realOp.partitionPrefix()+"/9/200005", getStringForMessage(info))
		}
		// "invalid data"
		{
			realOp.etcdClient.Put(context.Background(), realOp.partitionPrefix()+"/9/200009", "foo")
		}
		// default partition
		{
			info := &pb.PartitionInfo{
				State:         pb.PartitionState_PartitionCreated,
				PartitionName: realOp.defaultPartitionName,
			}
			realOp.etcdClient.Put(context.Background(), realOp.partitionPrefix()+"/9/200003", getStringForMessage(info))
		}
		// filled partition

		{
			info := &pb.PartitionInfo{
				State:         pb.PartitionState_PartitionCreated,
				PartitionName: "test345",
			}
			realOp.etcdClient.Put(context.Background(), realOp.partitionPrefix()+"/9/200004", getStringForMessage(info))
		}
		// success
		{
			info := &pb.PartitionInfo{
				State:         pb.PartitionState_PartitionCreated,
				PartitionName: "foo",
			}
			realOp.etcdClient.Put(context.Background(), realOp.partitionPrefix()+"/9/200005", getStringForMessage(info))

			time.Sleep(500 * time.Millisecond)
			assert.Eventually(t, func() bool {
				return success.Load()
			}, time.Second, time.Millisecond*100)
		}
	}

	// get all partition
	{
		partitions, err := etcdOp.GetAllPartition(context.Background(), func(pi *pb.PartitionInfo) bool {
			if strings.Contains(pi.PartitionName, "test") {
				return true
			}
			return false
		})
		assert.NoError(t, err)
		assert.Len(t, partitions, 1)
		assert.EqualValues(t, "foo", partitions[0].PartitionName)
	}

	// get collection id from partition id
	{
		{
			id := realOp.getCollectionIDFromPartitionKey(realOp.partitionPrefix() + "/123456/9001")
			assert.EqualValues(t, 123456, id)
		}
		{
			id := realOp.getCollectionIDFromPartitionKey(realOp.partitionPrefix() + "/123456")
			assert.EqualValues(t, 0, id)
		}
		{
			id := realOp.getCollectionIDFromPartitionKey(realOp.partitionPrefix() + "/nonumber/9001")
			assert.EqualValues(t, 0, id)
		}
	}

	// get collection name by id
	{
		{
			collectionName := etcdOp.GetCollectionNameByID(context.Background(), 900000)
			assert.Equal(t, "", collectionName)
		}
		{
			realOp.etcdClient.Put(context.Background(), realOp.collectionPrefix()+"/1/800002", "foo")
			collectionName := etcdOp.GetCollectionNameByID(context.Background(), 800002)
			assert.Equal(t, "", collectionName)
		}
	}
}

func getStringForMessage(m proto.Message) string {
	b, _ := proto.Marshal(m)
	return string(b)
}