package neuron

import "testing"

func TestNeuronGetOutput(t *testing.T) {
  var neuron1 = new(Neuron)

  want := 1.0
  neuron1.SetInputfunction(func([]Synapse) float64 { return want})
  neuron1.SetActivationfunction(IdentityFunction)
  neuron1.SetOutputfunction(IdentityFunction)

  got := neuron1.GetOutput()
  if got != want {
    t.Errorf("(calc _output) neuron.GetOutput() == %v, want %v", got, want)
  }

  neuron1.SetInputfunction(func([]Synapse)float64 { return 42.0 })
  if got != want {
    t.Errorf("neuron.GetOutput ignored _oCalculated; want %v, got %v", want, got)
  }
}
