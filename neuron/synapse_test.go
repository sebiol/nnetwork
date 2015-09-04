package neuron

import "testing"

func TestConnectingSynapseValue(t *testing.T) {
  var neuron1 = new(Neuron)

  neuronOutput := 1.0
  synapseWeight := 0.5

  neuron1.SetInputfunction(func([]Synapse) float64 { return neuronOutput })
  neuron1.SetActivationfunction(IdentityFunction)
  neuron1.SetOutputfunction(IdentityFunction)

  want := neuronOutput

  synapse := &ConnectingSynapse{ synapseWeight, neuron1 }

  got := synapse.Value()
  if got != want {
    t.Errorf("Synapse.Value() == %v, want %v", got, want)
  }
}

func TestConnectingSynapseWeightedValue(t *testing.T) {
  var neuron1 = new(Neuron)

  neuronOutput := 1.0
  synapseWeight := 0.5

  neuron1.SetInputfunction(func([]Synapse) float64 { return neuronOutput })
  neuron1.SetActivationfunction(IdentityFunction)
  neuron1.SetOutputfunction(IdentityFunction)

  want := neuronOutput * synapseWeight

  synapse := &ConnectingSynapse{ synapseWeight, neuron1 }

  got := synapse.WeightedValue()
  if got != want {
    t.Errorf("Synapse.WeightedValue() == %v, want %v", got, want)
  }
}

func TestInhibitingSynapseValue(t *testing.T) {
  synValue  := 1.0
  want := synValue

  synapse := &InhibitingSynapse{ 0.5 }
  got := synapse.Value()
  if got != want {
    t.Errorf("Synapse.Value() == %v, want %v", got, want)
  }
}

func TestInhibitingSynapseWeightedValue(t *testing.T) {
  synValue  := 1.0
  synWeight := 0.5
  want := synValue * synWeight

  synapse := &InhibitingSynapse{ synWeight }
  got := synapse.WeightedValue()
  if got != want {
    t.Errorf("Synapse.WeightedValue() == %v, want %v", got, want)
  }
}

func TestSetWeight(t *testing.T) {
  synapse1 := &ConnectingSynapse{ 1.0, nil }
  synapse2 := &InhibitingSynapse{ 1.0 }

  want := 2.33
  synapse1.SetWeight(want)
  synapse2.SetWeight(want)

  if synapse1.Weight() != want {
    t.Errorf("ConnectingSynapse.Weight() == %v, want %v", synapse1.Weight, want)
  }

  if synapse2.Weight() != want {
    t.Errorf("InhibitingSynapse.Weight() == %v, want %v", synapse2.Weight, want)
  }
}
