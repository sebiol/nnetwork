package nnetwork

import "testing"

func TestNeuronGetOutput(t *testing.T) {
	var neuron1 = new(Neuron)

	want := 1.0
	neuron1.SetInputfunction(func([]Synapse) float64 { return want })
	neuron1.SetActivationfunction(IdentityFunction)
	neuron1.SetOutputfunction(IdentityFunction)

	got := neuron1.GetOutput()
	if got != want {
		t.Errorf("(calc _output) neuron.GetOutput() == %v, want %v", got, want)
	}

	neuron1.SetInputfunction(func([]Synapse) float64 { return 42.0 })
	got = neuron1.GetOutput()
	if got != want {
		t.Errorf("neuron.GetOutput ignored _oCalculated; want %v, got %v", want, got)
	}
}

func TestNeuronLinkedToInputArray(t *testing.T) {
	var neuron1 = new(Neuron)

	testInputs := [][]float64{{1.0}, {2.0}, {3.0}}
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

func TestNeuronWithSynapse(t *testing.T) {
	inputArray := []float64{1.0, 2.0}

	inputNeuron1 := Neuron{
		input:      func([]Synapse) float64 { return inputArray[0] },
		activation: IdentityFunction,
		output:     IdentityFunction,
	}
	inputNeuron2 := Neuron{
		input:      func([]Synapse) float64 { return inputArray[1] },
		activation: IdentityFunction,
		output:     IdentityFunction,
	}

	outputNeuron := Neuron{
		input:      NetInputFunction,
		activation: IdentityFunction,
		output:     IdentityFunction,
	}
	synapses := make([]Synapse, 2, 2)
	synapses[0] = &ConnectingSynapse{
		neuron: &inputNeuron1,
		weight: 1.0,
	}
	synapses[1] = &ConnectingSynapse{
		neuron: &inputNeuron2,
		weight: 2.0,
	}
	outputNeuron.SetSynapses(synapses)

	want := 5.0
	got := outputNeuron.GetOutput()
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}

	//adjust the weight of synapse1
	want = 6.0
	synapses = outputNeuron.GetSynapses()
	//alt .GetSynapses()[i].SetWeight(x)
	synapses[0].SetWeight(2.0)
	outputNeuron.SetSynapses(synapses)
	outputNeuron.ResetOutput()
	got = outputNeuron.GetOutput()
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}

}
