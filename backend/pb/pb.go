package pb

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func StartPocketBase(addr string) *pocketbase.PocketBase {
	app := pocketbase.New()

	app.RootCmd.SetArgs([]string{
		"serve",
		"--http=" + addr,
	})

	go func() {
		log.Printf("▶ PocketBase: 开始启动 (%s)", addr)
		if err := app.Start(); err != nil {
			log.Fatalf("错误: %v", err)
		}
	}()

	return app
}
