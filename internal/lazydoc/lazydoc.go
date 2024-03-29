/*
Copyright © 2024 Vino <vino0908@outlook.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package lazydoc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vinoMamba/lazydoc/internal/pkg/known"
	"github.com/vinoMamba/lazydoc/internal/pkg/log"
	"github.com/vinoMamba/lazydoc/internal/pkg/middleware"
	"github.com/vinoMamba/lazydoc/pkg/id"
	"github.com/vinoMamba/lazydoc/pkg/token"
)

var cfgFile string

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "lazydoc",
		Short: "lazydoc",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.InitLogger(logOptions())
			defer log.Sync()

			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				log.Infow("arg", "arg", arg)
			}
			return nil
		},
	}
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lazydoc.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	return rootCmd
}

func run() error {
	// 初始化snowflake id
	if err := id.Init("2024-01-01 00:00:00", 1); err != nil {
		return err
	}

	if err := initStore(); err != nil {
		return err
	}

	token.Init(viper.GetString("jwt.key"), known.XUserInfoKey)

	gin.SetMode(viper.GetString("gin.mode"))
	g := gin.New()
	g.Use(
		middleware.Cors(),
		middleware.RequestId(),
	)

	if err := registerAllApis(g); err != nil {
		return err
	}

	httpsrv := &http.Server{Addr: viper.GetString("gin.port"), Handler: g}
	log.Infow("Server success", "addr", viper.GetString("gin.port"))

	if err := httpsrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalw("ListenAndServe error", "err", err)
	}
	return nil
}
