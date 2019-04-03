package pkg

type Clusters struct {
	Cluster struct {
		CertificateAuthorityData string `yaml:"certificate-authority-data"`
		Server                   string `yaml:"server"`
	} `yaml:"cluster"`
	Name string `yaml:"name"`
}

type Contexts struct {
	Context struct {
		Cluster   string `yaml:"cluster"`
		Namespace string `yaml:"namespace"`
		User      string `yaml:"user"`
	} `yaml:"context"`
	Name string `yaml:"name"`
}

type Users struct {
	Name string `yaml:"name"`
	User struct {
		ClientCertificateData string `yaml:"client-certificate-data"`
		ClientKeyData         string `yaml:"client-key-data"`
		Password              string `yaml:"password"`
		Username              string `yaml:"username"`
	} `yaml:"user"`
}

// Config ...
type Config struct {
	APIVersion     string     `yaml:"apiVersion"`
	Clusters       []Clusters `yaml:"clusters"`
	Contexts       []Contexts `yaml:"contexts"`
	CurrentContext string     `yaml:"current-context"`
	Kind           string     `yaml:"kind"`
	Preferences    struct {
	} `yaml:"preferences"`
	Users []Users `yaml:"users"`
}
