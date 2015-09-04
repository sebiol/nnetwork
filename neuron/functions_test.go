package neuron

import "testing"

func TestIdentityFunction(t *testing.T) {
  cases := []struct {
    in, want float64
  }{
    {5.0, 5.0},
    {-0.7, -0.7},
    {0, 0},
  }

  for _, c := range cases {
    got := IdentityFunction(c.in)
    if got != c.want {
      t.Errorf("IdentityFunction(%v) == %v, want %v", c.in, got, c.want)
    }
  }
}

func TestBipolarThresholdFunction(t *testing.T) {
  cases := []struct {
    in, want float64
  }{
    {0, 1},
    {-0.7, -1},
    {0.001, 1},
    {22, 1},
    {-120312, -1},
  }

  for _, c := range cases {
    got := BipolarThresholdFunction(c.in)
    if got != c.want {
      t.Errorf("IdentityFunction(%v) == %v, want %v", c.in, got, c.want)
    }
  }
}

func TestNetInputFunction(t *testing.T) {
  neuron1 := Neuron{_output: 2, _oCalculated: true}
  neuron2 := Neuron{_output: 3, _oCalculated: true}
  synapse1 := ConnectingSynapse{ 1, &neuron1 }
  synapse2 := ConnectingSynapse{ 0.5, &neuron2 }
  synapses := make([]Synapse, 2)
  synapses[0] = &synapse1
  synapses[1] = &synapse2

  want := 2 + 0.5 * 3
  got := NetInputFunction(synapses)
  if got != want {
    t.Errorf("NetInputFunction == %v, want %v", got, want)
  }
}
