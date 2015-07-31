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
