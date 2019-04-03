package pkg

// Rename ...
func (c *Config) Rename() {
	cnun := make(map[string]string)
	for i, ctx := range c.Contexts {
		un := ctx.Context.User
		cnun[un] = ctx.Context.Cluster
		c.Contexts[i].Context.User = ctx.Context.Cluster
	}

	for i, user := range c.Users {
		if key, found := cnun[user.Name]; found {
			c.Users[i].Name = key
		}
	}
}
