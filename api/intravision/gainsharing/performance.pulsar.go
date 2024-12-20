// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package gainsharing

import (
	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Performance        protoreflect.MessageDescriptor
	fd_Performance_id     protoreflect.FieldDescriptor
	fd_Performance_mid    protoreflect.FieldDescriptor
	fd_Performance_tid    protoreflect.FieldDescriptor
	fd_Performance_wager  protoreflect.FieldDescriptor
	fd_Performance_reward protoreflect.FieldDescriptor
	fd_Performance_earner protoreflect.FieldDescriptor
)

func init() {
	file_intravision_gainsharing_performance_proto_init()
	md_Performance = File_intravision_gainsharing_performance_proto.Messages().ByName("Performance")
	fd_Performance_id = md_Performance.Fields().ByName("id")
	fd_Performance_mid = md_Performance.Fields().ByName("mid")
	fd_Performance_tid = md_Performance.Fields().ByName("tid")
	fd_Performance_wager = md_Performance.Fields().ByName("wager")
	fd_Performance_reward = md_Performance.Fields().ByName("reward")
	fd_Performance_earner = md_Performance.Fields().ByName("earner")
}

var _ protoreflect.Message = (*fastReflection_Performance)(nil)

type fastReflection_Performance Performance

func (x *Performance) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Performance)(x)
}

func (x *Performance) slowProtoReflect() protoreflect.Message {
	mi := &file_intravision_gainsharing_performance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Performance_messageType fastReflection_Performance_messageType
var _ protoreflect.MessageType = fastReflection_Performance_messageType{}

type fastReflection_Performance_messageType struct{}

func (x fastReflection_Performance_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Performance)(nil)
}
func (x fastReflection_Performance_messageType) New() protoreflect.Message {
	return new(fastReflection_Performance)
}
func (x fastReflection_Performance_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Performance
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Performance) Descriptor() protoreflect.MessageDescriptor {
	return md_Performance
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Performance) Type() protoreflect.MessageType {
	return _fastReflection_Performance_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Performance) New() protoreflect.Message {
	return new(fastReflection_Performance)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Performance) Interface() protoreflect.ProtoMessage {
	return (*Performance)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Performance) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_Performance_id, value) {
			return
		}
	}
	if x.Mid != int32(0) {
		value := protoreflect.ValueOfInt32(x.Mid)
		if !f(fd_Performance_mid, value) {
			return
		}
	}
	if x.Tid != "" {
		value := protoreflect.ValueOfString(x.Tid)
		if !f(fd_Performance_tid, value) {
			return
		}
	}
	if x.Wager != nil {
		value := protoreflect.ValueOfMessage(x.Wager.ProtoReflect())
		if !f(fd_Performance_wager, value) {
			return
		}
	}
	if x.Reward != nil {
		value := protoreflect.ValueOfMessage(x.Reward.ProtoReflect())
		if !f(fd_Performance_reward, value) {
			return
		}
	}
	if x.Earner != "" {
		value := protoreflect.ValueOfString(x.Earner)
		if !f(fd_Performance_earner, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Performance) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "intravision.gainsharing.Performance.id":
		return x.Id != uint64(0)
	case "intravision.gainsharing.Performance.mid":
		return x.Mid != int32(0)
	case "intravision.gainsharing.Performance.tid":
		return x.Tid != ""
	case "intravision.gainsharing.Performance.wager":
		return x.Wager != nil
	case "intravision.gainsharing.Performance.reward":
		return x.Reward != nil
	case "intravision.gainsharing.Performance.earner":
		return x.Earner != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.gainsharing.Performance"))
		}
		panic(fmt.Errorf("message intravision.gainsharing.Performance does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Performance) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "intravision.gainsharing.Performance.id":
		x.Id = uint64(0)
	case "intravision.gainsharing.Performance.mid":
		x.Mid = int32(0)
	case "intravision.gainsharing.Performance.tid":
		x.Tid = ""
	case "intravision.gainsharing.Performance.wager":
		x.Wager = nil
	case "intravision.gainsharing.Performance.reward":
		x.Reward = nil
	case "intravision.gainsharing.Performance.earner":
		x.Earner = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.gainsharing.Performance"))
		}
		panic(fmt.Errorf("message intravision.gainsharing.Performance does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Performance) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "intravision.gainsharing.Performance.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "intravision.gainsharing.Performance.mid":
		value := x.Mid
		return protoreflect.ValueOfInt32(value)
	case "intravision.gainsharing.Performance.tid":
		value := x.Tid
		return protoreflect.ValueOfString(value)
	case "intravision.gainsharing.Performance.wager":
		value := x.Wager
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "intravision.gainsharing.Performance.reward":
		value := x.Reward
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "intravision.gainsharing.Performance.earner":
		value := x.Earner
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.gainsharing.Performance"))
		}
		panic(fmt.Errorf("message intravision.gainsharing.Performance does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Performance) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "intravision.gainsharing.Performance.id":
		x.Id = value.Uint()
	case "intravision.gainsharing.Performance.mid":
		x.Mid = int32(value.Int())
	case "intravision.gainsharing.Performance.tid":
		x.Tid = value.Interface().(string)
	case "intravision.gainsharing.Performance.wager":
		x.Wager = value.Message().Interface().(*v1beta1.Coin)
	case "intravision.gainsharing.Performance.reward":
		x.Reward = value.Message().Interface().(*v1beta1.Coin)
	case "intravision.gainsharing.Performance.earner":
		x.Earner = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.gainsharing.Performance"))
		}
		panic(fmt.Errorf("message intravision.gainsharing.Performance does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Performance) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "intravision.gainsharing.Performance.wager":
		if x.Wager == nil {
			x.Wager = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.Wager.ProtoReflect())
	case "intravision.gainsharing.Performance.reward":
		if x.Reward == nil {
			x.Reward = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.Reward.ProtoReflect())
	case "intravision.gainsharing.Performance.id":
		panic(fmt.Errorf("field id of message intravision.gainsharing.Performance is not mutable"))
	case "intravision.gainsharing.Performance.mid":
		panic(fmt.Errorf("field mid of message intravision.gainsharing.Performance is not mutable"))
	case "intravision.gainsharing.Performance.tid":
		panic(fmt.Errorf("field tid of message intravision.gainsharing.Performance is not mutable"))
	case "intravision.gainsharing.Performance.earner":
		panic(fmt.Errorf("field earner of message intravision.gainsharing.Performance is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.gainsharing.Performance"))
		}
		panic(fmt.Errorf("message intravision.gainsharing.Performance does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Performance) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "intravision.gainsharing.Performance.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "intravision.gainsharing.Performance.mid":
		return protoreflect.ValueOfInt32(int32(0))
	case "intravision.gainsharing.Performance.tid":
		return protoreflect.ValueOfString("")
	case "intravision.gainsharing.Performance.wager":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "intravision.gainsharing.Performance.reward":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "intravision.gainsharing.Performance.earner":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.gainsharing.Performance"))
		}
		panic(fmt.Errorf("message intravision.gainsharing.Performance does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Performance) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in intravision.gainsharing.Performance", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Performance) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Performance) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Performance) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Performance) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Performance)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		if x.Mid != 0 {
			n += 1 + runtime.Sov(uint64(x.Mid))
		}
		l = len(x.Tid)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Wager != nil {
			l = options.Size(x.Wager)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Reward != nil {
			l = options.Size(x.Reward)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Earner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Performance)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Earner) > 0 {
			i -= len(x.Earner)
			copy(dAtA[i:], x.Earner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Earner)))
			i--
			dAtA[i] = 0x32
		}
		if x.Reward != nil {
			encoded, err := options.Marshal(x.Reward)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if x.Wager != nil {
			encoded, err := options.Marshal(x.Wager)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Tid) > 0 {
			i -= len(x.Tid)
			copy(dAtA[i:], x.Tid)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Tid)))
			i--
			dAtA[i] = 0x1a
		}
		if x.Mid != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Mid))
			i--
			dAtA[i] = 0x10
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Performance)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Performance: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Performance: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
				}
				x.Mid = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Mid |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tid", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Tid = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Wager", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Wager == nil {
					x.Wager = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Wager); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Reward == nil {
					x.Reward = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Reward); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Earner", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Earner = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: intravision/gainsharing/performance.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Performance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mid    int32         `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	Tid    string        `protobuf:"bytes,3,opt,name=tid,proto3" json:"tid,omitempty"`
	Wager  *v1beta1.Coin `protobuf:"bytes,4,opt,name=wager,proto3" json:"wager,omitempty"`
	Reward *v1beta1.Coin `protobuf:"bytes,5,opt,name=reward,proto3" json:"reward,omitempty"`
	Earner string        `protobuf:"bytes,6,opt,name=earner,proto3" json:"earner,omitempty"`
}

func (x *Performance) Reset() {
	*x = Performance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intravision_gainsharing_performance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Performance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Performance) ProtoMessage() {}

// Deprecated: Use Performance.ProtoReflect.Descriptor instead.
func (*Performance) Descriptor() ([]byte, []int) {
	return file_intravision_gainsharing_performance_proto_rawDescGZIP(), []int{0}
}

func (x *Performance) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Performance) GetMid() int32 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *Performance) GetTid() string {
	if x != nil {
		return x.Tid
	}
	return ""
}

func (x *Performance) GetWager() *v1beta1.Coin {
	if x != nil {
		return x.Wager
	}
	return nil
}

func (x *Performance) GetReward() *v1beta1.Coin {
	if x != nil {
		return x.Reward
	}
	return nil
}

func (x *Performance) GetEarner() string {
	if x != nil {
		return x.Earner
	}
	return ""
}

var File_intravision_gainsharing_performance_proto protoreflect.FileDescriptor

var file_intravision_gainsharing_performance_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x61,
	0x69, 0x6e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x74,
	0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x69, 0x6e, 0x73, 0x68, 0x61,
	0x72, 0x69, 0x6e, 0x67, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x0b, 0x50,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x35,
	0x0a, 0x05, 0x77, 0x61, 0x67, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x05,
	0x77, 0x61, 0x67, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x42, 0xe9, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x69, 0x6e, 0x73,
	0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x10, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x67, 0x6f, 0x2f,
	0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x61, 0x69, 0x6e,
	0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0xa2, 0x02, 0x03, 0x49, 0x47, 0x58, 0xaa, 0x02, 0x17,
	0x49, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x61, 0x69, 0x6e,
	0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x17, 0x49, 0x6e, 0x74, 0x72, 0x61, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x61, 0x69, 0x6e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e,
	0x67, 0xe2, 0x02, 0x23, 0x49, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5c,
	0x47, 0x61, 0x69, 0x6e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x49, 0x6e, 0x74, 0x72, 0x61, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x47, 0x61, 0x69, 0x6e, 0x73, 0x68, 0x61, 0x72, 0x69,
	0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_intravision_gainsharing_performance_proto_rawDescOnce sync.Once
	file_intravision_gainsharing_performance_proto_rawDescData = file_intravision_gainsharing_performance_proto_rawDesc
)

func file_intravision_gainsharing_performance_proto_rawDescGZIP() []byte {
	file_intravision_gainsharing_performance_proto_rawDescOnce.Do(func() {
		file_intravision_gainsharing_performance_proto_rawDescData = protoimpl.X.CompressGZIP(file_intravision_gainsharing_performance_proto_rawDescData)
	})
	return file_intravision_gainsharing_performance_proto_rawDescData
}

var file_intravision_gainsharing_performance_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_intravision_gainsharing_performance_proto_goTypes = []interface{}{
	(*Performance)(nil),  // 0: intravision.gainsharing.Performance
	(*v1beta1.Coin)(nil), // 1: cosmos.base.v1beta1.Coin
}
var file_intravision_gainsharing_performance_proto_depIdxs = []int32{
	1, // 0: intravision.gainsharing.Performance.wager:type_name -> cosmos.base.v1beta1.Coin
	1, // 1: intravision.gainsharing.Performance.reward:type_name -> cosmos.base.v1beta1.Coin
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_intravision_gainsharing_performance_proto_init() }
func file_intravision_gainsharing_performance_proto_init() {
	if File_intravision_gainsharing_performance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_intravision_gainsharing_performance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Performance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_intravision_gainsharing_performance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_intravision_gainsharing_performance_proto_goTypes,
		DependencyIndexes: file_intravision_gainsharing_performance_proto_depIdxs,
		MessageInfos:      file_intravision_gainsharing_performance_proto_msgTypes,
	}.Build()
	File_intravision_gainsharing_performance_proto = out.File
	file_intravision_gainsharing_performance_proto_rawDesc = nil
	file_intravision_gainsharing_performance_proto_goTypes = nil
	file_intravision_gainsharing_performance_proto_depIdxs = nil
}
