//go:build integration

package integration

import (
	"context"
	"testing"

	"github.com/dotdak/exchange-system/handler"
	"github.com/dotdak/exchange-system/infrastructure"
	"github.com/dotdak/exchange-system/repo"
	"github.com/stretchr/testify/suite"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"gorm.io/gorm"
)

type MySqlRepositoryTestSuite struct {
	db        *gorm.DB
	buyRepo   repo.BuyRepo
	wagerRepo repo.WagerRepo
	handler   handler.Handler
	docker    tc.ComposeStack
	suite.Suite
}

func (p *MySqlRepositoryTestSuite) SetupSuite() {
	db, err := infrastructure.NewMysqlDb(&infrastructure.DbConfig{
		Username:     "test",
		Password:     "test",
		Host:         "localhost",
		Port:         "3306",
		DatabaseName: "exchange-test",
	})
	if err != nil {
		panic(err)
	}

	p.db = db
	p.buyRepo = repo.NewBuyRepo(db)
	p.wagerRepo = repo.NewWagerRepo(db)
	p.handler = handler.NewHandler(p.wagerRepo, p.buyRepo)
}

func TestMySqlRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &MySqlRepositoryTestSuite{})
}

func (p *MySqlRepositoryTestSuite) SetupSubTest() {
	Truncate(p.db)
}

func (p *MySqlRepositoryTestSuite) TearDownAllSuite() {
	Teardown(p.db)
	p.docker.Down(context.Background())
}
