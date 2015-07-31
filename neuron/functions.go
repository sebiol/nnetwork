package neuron

type InputFunction func(*Neuron) float64
type ActivationFunction func(float64) float64
type OutputFunction func(float64) float64

