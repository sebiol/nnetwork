package nnetwork

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

func TestBinaryThresholdFunction(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{-0.7, 0},
		{0.001, 0},
		{22, 1},
		{0.51, 1},
		{-120312, 0},
	}

	for _, c := range cases {
		got := BinaryThresholdFunction(c.in)
		if got != c.want {
			t.Errorf("BinaryThresholdFunction(%v) == %v, want %v", c.in, got, c.want)
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
			t.Errorf("BipolarThresholdFunction(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestBinaryLogisticLinearFunction(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{5.0, 1.0},
		{-0.7, 0},
		{0, 0},
		{0.1, 0.1},
		{-5, 0},
		{0.99, 0.99},
	}

	for _, c := range cases {
		got := BinaryLogisticLinearFunction(c.in)
		if got != c.want {
			t.Errorf("BinaryLogisticLinearFunction(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestBipolarLogisticLinearFunction(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{5.0, 1.0},
		{-5, -1},
		{-0.95, -0.95},
		{-0.7, -0.7},
		{0, 0},
		{0.1, 0.1},
		{0.99, 0.99},
	}

	for _, c := range cases {
		got := BipolarLogisticLinearFunction(c.in)
		if got != c.want {
			t.Errorf("BipolarLogisticLinearFunction(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestLogisticalFunction(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0.5},
		{100, 1},
	}

	for _, c := range cases {
		got := GeneralLogisticallFunction(c.in, 1, -0.5, 0)
		if got != c.want {
			t.Errorf("BipolarLogisticLinearFunction(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestNetInputFunction(t *testing.T) {
	neuron1 := Neuron{_output: 2, _oCalculated: true}
	neuron2 := Neuron{_output: 3, _oCalculated: true}
	synapse1 := ConnectingSynapse{1, &neuron1}
	synapse2 := ConnectingSynapse{0.5, &neuron2}
	synapses := make([]Synapse, 2)
	synapses[0] = &synapse1
	synapses[1] = &synapse2

	want := 2 + 0.5*3
	got := NetInputFunction(synapses)
	if got != want {
		t.Errorf("NetInputFunction == %v, want %v", got, want)
	}
}
