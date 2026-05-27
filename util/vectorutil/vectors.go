package vectorutil

// Mean of a float32 vector.
func Mean(vector []float32) float32 { _ = "STUB: not implemented"; return 0 }

// SoftMax take a vector and calculate softmax scores of its values.
func SoftMax(vector []float32) []float32 { _ = "STUB: not implemented"; return nil }

func SumSlice(s []float64) float64 { _ = "STUB: not implemented"; return 0 }

// ArgMax find both index of max value in s and max value.
func ArgMax(s []float32) (int, float32, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func Sigmoid(s []float32) []float32 { _ = "STUB: not implemented"; return nil }

// Norm of a vector.
func Norm(v []float32, p int) float64 { _ = "STUB: not implemented"; return 0 }

// Normalize single vector according to: https://pytorch.org/docs/stable/generated/torch.nn.functional.normalize.html
func Normalize(embedding []float32, p int) []float32 { _ = "STUB: not implemented"; return nil }
