// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/test_suite.proto

package resultstore // import "google.golang.org/genproto/googleapis/devtools/resultstore/v2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The result of running a test case.
type TestCase_Result int32

const (
	// The implicit default enum value. Do not use.
	TestCase_RESULT_UNSPECIFIED TestCase_Result = 0
	// Test case ran to completion. Look for failures or errors to determine
	// whether it passed, failed, or errored.
	TestCase_COMPLETED TestCase_Result = 1
	// Test case started but did not complete because the test harness received
	// a signal and decided to stop running tests.
	TestCase_INTERRUPTED TestCase_Result = 2
	// Test case was not started because the test harness received a SIGINT or
	// timed out.
	TestCase_CANCELLED TestCase_Result = 3
	// Test case was not run because the user or process running the test
	// specified a filter that excluded this test case.
	TestCase_FILTERED TestCase_Result = 4
	// Test case was not run to completion because the test case decided it
	// should not be run (eg. due to a failed assumption in a JUnit4 test).
	// Per-test setup or tear-down may or may not have run.
	TestCase_SKIPPED TestCase_Result = 5
	// The test framework did not run the test case because it was labeled as
	// suppressed.  Eg. if someone temporarily disables a failing test.
	TestCase_SUPPRESSED TestCase_Result = 6
)

var TestCase_Result_name = map[int32]string{
	0: "RESULT_UNSPECIFIED",
	1: "COMPLETED",
	2: "INTERRUPTED",
	3: "CANCELLED",
	4: "FILTERED",
	5: "SKIPPED",
	6: "SUPPRESSED",
}
var TestCase_Result_value = map[string]int32{
	"RESULT_UNSPECIFIED": 0,
	"COMPLETED":          1,
	"INTERRUPTED":        2,
	"CANCELLED":          3,
	"FILTERED":           4,
	"SKIPPED":            5,
	"SUPPRESSED":         6,
}

func (x TestCase_Result) String() string {
	return proto.EnumName(TestCase_Result_name, int32(x))
}
func (TestCase_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_test_suite_9672634cf69bbe80, []int{2, 0}
}

// The result of running a test suite, as reported in a <testsuite> element of
// an XML log.
type TestSuite struct {
	// The full name of this suite, as reported in the name attribute. For Java
	// tests, this is normally the fully qualified class name. Eg.
	// "com.google.common.hash.BloomFilterTest".
	SuiteName string `protobuf:"bytes,1,opt,name=suite_name,json=suiteName,proto3" json:"suite_name,omitempty"`
	// The results of the test cases and test suites contained in this suite,
	// as reported in the <testcase> and <testsuite> elements contained within
	// this <testsuite>.
	Tests []*Test `protobuf:"bytes,2,rep,name=tests,proto3" json:"tests,omitempty"`
	// Failures reported in <failure> elements within this <testsuite>.
	Failures []*TestFailure `protobuf:"bytes,3,rep,name=failures,proto3" json:"failures,omitempty"`
	// Errors reported in <error> elements within this <testsuite>.
	Errors []*TestError `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
	// The timing for the entire TestSuite, as reported by the time attribute.
	Timing *Timing `protobuf:"bytes,6,opt,name=timing,proto3" json:"timing,omitempty"`
	// Arbitrary name-value pairs, as reported in custom attributes or in a
	// <properties> element within this <testsuite>. Multiple properties are
	// allowed with the same key. Properties will be returned in lexicographical
	// order by key.
	Properties []*Property `protobuf:"bytes,7,rep,name=properties,proto3" json:"properties,omitempty"`
	// Files produced by this test suite, as reported by undeclared output
	// annotations.
	// The file IDs must be unique within this list. Duplicate file IDs will
	// result in an error. Files will be returned in lexicographical order by ID.
	Files                []*File  `protobuf:"bytes,8,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestSuite) Reset()         { *m = TestSuite{} }
func (m *TestSuite) String() string { return proto.CompactTextString(m) }
func (*TestSuite) ProtoMessage()    {}
func (*TestSuite) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_suite_9672634cf69bbe80, []int{0}
}
func (m *TestSuite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestSuite.Unmarshal(m, b)
}
func (m *TestSuite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestSuite.Marshal(b, m, deterministic)
}
func (dst *TestSuite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestSuite.Merge(dst, src)
}
func (m *TestSuite) XXX_Size() int {
	return xxx_messageInfo_TestSuite.Size(m)
}
func (m *TestSuite) XXX_DiscardUnknown() {
	xxx_messageInfo_TestSuite.DiscardUnknown(m)
}

var xxx_messageInfo_TestSuite proto.InternalMessageInfo

func (m *TestSuite) GetSuiteName() string {
	if m != nil {
		return m.SuiteName
	}
	return ""
}

func (m *TestSuite) GetTests() []*Test {
	if m != nil {
		return m.Tests
	}
	return nil
}

func (m *TestSuite) GetFailures() []*TestFailure {
	if m != nil {
		return m.Failures
	}
	return nil
}

func (m *TestSuite) GetErrors() []*TestError {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *TestSuite) GetTiming() *Timing {
	if m != nil {
		return m.Timing
	}
	return nil
}

func (m *TestSuite) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *TestSuite) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

// The result of running a test case or test suite. JUnit3 TestDecorators are
// represented as a TestSuite with a single test.
type Test struct {
	// Either a TestCase of a TestSuite
	//
	// Types that are valid to be assigned to TestType:
	//	*Test_TestCase
	//	*Test_TestSuite
	TestType             isTest_TestType `protobuf_oneof:"test_type"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_suite_9672634cf69bbe80, []int{1}
}
func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (dst *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(dst, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

type isTest_TestType interface {
	isTest_TestType()
}

type Test_TestCase struct {
	TestCase *TestCase `protobuf:"bytes,1,opt,name=test_case,json=testCase,proto3,oneof"`
}

type Test_TestSuite struct {
	TestSuite *TestSuite `protobuf:"bytes,2,opt,name=test_suite,json=testSuite,proto3,oneof"`
}

func (*Test_TestCase) isTest_TestType() {}

func (*Test_TestSuite) isTest_TestType() {}

func (m *Test) GetTestType() isTest_TestType {
	if m != nil {
		return m.TestType
	}
	return nil
}

func (m *Test) GetTestCase() *TestCase {
	if x, ok := m.GetTestType().(*Test_TestCase); ok {
		return x.TestCase
	}
	return nil
}

func (m *Test) GetTestSuite() *TestSuite {
	if x, ok := m.GetTestType().(*Test_TestSuite); ok {
		return x.TestSuite
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Test) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Test_OneofMarshaler, _Test_OneofUnmarshaler, _Test_OneofSizer, []interface{}{
		(*Test_TestCase)(nil),
		(*Test_TestSuite)(nil),
	}
}

func _Test_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Test)
	// test_type
	switch x := m.TestType.(type) {
	case *Test_TestCase:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TestCase); err != nil {
			return err
		}
	case *Test_TestSuite:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TestSuite); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Test.TestType has unexpected type %T", x)
	}
	return nil
}

func _Test_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Test)
	switch tag {
	case 1: // test_type.test_case
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TestCase)
		err := b.DecodeMessage(msg)
		m.TestType = &Test_TestCase{msg}
		return true, err
	case 2: // test_type.test_suite
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TestSuite)
		err := b.DecodeMessage(msg)
		m.TestType = &Test_TestSuite{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Test_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Test)
	// test_type
	switch x := m.TestType.(type) {
	case *Test_TestCase:
		s := proto.Size(x.TestCase)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Test_TestSuite:
		s := proto.Size(x.TestSuite)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The result of running a test case, as reported in a <testcase> element of
// an XML log.
type TestCase struct {
	// The name of the test case, as reported in the name attribute. For Java,
	// this is normally the method name. Eg. "testBasic".
	CaseName string `protobuf:"bytes,1,opt,name=case_name,json=caseName,proto3" json:"case_name,omitempty"`
	// The name of the class in which the test case was defined, as reported in
	// the classname attribute. For Java, this is normally the fully qualified
	// class name. Eg. "com.google.common.hash.BloomFilterTest".
	ClassName string `protobuf:"bytes,2,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	// An enum reported in the result attribute that is used in conjunction with
	// failures and errors below to report the outcome.
	Result TestCase_Result `protobuf:"varint,3,opt,name=result,proto3,enum=google.devtools.resultstore.v2.TestCase_Result" json:"result,omitempty"`
	// Failures reported in <failure> elements within this <testcase>.
	Failures []*TestFailure `protobuf:"bytes,4,rep,name=failures,proto3" json:"failures,omitempty"`
	// Errors reported in <error> elements within this <testcase>.
	Errors []*TestError `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty"`
	// The timing for the TestCase, as reported by the time attribute.
	Timing *Timing `protobuf:"bytes,7,opt,name=timing,proto3" json:"timing,omitempty"`
	// Arbitrary name-value pairs, as reported in custom attributes or in a
	// <properties> element within this <testcase>. Multiple properties are
	// allowed with the same key. Properties will be returned in lexicographical
	// order by key.
	Properties []*Property `protobuf:"bytes,8,rep,name=properties,proto3" json:"properties,omitempty"`
	// Files produced by this test case, as reported by undeclared output
	// annotations.
	// The file IDs must be unique within this list. Duplicate file IDs will
	// result in an error. Files will be returned in lexicographical order by ID.
	Files                []*File  `protobuf:"bytes,9,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestCase) Reset()         { *m = TestCase{} }
func (m *TestCase) String() string { return proto.CompactTextString(m) }
func (*TestCase) ProtoMessage()    {}
func (*TestCase) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_suite_9672634cf69bbe80, []int{2}
}
func (m *TestCase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestCase.Unmarshal(m, b)
}
func (m *TestCase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestCase.Marshal(b, m, deterministic)
}
func (dst *TestCase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestCase.Merge(dst, src)
}
func (m *TestCase) XXX_Size() int {
	return xxx_messageInfo_TestCase.Size(m)
}
func (m *TestCase) XXX_DiscardUnknown() {
	xxx_messageInfo_TestCase.DiscardUnknown(m)
}

var xxx_messageInfo_TestCase proto.InternalMessageInfo

func (m *TestCase) GetCaseName() string {
	if m != nil {
		return m.CaseName
	}
	return ""
}

func (m *TestCase) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *TestCase) GetResult() TestCase_Result {
	if m != nil {
		return m.Result
	}
	return TestCase_RESULT_UNSPECIFIED
}

func (m *TestCase) GetFailures() []*TestFailure {
	if m != nil {
		return m.Failures
	}
	return nil
}

func (m *TestCase) GetErrors() []*TestError {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *TestCase) GetTiming() *Timing {
	if m != nil {
		return m.Timing
	}
	return nil
}

func (m *TestCase) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *TestCase) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

// Represents a violated assertion, as reported in a <failure> element within a
// <testcase>. Some languages allow assertions to be made without stopping the
// test case when they're violated, leading to multiple TestFailures. For Java,
// multiple TestFailures are used to represent a chained exception.
type TestFailure struct {
	// The exception message reported in the message attribute. Typically short,
	// but may be multi-line. Eg. "Expected 'foo' but was 'bar'".
	FailureMessage string `protobuf:"bytes,1,opt,name=failure_message,json=failureMessage,proto3" json:"failure_message,omitempty"`
	// The type of the exception being thrown, reported in the type attribute.
	// Eg: "org.junit.ComparisonFailure"
	ExceptionType string `protobuf:"bytes,2,opt,name=exception_type,json=exceptionType,proto3" json:"exception_type,omitempty"`
	// The stack trace reported as the content of the <failure> element, often in
	// a CDATA block. This contains one line for each stack frame, each including
	// a method/function name, a class/file name, and a line number. Most recent
	// call is usually first, but not for Python stack traces. May contain the
	// exception_type and message.
	StackTrace           string   `protobuf:"bytes,3,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestFailure) Reset()         { *m = TestFailure{} }
func (m *TestFailure) String() string { return proto.CompactTextString(m) }
func (*TestFailure) ProtoMessage()    {}
func (*TestFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_suite_9672634cf69bbe80, []int{3}
}
func (m *TestFailure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestFailure.Unmarshal(m, b)
}
func (m *TestFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestFailure.Marshal(b, m, deterministic)
}
func (dst *TestFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestFailure.Merge(dst, src)
}
func (m *TestFailure) XXX_Size() int {
	return xxx_messageInfo_TestFailure.Size(m)
}
func (m *TestFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_TestFailure.DiscardUnknown(m)
}

var xxx_messageInfo_TestFailure proto.InternalMessageInfo

func (m *TestFailure) GetFailureMessage() string {
	if m != nil {
		return m.FailureMessage
	}
	return ""
}

func (m *TestFailure) GetExceptionType() string {
	if m != nil {
		return m.ExceptionType
	}
	return ""
}

func (m *TestFailure) GetStackTrace() string {
	if m != nil {
		return m.StackTrace
	}
	return ""
}

// Represents an exception that prevented a test case from completing, as
// reported in an <error> element within a <testcase>. For Java, multiple
// TestErrors are used to represent a chained exception.
type TestError struct {
	// The exception message, as reported in the message attribute. Typically
	// short, but may be multi-line. Eg. "argument cannot be null".
	ErrorMessage string `protobuf:"bytes,1,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// The type of the exception being thrown, reported in the type attribute.
	// For Java, this is a fully qualified Throwable class name.
	// Eg: "java.lang.IllegalArgumentException"
	ExceptionType string `protobuf:"bytes,2,opt,name=exception_type,json=exceptionType,proto3" json:"exception_type,omitempty"`
	// The stack trace reported as the content of the <error> element, often in
	// a CDATA block. This contains one line for each stack frame, each including
	// a method/function name, a class/file name, and a line number. Most recent
	// call is usually first, but not for Python stack traces. May contain the
	// exception_type and message.
	StackTrace           string   `protobuf:"bytes,3,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestError) Reset()         { *m = TestError{} }
func (m *TestError) String() string { return proto.CompactTextString(m) }
func (*TestError) ProtoMessage()    {}
func (*TestError) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_suite_9672634cf69bbe80, []int{4}
}
func (m *TestError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestError.Unmarshal(m, b)
}
func (m *TestError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestError.Marshal(b, m, deterministic)
}
func (dst *TestError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestError.Merge(dst, src)
}
func (m *TestError) XXX_Size() int {
	return xxx_messageInfo_TestError.Size(m)
}
func (m *TestError) XXX_DiscardUnknown() {
	xxx_messageInfo_TestError.DiscardUnknown(m)
}

var xxx_messageInfo_TestError proto.InternalMessageInfo

func (m *TestError) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *TestError) GetExceptionType() string {
	if m != nil {
		return m.ExceptionType
	}
	return ""
}

func (m *TestError) GetStackTrace() string {
	if m != nil {
		return m.StackTrace
	}
	return ""
}

func init() {
	proto.RegisterType((*TestSuite)(nil), "google.devtools.resultstore.v2.TestSuite")
	proto.RegisterType((*Test)(nil), "google.devtools.resultstore.v2.Test")
	proto.RegisterType((*TestCase)(nil), "google.devtools.resultstore.v2.TestCase")
	proto.RegisterType((*TestFailure)(nil), "google.devtools.resultstore.v2.TestFailure")
	proto.RegisterType((*TestError)(nil), "google.devtools.resultstore.v2.TestError")
	proto.RegisterEnum("google.devtools.resultstore.v2.TestCase_Result", TestCase_Result_name, TestCase_Result_value)
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/test_suite.proto", fileDescriptor_test_suite_9672634cf69bbe80)
}

var fileDescriptor_test_suite_9672634cf69bbe80 = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0xb6, 0xcb, 0x92, 0x93, 0xad, 0xab, 0x7c, 0x81, 0xa2, 0xa1, 0x41, 0x55, 0xfe,
	0x75, 0x9a, 0x94, 0x48, 0xe5, 0x0e, 0x24, 0xa4, 0xad, 0x75, 0xb7, 0x42, 0x57, 0x22, 0x27, 0xbd,
	0xe1, 0xa6, 0x0a, 0xc1, 0x8b, 0x22, 0x92, 0x3a, 0xc4, 0xde, 0xb4, 0x71, 0xc1, 0xf3, 0xf0, 0x08,
	0xbc, 0x0c, 0xef, 0x82, 0x62, 0xa7, 0xa3, 0x20, 0x41, 0x8b, 0xe8, 0x5d, 0xf3, 0xd9, 0xbf, 0x73,
	0x3e, 0xfb, 0x7c, 0xaa, 0xc1, 0x8d, 0x19, 0x8b, 0x53, 0xea, 0x7e, 0xa0, 0xd7, 0x82, 0xb1, 0x94,
	0xbb, 0x05, 0xe5, 0x57, 0xa9, 0xe0, 0x82, 0x15, 0xd4, 0xbd, 0xee, 0xb9, 0x82, 0x72, 0x31, 0xe3,
	0x57, 0x89, 0xa0, 0x4e, 0x5e, 0x30, 0xc1, 0xd0, 0x03, 0x05, 0x38, 0x0b, 0xc0, 0x59, 0x02, 0x9c,
	0xeb, 0xde, 0xc1, 0xf1, 0x8a, 0x82, 0x11, 0xcb, 0x32, 0x36, 0x57, 0xc5, 0x0e, 0x8e, 0x56, 0x6c,
	0xbe, 0x4c, 0xd2, 0xaa, 0x6f, 0xe7, 0x5b, 0x1d, 0xcc, 0x80, 0x72, 0xe1, 0x97, 0x5e, 0xd0, 0x21,
	0x80, 0x34, 0x35, 0x9b, 0x87, 0x19, 0xb5, 0xb5, 0xb6, 0xd6, 0x35, 0x89, 0x29, 0x95, 0x49, 0x98,
	0x51, 0xf4, 0x02, 0xb6, 0x4b, 0xe3, 0xdc, 0xae, 0xb5, 0xeb, 0x5d, 0xab, 0xf7, 0xd8, 0xf9, 0xbb,
	0x69, 0xa7, 0x2c, 0x4c, 0x14, 0x82, 0xce, 0xc0, 0xb8, 0x0c, 0x93, 0xf4, 0xaa, 0xa0, 0xdc, 0xae,
	0x4b, 0xfc, 0x78, 0x1d, 0x7c, 0xa8, 0x18, 0x72, 0x07, 0xa3, 0x13, 0xd0, 0x69, 0x51, 0xb0, 0x82,
	0xdb, 0x0d, 0x59, 0xe6, 0x68, 0x9d, 0x32, 0xb8, 0x24, 0x48, 0x05, 0xa2, 0x57, 0xa0, 0x8b, 0x24,
	0x4b, 0xe6, 0xb1, 0xad, 0xb7, 0xb5, 0xae, 0xd5, 0x7b, 0xba, 0xb2, 0x84, 0xdc, 0x4d, 0x2a, 0x0a,
	0x9d, 0x03, 0xe4, 0x05, 0xcb, 0x69, 0x21, 0x12, 0xca, 0xed, 0x1d, 0x69, 0xa3, 0xbb, 0xaa, 0x86,
	0xa7, 0x88, 0x5b, 0xb2, 0xc4, 0x96, 0x37, 0x5a, 0x0e, 0x83, 0xdb, 0xc6, 0x7a, 0x37, 0x3a, 0x4c,
	0x52, 0x4a, 0x14, 0xd2, 0xf9, 0xaa, 0x41, 0xa3, 0x3c, 0x1b, 0x3a, 0x03, 0x53, 0xe6, 0x29, 0x0a,
	0xb9, 0x1a, 0xda, 0x1a, 0x6e, 0x4a, 0xb0, 0x1f, 0x72, 0x7a, 0xbe, 0x45, 0x0c, 0x51, 0xfd, 0x46,
	0xaf, 0x01, 0x7e, 0x06, 0xd3, 0xae, 0xc9, 0x4a, 0x6b, 0x5d, 0xaf, 0x4c, 0xcf, 0xf9, 0x16, 0x91,
	0x3e, 0xe4, 0xc7, 0xa9, 0x55, 0x99, 0x12, 0xb7, 0x39, 0xed, 0x7c, 0x6f, 0x80, 0xb1, 0xe8, 0x88,
	0xee, 0x83, 0x59, 0x3a, 0x5d, 0xce, 0x98, 0x51, 0x0a, 0x32, 0x62, 0x87, 0x00, 0x51, 0x1a, 0x72,
	0xae, 0x56, 0x6b, 0x2a, 0x81, 0x52, 0x91, 0xcb, 0x67, 0xa0, 0xab, 0xf6, 0x76, 0xbd, 0xad, 0x75,
	0x9b, 0x3d, 0x77, 0xdd, 0x73, 0x3a, 0x44, 0xea, 0xa4, 0xc2, 0x7f, 0x89, 0x63, 0x63, 0x33, 0x71,
	0xdc, 0xfe, 0xff, 0x38, 0xee, 0x6c, 0x20, 0x8e, 0xc6, 0x26, 0xe2, 0x68, 0xfe, 0x7b, 0x1c, 0x3f,
	0x83, 0xae, 0xee, 0x18, 0xdd, 0x03, 0x44, 0xb0, 0x3f, 0x1d, 0x07, 0xb3, 0xe9, 0xc4, 0xf7, 0x70,
	0x7f, 0x34, 0x1c, 0xe1, 0x41, 0x6b, 0x0b, 0xed, 0x81, 0xd9, 0x7f, 0x7b, 0xe1, 0x8d, 0x71, 0x80,
	0x07, 0x2d, 0x0d, 0xed, 0x83, 0x35, 0x9a, 0x04, 0x98, 0x90, 0xa9, 0x57, 0x0a, 0x35, 0xb9, 0x7e,
	0x32, 0xe9, 0xe3, 0xf1, 0x18, 0x0f, 0x5a, 0x75, 0xb4, 0x0b, 0xc6, 0x70, 0x34, 0x0e, 0x30, 0xc1,
	0x83, 0x56, 0x03, 0x59, 0xb0, 0xe3, 0xbf, 0x19, 0x79, 0x1e, 0x1e, 0xb4, 0xb6, 0x51, 0x13, 0xc0,
	0x9f, 0x7a, 0x1e, 0xc1, 0xbe, 0x8f, 0x07, 0x2d, 0xbd, 0xf3, 0x05, 0xac, 0xa5, 0xe9, 0xa0, 0x67,
	0xb0, 0x5f, 0xcd, 0x67, 0x96, 0x51, 0xce, 0xc3, 0x78, 0x91, 0xb3, 0x66, 0x25, 0x5f, 0x28, 0x15,
	0x3d, 0x81, 0x26, 0xbd, 0x89, 0x68, 0x2e, 0x12, 0x36, 0x97, 0x49, 0xad, 0x12, 0xb7, 0x77, 0xa7,
	0x06, 0xb7, 0x39, 0x45, 0x0f, 0xc1, 0xe2, 0x22, 0x8c, 0x3e, 0xce, 0x44, 0x11, 0x46, 0x54, 0x46,
	0xcf, 0x24, 0x20, 0xa5, 0xa0, 0x54, 0x3a, 0x37, 0xea, 0x4f, 0x54, 0x8e, 0x15, 0x3d, 0x82, 0x3d,
	0x39, 0xd8, 0xdf, 0x7a, 0xef, 0x4a, 0x71, 0xc3, 0x9d, 0x4f, 0x3f, 0x41, 0x27, 0x62, 0xd9, 0x8a,
	0x39, 0x79, 0xda, 0xbb, 0x51, 0xb5, 0x23, 0x66, 0x69, 0x38, 0x8f, 0x1d, 0x56, 0xc4, 0x6e, 0x4c,
	0xe7, 0xf2, 0x0d, 0xa8, 0xde, 0xaa, 0x30, 0x4f, 0xf8, 0x9f, 0x5e, 0x8c, 0x97, 0x4b, 0x9f, 0xef,
	0x75, 0x49, 0x3d, 0xff, 0x11, 0x00, 0x00, 0xff, 0xff, 0x38, 0x57, 0x5a, 0x20, 0xe4, 0x06, 0x00,
	0x00,
}
