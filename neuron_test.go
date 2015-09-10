package nnetwork

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
  got = neuron1.GetOutput()
  if got != want {
    t.Errorf("neuron.GetOutput ignored _oCalculated; want %v, got %v", want, got)
  }
}

func TestNeuronLinkedToInputArray(t *testing.T) {
  var neuron1 = new(Neuron)

  testInputs := [][]float64 { {1.0}, {2.0}, {3.0}}
  var input []float64

  neuron1.SetInputfunction(func([]Synapse) float64 { return input[0] })
  neuron1.SetActivationfunction(IdentityFunction)
  neuron1.SetOutputfunction(IdentityFunction)

  for _, c := range testInputs {
    input = c
    got := neuron1.GetOutput()
    neuron1.ResetOutput()
    if got != c[0] {
      t.Errorf("Expected %v, got %v", c[0], got)
    }
  }
}
