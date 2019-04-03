package pkg

import "fmt"

// Merge ...
func Merge(a, b *Config) (*Config, error) {
	cluster, conflictName := mergeClusterWithoutConflict(a.Clusters, b.Clusters)
	if conflictName != "" {
		return nil, fmt.Errorf(`both configs have cluster with name "%s"`, conflictName)
	}

	context, conflictName := mergeContextWithoutConflict(a.Contexts, b.Contexts)
	if conflictName != "" {
		return nil, fmt.Errorf(`both configs have context with name "%s"`, conflictName)
	}

	user, conflictName := mergeUserWithoutConflict(a.Users, b.Users)
	if conflictName != "" {
		return nil, fmt.Errorf(`both configs have user with name "%s"`, conflictName)
	}

	config := &Config{
		APIVersion:     a.APIVersion,
		Clusters:       cluster,
		Contexts:       context,
		CurrentContext: a.CurrentContext,
		Kind:           a.Kind,
		Preferences:    a.Preferences,
		Users:          user,
	}
	return config, nil
}

func mergeClusterWithoutConflict(a, b []Clusters) ([]Clusters, string) {
	clusters := make([]Clusters, 0)
	cluserName := make(map[string]bool)

	for _, cl := range a {
		if _, found := cluserName[cl.Name]; found {
			return nil, cl.Name
		}
		clusters = append(clusters, cl)
		cluserName[cl.Name] = true
	}

	for _, cl := range b {
		if _, found := cluserName[cl.Name]; found {
			return nil, cl.Name
		}
		clusters = append(clusters, cl)
		cluserName[cl.Name] = true
	}

	return clusters, ""
}

func mergeContextWithoutConflict(a, b []Contexts) ([]Contexts, string) {
	contexts := make([]Contexts, 0)
	contextName := make(map[string]bool)

	for _, cn := range a {
		if _, found := contextName[cn.Name]; found {
			return nil, cn.Name
		}
		contexts = append(contexts, cn)
		contextName[cn.Name] = true
	}

	for _, cn := range b {
		if _, found := contextName[cn.Name]; found {
			return nil, cn.Name
		}
		contexts = append(contexts, cn)
		contextName[cn.Name] = true
	}

	return contexts, ""
}

func mergeUserWithoutConflict(a, b []Users) ([]Users, string) {
	users := make([]Users, 0)
	userName := make(map[string]bool)

	for _, u := range a {
		if _, found := userName[u.Name]; found {
			return nil, u.Name
		}
		users = append(users, u)
		userName[u.Name] = true
	}

	for _, u := range b {
		if _, found := userName[u.Name]; found {
			return nil, u.Name
		}
		users = append(users, u)
		userName[u.Name] = true
	}

	return users, ""
}
