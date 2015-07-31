package neuron

type Synapse struct {
  weight    float64
  neuron   *Neuron
}

func (s *Synapse) SetNeuron(neuron *Neuron) {
  s.neuron = neuron
}

func (s *Synapse) GetWeight() {
  return s.weight
}

func (s *Synapse) SetWeight(weight float64) {
  s.weight = weight
}

func (s *Synapse) Value() float64 {
  return s.neuron.GetOutput() * s.weight
}
