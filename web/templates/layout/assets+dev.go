//go:build !PROD

package layout

import (
	"net/url"

	"github.com/moroz/kashanka/config"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AssetEntryPoint() Node {
	return Group{
		Script(Type("module"), Src(config.ViteDevServer+"/@vite/client")),
		Script(Type("module"), Src(config.ViteDevServer+"/src/main.ts")),
	}
}

func ScriptEntryPoint(entrypoint string) Node {
	src, _ := url.JoinPath(config.ViteDevServer, entrypoint)
	return Script(Type("module"), Src(src))
}
