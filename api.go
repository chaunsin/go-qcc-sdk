package api

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Key       string        `json:"key" yaml:"key"`
	SecretKey string        `json:"secretKey" yaml:"secretKey"`
	Location  string        `json:"location" yaml:"location"` // default: time.Local
	Debug     bool          `json:"debug" yaml:"debug"`
	Timeout   time.Duration `json:"timeout" yaml:"timeout"` // A Timeout of zero means no timeout
}

type Api struct {
	cfg *Config
	cli *resty.Client
	loc *time.Location
}

func New(cfg *Config) *Api {
	var (
		cli = resty.New()
		a   = Api{
			cfg: cfg,
			loc: time.Local,
		}
	)
	cli.SetDebug(cfg.Debug)
	cli.SetTimeout(cfg.Timeout)
	cli.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if cfg.Location != "" {
		loc, err := time.LoadLocation(cfg.Location)
		if err != nil {
			panic(err)
		}
		a.loc = loc
	}
	return &a
}

func NewClient(cfg *Config, cli *resty.Client) *Api {
	var a = Api{
		cfg: cfg,
		loc: time.Local,
		cli: cli,
	}

	a.cli.SetDebug(cfg.Debug)
	a.cli.SetTimeout(cfg.Timeout)
	if cfg.Location != "" {
		loc, err := time.LoadLocation(cfg.Location)
		if err != nil {
			panic(err)
		}
		a.loc = loc
	}
	return &a
}

func NewFromFile(filename string) *Api {
	cfg, err := LoadConfig(filename)
	if err != nil {
		panic(fmt.Sprintf("load config from file [%s] error: %v", filename, err))
	}
	return New(cfg)
}

func (a *Api) GetClient() *resty.Client {
	return a.cli
}

func (a *Api) auth() (string, string, error) {
	return GenToken(a.cfg.Key, a.cfg.SecretKey, a.loc)
}

// LoadConfig load config from file
func LoadConfig(filename string) (*Config, error) {
	if filename == "" {
		return nil, errors.New("filename is empty")
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	switch filepath.Ext(filename) {
	case ".yaml", ".yml":
		err = yaml.NewDecoder(file).Decode(&cfg)
	case ".json":
		err = json.NewDecoder(file).Decode(&cfg)
	default:
		return nil, errors.New("unsupported file extension types")
	}
	return &cfg, nil
}

// GenToken https://openapi.qcc.com/services/after/code
// return:
// - token
// - timestamp(秒级)
// - error
func GenToken(key, secretKey string, loc ...*time.Location) (string, string, error) {
	var l = time.Local
	if len(loc) > 0 {
		l = loc[0]
	}
	var (
		unix = time.Now().In(l).Unix()
		t    = fmt.Sprintf("%s%d%s", key, unix, secretKey)
		m    = md5.New()
	)
	if _, err := io.WriteString(m, t); err != nil {
		return "", "", err
	}
	token := strings.ToUpper(hex.EncodeToString(m.Sum(nil)))
	return token, fmt.Sprintf("%d", unix), nil
}
