package github

import (
	"context"

	"github.com/Improwised/GPAT/config"
	"github.com/Improwised/GPAT/database"
	"github.com/Improwised/GPAT/models"
	"github.com/shurcooL/githubv4"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)


const (
	DEBUG = "DEBUG"
	INFO = "INFO"
	ERROR = "ERROR"
	WARNING = "WARNING"
)

type GithubService struct {
	client *githubv4.Client
	model  *models.Queries
	config config.AppConfig
	ctx    context.Context
	logger *zap.Logger
}

func NewGithubService(cfg config.AppConfig, logger *zap.Logger) (*GithubService, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.Github.Token},
	)
	db, err := database.Connect(cfg.DB)
	if err != nil {
		return &GithubService{}, err
	}
	client := githubv4.NewClient(oauth2.NewClient(ctx, ts))
	model := models.New(db)
	if err != nil {
		return &GithubService{}, err
	}
	return &GithubService{
		client: client,
		model:  model,
		ctx:    ctx,
		config: cfg,
		logger: logger,
	}, nil
}
