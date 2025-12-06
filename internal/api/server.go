package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redhander/go-eshop/config"
	"github.com/redhander/go-eshop/internal/db/repository"
	"github.com/redhander/go-eshop/internal/dto"
	"github.com/redhander/go-eshop/internal/worker"
	"github.com/redhander/go-eshop/pkg/auth"
	cachesrv "github.com/redhander/go-eshop/pkg/cache"
	"github.com/redhander/go-eshop/pkg/payment"
	"github.com/redhander/go-eshop/pkg/upload"
)

// gin-swagger middleware
// swagger embed files

// @BasePath /api/v1
// @title           E-Commerce API
// @description     This is a sample server for a simple e-commerce API.
// @BasePath /api/v1
// @host      localhost:4000
type Server struct {
	config          config.Config
	router          *gin.Engine
	repo            repository.Repository
	tokenGenerator  auth.TokenGenerator
	uploadService   upload.CdnUploader
	paymentSrv      *payment.PaymentManager
	cacheSrv        cachesrv.CacheContainer
	taskDistributor worker.TaskDistributor
}

func NewAPI(
	cfg config.Config,
	repo repository.Repository,
	cachesrv cachesrv.CacheContainer,
	taskDistributor worker.TaskDistributor,
	uploadService upload.CdnUploader,
	paymentSrv *payment.PaymentManager,
) (*Server, error) {
	tokenGenerator, err := auth.NewJwtGenerator(cfg.SymmetricKey)
	if err != nil {
		return nil, err
	}
	server := &Server{
		tokenGenerator:  tokenGenerator,
		repo:            repo,
		config:          cfg,
		taskDistributor: taskDistributor,
		uploadService:   uploadService,
		cacheSrv:        cachesrv,
		paymentSrv:      paymentSrv,
	}
	server.initializeRouter()
	return server, nil
}

func (s *Server) Server(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: s.router.Handler(),
	}
}

type DashboardData struct {
	Categories  []dto.CategoryDetail `json:"categories"`
	Collections []dto.CategoryDetail `json:"collections"`
}
