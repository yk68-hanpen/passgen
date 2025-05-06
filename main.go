package main

import "passgen/cmd"

// バージョン情報（ビルド時に設定される）
var version = "dev"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
