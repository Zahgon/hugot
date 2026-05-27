package options

type Options struct {
	BackendOptions any
	ORTOptions     *OrtOptions
	GoMLXOptions   *GoMLXOptions
	Destroy        func() error
	Backend        string
}

func Defaults() *Options { _ = "STUB: not implemented"; return nil }

func getDefaultLibraryPaths() (string, string, string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

type GraphOptimizationLevel int

const (
	GraphOptimizationLevelDisableAll     GraphOptimizationLevel = 0
	GraphOptimizationLevelEnableBasic    GraphOptimizationLevel = 1
	GraphOptimizationLevelEnableExtended GraphOptimizationLevel = 2
	GraphOptimizationLevelEnableAll      GraphOptimizationLevel = 99
)

type LoggingLevel int

const (
	LoggingLevelVerbose LoggingLevel = 0
	LoggingLevelInfo    LoggingLevel = 1
	LoggingLevelWarning LoggingLevel = 2
	LoggingLevelError   LoggingLevel = 3
	LoggingLevelFatal   LoggingLevel = 4
)

type OrtOptions struct {
	LibraryPath             *string
	LibraryDir              *string
	Telemetry               *bool
	IntraOpNumThreads       *int
	InterOpNumThreads       *int
	CPUMemArena             *bool
	MemPattern              *bool
	ParallelExecutionMode   *bool
	IntraOpSpinning         *bool
	InterOpSpinning         *bool
	LogSeverityLevel        *LoggingLevel
	EnvLoggingLevel         *LoggingLevel
	GraphOptimizationLevel  *GraphOptimizationLevel
	CudaOptions             map[string]string
	CoreMLOptions           map[string]string
	DirectMLOptions         *int
	OpenVINOOptions         map[string]string
	TensorRTOptions         map[string]string
	NvTensorRTRTXOptions    map[string]string
	ExtraExecutionProviders []ExtraExecutionProvider
	OptimizedModelFilePath  *string
	ProfilingEnabled        *bool
	ProfilingFilePrefix     *string
	UseEngine               bool
}

type ExtraExecutionProvider struct {
	Name    string
	Options map[string]string
}
type GoMLXOptions struct {
	// BatchBuckets defines the bucket sizes for batch dimension padding.
	// Coarse bucketing reduces JIT cache pressure by limiting unique shapes.
	// Default: []int{1, 8, 32}
	BatchBuckets []int
	// SequenceBuckets defines the bucket sizes for sequence length padding.
	// Coarse bucketing reduces JIT cache pressure by limiting unique shapes.
	// Default: []int{32, 128, 512}
	SequenceBuckets []int
	Cuda            bool
	XLA             bool
	TPU             bool
}

// WithOption is the interface for all option functions.
type WithOption func(o *Options) error

// WithOnnxLibraryPath (ORT only) Use this function to set the path to the "libonnxuntime.so", "libonnxuntime.dylib" or "onnxruntime.dll" files.
func WithOnnxLibraryPath(ortLibraryPath string) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithTelemetry (ORT only) Enables telemetry events for the onnxbackend environment. Default is off.
func WithTelemetry() WithOption { _ = "STUB: not implemented"; return *new(WithOption) }

// WithIntraOpNumThreads (ORT only) Sets the number of threads used to parallelize execution within onnxbackend
// graph nodes. If unspecified, onnxbackend uses the number of physical CPU cores.
func WithIntraOpNumThreads(numThreads int) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithInterOpNumThreads (ORT only) Sets the number of threads used to parallelize execution across separate
// onnxbackend graph nodes. If unspecified, onnxbackend uses the number of physical CPU cores.
func WithInterOpNumThreads(numThreads int) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithCPUMemArena (ORT only) Enable/Disable the usage of the memory arena on CPU.
// Arena may pre-allocate memory for future usage. Default is true.
func WithCPUMemArena(enable bool) WithOption { _ = "STUB: not implemented"; return *new(WithOption) }

// WithMemPattern (ORT only) Enable/Disable the memory pattern optimization.
// If this is enabled memory is preallocated if all shapes are known. Default is true.
func WithMemPattern(enable bool) WithOption { _ = "STUB: not implemented"; return *new(WithOption) }

// WithExecutionMode sets the parallel execution mode for the ORT backend. Returns an error if the backend is not ORT.
func WithExecutionMode(parallel bool) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithIntraOpSpinning configures whether intra-op spinning is enabled for the ORT backend.
// It returns an error if used with a backend other than ORT.
func WithIntraOpSpinning(spinning bool) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithInterOpSpinning sets the spinning behavior for inter-op threads when the backend is ORT.
// It returns an error if used with a backend other than ORT.
func WithInterOpSpinning(spinning bool) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithCuda Use this function to set the options for CUDA provider.
// It takes a map of CUDA parameters as input.
// The options will be applied to the OrtOptions or GoMLXOptions struct, depending on your current backend.
func WithCuda(options map[string]string) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithTPU (XLA only) Use this function to enable TPU acceleration for the XLA backend.
// Requires libtpu.so to be available (pre-installed on GKE TPU nodes).
// Set PJRT_PLUGIN_LIBRARY_PATH to the directory containing pjrt_plugin_tpu.so or libtpu.so.
func WithTPU() WithOption { _ = "STUB: not implemented"; return *new(WithOption) }

// WithGoMLXBatchBuckets (XLA and GO only) sets the bucket sizes for batch dimension padding.
// Inputs are padded to the smallest bucket >= their batch size.
// Fewer/coarser buckets reduce JIT cache pressure but increase padding overhead.
// Default is []int{1, 8, 32}.
// IMPORTANT: Ensure MaxCache >= len(BatchBuckets) * len(SequenceBuckets).
func WithGoMLXBatchBuckets(buckets []int) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithGoMLXSequenceBuckets (XLA and GO only) sets the bucket sizes for sequence length padding.
// Inputs are padded to the smallest bucket >= their sequence length.
// Fewer/coarser buckets reduce JIT cache pressure but increase padding overhead.
// Default is []int{32, 128, 512}.
// IMPORTANT: Ensure MaxCache >= len(BatchBuckets) * len(SequenceBuckets).
func WithGoMLXSequenceBuckets(buckets []int) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithCoreML (ORT only) Use this function to set the CoreML options flags for the ONNX backend configuration.
// The `flags` parameter represents the CoreML options flags.
// The `o.CoreMLOptions` field in `OrtOptions` struct will be set to the provided flags parameter.
func WithCoreML(flags map[string]string) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithDirectML (ORT only) Use this function to set the DirectML device ID for the
// onnxbackend. By default, this option is not set.
func WithDirectML(deviceID int) WithOption { _ = "STUB: not implemented"; return *new(WithOption) }

// WithOpenVINO (ORT only) Use this function to set the OpenVINO options for the OpenVINO execution provider.
// The options parameter should be a map of string keys and string values, representing the configuration options.
// For each key-value pair in the map, the specified option will be set in the OpenVINO execution provider.
// Example usage: WithOpenVINO(map[string]string{"device_type": "CPU", "num_threads": "4"})
// This will configure the OpenVINO execution provider to use CPU as the device type and set the number of threads to 4.
func WithOpenVINO(options map[string]string) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithTensorRT (ORT only) Use this function to set the options for the TensorRT provider.
// The options parameter should be a pointer to an instance of TensorRTProviderOptions.
// By default, the options will be nil and the TensorRT provider will not be used.
// Example usage:
//
//	WithTensorRT(map[string]string{"device_id": "0"})
//
// Note: For the TensorRT provider to work, the onnxbackend library must be built with TensorRT support.
func WithTensorRT(options map[string]string) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithNvTensorRTRTX (ORT only) Use this function to set the options for the NvTensorRTRTX provider.
// The options parameter should be a pointer to an instance of NvTensorRTRTXOptions.
// By default, the options will be nil and the NvTensorRTRTX provider will not be used.
// Example usage:
//
//	WithNvTensorRTRTX(map[string]string{"device_id": "0"})
//
// Note: For the TensorRT provider to work, the onnxbackend library must be built with TensorRT support.
func WithNvTensorRTRTX(options map[string]string) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithLogSeverityLevel (ORT only) Sets the log severity level for the session.
func WithLogSeverityLevel(level LoggingLevel) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithEnvLoggingLevel (ORT only) Sets the log severity level for the environment.
func WithEnvLoggingLevel(level LoggingLevel) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithGraphOptimizationLevel (ORT only) Sets the graph optimization level for the session.
func WithGraphOptimizationLevel(level GraphOptimizationLevel) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithExtraExecutionProvider (ORT only) Adds an extra execution provider to the session.
func WithExtraExecutionProvider(name string, options map[string]string) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithOptimizedModelFilePath (ORT only) Sets the path to the optimized model file.
func WithOptimizedModelFilePath(path string) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithProfiling (ORT only) Enables or disables profiling.
func WithProfiling(enabled bool, filePrefix string) WithOption {
	_ = "STUB: not implemented"
	return *new(WithOption)
}

// WithGenerativeEngine for generative models, uses an ORT Gen AI Engine for dynamic batching and concurrent request support.
// Note: currently does not support image tensors in the upstream project.
func WithGenerativeEngine() WithOption { _ = "STUB: not implemented"; return *new(WithOption) }
