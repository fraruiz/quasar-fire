package usecases

type TopSecretSplitCreator struct {
}

func NewTopSecretSplitCreator() TopSecretSplitCreator {
	return TopSecretSplitCreator{}
}

func (creator TopSecretSplitCreator) Create(name string, distance float64, messages []string) error {
	return nil
}
