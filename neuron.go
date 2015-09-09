package nnetwork

type Neuron struct {
  _output       float64
  _oCalculated  bool
  input         InputFunction
  activation    ActivationFunction
  output        OutputFunction
  synapses      []Synapse
}

func (n *Neuron) SetInputfunction(fp InputFunction) {
  n.input = fp
}

func (n *Neuron) SetActivationfunction(fp ActivationFunction) {
  n.activation = fp
}

func (n *Neuron) SetOutputfunction(fp OutputFunction) {
  n.output = fp
}

func (n *Neuron) SetSynapses(syn []Synapse) {
  n.synapses = syn
}

func (n *Neuron) ResetOutput() {
  n._output = 0
  n._oCalculated = false
}

func (n *Neuron) GetOutput() float64 {
  if !n._oCalculated {
    n._output = n.output( n.activation( n.input(n.synapses)  ) )
    n._oCalculated = true
  }
  return n._output
}
