package neuron

type Synapse struct {
  Weight    float64
  Neuron   *Neuron
}

func (s *Synapse) Value() float64 {
  return s.Neuron.GetOutput() * s.Weight
}
