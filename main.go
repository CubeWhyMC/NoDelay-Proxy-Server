package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/layou233/ZBProxy/common"
	"github.com/layou233/ZBProxy/config"
	"github.com/layou233/ZBProxy/console"
	"github.com/layou233/ZBProxy/service"
	"github.com/layou233/ZBProxy/version"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

func main() {
	log.SetOutput(color.Output)
	console.SetTitle(fmt.Sprintf("NoDelay %v | Running...", version.Version))
	color.HiGreen("Welcome to NoDelay %s (%s)!\n", version.Version, version.CommitHash)
	color.HiBlack("Build Information: %s, %s/%s, CGO %s\n",
		runtime.Version(), runtime.GOOS, runtime.GOARCH, common.CGOHint)

	resp, err := http.Get("http://whitelist.hln-network.xyz/NoDelay/NoDelay.php")
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	if string(body) == "true" {
		color.HiGreen("You are currently allowed to use NoDelay: your IP is authorized.")
		config.LoadConfig()
		service.Listeners = make([]net.Listener, 0, len(config.Config.Services))

		for _, s := range config.Config.Services {
			go service.StartNewService(s)
		}

		{
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				log.Panic(err)
			}
			defer watcher.Close()
			err = config.MonitorConfig(watcher)
			if err != nil {
				log.Panic("Config Reload Error : ", err)
			}
		}

		{
			osSignals := make(chan os.Signal, 1)
			signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
			<-osSignals

			for _, listener := range service.Listeners {
				if listener != nil {
					listener.Close()
				}
			}
		}
	} else {
		color.HiRed("You are currently not allowed to run NoDelay: your IP is not authorized.")
	}
}