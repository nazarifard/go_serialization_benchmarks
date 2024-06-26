// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package pulsar

import (
	binary "encoding/binary"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	math "math"
	reflect "reflect"
	sync "sync"
)

var (
	md_PulsarBufA          protoreflect.MessageDescriptor
	fd_PulsarBufA_name     protoreflect.FieldDescriptor
	fd_PulsarBufA_birthDay protoreflect.FieldDescriptor
	fd_PulsarBufA_phone    protoreflect.FieldDescriptor
	fd_PulsarBufA_siblings protoreflect.FieldDescriptor
	fd_PulsarBufA_spouse   protoreflect.FieldDescriptor
	fd_PulsarBufA_money    protoreflect.FieldDescriptor
)

func init() {
	file_structdef_pulsar_proto_init()
	md_PulsarBufA = File_structdef_pulsar_proto.Messages().ByName("PulsarBufA")
	fd_PulsarBufA_name = md_PulsarBufA.Fields().ByName("name")
	fd_PulsarBufA_birthDay = md_PulsarBufA.Fields().ByName("birthDay")
	fd_PulsarBufA_phone = md_PulsarBufA.Fields().ByName("phone")
	fd_PulsarBufA_siblings = md_PulsarBufA.Fields().ByName("siblings")
	fd_PulsarBufA_spouse = md_PulsarBufA.Fields().ByName("spouse")
	fd_PulsarBufA_money = md_PulsarBufA.Fields().ByName("money")
}

var _ protoreflect.Message = (*fastReflection_PulsarBufA)(nil)

type fastReflection_PulsarBufA PulsarBufA

func (x *PulsarBufA) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PulsarBufA)(x)
}

func (x *PulsarBufA) slowProtoReflect() protoreflect.Message {
	mi := &file_structdef_pulsar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PulsarBufA_messageType fastReflection_PulsarBufA_messageType
var _ protoreflect.MessageType = fastReflection_PulsarBufA_messageType{}

type fastReflection_PulsarBufA_messageType struct{}

func (x fastReflection_PulsarBufA_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PulsarBufA)(nil)
}
func (x fastReflection_PulsarBufA_messageType) New() protoreflect.Message {
	return new(fastReflection_PulsarBufA)
}
func (x fastReflection_PulsarBufA_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PulsarBufA
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PulsarBufA) Descriptor() protoreflect.MessageDescriptor {
	return md_PulsarBufA
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PulsarBufA) Type() protoreflect.MessageType {
	return _fastReflection_PulsarBufA_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PulsarBufA) New() protoreflect.Message {
	return new(fastReflection_PulsarBufA)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PulsarBufA) Interface() protoreflect.ProtoMessage {
	return (*PulsarBufA)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PulsarBufA) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_PulsarBufA_name, value) {
			return
		}
	}
	if x.BirthDay != int64(0) {
		value := protoreflect.ValueOfInt64(x.BirthDay)
		if !f(fd_PulsarBufA_birthDay, value) {
			return
		}
	}
	if x.Phone != "" {
		value := protoreflect.ValueOfString(x.Phone)
		if !f(fd_PulsarBufA_phone, value) {
			return
		}
	}
	if x.Siblings != int32(0) {
		value := protoreflect.ValueOfInt32(x.Siblings)
		if !f(fd_PulsarBufA_siblings, value) {
			return
		}
	}
	if x.Spouse != false {
		value := protoreflect.ValueOfBool(x.Spouse)
		if !f(fd_PulsarBufA_spouse, value) {
			return
		}
	}
	if x.Money != float64(0) || math.Signbit(x.Money) {
		value := protoreflect.ValueOfFloat64(x.Money)
		if !f(fd_PulsarBufA_money, value) {
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
func (x *fastReflection_PulsarBufA) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "goserbench.PulsarBufA.name":
		return x.Name != ""
	case "goserbench.PulsarBufA.birthDay":
		return x.BirthDay != int64(0)
	case "goserbench.PulsarBufA.phone":
		return x.Phone != ""
	case "goserbench.PulsarBufA.siblings":
		return x.Siblings != int32(0)
	case "goserbench.PulsarBufA.spouse":
		return x.Spouse != false
	case "goserbench.PulsarBufA.money":
		return x.Money != float64(0) || math.Signbit(x.Money)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goserbench.PulsarBufA"))
		}
		panic(fmt.Errorf("message goserbench.PulsarBufA does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PulsarBufA) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "goserbench.PulsarBufA.name":
		x.Name = ""
	case "goserbench.PulsarBufA.birthDay":
		x.BirthDay = int64(0)
	case "goserbench.PulsarBufA.phone":
		x.Phone = ""
	case "goserbench.PulsarBufA.siblings":
		x.Siblings = int32(0)
	case "goserbench.PulsarBufA.spouse":
		x.Spouse = false
	case "goserbench.PulsarBufA.money":
		x.Money = float64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goserbench.PulsarBufA"))
		}
		panic(fmt.Errorf("message goserbench.PulsarBufA does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PulsarBufA) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "goserbench.PulsarBufA.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "goserbench.PulsarBufA.birthDay":
		value := x.BirthDay
		return protoreflect.ValueOfInt64(value)
	case "goserbench.PulsarBufA.phone":
		value := x.Phone
		return protoreflect.ValueOfString(value)
	case "goserbench.PulsarBufA.siblings":
		value := x.Siblings
		return protoreflect.ValueOfInt32(value)
	case "goserbench.PulsarBufA.spouse":
		value := x.Spouse
		return protoreflect.ValueOfBool(value)
	case "goserbench.PulsarBufA.money":
		value := x.Money
		return protoreflect.ValueOfFloat64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goserbench.PulsarBufA"))
		}
		panic(fmt.Errorf("message goserbench.PulsarBufA does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_PulsarBufA) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "goserbench.PulsarBufA.name":
		x.Name = value.Interface().(string)
	case "goserbench.PulsarBufA.birthDay":
		x.BirthDay = value.Int()
	case "goserbench.PulsarBufA.phone":
		x.Phone = value.Interface().(string)
	case "goserbench.PulsarBufA.siblings":
		x.Siblings = int32(value.Int())
	case "goserbench.PulsarBufA.spouse":
		x.Spouse = value.Bool()
	case "goserbench.PulsarBufA.money":
		x.Money = value.Float()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goserbench.PulsarBufA"))
		}
		panic(fmt.Errorf("message goserbench.PulsarBufA does not contain field %s", fd.FullName()))
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
func (x *fastReflection_PulsarBufA) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goserbench.PulsarBufA.name":
		panic(fmt.Errorf("field name of message goserbench.PulsarBufA is not mutable"))
	case "goserbench.PulsarBufA.birthDay":
		panic(fmt.Errorf("field birthDay of message goserbench.PulsarBufA is not mutable"))
	case "goserbench.PulsarBufA.phone":
		panic(fmt.Errorf("field phone of message goserbench.PulsarBufA is not mutable"))
	case "goserbench.PulsarBufA.siblings":
		panic(fmt.Errorf("field siblings of message goserbench.PulsarBufA is not mutable"))
	case "goserbench.PulsarBufA.spouse":
		panic(fmt.Errorf("field spouse of message goserbench.PulsarBufA is not mutable"))
	case "goserbench.PulsarBufA.money":
		panic(fmt.Errorf("field money of message goserbench.PulsarBufA is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goserbench.PulsarBufA"))
		}
		panic(fmt.Errorf("message goserbench.PulsarBufA does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PulsarBufA) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goserbench.PulsarBufA.name":
		return protoreflect.ValueOfString("")
	case "goserbench.PulsarBufA.birthDay":
		return protoreflect.ValueOfInt64(int64(0))
	case "goserbench.PulsarBufA.phone":
		return protoreflect.ValueOfString("")
	case "goserbench.PulsarBufA.siblings":
		return protoreflect.ValueOfInt32(int32(0))
	case "goserbench.PulsarBufA.spouse":
		return protoreflect.ValueOfBool(false)
	case "goserbench.PulsarBufA.money":
		return protoreflect.ValueOfFloat64(float64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goserbench.PulsarBufA"))
		}
		panic(fmt.Errorf("message goserbench.PulsarBufA does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PulsarBufA) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in goserbench.PulsarBufA", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PulsarBufA) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PulsarBufA) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_PulsarBufA) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PulsarBufA) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PulsarBufA)
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
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.BirthDay != 0 {
			n += 1 + runtime.Sov(uint64(x.BirthDay))
		}
		l = len(x.Phone)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Siblings != 0 {
			n += 1 + runtime.Sov(uint64(x.Siblings))
		}
		if x.Spouse {
			n += 2
		}
		if x.Money != 0 || math.Signbit(x.Money) {
			n += 9
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
		x := input.Message.Interface().(*PulsarBufA)
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
		if x.Money != 0 || math.Signbit(x.Money) {
			i -= 8
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(x.Money))))
			i--
			dAtA[i] = 0x31
		}
		if x.Spouse {
			i--
			if x.Spouse {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x28
		}
		if x.Siblings != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Siblings))
			i--
			dAtA[i] = 0x20
		}
		if len(x.Phone) > 0 {
			i -= len(x.Phone)
			copy(dAtA[i:], x.Phone)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Phone)))
			i--
			dAtA[i] = 0x1a
		}
		if x.BirthDay != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BirthDay))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
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
		x := input.Message.Interface().(*PulsarBufA)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PulsarBufA: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PulsarBufA: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BirthDay", wireType)
				}
				x.BirthDay = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BirthDay |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Phone", wireType)
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
				x.Phone = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Siblings", wireType)
				}
				x.Siblings = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Siblings |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Spouse", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Spouse = bool(v != 0)
			case 6:
				if wireType != 1 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Money", wireType)
				}
				var v uint64
				if (iNdEx + 8) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				x.Money = float64(math.Float64frombits(v))
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
// 	protoc        v3.21.12
// source: structdef-pulsar.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PulsarBufA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BirthDay int64   `protobuf:"varint,2,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	Phone    string  `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Siblings int32   `protobuf:"varint,4,opt,name=siblings,proto3" json:"siblings,omitempty"`
	Spouse   bool    `protobuf:"varint,5,opt,name=spouse,proto3" json:"spouse,omitempty"`
	Money    float64 `protobuf:"fixed64,6,opt,name=money,proto3" json:"money,omitempty"`
}

func (x *PulsarBufA) Reset() {
	*x = PulsarBufA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_structdef_pulsar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PulsarBufA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PulsarBufA) ProtoMessage() {}

// Deprecated: Use PulsarBufA.ProtoReflect.Descriptor instead.
func (*PulsarBufA) Descriptor() ([]byte, []int) {
	return file_structdef_pulsar_proto_rawDescGZIP(), []int{0}
}

func (x *PulsarBufA) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PulsarBufA) GetBirthDay() int64 {
	if x != nil {
		return x.BirthDay
	}
	return 0
}

func (x *PulsarBufA) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *PulsarBufA) GetSiblings() int32 {
	if x != nil {
		return x.Siblings
	}
	return 0
}

func (x *PulsarBufA) GetSpouse() bool {
	if x != nil {
		return x.Spouse
	}
	return false
}

func (x *PulsarBufA) GetMoney() float64 {
	if x != nil {
		return x.Money
	}
	return 0
}

var File_structdef_pulsar_proto protoreflect.FileDescriptor

var file_structdef_pulsar_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x64, 0x65, 0x66, 0x2d, 0x70, 0x75, 0x6c, 0x73,
	0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x73, 0x65, 0x72, 0x62,
	0x65, 0x6e, 0x63, 0x68, 0x22, 0x9c, 0x01, 0x0a, 0x0a, 0x50, 0x75, 0x6c, 0x73, 0x61, 0x72, 0x42,
	0x75, 0x66, 0x41, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x44, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x44, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x62,
	0x6c, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x69, 0x62,
	0x6c, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x6f, 0x75, 0x73, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x70, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x63, 0x74, 0x68, 0x6f, 0x6d, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x5f,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x65,
	0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x3b, 0x67, 0x6f, 0x73, 0x65, 0x72, 0x62, 0x65,
	0x6e, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_structdef_pulsar_proto_rawDescOnce sync.Once
	file_structdef_pulsar_proto_rawDescData = file_structdef_pulsar_proto_rawDesc
)

func file_structdef_pulsar_proto_rawDescGZIP() []byte {
	file_structdef_pulsar_proto_rawDescOnce.Do(func() {
		file_structdef_pulsar_proto_rawDescData = protoimpl.X.CompressGZIP(file_structdef_pulsar_proto_rawDescData)
	})
	return file_structdef_pulsar_proto_rawDescData
}

var file_structdef_pulsar_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_structdef_pulsar_proto_goTypes = []interface{}{
	(*PulsarBufA)(nil), // 0: goserbench.PulsarBufA
}
var file_structdef_pulsar_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_structdef_pulsar_proto_init() }
func file_structdef_pulsar_proto_init() {
	if File_structdef_pulsar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_structdef_pulsar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PulsarBufA); i {
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
			RawDescriptor: file_structdef_pulsar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_structdef_pulsar_proto_goTypes,
		DependencyIndexes: file_structdef_pulsar_proto_depIdxs,
		MessageInfos:      file_structdef_pulsar_proto_msgTypes,
	}.Build()
	File_structdef_pulsar_proto = out.File
	file_structdef_pulsar_proto_rawDesc = nil
	file_structdef_pulsar_proto_goTypes = nil
	file_structdef_pulsar_proto_depIdxs = nil
}
