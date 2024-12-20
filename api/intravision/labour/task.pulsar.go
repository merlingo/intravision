// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package labour

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Task            protoreflect.MessageDescriptor
	fd_Task_id         protoreflect.FieldDescriptor
	fd_Task_taskId     protoreflect.FieldDescriptor
	fd_Task_assigner   protoreflect.FieldDescriptor
	fd_Task_state      protoreflect.FieldDescriptor
	fd_Task_beginTask  protoreflect.FieldDescriptor
	fd_Task_deadline   protoreflect.FieldDescriptor
	fd_Task_finishTask protoreflect.FieldDescriptor
	fd_Task_wager      protoreflect.FieldDescriptor
)

func init() {
	file_intravision_labour_task_proto_init()
	md_Task = File_intravision_labour_task_proto.Messages().ByName("Task")
	fd_Task_id = md_Task.Fields().ByName("id")
	fd_Task_taskId = md_Task.Fields().ByName("taskId")
	fd_Task_assigner = md_Task.Fields().ByName("assigner")
	fd_Task_state = md_Task.Fields().ByName("state")
	fd_Task_beginTask = md_Task.Fields().ByName("beginTask")
	fd_Task_deadline = md_Task.Fields().ByName("deadline")
	fd_Task_finishTask = md_Task.Fields().ByName("finishTask")
	fd_Task_wager = md_Task.Fields().ByName("wager")
}

var _ protoreflect.Message = (*fastReflection_Task)(nil)

type fastReflection_Task Task

func (x *Task) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Task)(x)
}

func (x *Task) slowProtoReflect() protoreflect.Message {
	mi := &file_intravision_labour_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Task_messageType fastReflection_Task_messageType
var _ protoreflect.MessageType = fastReflection_Task_messageType{}

type fastReflection_Task_messageType struct{}

func (x fastReflection_Task_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Task)(nil)
}
func (x fastReflection_Task_messageType) New() protoreflect.Message {
	return new(fastReflection_Task)
}
func (x fastReflection_Task_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Task
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Task) Descriptor() protoreflect.MessageDescriptor {
	return md_Task
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Task) Type() protoreflect.MessageType {
	return _fastReflection_Task_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Task) New() protoreflect.Message {
	return new(fastReflection_Task)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Task) Interface() protoreflect.ProtoMessage {
	return (*Task)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Task) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_Task_id, value) {
			return
		}
	}
	if x.TaskId != "" {
		value := protoreflect.ValueOfString(x.TaskId)
		if !f(fd_Task_taskId, value) {
			return
		}
	}
	if x.Assigner != "" {
		value := protoreflect.ValueOfString(x.Assigner)
		if !f(fd_Task_assigner, value) {
			return
		}
	}
	if x.State != int32(0) {
		value := protoreflect.ValueOfInt32(x.State)
		if !f(fd_Task_state, value) {
			return
		}
	}
	if x.BeginTask != uint64(0) {
		value := protoreflect.ValueOfUint64(x.BeginTask)
		if !f(fd_Task_beginTask, value) {
			return
		}
	}
	if x.Deadline != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Deadline)
		if !f(fd_Task_deadline, value) {
			return
		}
	}
	if x.FinishTask != uint64(0) {
		value := protoreflect.ValueOfUint64(x.FinishTask)
		if !f(fd_Task_finishTask, value) {
			return
		}
	}
	if x.Wager != "" {
		value := protoreflect.ValueOfString(x.Wager)
		if !f(fd_Task_wager, value) {
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
func (x *fastReflection_Task) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "intravision.labour.Task.id":
		return x.Id != uint64(0)
	case "intravision.labour.Task.taskId":
		return x.TaskId != ""
	case "intravision.labour.Task.assigner":
		return x.Assigner != ""
	case "intravision.labour.Task.state":
		return x.State != int32(0)
	case "intravision.labour.Task.beginTask":
		return x.BeginTask != uint64(0)
	case "intravision.labour.Task.deadline":
		return x.Deadline != uint64(0)
	case "intravision.labour.Task.finishTask":
		return x.FinishTask != uint64(0)
	case "intravision.labour.Task.wager":
		return x.Wager != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.labour.Task"))
		}
		panic(fmt.Errorf("message intravision.labour.Task does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Task) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "intravision.labour.Task.id":
		x.Id = uint64(0)
	case "intravision.labour.Task.taskId":
		x.TaskId = ""
	case "intravision.labour.Task.assigner":
		x.Assigner = ""
	case "intravision.labour.Task.state":
		x.State = int32(0)
	case "intravision.labour.Task.beginTask":
		x.BeginTask = uint64(0)
	case "intravision.labour.Task.deadline":
		x.Deadline = uint64(0)
	case "intravision.labour.Task.finishTask":
		x.FinishTask = uint64(0)
	case "intravision.labour.Task.wager":
		x.Wager = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.labour.Task"))
		}
		panic(fmt.Errorf("message intravision.labour.Task does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Task) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "intravision.labour.Task.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "intravision.labour.Task.taskId":
		value := x.TaskId
		return protoreflect.ValueOfString(value)
	case "intravision.labour.Task.assigner":
		value := x.Assigner
		return protoreflect.ValueOfString(value)
	case "intravision.labour.Task.state":
		value := x.State
		return protoreflect.ValueOfInt32(value)
	case "intravision.labour.Task.beginTask":
		value := x.BeginTask
		return protoreflect.ValueOfUint64(value)
	case "intravision.labour.Task.deadline":
		value := x.Deadline
		return protoreflect.ValueOfUint64(value)
	case "intravision.labour.Task.finishTask":
		value := x.FinishTask
		return protoreflect.ValueOfUint64(value)
	case "intravision.labour.Task.wager":
		value := x.Wager
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.labour.Task"))
		}
		panic(fmt.Errorf("message intravision.labour.Task does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Task) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "intravision.labour.Task.id":
		x.Id = value.Uint()
	case "intravision.labour.Task.taskId":
		x.TaskId = value.Interface().(string)
	case "intravision.labour.Task.assigner":
		x.Assigner = value.Interface().(string)
	case "intravision.labour.Task.state":
		x.State = int32(value.Int())
	case "intravision.labour.Task.beginTask":
		x.BeginTask = value.Uint()
	case "intravision.labour.Task.deadline":
		x.Deadline = value.Uint()
	case "intravision.labour.Task.finishTask":
		x.FinishTask = value.Uint()
	case "intravision.labour.Task.wager":
		x.Wager = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.labour.Task"))
		}
		panic(fmt.Errorf("message intravision.labour.Task does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Task) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "intravision.labour.Task.id":
		panic(fmt.Errorf("field id of message intravision.labour.Task is not mutable"))
	case "intravision.labour.Task.taskId":
		panic(fmt.Errorf("field taskId of message intravision.labour.Task is not mutable"))
	case "intravision.labour.Task.assigner":
		panic(fmt.Errorf("field assigner of message intravision.labour.Task is not mutable"))
	case "intravision.labour.Task.state":
		panic(fmt.Errorf("field state of message intravision.labour.Task is not mutable"))
	case "intravision.labour.Task.beginTask":
		panic(fmt.Errorf("field beginTask of message intravision.labour.Task is not mutable"))
	case "intravision.labour.Task.deadline":
		panic(fmt.Errorf("field deadline of message intravision.labour.Task is not mutable"))
	case "intravision.labour.Task.finishTask":
		panic(fmt.Errorf("field finishTask of message intravision.labour.Task is not mutable"))
	case "intravision.labour.Task.wager":
		panic(fmt.Errorf("field wager of message intravision.labour.Task is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.labour.Task"))
		}
		panic(fmt.Errorf("message intravision.labour.Task does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Task) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "intravision.labour.Task.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "intravision.labour.Task.taskId":
		return protoreflect.ValueOfString("")
	case "intravision.labour.Task.assigner":
		return protoreflect.ValueOfString("")
	case "intravision.labour.Task.state":
		return protoreflect.ValueOfInt32(int32(0))
	case "intravision.labour.Task.beginTask":
		return protoreflect.ValueOfUint64(uint64(0))
	case "intravision.labour.Task.deadline":
		return protoreflect.ValueOfUint64(uint64(0))
	case "intravision.labour.Task.finishTask":
		return protoreflect.ValueOfUint64(uint64(0))
	case "intravision.labour.Task.wager":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: intravision.labour.Task"))
		}
		panic(fmt.Errorf("message intravision.labour.Task does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Task) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in intravision.labour.Task", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Task) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Task) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Task) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Task) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Task)
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
		l = len(x.TaskId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Assigner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.State != 0 {
			n += 1 + runtime.Sov(uint64(x.State))
		}
		if x.BeginTask != 0 {
			n += 1 + runtime.Sov(uint64(x.BeginTask))
		}
		if x.Deadline != 0 {
			n += 1 + runtime.Sov(uint64(x.Deadline))
		}
		if x.FinishTask != 0 {
			n += 1 + runtime.Sov(uint64(x.FinishTask))
		}
		l = len(x.Wager)
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
		x := input.Message.Interface().(*Task)
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
		if len(x.Wager) > 0 {
			i -= len(x.Wager)
			copy(dAtA[i:], x.Wager)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Wager)))
			i--
			dAtA[i] = 0x42
		}
		if x.FinishTask != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.FinishTask))
			i--
			dAtA[i] = 0x38
		}
		if x.Deadline != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Deadline))
			i--
			dAtA[i] = 0x30
		}
		if x.BeginTask != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BeginTask))
			i--
			dAtA[i] = 0x28
		}
		if x.State != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.State))
			i--
			dAtA[i] = 0x20
		}
		if len(x.Assigner) > 0 {
			i -= len(x.Assigner)
			copy(dAtA[i:], x.Assigner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Assigner)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.TaskId) > 0 {
			i -= len(x.TaskId)
			copy(dAtA[i:], x.TaskId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TaskId)))
			i--
			dAtA[i] = 0x12
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
		x := input.Message.Interface().(*Task)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Task: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Task: illegal tag %d (wire type %d)", fieldNum, wire)
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
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TaskId", wireType)
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
				x.TaskId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Assigner", wireType)
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
				x.Assigner = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
				}
				x.State = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.State |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BeginTask", wireType)
				}
				x.BeginTask = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BeginTask |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
				}
				x.Deadline = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Deadline |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 7:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FinishTask", wireType)
				}
				x.FinishTask = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.FinishTask |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Wager", wireType)
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
				x.Wager = string(dAtA[iNdEx:postIndex])
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
// source: intravision/labour/task.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TaskId     string `protobuf:"bytes,2,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Assigner   string `protobuf:"bytes,3,opt,name=assigner,proto3" json:"assigner,omitempty"`
	State      int32  `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
	BeginTask  uint64 `protobuf:"varint,5,opt,name=beginTask,proto3" json:"beginTask,omitempty"`
	Deadline   uint64 `protobuf:"varint,6,opt,name=deadline,proto3" json:"deadline,omitempty"`
	FinishTask uint64 `protobuf:"varint,7,opt,name=finishTask,proto3" json:"finishTask,omitempty"`
	Wager      string `protobuf:"bytes,8,opt,name=wager,proto3" json:"wager,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intravision_labour_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_intravision_labour_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Task) GetAssigner() string {
	if x != nil {
		return x.Assigner
	}
	return ""
}

func (x *Task) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *Task) GetBeginTask() uint64 {
	if x != nil {
		return x.BeginTask
	}
	return 0
}

func (x *Task) GetDeadline() uint64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

func (x *Task) GetFinishTask() uint64 {
	if x != nil {
		return x.FinishTask
	}
	return 0
}

func (x *Task) GetWager() string {
	if x != nil {
		return x.Wager
	}
	return ""
}

var File_intravision_labour_task_proto protoreflect.FileDescriptor

var file_intravision_labour_task_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x61,
	0x62, 0x6f, 0x75, 0x72, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6c, 0x61, 0x62,
	0x6f, 0x75, 0x72, 0x22, 0xd0, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54,
	0x61, 0x73, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x77, 0x61, 0x67, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x77, 0x61, 0x67, 0x65, 0x72, 0x42, 0xc4, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6c, 0x61, 0x62, 0x6f, 0x75,
	0x72, 0x42, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x72, 0x6c, 0x69,
	0x6e, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x6c, 0x61, 0x62, 0x6f, 0x75, 0x72, 0xa2, 0x02, 0x03, 0x49, 0x4c, 0x58, 0xaa, 0x02, 0x12, 0x49,
	0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x6f, 0x75,
	0x72, 0xca, 0x02, 0x12, 0x49, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5c,
	0x4c, 0x61, 0x62, 0x6f, 0x75, 0x72, 0xe2, 0x02, 0x1e, 0x49, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x4c, 0x61, 0x62, 0x6f, 0x75, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x49, 0x6e, 0x74, 0x72, 0x61, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x4c, 0x61, 0x62, 0x6f, 0x75, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_intravision_labour_task_proto_rawDescOnce sync.Once
	file_intravision_labour_task_proto_rawDescData = file_intravision_labour_task_proto_rawDesc
)

func file_intravision_labour_task_proto_rawDescGZIP() []byte {
	file_intravision_labour_task_proto_rawDescOnce.Do(func() {
		file_intravision_labour_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_intravision_labour_task_proto_rawDescData)
	})
	return file_intravision_labour_task_proto_rawDescData
}

var file_intravision_labour_task_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_intravision_labour_task_proto_goTypes = []interface{}{
	(*Task)(nil), // 0: intravision.labour.Task
}
var file_intravision_labour_task_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_intravision_labour_task_proto_init() }
func file_intravision_labour_task_proto_init() {
	if File_intravision_labour_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_intravision_labour_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
			RawDescriptor: file_intravision_labour_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_intravision_labour_task_proto_goTypes,
		DependencyIndexes: file_intravision_labour_task_proto_depIdxs,
		MessageInfos:      file_intravision_labour_task_proto_msgTypes,
	}.Build()
	File_intravision_labour_task_proto = out.File
	file_intravision_labour_task_proto_rawDesc = nil
	file_intravision_labour_task_proto_goTypes = nil
	file_intravision_labour_task_proto_depIdxs = nil
}
