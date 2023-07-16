//go:build integration

package integration

import (
	"os"
	"testing"

	"github.com/dotdak/exchange-system/handler"
	"github.com/dotdak/exchange-system/infrastructure"
	"github.com/dotdak/exchange-system/pkg/utils"
	"github.com/dotdak/exchange-system/repo"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type MySqlRepositoryTestSuite struct {
	db        *gorm.DB
	buyRepo   repo.BuyRepo
	wagerRepo repo.WagerRepo
	handler   handler.Handler
	suite.Suite
}

func (p *MySqlRepositoryTestSuite) SetupSuite() {
	db, err := infrastructure.NewMysqlDb(&infrastructure.DbConfig{
		Username:     utils.Any(os.Getenv("TEST_DB_USER"), "test"),
		Password:     utils.Any(os.Getenv("TEST_DB_PASSWORD"), "test"),
		Host:         utils.Any(os.Getenv("TEST_DB_HOST"), "db"),
		Port:         utils.Any(os.Getenv("TEST_DB_PORT"), "3306"),
		DatabaseName: utils.Any(os.Getenv("TEST_DB_NAME"), "exchange-test"),
	})
	if err != nil {
		panic(err)
	}
	logger := infrastructure.NullLogger()
	p.db = db
	p.buyRepo = repo.NewBuyRepo(db, logger)
	p.wagerRepo = repo.NewWagerRepo(db, logger)
	p.handler = handler.NewHandler(p.wagerRepo, p.buyRepo, logger)
}

func TestMySqlRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &MySqlRepositoryTestSuite{})
}

func (p *MySqlRepositoryTestSuite) SetupSubTest() {
	Truncate(p.db)
}

func (p *MySqlRepositoryTestSuite) TearDownAllSuite() {
	Teardown(p.db)
}
