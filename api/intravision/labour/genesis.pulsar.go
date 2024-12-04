// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package labour

import (
	_ "cosmossdk.io/api/amino"
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

var _ protoreflect.List = (*_GenesisState_2_list)(nil)

type _GenesisState_2_list struct {
	list *[]*Task
}

func (x *_GenesisState_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Task)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Task)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_2_list) AppendMutable() protoreflect.Value {
	v := new(Task)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_2_list) NewElement() protoreflect.Value {
	v := new(Task)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_2_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_4_list)(nil)

type _GenesisState_4_list struct {
	list *[]*Activity
}

func (x *_GenesisState_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Activity)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Activity)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_4_list) AppendMutable() protoreflect.Value {
	v := new(Activity)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_4_list) NewElement() protoreflect.Value {
	v := new(Activity)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_4_list) IsValid() bool {
	return x.list != nil
}

var (
	md_GenesisState               protoreflect.MessageDescriptor
	fd_GenesisState_params        protoreflect.FieldDescriptor
	fd_GenesisState_taskList      protoreflect.FieldDescriptor
	fd_GenesisState_taskCount     protoreflect.FieldDescriptor
	fd_GenesisState_activityList  protoreflect.FieldDescriptor
	fd_GenesisState_activityCount protoreflect.FieldDescriptor
)

func init() {
	file_intravision_labour_genesis_proto_init()
	md_GenesisState = File_intravision_labour_genesis_proto.Messages().ByName("GenesisState")
	fd_GenesisState_params = md_GenesisState.Fields().ByName("params")
	fd_GenesisState_taskList = md_GenesisState.Fields().ByName("taskList")
	fd_GenesisState_taskCount = md_GenesisState.Fields().ByName("taskCount")
	fd_GenesisState_activityList = md_GenesisState.Fields().ByName("activityList")
	fd_GenesisState_activityCount = md_GenesisState.Fields().ByName("activityCount")
}

var _ protoreflect.Message = (*fastReflection_GenesisState)(nil)

type fastReflection_GenesisState GenesisState

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GenesisState)(x)
}

func (x *GenesisState) slowProtoReflect() protoreflect.Message {
	mi := &file_intravision_labour_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GenesisState_messageType fastReflection_GenesisState_messageType
var _ protoreflect.MessageType = fastReflection_GenesisState_messageType{}

type fastReflection_GenesisState_messageType struct{}

func (x fastReflection_GenesisState_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GenesisState)(nil)
}
func (x fastReflection_GenesisState_messageType) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}
func (x fastReflection_GenesisState_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GenesisState) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GenesisState) Type() protoreflect.MessageType {
	return _fastReflection_GenesisState_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GenesisState) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GenesisState) Interface() protoreflect.ProtoMessage {
	return (*GenesisState)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GenesisState) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Params != nil {
		value := protoreflect.ValueOfMessage(x.Params.ProtoReflect())
		if !f(fd_GenesisState_params, value) {
			return
		}
	}
	if len(x.TaskList) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_2_list{list: &x.TaskList})
		if !f(fd_GenesisState_taskList, value) {
			return
		}
	}
	if x.TaskCount != uint64(0) {
		value := protoreflect.ValueOfUint64(x.TaskCount)
		if !f(fd_GenesisState_taskCount, value) {
			return
		}
	}
	if len(x.ActivityList) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_4_list{list: &x.ActivityList})
		if !f(fd_GenesisState_activityList, value) {
			return
		}
	}
	if x.ActivityCount != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ActivityCount)
		if !f(fd_GenesisState_activityCount, value) {
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
func (x *fastReflection_GenesisState) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "intravision.labour.GenesisState.params":
		return x.Params != nil
	case "intravision.labour.GenesisState.taskList":
		return len(x.TaskList) != 0
	case "intravision.labour.GenesisState.taskCount":
		return x.TaskCount != uint64(0)
	case "intravision.labour.GenesisState.activityList":
		return len(x.ActivityList) != 0
	case "intravision.labour.GenesisState.activityCount":
		return x.ActivityCount != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.labour.GenesisState"))
		}
		panic(fmt.Errorf("message intravision.labour.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "intravision.labour.GenesisState.params":
		x.Params = nil
	case "intravision.labour.GenesisState.taskList":
		x.TaskList = nil
	case "intravision.labour.GenesisState.taskCount":
		x.TaskCount = uint64(0)
	case "intravision.labour.GenesisState.activityList":
		x.ActivityList = nil
	case "intravision.labour.GenesisState.activityCount":
		x.ActivityCount = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.labour.GenesisState"))
		}
		panic(fmt.Errorf("message intravision.labour.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GenesisState) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "intravision.labour.GenesisState.params":
		value := x.Params
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "intravision.labour.GenesisState.taskList":
		if len(x.TaskList) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_2_list{})
		}
		listValue := &_GenesisState_2_list{list: &x.TaskList}
		return protoreflect.ValueOfList(listValue)
	case "intravision.labour.GenesisState.taskCount":
		value := x.TaskCount
		return protoreflect.ValueOfUint64(value)
	case "intravision.labour.GenesisState.activityList":
		if len(x.ActivityList) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_4_list{})
		}
		listValue := &_GenesisState_4_list{list: &x.ActivityList}
		return protoreflect.ValueOfList(listValue)
	case "intravision.labour.GenesisState.activityCount":
		value := x.ActivityCount
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.labour.GenesisState"))
		}
		panic(fmt.Errorf("message intravision.labour.GenesisState does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GenesisState) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "intravision.labour.GenesisState.params":
		x.Params = value.Message().Interface().(*Params)
	case "intravision.labour.GenesisState.taskList":
		lv := value.List()
		clv := lv.(*_GenesisState_2_list)
		x.TaskList = *clv.list
	case "intravision.labour.GenesisState.taskCount":
		x.TaskCount = value.Uint()
	case "intravision.labour.GenesisState.activityList":
		lv := value.List()
		clv := lv.(*_GenesisState_4_list)
		x.ActivityList = *clv.list
	case "intravision.labour.GenesisState.activityCount":
		x.ActivityCount = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.labour.GenesisState"))
		}
		panic(fmt.Errorf("message intravision.labour.GenesisState does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GenesisState) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "intravision.labour.GenesisState.params":
		if x.Params == nil {
			x.Params = new(Params)
		}
		return protoreflect.ValueOfMessage(x.Params.ProtoReflect())
	case "intravision.labour.GenesisState.taskList":
		if x.TaskList == nil {
			x.TaskList = []*Task{}
		}
		value := &_GenesisState_2_list{list: &x.TaskList}
		return protoreflect.ValueOfList(value)
	case "intravision.labour.GenesisState.activityList":
		if x.ActivityList == nil {
			x.ActivityList = []*Activity{}
		}
		value := &_GenesisState_4_list{list: &x.ActivityList}
		return protoreflect.ValueOfList(value)
	case "intravision.labour.GenesisState.taskCount":
		panic(fmt.Errorf("field taskCount of message intravision.labour.GenesisState is not mutable"))
	case "intravision.labour.GenesisState.activityCount":
		panic(fmt.Errorf("field activityCount of message intravision.labour.GenesisState is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.labour.GenesisState"))
		}
		panic(fmt.Errorf("message intravision.labour.GenesisState does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GenesisState) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "intravision.labour.GenesisState.params":
		m := new(Params)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "intravision.labour.GenesisState.taskList":
		list := []*Task{}
		return protoreflect.ValueOfList(&_GenesisState_2_list{list: &list})
	case "intravision.labour.GenesisState.taskCount":
		return protoreflect.ValueOfUint64(uint64(0))
	case "intravision.labour.GenesisState.activityList":
		list := []*Activity{}
		return protoreflect.ValueOfList(&_GenesisState_4_list{list: &list})
	case "intravision.labour.GenesisState.activityCount":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.labour.GenesisState"))
		}
		panic(fmt.Errorf("message intravision.labour.GenesisState does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GenesisState) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in intravision.labour.GenesisState", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GenesisState) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GenesisState) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GenesisState) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GenesisState)
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
		if x.Params != nil {
			l = options.Size(x.Params)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.TaskList) > 0 {
			for _, e := range x.TaskList {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.TaskCount != 0 {
			n += 1 + runtime.Sov(uint64(x.TaskCount))
		}
		if len(x.ActivityList) > 0 {
			for _, e := range x.ActivityList {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.ActivityCount != 0 {
			n += 1 + runtime.Sov(uint64(x.ActivityCount))
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
		x := input.Message.Interface().(*GenesisState)
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
		if x.ActivityCount != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ActivityCount))
			i--
			dAtA[i] = 0x28
		}
		if len(x.ActivityList) > 0 {
			for iNdEx := len(x.ActivityList) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.ActivityList[iNdEx])
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
		}
		if x.TaskCount != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TaskCount))
			i--
			dAtA[i] = 0x18
		}
		if len(x.TaskList) > 0 {
			for iNdEx := len(x.TaskList) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.TaskList[iNdEx])
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
				dAtA[i] = 0x12
			}
		}
		if x.Params != nil {
			encoded, err := options.Marshal(x.Params)
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
			dAtA[i] = 0xa
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
		x := input.Message.Interface().(*GenesisState)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
				if x.Params == nil {
					x.Params = &Params{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Params); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TaskList", wireType)
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
				x.TaskList = append(x.TaskList, &Task{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.TaskList[len(x.TaskList)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TaskCount", wireType)
				}
				x.TaskCount = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TaskCount |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ActivityList", wireType)
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
				x.ActivityList = append(x.ActivityList, &Activity{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ActivityList[len(x.ActivityList)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ActivityCount", wireType)
				}
				x.ActivityCount = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ActivityCount |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
// source: intravision/labour/genesis.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GenesisState defines the labour module's genesis state.
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params defines all the parameters of the module.
	Params        *Params     `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	TaskList      []*Task     `protobuf:"bytes,2,rep,name=taskList,proto3" json:"taskList,omitempty"`
	TaskCount     uint64      `protobuf:"varint,3,opt,name=taskCount,proto3" json:"taskCount,omitempty"`
	ActivityList  []*Activity `protobuf:"bytes,4,rep,name=activityList,proto3" json:"activityList,omitempty"`
	ActivityCount uint64      `protobuf:"varint,5,opt,name=activityCount,proto3" json:"activityCount,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intravision_labour_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_intravision_labour_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenesisState) GetTaskList() []*Task {
	if x != nil {
		return x.TaskList
	}
	return nil
}

func (x *GenesisState) GetTaskCount() uint64 {
	if x != nil {
		return x.TaskCount
	}
	return 0
}

func (x *GenesisState) GetActivityList() []*Activity {
	if x != nil {
		return x.ActivityList
	}
	return nil
}

func (x *GenesisState) GetActivityCount() uint64 {
	if x != nil {
		return x.ActivityCount
	}
	return 0
}

var File_intravision_labour_genesis_proto protoreflect.FileDescriptor

var file_intravision_labour_genesis_proto_rawDesc = []byte{
	0x0a, 0x20, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x61,
	0x62, 0x6f, 0x75, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x6c, 0x61, 0x62, 0x6f, 0x75, 0x72, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d,
	0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x61, 0x62,
	0x6f, 0x75, 0x72, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x61,
	0x62, 0x6f, 0x75, 0x72, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x61, 0x62,
	0x6f, 0x75, 0x72, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x95, 0x02, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x6c, 0x61, 0x62, 0x6f, 0x75, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42,
	0x09, 0xc8, 0xde, 0x1f, 0x00, 0xa8, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x3a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x6c, 0x61, 0x62, 0x6f, 0x75, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x04,
	0xc8, 0xde, 0x1f, 0x00, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x0c,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x6c, 0x61, 0x62, 0x6f, 0x75, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xc7, 0x01, 0x0a, 0x16, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6c,
	0x61, 0x62, 0x6f, 0x75, 0x72, 0x42, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x61,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x61,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x61, 0x62, 0x6f, 0x75, 0x72, 0xa2, 0x02, 0x03,
	0x49, 0x4c, 0x58, 0xaa, 0x02, 0x12, 0x49, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x6f, 0x75, 0x72, 0xca, 0x02, 0x12, 0x49, 0x6e, 0x74, 0x72, 0x61,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x4c, 0x61, 0x62, 0x6f, 0x75, 0x72, 0xe2, 0x02, 0x1e,
	0x49, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x4c, 0x61, 0x62, 0x6f,
	0x75, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x13, 0x49, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x4c, 0x61,
	0x62, 0x6f, 0x75, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_intravision_labour_genesis_proto_rawDescOnce sync.Once
	file_intravision_labour_genesis_proto_rawDescData = file_intravision_labour_genesis_proto_rawDesc
)

func file_intravision_labour_genesis_proto_rawDescGZIP() []byte {
	file_intravision_labour_genesis_proto_rawDescOnce.Do(func() {
		file_intravision_labour_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_intravision_labour_genesis_proto_rawDescData)
	})
	return file_intravision_labour_genesis_proto_rawDescData
}

var file_intravision_labour_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_intravision_labour_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil), // 0: intravision.labour.GenesisState
	(*Params)(nil),       // 1: intravision.labour.Params
	(*Task)(nil),         // 2: intravision.labour.Task
	(*Activity)(nil),     // 3: intravision.labour.Activity
}
var file_intravision_labour_genesis_proto_depIdxs = []int32{
	1, // 0: intravision.labour.GenesisState.params:type_name -> intravision.labour.Params
	2, // 1: intravision.labour.GenesisState.taskList:type_name -> intravision.labour.Task
	3, // 2: intravision.labour.GenesisState.activityList:type_name -> intravision.labour.Activity
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_intravision_labour_genesis_proto_init() }
func file_intravision_labour_genesis_proto_init() {
	if File_intravision_labour_genesis_proto != nil {
		return
	}
	file_intravision_labour_params_proto_init()
	file_intravision_labour_task_proto_init()
	file_intravision_labour_activity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_intravision_labour_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
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
			RawDescriptor: file_intravision_labour_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_intravision_labour_genesis_proto_goTypes,
		DependencyIndexes: file_intravision_labour_genesis_proto_depIdxs,
		MessageInfos:      file_intravision_labour_genesis_proto_msgTypes,
	}.Build()
	File_intravision_labour_genesis_proto = out.File
	file_intravision_labour_genesis_proto_rawDesc = nil
	file_intravision_labour_genesis_proto_goTypes = nil
	file_intravision_labour_genesis_proto_depIdxs = nil
}
