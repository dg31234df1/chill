package funcutil

import (
	"fmt"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/proto/commonpb"
	"github.com/milvus-io/milvus/internal/proto/milvuspb"
	"go.uber.org/zap"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	GlobalResourceName = "*"
)

func GetVersion(m proto.GeneratedMessage) (string, error) {
	md, _ := descriptor.MessageDescriptorProto(m)
	if md == nil {
		log.Error("MessageDescriptorProto result is nil")
		return "", fmt.Errorf("MessageDescriptorProto result is nil")
	}
	extObj, err := proto.GetExtension(md.Options, milvuspb.E_MilvusExtObj)
	if err != nil {
		log.Error("GetExtension fail", zap.Error(err))
		return "", err
	}
	version := extObj.(*milvuspb.MilvusExt).Version
	log.Debug("GetVersion success", zap.String("version", version))
	return version, nil
}

func GetPrivilegeExtObj(m proto.GeneratedMessage) (commonpb.PrivilegeExt, error) {
	_, md := descriptor.MessageDescriptorProto(m)
	if md == nil {
		log.Error("MessageDescriptorProto result is nil")
		return commonpb.PrivilegeExt{}, fmt.Errorf("MessageDescriptorProto result is nil")
	}

	extObj, err := proto.GetExtension(md.Options, commonpb.E_PrivilegeExtObj)
	if err != nil {
		log.Error("GetExtension fail", zap.Error(err))
		return commonpb.PrivilegeExt{}, err
	}
	privilegeExt := extObj.(*commonpb.PrivilegeExt)
	log.Debug("GetPrivilegeExtObj success", zap.String("resource_type", privilegeExt.ObjectType.String()), zap.String("resource_privilege", privilegeExt.ObjectPrivilege.String()))
	return commonpb.PrivilegeExt{
		ObjectType:      privilegeExt.ObjectType,
		ObjectPrivilege: privilegeExt.ObjectPrivilege,
		ObjectNameIndex: privilegeExt.ObjectNameIndex,
	}, nil
}

func GetResourceName(m proto.GeneratedMessage, index int32) string {
	if index <= 0 {
		return GlobalResourceName
	}
	msg := proto.MessageReflect(proto.MessageV1(m))
	msgDesc := msg.Descriptor()
	return msg.Get(msgDesc.Fields().ByNumber(protoreflect.FieldNumber(index))).String()
}

func PolicyForPrivilege(roleName string, objectType string, objectName string, privilege string) string {
	return fmt.Sprintf(`{"PType":"p","V0":"%s","V1":"%s","V2":"%s"}`, roleName, PolicyForResource(objectType, objectName), privilege)
}

func PolicyForResource(objectType string, objectName string) string {
	return fmt.Sprintf("%s-%s", objectType, objectName)
}
