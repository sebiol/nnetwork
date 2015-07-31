package neuron

import "testing"

func TestSynapseValue(t *testing.T) {
  var neuron1 = new(Neuron)

  neuronOutput := 1.0
  synapseWeight := 0.5

  neuron1.SetInputfunction(func([]Synapse) float64 { return neuronOutput })
  neuron1.SetActivationfunction(IdentityFunction)
  neuron1.SetOutputfunction(IdentityFunction)

  want := neuronOutput * synapseWeight

  synapse := &Synapse{ synapseWeight, neuron1 }

  got := synapse.Value()
  if got != want {
    t.Errorf("Synapse.Value() == %v, want %v", got, want)
  }
}
