package command

type ParameterGate struct {
}

func NewParameterGate() (*ParameterGate, error) {

	return &ParameterGate{}, nil
}

func (parameterGate *ParameterGate) LowerParameterLimit(amount int) error {

	return nil
}
