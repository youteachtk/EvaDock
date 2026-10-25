//go:build eva_assets

package main

import "embed"

//go:embed assets/independientes/*.png
var evaAssets embed.FS

func bundledAsset(name string) ([]byte, error) { return evaAssets.ReadFile("assets/" + name) }