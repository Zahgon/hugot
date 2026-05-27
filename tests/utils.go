package testutil

import (
	"testing"

	"github.com/knights-analytics/hugot"
	"github.com/knights-analytics/hugot/pipelines"
)

const ModelsFolder = "../../models/"
const TestCasesFolder = "../../testcases/"

// test download validation

func TestDownloadValidation(t *testing.T) { _ = "STUB: not implemented"; return }

// a model with the required files in a subfolder should not error

// a model without tokenizer.json or .onnx model should error

// FEATURE EXTRACTION

func FeatureExtractionPipeline(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// test 'robert smith'

// test ['robert smith junior', 'francis ford coppola']

// determinism test to make sure embeddings of a string are not influenced by other strings in the batch

// these vectors should be the same

// test normalization

// test getting output by name

func FeatureExtractionPipelineValidation(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// Text classification

func TextClassificationPipeline(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// check PrintStatistics

func TextClassificationPipelineMulti(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// check GetStatistics

func TextClassificationPipelineValidation(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// Zero shot

func ZeroShotClassificationPipeline(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// Gets overridden per test, but included for coverage

func ZeroShotClassificationPipelineValidation(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// Token classification

func TokenClassificationPipeline(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// Split-words enabled pipeline

// Expect same entities as the simple aggregation for the equivalent sentence for split words

// Non-split input with double space changes raw offsets.

// Compare first sequence: entity words may match, but offsets should differ due to space normalization.

// Pre-tokenization splitting on 'X': expect different entity results from the non-split case

// Expect split-words to detect 'Berlin' as an entity, while non-split should not because of the confusing X characters.

func TokenClassificationPipelineValidation(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// Cross Encoder

func CrossEncoderPipeline(t *testing.T, session *hugot.Session) { _ = "STUB: not implemented"; return }

func CrossEncoderPipelineValidation(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// 1. Test: output dims length != 2

// 2. Test: output dims second dim != 1

// 3. Test: more than one dynamic dim (-1)

// Image classification test using HuggingFace SqueezeNet and a sample image.
func ImageClassificationPipeline(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

func ImageClassificationPipelineValidation(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// object detection

func ObjectDetectionPipeline(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// Use a simple cat image similar to classification test style

// Find a detection labeled cat (COCO index 15)

// basic box sanity

// score should be reasonable

// fall back to checking top detection label for debug

func ObjectDetectionPipelineValidation(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// Corrupt the primary image input to have invalid dims

// If a mask input exists, make it invalid length to trigger error

// invalid

// Rename outputs so inference of boxes/scores fails

// No same name

func NoSameNamePipeline(t *testing.T, session *hugot.Session) { _ = "STUB: not implemented"; return }

func DestroyPipelines(t *testing.T, session *hugot.Session) { _ = "STUB: not implemented"; return }

// Text Generation.
func TextGenerationPipeline(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// Configure the text generation pipeline

// Create the pipeline

// Execute tests

// streaming test

// tools test

// Lark grammar that constrains both <tool_call> blocks to valid JSON.
//
// Requirements:
//   - <tool_call> and </tool_call> must be "special": false in tokenizer.json so that
//     llguidance can match them as regular byte sequences. Marking them "special": true
//     promotes them to control tokens that the guidance system cannot byte-force, causing
//     "token doesn't satisfy the grammar" errors.
//   - The grammar expects exactly two tool calls for this query.

// Create the pipeline

// Two minimal Hermes-style tool definitions.

func TextGenerationPipelineValidation(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// Configure the text generation pipeline

// QUESTION ANSWERING

func QuestionAnsweringPipeline(t *testing.T, session *hugot.Session) {
	_ = "STUB: not implemented"
	return
}

// Context is a JSON document; questions target specific property values.

// WithTopKAnswers(2): verify that 2 ranked answers are returned per input and are ordered by score.

// TABULAR

func TabularPipeline(t *testing.T, session *hugot.Session) { _ = "STUB: not implemented"; return }

// Iris classification for an example

// Thread safety

func ThreadSafety(t *testing.T, session *hugot.Session, numEmbeddings int) {
	_ = "STUB: not implemented"
	return
}

// Utilities

func checkClassificationOutput(t *testing.T, inputResult []pipelines.ClassificationOutput, inputExpected []pipelines.ClassificationOutput) {
	_ = "STUB: not implemented"
	return
}

// Returns an error if any element between a and b don't match.
func floatsEqual(a, b []float32) error { _ = "STUB: not implemented"; return nil }

// Arbitrarily chosen precision. Large enough not to be affected by quantization

func almostEqual(a, b float64) bool { _ = "STUB: not implemented"; return false }

func CheckT(t *testing.T, err error) { _ = "STUB: not implemented"; return }

func printTokenEntities(o *pipelines.TokenClassificationOutput) { _ = "STUB: not implemented"; return }

type toolCall struct {
	Name      string         `json:"name"`
	Arguments map[string]any `json:"arguments"`
}

// parseToolCalls extracts all <tool_call>...</tool_call> blocks from s and unmarshals
// each as a toolCall. Uses json.Decoder so that a trailing stray `}` (a common
// int4-quantisation artefact) does not cause the block to be skipped — Decode reads
// exactly one JSON value and stops, leaving trailing garbage unread.
func parseToolCalls(s string) []toolCall { _ = "STUB: not implemented"; return nil }

func CheckToolCalls(t *testing.T, output string) { _ = "STUB: not implemented"; return }
