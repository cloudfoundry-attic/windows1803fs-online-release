package fakes

type Config struct {
	WriteCall struct {
		CallCount int
		Returns   struct {
			Error error
		}
	}
}

func (c *Config) Write() error {
	c.WriteCall.CallCount++

	return c.WriteCall.Returns.Error
}
