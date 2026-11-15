//go:build PROD

package layout

import (
	"encoding/json/v2"
	"log"
	"os"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

const ManifestPath = "assets/dist/.vite/manifest.json"

var manifest ViteManifest

func AssetEntryPoint() Node {
	entry := manifest["src/main.ts"]

	return Group{
		Script(Type("module"), Src("/"+entry.File)),
		Map(entry.Css, func(css string) Node {
			return Link(Rel("stylesheet"), Href("/"+css))
		}),
	}
}

type ViteManifest map[string]ViteManifestObject

type ViteManifestObject struct {
	File    string
	Name    string
	Src     string
	IsEntry bool
	Css     []string
}

func init() {
	bytes, err := os.ReadFile(ManifestPath)
	if err != nil {
		log.Fatalf("Failed to read Vite asset manifest file: %s", err)
	}

	var decoded ViteManifest
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		log.Fatalf("Failed to decode Vite asset manifest file: %s", err)
	}

	if _, ok := decoded["src/main.ts"]; !ok {
		log.Fatalf("The decoded Vite manifest does not contain an entry point for src/main.ts")
	}

	manifest = decoded
}
