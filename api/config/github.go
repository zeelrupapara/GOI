package config

// DBConfig type of db config object
type GithubConfig struct {
	Token        string `required:"true" envconfig:"GITHUB_TOKEN" validate:"required"`
}
