package neuron

type InputFunction func(*Neuron) float64
type ActivationFunction func(float64) float64
type OutputFunction func(float64) float64

func IdentityFunction(input float64) float64 {
  return input
}

func BipolarThresholdFunction(input float64) float64 {
  if input < 0 {
    return -1
  } else {
    return 1
  }
}
