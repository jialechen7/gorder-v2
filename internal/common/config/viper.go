package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	if err := NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

var once sync.Once

func NewViperConfig() (err error) {
	once.Do(func() {
		err = newViperConfig()
	})
	return
}

func newViperConfig() error {
	relPath, err := getRelativePathFromCaller()
	if err != nil {
		return err
	}
	viper.SetConfigName("global")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(relPath)
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

func getRelativePathFromCaller() (relPath string, err error) {
	callerPwd, err := os.Getwd()
	if err != nil {
		return
	}
	_, here, _, _ := runtime.Caller(0)
	relPath, err = filepath.Rel(callerPwd, filepath.Dir(here))
	if err != nil {
		return
	}
	logrus.Infof("pwd: %s, relative path: %s", callerPwd, relPath)
	return
}
