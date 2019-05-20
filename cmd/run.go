package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var scheme, key, cert string
var returnStatusCode, port int

type responseYesSir struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the mock API server",
	Long:  `Start serving!`,
	Run: func(cmd *cobra.Command, args []string) {

		log.SetFormatter(&log.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
		log.SetOutput(os.Stdout)

		//Capture signals to correctly close channels and connections
		sigc := make(chan os.Signal, 1)
		var srv http.Server

		signal.Notify(sigc,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)

		go func() {
			s := <-sigc
			switch s {
			case os.Interrupt:
				log.Debug("Interrupt received. Closing connections...")
				ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
				srv.Shutdown(ctx)
				os.Exit(0)

			case syscall.SIGTERM:
				log.Debug("Termination received. Closing connections...")
				ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
				srv.Shutdown(ctx)
				os.Exit(-1)
			default:
				log.Debug("Closing connections...")
				ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
				srv.Shutdown(ctx)
				os.Exit(-2)
			}
		}()

		//Create Router for REST interface
		muxRouter := mux.NewRouter()
		muxRouter.PathPrefix("/").HandlerFunc(getInfo)

		//Starting up server
		log.Info("Starting up " + ApplicationName + " on port " + fmt.Sprintf("%v", port))

		if scheme == "https" {
			log.Error(http.ListenAndServeTLS(fmt.Sprintf(":%v", port), cert, key, muxRouter))
		} else {
			srv.Addr = fmt.Sprintf(":%v", port)
			srv.Handler = muxRouter
			if err := srv.ListenAndServe(); err != nil {
				log.Error("%s\n" + fmt.Sprintf("%v", err))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&scheme, "scheme", "s", "http", "Scheme http|https")
	runCmd.Flags().IntVarP(&returnStatusCode, "return", "r", 200, "HTTP return code (200,404,500)")
	runCmd.Flags().IntVarP(&port, "port", "p", 8888, "Port to listen on")
	runCmd.Flags().StringVarP(&cert, "cert", "c", "", "Path to the server TLS certificate file (only for https)")
	runCmd.Flags().StringVarP(&key, "key", "k", "", "Path to the server TLS certificate key file (only for https)")

	if scheme != "https" && scheme != "http" {
		log.Error("Scheme needs to be http or https")
		os.Exit(-1)
	}

	if scheme == "https" && (cert == "" || key == "") {
		log.Error("Need a certificate and a key for https!")
		os.Exit(-1)
	}
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	connData := fmt.Sprintf("%v:%v %v", r.RemoteAddr, r.RequestURI, r.Method)
	callback := r.URL.Query().Get("callback")

	log.Info("Request from " + connData)
	req, _ := httputil.DumpRequest(r, true)
	log.Info(fmt.Sprintf("%s", req))

	//Send them over the connection
	switch returnStatusCode {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 500:
		w.WriteHeader(http.StatusInternalServerError)
	case 404:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusOK)
		returnStatusCode = 200
	}

	stat := responseYesSir{
		Message: fmt.Sprintf("%s", req),
		Status:  returnStatusCode,
	}

	resp, _ := json.Marshal(stat)

	if callback != "" {
		fmt.Fprintf(w, "%s(%s)", callback, resp)
	} else {
		w.Write(resp)
	}
}
