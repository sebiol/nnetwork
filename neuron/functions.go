package neuron

type InputFunction func([]Synapse) float64
type ActivationFunction func(float64) float64
type OutputFunction func(float64) float64

func NetInputFunction(input []Synapse) float64{
  var tmpOut float64
  for _, syn := range input {
    tmpOut += syn.Value()
  }
  return tmpOut
}

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
