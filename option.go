package android

type Option func(Runtime) error

func SetOption(v Runtime, options ...Option) error {
	for _, option := range options {
		if err := option(v); err != nil {
			return err
		}
	}
	return nil
}
