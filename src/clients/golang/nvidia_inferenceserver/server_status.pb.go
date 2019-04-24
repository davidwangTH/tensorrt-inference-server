// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/core/server_status.proto

package nvidia_inferenceserver

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//@@
//@@.. cpp:enum:: ModelReadyState
//@@
//@@   Readiness status for models.
//@@
type ModelReadyState int32

const (
	//@@  .. cpp:enumerator:: ModelReadyState::MODEL_UNKNOWN = 0
	//@@
	//@@     The model is in an unknown state. The model is not available for
	//@@     inferencing.
	//@@
	ModelReadyState_MODEL_UNKNOWN ModelReadyState = 0
	//@@  .. cpp:enumerator:: ModelReadyState::MODEL_READY = 1
	//@@
	//@@     The model is ready and available for inferencing.
	//@@
	ModelReadyState_MODEL_READY ModelReadyState = 1
	//@@  .. cpp:enumerator:: ModelReadyState::MODEL_UNAVAILABLE = 2
	//@@
	//@@     The model is unavailable, indicating that the model failed to
	//@@     load or has been implicitly or explicitly unloaded. The model is
	//@@     not available for inferencing.
	//@@
	ModelReadyState_MODEL_UNAVAILABLE ModelReadyState = 2
	//@@  .. cpp:enumerator:: ModelReadyState::MODEL_LOADING = 3
	//@@
	//@@     The model is being loaded by the inference server. The model is
	//@@     not available for inferencing.
	//@@
	ModelReadyState_MODEL_LOADING ModelReadyState = 3
	//@@  .. cpp:enumerator:: ModelReadyState::MODEL_UNLOADING = 4
	//@@
	//@@     The model is being unloaded by the inference server. The model is
	//@@     not available for inferencing.
	//@@
	ModelReadyState_MODEL_UNLOADING ModelReadyState = 4
)

var ModelReadyState_name = map[int32]string{
	0: "MODEL_UNKNOWN",
	1: "MODEL_READY",
	2: "MODEL_UNAVAILABLE",
	3: "MODEL_LOADING",
	4: "MODEL_UNLOADING",
}

var ModelReadyState_value = map[string]int32{
	"MODEL_UNKNOWN":     0,
	"MODEL_READY":       1,
	"MODEL_UNAVAILABLE": 2,
	"MODEL_LOADING":     3,
	"MODEL_UNLOADING":   4,
}

func (x ModelReadyState) String() string {
	return proto.EnumName(ModelReadyState_name, int32(x))
}

func (ModelReadyState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67e9cedab6951959, []int{0}
}

//@@
//@@.. cpp:enum:: ServerReadyState
//@@
//@@   Readiness status for the inference server.
//@@
type ServerReadyState int32

const (
	//@@  .. cpp:enumerator:: ServerReadyState::SERVER_INVALID = 0
	//@@
	//@@     The server is in an invalid state and will likely not
	//@@     response correctly to any requests.
	//@@
	ServerReadyState_SERVER_INVALID ServerReadyState = 0
	//@@  .. cpp:enumerator:: ServerReadyState::SERVER_INITIALIZING = 1
	//@@
	//@@     The server is initializing.
	//@@
	ServerReadyState_SERVER_INITIALIZING ServerReadyState = 1
	//@@  .. cpp:enumerator:: ServerReadyState::SERVER_READY = 2
	//@@
	//@@     The server is ready and accepting requests.
	//@@
	ServerReadyState_SERVER_READY ServerReadyState = 2
	//@@  .. cpp:enumerator:: ServerReadyState::SERVER_EXITING = 3
	//@@
	//@@     The server is exiting and will not respond to requests.
	//@@
	ServerReadyState_SERVER_EXITING ServerReadyState = 3
	//@@  .. cpp:enumerator:: ServerReadyState::SERVER_FAILED_TO_INITIALIZE = 10
	//@@
	//@@     The server did not initialize correctly. Most requests will fail.
	//@@
	ServerReadyState_SERVER_FAILED_TO_INITIALIZE ServerReadyState = 10
)

var ServerReadyState_name = map[int32]string{
	0:  "SERVER_INVALID",
	1:  "SERVER_INITIALIZING",
	2:  "SERVER_READY",
	3:  "SERVER_EXITING",
	10: "SERVER_FAILED_TO_INITIALIZE",
}

var ServerReadyState_value = map[string]int32{
	"SERVER_INVALID":              0,
	"SERVER_INITIALIZING":         1,
	"SERVER_READY":                2,
	"SERVER_EXITING":              3,
	"SERVER_FAILED_TO_INITIALIZE": 10,
}

func (x ServerReadyState) String() string {
	return proto.EnumName(ServerReadyState_name, int32(x))
}

func (ServerReadyState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67e9cedab6951959, []int{1}
}

//@@
//@@.. cpp:var:: message StatDuration
//@@
//@@   Statistic collecting a duration metric.
//@@
type StatDuration struct {
	//@@  .. cpp:var:: uint64 count
	//@@
	//@@     Cumulative number of times this metric occurred.
	//@@
	Count uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	//@@  .. cpp:var:: uint64 total_time_ns
	//@@
	//@@     Total collected duration of this metric in nanoseconds.
	//@@
	TotalTimeNs          uint64   `protobuf:"varint,2,opt,name=total_time_ns,json=totalTimeNs,proto3" json:"total_time_ns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatDuration) Reset()         { *m = StatDuration{} }
func (m *StatDuration) String() string { return proto.CompactTextString(m) }
func (*StatDuration) ProtoMessage()    {}
func (*StatDuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e9cedab6951959, []int{0}
}

func (m *StatDuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatDuration.Unmarshal(m, b)
}
func (m *StatDuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatDuration.Marshal(b, m, deterministic)
}
func (m *StatDuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatDuration.Merge(m, src)
}
func (m *StatDuration) XXX_Size() int {
	return xxx_messageInfo_StatDuration.Size(m)
}
func (m *StatDuration) XXX_DiscardUnknown() {
	xxx_messageInfo_StatDuration.DiscardUnknown(m)
}

var xxx_messageInfo_StatDuration proto.InternalMessageInfo

func (m *StatDuration) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *StatDuration) GetTotalTimeNs() uint64 {
	if m != nil {
		return m.TotalTimeNs
	}
	return 0
}

//@@
//@@.. cpp:var:: message StatusRequestStats
//@@
//@@   Statistics collected for Status requests.
//@@
type StatusRequestStats struct {
	//@@  .. cpp:var:: StatDuration success
	//@@
	//@@     Total time required to handle successful Status requests, not
	//@@     including HTTP or gRPC endpoint termination time.
	//@@
	Success              *StatDuration `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StatusRequestStats) Reset()         { *m = StatusRequestStats{} }
func (m *StatusRequestStats) String() string { return proto.CompactTextString(m) }
func (*StatusRequestStats) ProtoMessage()    {}
func (*StatusRequestStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e9cedab6951959, []int{1}
}

func (m *StatusRequestStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequestStats.Unmarshal(m, b)
}
func (m *StatusRequestStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequestStats.Marshal(b, m, deterministic)
}
func (m *StatusRequestStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequestStats.Merge(m, src)
}
func (m *StatusRequestStats) XXX_Size() int {
	return xxx_messageInfo_StatusRequestStats.Size(m)
}
func (m *StatusRequestStats) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequestStats.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequestStats proto.InternalMessageInfo

func (m *StatusRequestStats) GetSuccess() *StatDuration {
	if m != nil {
		return m.Success
	}
	return nil
}

//@@
//@@.. cpp:var:: message ProfileRequestStats
//@@
//@@   Statistics collected for Profile requests.
//@@
type ProfileRequestStats struct {
	//@@  .. cpp:var:: StatDuration success
	//@@
	//@@     Total time required to handle successful Profile requests, not
	//@@     including HTTP or gRPC endpoint termination time.
	//@@
	Success              *StatDuration `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ProfileRequestStats) Reset()         { *m = ProfileRequestStats{} }
func (m *ProfileRequestStats) String() string { return proto.CompactTextString(m) }
func (*ProfileRequestStats) ProtoMessage()    {}
func (*ProfileRequestStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e9cedab6951959, []int{2}
}

func (m *ProfileRequestStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileRequestStats.Unmarshal(m, b)
}
func (m *ProfileRequestStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileRequestStats.Marshal(b, m, deterministic)
}
func (m *ProfileRequestStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileRequestStats.Merge(m, src)
}
func (m *ProfileRequestStats) XXX_Size() int {
	return xxx_messageInfo_ProfileRequestStats.Size(m)
}
func (m *ProfileRequestStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileRequestStats.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileRequestStats proto.InternalMessageInfo

func (m *ProfileRequestStats) GetSuccess() *StatDuration {
	if m != nil {
		return m.Success
	}
	return nil
}

//@@
//@@.. cpp:var:: message HealthRequestStats
//@@
//@@   Statistics collected for Health requests.
//@@
type HealthRequestStats struct {
	//@@  .. cpp:var:: StatDuration success
	//@@
	//@@     Total time required to handle successful Health requests, not
	//@@     including HTTP or gRPC endpoint termination time.
	//@@
	Success              *StatDuration `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HealthRequestStats) Reset()         { *m = HealthRequestStats{} }
func (m *HealthRequestStats) String() string { return proto.CompactTextString(m) }
func (*HealthRequestStats) ProtoMessage()    {}
func (*HealthRequestStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e9cedab6951959, []int{3}
}

func (m *HealthRequestStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequestStats.Unmarshal(m, b)
}
func (m *HealthRequestStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequestStats.Marshal(b, m, deterministic)
}
func (m *HealthRequestStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequestStats.Merge(m, src)
}
func (m *HealthRequestStats) XXX_Size() int {
	return xxx_messageInfo_HealthRequestStats.Size(m)
}
func (m *HealthRequestStats) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequestStats.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequestStats proto.InternalMessageInfo

func (m *HealthRequestStats) GetSuccess() *StatDuration {
	if m != nil {
		return m.Success
	}
	return nil
}

//@@
//@@.. cpp:var:: message InferRequestStats
//@@
//@@   Statistics collected for Infer requests.
//@@
type InferRequestStats struct {
	//@@  .. cpp:var:: StatDuration success
	//@@
	//@@     Total time required to handle successful Infer requests, not
	//@@     including HTTP or gRPC endpoint termination time.
	//@@
	Success *StatDuration `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
	//@@  .. cpp:var:: StatDuration failed
	//@@
	//@@     Total time required to handle failed Infer requests, not
	//@@     including HTTP or gRPC endpoint termination time.
	//@@
	Failed *StatDuration `protobuf:"bytes,2,opt,name=failed,proto3" json:"failed,omitempty"`
	//@@  .. cpp:var:: StatDuration compute
	//@@
	//@@     Time required to run inferencing for an inference request;
	//@@     including time copying input tensors to GPU memory, time
	//@@     executing the model, and time copying output tensors from GPU
	//@@     memory.
	//@@
	Compute *StatDuration `protobuf:"bytes,3,opt,name=compute,proto3" json:"compute,omitempty"`
	//@@  .. cpp:var:: StatDuration queue
	//@@
	//@@     Time an inference request waits in scheduling queue for an
	//@@     available model instance.
	//@@
	Queue                *StatDuration `protobuf:"bytes,4,opt,name=queue,proto3" json:"queue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *InferRequestStats) Reset()         { *m = InferRequestStats{} }
func (m *InferRequestStats) String() string { return proto.CompactTextString(m) }
func (*InferRequestStats) ProtoMessage()    {}
func (*InferRequestStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e9cedab6951959, []int{4}
}

func (m *InferRequestStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferRequestStats.Unmarshal(m, b)
}
func (m *InferRequestStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferRequestStats.Marshal(b, m, deterministic)
}
func (m *InferRequestStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferRequestStats.Merge(m, src)
}
func (m *InferRequestStats) XXX_Size() int {
	return xxx_messageInfo_InferRequestStats.Size(m)
}
func (m *InferRequestStats) XXX_DiscardUnknown() {
	xxx_messageInfo_InferRequestStats.DiscardUnknown(m)
}

var xxx_messageInfo_InferRequestStats proto.InternalMessageInfo

func (m *InferRequestStats) GetSuccess() *StatDuration {
	if m != nil {
		return m.Success
	}
	return nil
}

func (m *InferRequestStats) GetFailed() *StatDuration {
	if m != nil {
		return m.Failed
	}
	return nil
}

func (m *InferRequestStats) GetCompute() *StatDuration {
	if m != nil {
		return m.Compute
	}
	return nil
}

func (m *InferRequestStats) GetQueue() *StatDuration {
	if m != nil {
		return m.Queue
	}
	return nil
}

//@@
//@@.. cpp:var:: message ModelVersionStatus
//@@
//@@   Status for a version of a model.
//@@
type ModelVersionStatus struct {
	//@@  .. cpp:var:: ModelReadyState ready_statue
	//@@
	//@@     Current readiness state for the model.
	//@@
	ReadyState ModelReadyState `protobuf:"varint,1,opt,name=ready_state,json=readyState,proto3,enum=nvidia.inferenceserver.ModelReadyState" json:"ready_state,omitempty"`
	//@@  .. cpp:var:: map<uint32, InferRequestStats> infer_stats
	//@@
	//@@     Inference statistics for the model, as a map from batch size
	//@@     to the statistics. A batch size will not occur in the map
	//@@     unless there has been at least one inference request of
	//@@     that batch size.
	//@@
	InferStats map[uint32]*InferRequestStats `protobuf:"bytes,2,rep,name=infer_stats,json=inferStats,proto3" json:"infer_stats,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//@@  .. cpp:var:: uint64 model_execution_count
	//@@
	//@@     Cumulative number of model executions performed for the
	//@@     model. A single model execution performs inferencing for
	//@@     the entire request batch and can perform inferencing for multiple
	//@@     requests if dynamic batching is enabled.
	//@@
	ModelExecutionCount uint64 `protobuf:"varint,3,opt,name=model_execution_count,json=modelExecutionCount,proto3" json:"model_execution_count,omitempty"`
	//@@  .. cpp:var:: uint64 model_inference_count
	//@@
	//@@     Cumulative number of model inferences performed for the
	//@@     model. Each inference in a batched request is counted as
	//@@     an individual inference.
	//@@
	ModelInferenceCount  uint64   `protobuf:"varint,4,opt,name=model_inference_count,json=modelInferenceCount,proto3" json:"model_inference_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelVersionStatus) Reset()         { *m = ModelVersionStatus{} }
func (m *ModelVersionStatus) String() string { return proto.CompactTextString(m) }
func (*ModelVersionStatus) ProtoMessage()    {}
func (*ModelVersionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e9cedab6951959, []int{5}
}

func (m *ModelVersionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelVersionStatus.Unmarshal(m, b)
}
func (m *ModelVersionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelVersionStatus.Marshal(b, m, deterministic)
}
func (m *ModelVersionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelVersionStatus.Merge(m, src)
}
func (m *ModelVersionStatus) XXX_Size() int {
	return xxx_messageInfo_ModelVersionStatus.Size(m)
}
func (m *ModelVersionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelVersionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ModelVersionStatus proto.InternalMessageInfo

func (m *ModelVersionStatus) GetReadyState() ModelReadyState {
	if m != nil {
		return m.ReadyState
	}
	return ModelReadyState_MODEL_UNKNOWN
}

func (m *ModelVersionStatus) GetInferStats() map[uint32]*InferRequestStats {
	if m != nil {
		return m.InferStats
	}
	return nil
}

func (m *ModelVersionStatus) GetModelExecutionCount() uint64 {
	if m != nil {
		return m.ModelExecutionCount
	}
	return 0
}

func (m *ModelVersionStatus) GetModelInferenceCount() uint64 {
	if m != nil {
		return m.ModelInferenceCount
	}
	return 0
}

//@@
//@@.. cpp:var:: message ModelStatus
//@@
//@@   Status for a model.
//@@
type ModelStatus struct {
	//@@  .. cpp:var:: ModelConfig config
	//@@
	//@@     The configuration for the model.
	//@@
	Config *ModelConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	//@@  .. cpp:var:: map<int64, ModelVersionStatus> version_status
	//@@
	//@@     Duration statistics for each version of the model, as a map
	//@@     from version to the status. A version will not occur in the map
	//@@     unless there has been at least one inference request of
	//@@     that model version. A version of -1 indicates the status is
	//@@     for requests for which the version could not be determined.
	//@@
	VersionStatus        map[int64]*ModelVersionStatus `protobuf:"bytes,2,rep,name=version_status,json=versionStatus,proto3" json:"version_status,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ModelStatus) Reset()         { *m = ModelStatus{} }
func (m *ModelStatus) String() string { return proto.CompactTextString(m) }
func (*ModelStatus) ProtoMessage()    {}
func (*ModelStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e9cedab6951959, []int{6}
}

func (m *ModelStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelStatus.Unmarshal(m, b)
}
func (m *ModelStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelStatus.Marshal(b, m, deterministic)
}
func (m *ModelStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelStatus.Merge(m, src)
}
func (m *ModelStatus) XXX_Size() int {
	return xxx_messageInfo_ModelStatus.Size(m)
}
func (m *ModelStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ModelStatus proto.InternalMessageInfo

func (m *ModelStatus) GetConfig() *ModelConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ModelStatus) GetVersionStatus() map[int64]*ModelVersionStatus {
	if m != nil {
		return m.VersionStatus
	}
	return nil
}

//@@
//@@.. cpp:var:: message ServerStatus
//@@
//@@   Status for the inference server.
//@@
type ServerStatus struct {
	//@@  .. cpp:var:: string id
	//@@
	//@@     The server's ID.
	//@@
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	//@@  .. cpp:var:: string version
	//@@
	//@@     The server's version.
	//@@
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	//@@  .. cpp:var:: ServerReadyState ready_state
	//@@
	//@@     Current readiness state for the server.
	//@@
	ReadyState ServerReadyState `protobuf:"varint,7,opt,name=ready_state,json=readyState,proto3,enum=nvidia.inferenceserver.ServerReadyState" json:"ready_state,omitempty"`
	//@@  .. cpp:var:: uint64 uptime_ns
	//@@
	//@@     Server uptime in nanoseconds.
	//@@
	UptimeNs uint64 `protobuf:"varint,3,opt,name=uptime_ns,json=uptimeNs,proto3" json:"uptime_ns,omitempty"`
	//@@  .. cpp:var:: map<string, ModelStatus> model_status
	//@@
	//@@     Status for each model, as a map from model name to the
	//@@     status.
	//@@
	ModelStatus map[string]*ModelStatus `protobuf:"bytes,4,rep,name=model_status,json=modelStatus,proto3" json:"model_status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//@@  .. cpp:var:: StatusRequestStats status_stats
	//@@
	//@@     Statistics for Status requests.
	//@@
	StatusStats *StatusRequestStats `protobuf:"bytes,5,opt,name=status_stats,json=statusStats,proto3" json:"status_stats,omitempty"`
	//@@  .. cpp:var:: ProfileRequestStats profile_stats
	//@@
	//@@     Statistics for Profile requests.
	//@@
	ProfileStats *ProfileRequestStats `protobuf:"bytes,6,opt,name=profile_stats,json=profileStats,proto3" json:"profile_stats,omitempty"`
	//@@  .. cpp:var:: HealthRequestStats health_stats
	//@@
	//@@     Statistics for Health requests.
	//@@
	HealthStats          *HealthRequestStats `protobuf:"bytes,8,opt,name=health_stats,json=healthStats,proto3" json:"health_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ServerStatus) Reset()         { *m = ServerStatus{} }
func (m *ServerStatus) String() string { return proto.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()    {}
func (*ServerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e9cedab6951959, []int{7}
}

func (m *ServerStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStatus.Unmarshal(m, b)
}
func (m *ServerStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStatus.Marshal(b, m, deterministic)
}
func (m *ServerStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStatus.Merge(m, src)
}
func (m *ServerStatus) XXX_Size() int {
	return xxx_messageInfo_ServerStatus.Size(m)
}
func (m *ServerStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStatus proto.InternalMessageInfo

func (m *ServerStatus) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServerStatus) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerStatus) GetReadyState() ServerReadyState {
	if m != nil {
		return m.ReadyState
	}
	return ServerReadyState_SERVER_INVALID
}

func (m *ServerStatus) GetUptimeNs() uint64 {
	if m != nil {
		return m.UptimeNs
	}
	return 0
}

func (m *ServerStatus) GetModelStatus() map[string]*ModelStatus {
	if m != nil {
		return m.ModelStatus
	}
	return nil
}

func (m *ServerStatus) GetStatusStats() *StatusRequestStats {
	if m != nil {
		return m.StatusStats
	}
	return nil
}

func (m *ServerStatus) GetProfileStats() *ProfileRequestStats {
	if m != nil {
		return m.ProfileStats
	}
	return nil
}

func (m *ServerStatus) GetHealthStats() *HealthRequestStats {
	if m != nil {
		return m.HealthStats
	}
	return nil
}

func init() {
	proto.RegisterEnum("nvidia.inferenceserver.ModelReadyState", ModelReadyState_name, ModelReadyState_value)
	proto.RegisterEnum("nvidia.inferenceserver.ServerReadyState", ServerReadyState_name, ServerReadyState_value)
	proto.RegisterType((*StatDuration)(nil), "nvidia.inferenceserver.StatDuration")
	proto.RegisterType((*StatusRequestStats)(nil), "nvidia.inferenceserver.StatusRequestStats")
	proto.RegisterType((*ProfileRequestStats)(nil), "nvidia.inferenceserver.ProfileRequestStats")
	proto.RegisterType((*HealthRequestStats)(nil), "nvidia.inferenceserver.HealthRequestStats")
	proto.RegisterType((*InferRequestStats)(nil), "nvidia.inferenceserver.InferRequestStats")
	proto.RegisterType((*ModelVersionStatus)(nil), "nvidia.inferenceserver.ModelVersionStatus")
	proto.RegisterMapType((map[uint32]*InferRequestStats)(nil), "nvidia.inferenceserver.ModelVersionStatus.InferStatsEntry")
	proto.RegisterType((*ModelStatus)(nil), "nvidia.inferenceserver.ModelStatus")
	proto.RegisterMapType((map[int64]*ModelVersionStatus)(nil), "nvidia.inferenceserver.ModelStatus.VersionStatusEntry")
	proto.RegisterType((*ServerStatus)(nil), "nvidia.inferenceserver.ServerStatus")
	proto.RegisterMapType((map[string]*ModelStatus)(nil), "nvidia.inferenceserver.ServerStatus.ModelStatusEntry")
}

func init() { proto.RegisterFile("src/core/server_status.proto", fileDescriptor_67e9cedab6951959) }

var fileDescriptor_67e9cedab6951959 = []byte{
	// 804 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x6d, 0x6f, 0xda, 0x48,
	0x10, 0x0e, 0xe6, 0x25, 0x61, 0xcc, 0x8b, 0xb3, 0x5c, 0xee, 0x10, 0x39, 0xe9, 0x4e, 0xe4, 0xa4,
	0xcb, 0xe5, 0x24, 0x22, 0x71, 0xba, 0xd3, 0x35, 0xad, 0xda, 0xba, 0xc1, 0x6d, 0xac, 0xf2, 0x12,
	0x2d, 0x84, 0xa6, 0xad, 0x2a, 0xcb, 0x35, 0x4b, 0x63, 0x15, 0x6c, 0x62, 0xaf, 0x51, 0xf3, 0x07,
	0xfa, 0x2d, 0x9f, 0xfb, 0xb3, 0xfa, 0x97, 0x2a, 0xef, 0xae, 0x89, 0x81, 0x80, 0x12, 0x29, 0x9f,
	0xc2, 0xce, 0xce, 0xf3, 0xcc, 0xf8, 0xd9, 0x67, 0x46, 0x81, 0x5f, 0x7d, 0xcf, 0x3a, 0xb4, 0x5c,
	0x8f, 0x1c, 0xfa, 0xc4, 0x9b, 0x12, 0xcf, 0xf0, 0xa9, 0x49, 0x03, 0xbf, 0x36, 0xf1, 0x5c, 0xea,
	0xa2, 0x9f, 0x9d, 0xa9, 0x3d, 0xb0, 0xcd, 0x9a, 0xed, 0x0c, 0x89, 0x47, 0x1c, 0x8b, 0xf0, 0xa4,
	0xca, 0xee, 0x0c, 0x35, 0x76, 0x07, 0x64, 0x64, 0x58, 0xae, 0x33, 0xb4, 0x3f, 0x71, 0x50, 0xf5,
	0x04, 0x72, 0x5d, 0x6a, 0xd2, 0x46, 0xe0, 0x99, 0xd4, 0x76, 0x1d, 0xf4, 0x13, 0xa4, 0x2d, 0x37,
	0x70, 0x68, 0x39, 0xf1, 0x7b, 0x62, 0x3f, 0x85, 0xf9, 0x01, 0x55, 0x21, 0x4f, 0x5d, 0x6a, 0x8e,
	0x0c, 0x6a, 0x8f, 0x89, 0xe1, 0xf8, 0x65, 0x89, 0xdd, 0xca, 0x2c, 0xd8, 0xb3, 0xc7, 0xa4, 0xed,
	0x57, 0x7b, 0x80, 0xba, 0xac, 0x1d, 0x4c, 0x2e, 0x03, 0xe2, 0xd3, 0xf0, 0xe0, 0xa3, 0xa7, 0xb0,
	0xe9, 0x07, 0x96, 0x45, 0x7c, 0x9f, 0x31, 0xca, 0xf5, 0x3f, 0x6a, 0xb7, 0xb7, 0x59, 0x8b, 0xb7,
	0x81, 0x23, 0x50, 0xf5, 0x0c, 0x4a, 0xa7, 0x9e, 0x3b, 0xb4, 0x47, 0xe4, 0x41, 0x69, 0x7b, 0x80,
	0x4e, 0x88, 0x39, 0xa2, 0x17, 0x0f, 0xca, 0x7a, 0x2d, 0xc1, 0xb6, 0x1e, 0x66, 0x3e, 0x24, 0x2b,
	0x7a, 0x02, 0x99, 0xa1, 0x69, 0x8f, 0xc8, 0x80, 0xa9, 0x7e, 0x57, 0xb8, 0xc0, 0x84, 0xd5, 0x2d,
	0x77, 0x3c, 0x09, 0x28, 0x29, 0x27, 0xef, 0x53, 0x5d, 0x80, 0xd0, 0x11, 0xa4, 0x2f, 0x03, 0x12,
	0x90, 0x72, 0xea, 0x1e, 0x68, 0x0e, 0xa9, 0x5e, 0x27, 0x01, 0xb5, 0x42, 0xcf, 0xf5, 0x89, 0xe7,
	0xdb, 0xae, 0xc3, 0xfd, 0x81, 0x4e, 0x40, 0xf6, 0x88, 0x39, 0xb8, 0x62, 0xf6, 0x25, 0x4c, 0x94,
	0x42, 0xfd, 0xcf, 0x55, 0xc4, 0x8c, 0x00, 0x87, 0xf9, 0x21, 0x9c, 0x60, 0xf0, 0x66, 0xbf, 0xd1,
	0x7b, 0x90, 0x59, 0x3a, 0x63, 0x0a, 0x5d, 0x99, 0xdc, 0x97, 0xeb, 0x47, 0x6b, 0x99, 0xe6, 0x5a,
	0xa9, 0xb1, 0xd7, 0x62, 0xcf, 0xa4, 0x39, 0xd4, 0xbb, 0xc2, 0x60, 0xcf, 0x02, 0xa8, 0x0e, 0x3b,
	0x7c, 0x60, 0xc8, 0x17, 0x62, 0x05, 0xe1, 0x77, 0x19, 0x7c, 0x34, 0x92, 0xcc, 0xfc, 0x25, 0x76,
	0xa9, 0x45, 0x77, 0xc7, 0x6c, 0x50, 0x66, 0x98, 0x59, 0x6d, 0x81, 0x49, 0xc5, 0x30, 0x7a, 0x74,
	0xc7, 0x30, 0x95, 0x0b, 0x28, 0x2e, 0xb4, 0x81, 0x14, 0x48, 0x7e, 0x26, 0x57, 0x4c, 0x99, 0x3c,
	0x0e, 0x7f, 0xa2, 0x67, 0x90, 0x9e, 0x9a, 0xa3, 0x80, 0x08, 0x0f, 0xfc, 0xb5, 0xea, 0x1b, 0x97,
	0xec, 0x87, 0x39, 0xee, 0x48, 0xfa, 0x3f, 0x51, 0xfd, 0x26, 0x81, 0xcc, 0x44, 0x10, 0x0f, 0xf1,
	0x18, 0x32, 0x7c, 0x19, 0x08, 0x63, 0xee, 0xad, 0x55, 0xee, 0x98, 0xa5, 0x62, 0x01, 0x41, 0x1f,
	0xa0, 0x30, 0xe5, 0x5a, 0x8a, 0x35, 0x24, 0xe4, 0xff, 0x6f, 0x2d, 0x89, 0xd0, 0x7d, 0xee, 0x15,
	0xb8, 0xf4, 0xf9, 0x69, 0x3c, 0x56, 0x19, 0x01, 0x5a, 0x4e, 0x8a, 0x0b, 0x93, 0xe4, 0xc2, 0x3c,
	0x9f, 0x17, 0xe6, 0xe0, 0xee, 0x8f, 0x1f, 0x57, 0xe6, 0x7b, 0x0a, 0x72, 0x5d, 0x96, 0x28, 0xa4,
	0x29, 0x80, 0x64, 0x0f, 0x58, 0x9d, 0x2c, 0x96, 0xec, 0x01, 0x2a, 0xc3, 0xa6, 0xe8, 0x8f, 0x15,
	0xca, 0xe2, 0xe8, 0x88, 0xf4, 0x79, 0x37, 0x6f, 0x32, 0x37, 0xef, 0xaf, 0x1c, 0x13, 0xf6, 0x67,
	0x85, 0x9d, 0x77, 0x21, 0x1b, 0x4c, 0xa2, 0x15, 0xcb, 0x5d, 0xb6, 0xc5, 0x03, 0x6d, 0x1f, 0x9d,
	0x43, 0x8e, 0x5b, 0x4b, 0xa8, 0x9d, 0x62, 0x6a, 0xff, 0xbb, 0xbe, 0x90, 0x90, 0x3b, 0x26, 0x3d,
	0x17, 0x5b, 0x1e, 0xc7, 0x6c, 0xd0, 0x82, 0x1c, 0xe7, 0x14, 0x63, 0x94, 0x5e, 0xaf, 0xe4, 0xf2,
	0x96, 0xc7, 0x32, 0xc7, 0xf3, 0xb9, 0x39, 0x85, 0xfc, 0x84, 0xaf, 0x6c, 0xc1, 0x97, 0x61, 0x7c,
	0x7f, 0xaf, 0xe2, 0xbb, 0x65, 0xbf, 0xe3, 0x9c, 0x60, 0xe0, 0x8c, 0x2d, 0xc8, 0x5d, 0xb0, 0x6d,
	0x2d, 0x08, 0xb7, 0xd6, 0x37, 0xb8, 0xbc, 0xd9, 0xb1, 0xcc, 0xf1, 0xec, 0x50, 0xb1, 0x40, 0x59,
	0x14, 0x24, 0x6e, 0xac, 0x2c, 0x37, 0xd6, 0xa3, 0x79, 0x63, 0xed, 0xdd, 0xc1, 0xd6, 0x31, 0x47,
	0x1d, 0x50, 0x28, 0x2e, 0x6c, 0x2e, 0xb4, 0x0d, 0xf9, 0x56, 0xa7, 0xa1, 0x35, 0x8d, 0xb3, 0xf6,
	0xeb, 0x76, 0xe7, 0x4d, 0x5b, 0xd9, 0x40, 0x45, 0x90, 0x79, 0x08, 0x6b, 0x6a, 0xe3, 0xad, 0x92,
	0x40, 0x3b, 0xb0, 0x1d, 0xe5, 0xa8, 0x7d, 0x55, 0x6f, 0xaa, 0x2f, 0x9a, 0x9a, 0x22, 0xdd, 0x40,
	0x9b, 0x1d, 0xb5, 0xa1, 0xb7, 0x5f, 0x29, 0x49, 0x54, 0x82, 0x62, 0x94, 0x19, 0x05, 0x53, 0x07,
	0x5f, 0x13, 0xa0, 0x2c, 0x5a, 0x0c, 0x21, 0x28, 0x74, 0x35, 0xdc, 0xd7, 0xb0, 0xa1, 0xb7, 0xfb,
	0x6a, 0x53, 0x6f, 0x28, 0x1b, 0xe8, 0x17, 0x28, 0xcd, 0x62, 0x7a, 0x4f, 0x57, 0x9b, 0xfa, 0xbb,
	0x90, 0x21, 0x81, 0x14, 0xc8, 0x89, 0x0b, 0xde, 0x92, 0x14, 0x83, 0x6b, 0xe7, 0x7a, 0x8f, 0x17,
	0xff, 0x0d, 0x76, 0x45, 0xec, 0xa5, 0xaa, 0x37, 0xb5, 0x86, 0xd1, 0xeb, 0xdc, 0x10, 0x69, 0x0a,
	0x7c, 0xcc, 0xb0, 0x7f, 0x2f, 0xfe, 0xf9, 0x11, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x7d, 0xdb, 0x0d,
	0xb3, 0x08, 0x00, 0x00,
}
