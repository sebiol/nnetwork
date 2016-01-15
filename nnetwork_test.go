package nnetwork

import "testing"

func TestNNetwork(t *testing.T) {
	network := new(NeuralNetwork)
	network.Run([]float64{1.0})
}

func TestNNetworkSetupErrors(t *testing.T) {
	network := new(NeuralNetwork)
	err := network.SetUp(nil)
	want := "No configuration given"
	if err.Error() != want {
		t.Errorf("Wrong error. Wanted '%s' got '%s'", want, err)
	}
	want = "Config needs to define at least input and output layer. Only 1 given"
	err = network.SetUp(make([]LayerConfig, 1))
	if err.Error() != want {
		t.Errorf("Wrong error. Wanted '%s' got '%s'", want, err)
	}
}
