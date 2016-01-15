package nnetwork

import "fmt"
import "errors"

type LayerConfig struct {
	// NeuronCount sets the amount of Neurons for the layer.
	NeuronCount uint16

	// StandardInputFunction holds the InputFunction to be set as default for
	// all Neurons in this layer.
	StandardInputFunction InputFunction

	// StandardActivationFunction holds the ActivationFunction to be set as default for
	// all Neurons in this layer.
	StandardActivationFunction ActivationFunction

	// StandardOutputFunction holds the OutputFunction to be set as default for
	// all Neurons in this layer.
	StandardOutputFunction OutputFunction

	// InhibitingSynapses controls the existance of InhibitingSynapses for the
	// Neurons in this layer.
	// True:  Each Neuron gets an InhibitingSynapse
	// False: No Neuron gets an InhibitingSynapse
	InhibitingSynapses bool
}

// NeuralNetwork
type NeuralNetwork struct {
	// input represents the last given input to the network.
	// input is set through the run method.
	// Each element in input is linked to a neuron in the input layer.
	input []float64

	// output represents the last calculated output of the netowrk.
	// output is set through the run method.
	output []float64

	// InternalState holds all Neurons of the NeuralNetwork separated by layer.
	// InternalState must be accessible to the learning algorithm
	InternalState [][]*Neuron
}

func (n *NeuralNetwork) GetInput() []float64 {
	return n.input
}

func (network *NeuralNetwork) Output() []float64 {
	return network.output
}

// Sets up the Network according to the provided configuration
// For each layer of the network is set up according to the corresponding
// entry in the configuration.
// config[0] is the input layer (handled differently)
// config[1:-2] hidden layer
// config[-1] is the output layer 
func (network *NeuralNetwork) SetUp(config []LayerConfig) error {
	if config == nil {
		return errors.New("No configuration given")
	}
	if len(config) < 2 {
		return fmt.Errorf("Config needs to define at least input and output layer. Only %v given", len(config))
	}

	var internalState [][]*Neuron
  //Create slices to hold layers of Neurons
	internalState = make([][]*Neuron, len(config))
	for i, _ := range internalState {
    //create slice per layer to hold Neurons
		internalState[i] = make([]*Neuron, config[i].NeuronCount)
	}

	//Set up the input layer
	//Input layer differs vastly as it ignores the assigned functions in config
	//and links the neuron to the input array
	for j, _ := range internalState[0] {
		a := j
		neuron := new(Neuron)
		//Set link to input array
		neuron.SetInputfunction(func([]Synapse) float64 { return network.input[a] })
		neuron.SetActivationfunction(IdentityFunction)
		neuron.SetOutputfunction(IdentityFunction)
		internalState[0][j] = neuron
	}

	config = config[1:]
	//Set up subsequent layers
	for i, layer := range internalState[1:] {
		for j, _ := range layer {
			neuron := new(Neuron)
			neuron.SetInputfunction(config[i].StandardInputFunction)
			neuron.SetActivationfunction(config[i].StandardActivationFunction)
			neuron.SetOutputfunction(config[i].StandardOutputFunction)

			//i references the element in the anonymous slice intState[1:] and runs
			//from 0 to len(intState) -1 while i = 0 references layer 1 in intState.
			previousLayer := internalState[i]
			//set up synapses for the neuron
			//create slice with space for optional inhibiting synapse
			synapses := make([]Synapse, len(previousLayer), len(previousLayer)+1)
			//create connecting synapses with neurons of previous layer
			for oIndex, oNeuron := range previousLayer {
				synapse := &ConnectingSynapse{
					neuron: oNeuron,
					weight: 0.0,
				}
				synapses[oIndex] = synapse
			}
			//optional: create inhibiting synapse for neuron
			if config[i].InhibitingSynapses {
				synapses = synapses[0:cap(synapses)]
				synapse := &InhibitingSynapse{
					weight: 0.0,
				}
				synapses[len(synapses)-1] = synapse
			}
			neuron.SetSynapses(synapses)

			internalState[i+1][j] = neuron
		}
	}
	network.InternalState = internalState
	return nil
}

// Calculates the output of the network for given input
func (network *NeuralNetwork) Run(netinput []float64) ([]float64, error) {
	if len(network.InternalState) < 1 {
		return nil, errors.New("Network not set up")
	}
	if len(netinput) < len(network.InternalState[0]) {
		return nil, errors.New("Input array to small")
	}
  // reset neuron output which is cached for efficiency (see Neuron)
	network.resetNetwork()
  // TODO: calculate result
	return network.output, nil
}

// Resets all Neurons in the Network
// Deletes cached results in Neurons, which are saved for efficiency
func (network *NeuralNetwork) resetNetwork() {
	for _, layer := range network.InternalState {
		for _, neuron := range layer {
			if neuron != nil {
				neuron.ResetOutput()
			}
		}
	}
}
