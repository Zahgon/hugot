//go:build !NODOWNLOAD

package hugot

import (
	"context"

	"github.com/gomlx/go-huggingface/hub"
)

// DownloadOptions is a struct of options that can be passed to DownloadModel.
type DownloadOptions struct {
	AuthToken             string
	OnnxFilePath          string
	ExternalDataPath      string
	Branch                string
	MaxRetries            int
	RetryInterval         int
	ConcurrentConnections int
	Verbose               bool
}

// NewDownloadOptions creates new DownloadOptions struct with default values.
// Override the values to specify different download options.
func NewDownloadOptions() DownloadOptions { _ = "STUB: not implemented"; return *new(DownloadOptions) }

// DownloadModel can be used to download a model directly from huggingface. Before the model is downloaded,
// validation occurs to ensure there is an .onnx and tokenizers.json file. Hugot only works with onnx models.
func DownloadModel(ctx context.Context, modelName string, destination string, options DownloadOptions) (string, error) {
	_ = "STUB: not implemented"
	// replicates code in hf downloader
	return "", nil
}

// make sure it's an onnx model with tokenizer

func ValidateDownloadedHFModel(repo *hub.Repo, options DownloadOptions) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
