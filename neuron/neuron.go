package neuron

type Neuron struct {
  input         InputFunction
  activation    ActivationFunction
  output        OutputFunction
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
