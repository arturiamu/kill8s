package cmd

import (
	"flag"
	"fmt"
	"kill8s/config"
	"kill8s/domain/repository"
	"kill8s/infrastructure/kube"
	"kill8s/infrastructure/persistence"
	"kill8s/server"
	"log"
	"os"
	"path/filepath"
)

func Run() {
	cfg, err := config.Load(".")
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	initComponents(cfg, kubeconfig)

	svr := server.NewRouter()
	if err = svr.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatalf("Error run server:%v", err)
	}
}

func initComponents(cfg *config.Config, path *string) {
	if cfg.MySQL != nil {
		persistence.InitMySQL(cfg)
	}
	if cfg.Redis != nil {
		persistence.InitRedis(cfg)
	}
	kube.InitK8s(path)
	repository.Init()
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}
