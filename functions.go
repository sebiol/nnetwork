package nnetwork

import "math"

type InputFunction func([]Synapse) float64
type ActivationFunction func(float64) float64
type OutputFunction func(float64) float64

func NetInputFunction(input []Synapse) float64 {
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

func GeneralLogisiticLinearFunction(input float64, thetaLower float64, thetaUpper, xLower float64, xUpper float64) float64 {
	if input < thetaLower {
		return xLower
	}
	if input > thetaUpper {
		return xUpper
	}

	return input
}

func BinaryLogisticLinearFunction(input float64) float64 {
	return GeneralLogisiticLinearFunction(input, 0, 1, 0, 1)
}

func BipolarLogisticLinearFunction(input float64) float64 {
	return GeneralLogisiticLinearFunction(input, -1, 1, -1, 1)
}

/*
 * Based on en.wikipedia.org/wiki/Logistic_function
 */
func GeneralLogisticallFunction(input float64, L float64, k float64, x0 float64) float64 {
	return (1 / (1 + math.Exp(-k*(input-x0))))
}
