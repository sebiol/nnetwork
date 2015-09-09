package nnetwork

type Synapse interface {
  Value()     float64
  WeightedValue()     float64
  Weight()    float64
  SetWeight(float64)
}

// ConnectingSynapse connects two neurons
type ConnectingSynapse struct {
  weight  float64
  neuron  *Neuron
}

// Each neuron (apart from- in most cases - the input layer) gets one inhibiting synapse.
// The InhibitingSynapse works as a threshold for the activation of the neuron.
// As a Synapse the InhibitingSynapse (Threshold) will be trained by the normal learning method
type InhibitingSynapse struct {
  weight  float64
}

func (s *ConnectingSynapse) Weight() float64 {
  return s.weight
}

func (s *ConnectingSynapse) SetWeight(weight float64) {
  s.weight = weight
}

func (s *ConnectingSynapse) Value() float64 {
  return s.neuron.GetOutput()
}

func (s *ConnectingSynapse) WeightedValue() float64 {
  return s.neuron.GetOutput() * s.Weight()
}

//InhibitingSynapse

func (s *InhibitingSynapse) Weight() float64 {
  return s.weight
}

func (s *InhibitingSynapse) SetWeight(weight float64) {
  s.weight = weight
}

func (s *InhibitingSynapse) Value() float64 {
  return 1.0
}

func (s *InhibitingSynapse) WeightedValue() float64 {
  return 1.0 * s.Weight()
}
