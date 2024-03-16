package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"time"
)

func Parse() Config {
	yml, err := os.ReadFile("resources/config.yml")

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yml, &config)

	if err != nil {
		panic(err)
	}

	return config
}

type Config struct {
	Email        EmailConfig
	GameFilters  GameFilters        `yaml:"gameFilters"`
	ExchangeRate ExchangeRateConfig `yaml:"exchangeRate"`
	GameTime     GameTimeConfig     `yaml:"gameTime"`
	FansFirst    FansFirstConfig    `yaml:"fansFirst"`
}

type EmailConfig struct {
	Address    string
	Password   string
	Recipients []string
}

type GameFilters struct {
	MaxPrice  int `yaml:"maxPrice"`
	Seats     int
	Days      []Date
	Teams     Teams
	Opponents []Opponent
}

type Teams struct {
	Flames    bool
	Wranglers bool
}

type Opponent string

const (
	SEA Opponent = "SEA"
)

type ExchangeRateConfig struct {
	BaseUrl string `yaml:"baseUrl"`
	ApiKey  string `yaml:"apiKey"`
}

type GameTimeConfig struct {
	BaseUrl string `yaml:"baseUrl"`
	BuyUrl  string `yaml:"buyUrl"`
}

type FansFirstConfig struct {
	BaseUrl string `yaml:"baseUrl"`
	BuyUrl  string `yaml:"buyUrl"`
}

type Date struct {
	Time time.Time
}

func (t *Date) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return nil
	}

	tt, err := time.Parse("2006-01-02", strings.TrimSpace(buf))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

func (t Date) MarshalYAML() (interface{}, error) {
	return t.Time.Format("2006-01-02"), nil
}
