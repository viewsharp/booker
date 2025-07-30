package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	UnspotToken string
	SpotID      string
	UnspotURL   string
}

var cfg Config

var rootCmd = &cobra.Command{
	Use:   "unspot-bot",
	Short: "Simple bot for booking places in unspot",
	Long: `Simple bot for booking places in unspot

Environment variables can be used instead of flags:
  UNSPOT_TOKEN  bearer token from unspot webpage
  SPOT_ID       UUID for spot from unspot
  UNSPOT_URL    url for self hosted installations`,
	Version: "1.0.0",
	RunE: func(cmd *cobra.Command, args []string) error {
		var errs []error

		// Проверка обязательных параметров
		if cfg.UnspotToken == "" {
			errs = append(errs, fmt.Errorf("unspot token is required (flag --unspot-token or env UNSPOT_TOKEN)"))
		}
		if cfg.SpotID == "" {
			errs = append(errs, fmt.Errorf("spot ID is required (flag --spot-id or env SPOT_ID)"))
		}

		return errors.Join(errs...)
	},
}

func init() {
	// Настройка флагов
	rootCmd.PersistentFlags().StringVarP(&cfg.UnspotURL, "unspot-url", "u", "", "url for self hosted installations (env: UNSPOT_URL)")
	rootCmd.PersistentFlags().StringVarP(&cfg.UnspotToken, "unspot-token", "t", "", "bearer token from unspot webpage (env: UNSPOT_TOKEN)")
	rootCmd.PersistentFlags().StringVarP(&cfg.SpotID, "spot-id", "s", "", "UUID for spot from unspot (env: SPOT_ID)")

	// Привязка переменных окружения
	viper.AutomaticEnv()
	viper.BindEnv("unspot_url")
	viper.BindEnv("unspot_token")
	viper.BindEnv("spot_id")
}

func main() {
	// Чтение значений из переменных окружения (если флаги не установлены)
	if cfg.UnspotURL == "" {
		cfg.UnspotURL = viper.GetString("unspot_url")
	}
	if cfg.UnspotToken == "" {
		cfg.UnspotToken = viper.GetString("unspot_token")
	}
	if cfg.SpotID == "" {
		cfg.SpotID = viper.GetString("spot_id")
	}

	// Выполнение команды
	if err := rootCmd.Execute(); err != nil {
		return
	}

	nextDate, err := GetNextDate()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = BookSeat(cfg, *nextDate)
	if err != nil {
		log.Fatal(err)
		return
	}
}
