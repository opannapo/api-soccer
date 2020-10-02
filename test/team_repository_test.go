package test

import (
	"app/app"
	"app/app/entities"
	"app/app/repository"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type SuiteTeamRepo struct {
	suite.Suite
	DB         *gorm.DB
	repository *repository.TeamRepository
	team       entities.TeamEntity
}

func (s *SuiteTeamRepo) SetupSuite() {
	app.SetupConfig("test")
	user := viper.GetString("database.user")
	password := viper.GetString("database.pass")
	dbname := viper.GetString("database.name")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open("mysql", connectionString)
	s.DB = db
	s.DB.LogMode(false)

	require.NoError(s.T(), err)

	s.repository = &repository.TeamRepository{Db: s.DB}
}

func TestSuiteTeamRepoInit(t *testing.T) {
	suite.Run(t, new(SuiteTeamRepo))
}

func (s *SuiteTeamRepo) Test_Match_Result_If_Exist() {
	expectIdResult := int64(1000)
	result, err := s.repository.GetById(int(expectIdResult))
	if !reflect.ValueOf(result).IsZero() {
		assert.Equal(s.T(), expectIdResult, result.ID)
		require.NoError(s.T(), err)
	} else {
		assert.Empty(s.T(), result)
		assert.Equal(s.T(), "record not found", err.Error())
	}
}

func (s *SuiteTeamRepo) Test_Match_Team_Player_If_Exist() {
	expectIdResult := int64(1)
	result, err := s.repository.GetById(int(expectIdResult))
	if !reflect.ValueOf(result).IsZero() {
		assert.Equal(s.T(), expectIdResult, result.ID)
		require.NoError(s.T(), err)
		if len(result.Players) > 0 {
			assert.Equal(s.T(), expectIdResult, result.Players[0].TeamId)
		}
	} else {
		emptyResult := entities.TeamEntity{
			ID:        0,
			Name:      "",
			Address:   "",
			CreatedAt: 0,
			UpdatedAt: 0,
			Players:   nil,
		}
		assert.Equal(s.T(), emptyResult, result)
		assert.Equal(s.T(), "record not found", err.Error())
	}
}

func (s *SuiteTeamRepo) Test_Match_Result_An_Array_If_Exist() {
	result, err := s.repository.GetAll()
	if !reflect.ValueOf(result).IsZero() {
		assert.Equal(s.T(), reflect.TypeOf([]*entities.TeamEntity{}), reflect.TypeOf(result))
		require.NoError(s.T(), err)
	} else {
		emptyResult := entities.TeamEntity{
			ID:        0,
			Name:      "",
			Address:   "",
			CreatedAt: 0,
			UpdatedAt: 0,
			Players:   nil,
		}
		assert.Equal(s.T(), emptyResult, result)
		assert.Error(s.T(), err)
		assert.Equal(s.T(), "record not found", err.Error())
	}
}
