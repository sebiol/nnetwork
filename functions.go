package nnetwork

type InputFunction func([]Synapse) float64
type ActivationFunction func(float64) float64
type OutputFunction func(float64) float64

func NetInputFunction(input []Synapse) float64{
  var tmpOut float64
  for _, syn := range input {
    tmpOut += syn.WeightedValue()
  }
  return tmpOut
}

func IdentityFunction(input float64) float64 {
  return input
}

func GeneralThresholdFUnction(input float64, threshold float64, x1 float64, x2 float64) float64 {
  if input < threshold {
    return x1
  } else {
    return x2
  }
}

func BinaryThresholdFunction(input float64) float64 {
  return GeneralThresholdFUnction(input, 0.5, 0, 1)
}

func BipolarThresholdFunction(input float64) float64 {
  return GeneralThresholdFUnction(input, 0, -1, 1)
}
