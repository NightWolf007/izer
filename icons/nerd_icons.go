// Code generated by gomplate; DO NOT EDIT.
// This file generated by gomplate at
// Sun, 15 Sep 2019 09:20:12 UTC
// using data from datafiles/nerd.json
package icons

import (
	"path/filepath"
)

var nerdDefaultIcon = Icon{"\ue612", 255}

var nerdIconsMap = map[string]Icon{
	".3gp":               Icon{"\uf001", 148},
	".DS_Store":          Icon{"\uf179", 8},
	".aac":               Icon{"\uf001", 148},
	".ai":                Icon{"\ue7b4", 3},
	".aif":               Icon{"\uf001", 148},
	".aifc":              Icon{"\uf001", 148},
	".aiff":              Icon{"\uf001", 148},
	".au":                Icon{"\uf001", 148},
	".avi":               Icon{"\uf03d", 198},
	".awk":               Icon{"\ue795", 148},
	".bash":              Icon{"\ue795", 148},
	".bashprofile":       Icon{"\ue615", 8},
	".bashrc":            Icon{"\ue615", 8},
	".bat":               Icon{"\ue615", 148},
	".bmp":               Icon{"\uf03e", 141},
	".bower.json":        Icon{"\ue61a", 208},
	".bowerrc":           Icon{"\ue61a", 208},
	".c":                 Icon{"\ufb70", 75},
	".caf":               Icon{"\uf001", 148},
	".cjsx":              Icon{"\ue7ba", 75},
	".clj":               Icon{"\ue76a", 148},
	".cljc":              Icon{"\ue76a", 148},
	".cljs":              Icon{"\ue76a", 148},
	".coffee":            Icon{"\ue61b", 3},
	".config":            Icon{"\ue615", 8},
	".cpp":               Icon{"\ufb71", 75},
	".cs":                Icon{"\uf81a", 75},
	".csh":               Icon{"\ue795", 148},
	".css":               Icon{"\ue614", 75},
	".csv":               Icon{"\uf0ce", 148},
	".ctp":               Icon{"\ue73d", 1},
	".d":                 Icon{"\ue7af", 1},
	".dart":              Icon{"\ue798", 75},
	".db":                Icon{"\uf1c0", 8},
	".diff":              Icon{"\ue728", 255},
	".direnv":            Icon{"\ue615", 8},
	".dropbox":           Icon{"\ue707", 75},
	".dump":              Icon{"\uf1c0", 8},
	".editorconfig":      Icon{"\ue615", 8},
	".edn":               Icon{"\ue76a", 75},
	".eex":               Icon{"\ue62d", 198},
	".ejs":               Icon{"\ue618", 3},
	".env":               Icon{"\ue615", 8},
	".erl":               Icon{"\ue7b1", 1},
	".es6":               Icon{"\ue60c", 3},
	".ex":                Icon{"\ue62d", 198},
	".exs":               Icon{"\ue62d", 198},
	".f#":                Icon{"\ue7a7", 75},
	".fish":              Icon{"\ue795", 148},
	".flac":              Icon{"\uf001", 148},
	".flv":               Icon{"\uf03d", 198},
	".fs":                Icon{"\ue7a7", 75},
	".fsi":               Icon{"\ue7a7", 75},
	".fsscript":          Icon{"\ue7a7", 75},
	".fsx":               Icon{"\ue7a7", 75},
	".gif":               Icon{"\uf03e", 141},
	".gitattributes":     Icon{"\uf09b", 255},
	".gitignore":         Icon{"\uf09b", 255},
	".gitkeep":           Icon{"\uf09b", 255},
	".go":                Icon{"\ue627", 75},
	".go-version":        Icon{"\ue615", 8},
	".groovy":            Icon{"\ue775", 75},
	".gvimrc":            Icon{"\ue62b", 148},
	".h":                 Icon{"\ufb70", 198},
	".hbs":               Icon{"\ue60f", 208},
	".hpp":               Icon{"\ufb71", 198},
	".hrl":               Icon{"\ue7b1", 198},
	".hs":                Icon{"\ue61f", 141},
	".html":              Icon{"\ue60e", 208},
	".ico":               Icon{"\ue623", 3},
	".ini":               Icon{"\ue615", 75},
	".java":              Icon{"\ue738", 1},
	".jl":                Icon{"\ue624", 141},
	".jpeg":              Icon{"\uf03e", 141},
	".jpg":               Icon{"\uf03e", 141},
	".js":                Icon{"\ue60c", 3},
	".jscsrc":            Icon{"\ue60c", 75},
	".jshintrc":          Icon{"\ue60c", 75},
	".json":              Icon{"\ue60b", 3},
	".jsx":               Icon{"\ue7ba", 75},
	".ksh":               Icon{"\ue795", 148},
	".l16":               Icon{"\uf001", 148},
	".leex":              Icon{"\ue62d", 198},
	".less":              Icon{"\ue614", 75},
	".lhs":               Icon{"\ue61f", 141},
	".lock":              Icon{"\uf023", 8},
	".lua":               Icon{"\ue620", 75},
	".m4a":               Icon{"\uf001", 148},
	".m4r":               Icon{"\uf001", 148},
	".md":                Icon{"\ue609", 75},
	".mdx":               Icon{"\ue609", 3},
	".ml":                Icon{"\u03bb", 208},
	".mm":                Icon{"\uf179", 255},
	".mogg":              Icon{"\uf001", 148},
	".mov":               Icon{"\uf03d", 198},
	".mp3":               Icon{"\uf001", 148},
	".mp4":               Icon{"\uf03d", 198},
	".mustache":          Icon{"\ue60f", 208},
	".npmignore":         Icon{"\ue616", 1},
	".nvimrc":            Icon{"\ue62b", 148},
	".oga":               Icon{"\uf001", 148},
	".ogg":               Icon{"\uf001", 148},
	".pcm":               Icon{"\uf001", 148},
	".pdf":               Icon{"\uf1c1", 1},
	".php":               Icon{"\ue608", 141},
	".pl":                Icon{"\ue769", 141},
	".pm":                Icon{"\ue769", 141},
	".png":               Icon{"\uf03e", 141},
	".pp":                Icon{"\uf499", 141},
	".ps1":               Icon{"\ue795", 148},
	".psb":               Icon{"\ue7b8", 75},
	".psd":               Icon{"\ue7b8", 75},
	".py":                Icon{"\ue606", 75},
	".pyc":               Icon{"\ue606", 75},
	".pyd":               Icon{"\ue606", 3},
	".pyo":               Icon{"\ue606", 3},
	".python-version":    Icon{"\ue615", 8},
	".rb":                Icon{"\ue605", 1},
	".rlib":              Icon{"\ue7a8", 8},
	".rs":                Icon{"\ue7a8", 8},
	".rss":               Icon{"\ue619", 208},
	".ruby-version":      Icon{"\ue615", 8},
	".sass":              Icon{"\ue603", 198},
	".sbt":               Icon{"\ue737", 75},
	".scala":             Icon{"\ue737", 1},
	".scss":              Icon{"\ue603", 198},
	".sh":                Icon{"\ue795", 148},
	".slim":              Icon{"\ue618", 1},
	".sln":               Icon{"\ue70c", 75},
	".slugignore":        Icon{"\ue615", 8},
	".sql":               Icon{"\uf1c0", 8},
	".stache":            Icon{"\ue60f", 208},
	".static":            Icon{"\ue615", 8},
	".styl":              Icon{"\ue600", 148},
	".suo":               Icon{"\ue70c", 75},
	".svg":               Icon{"\uf03e", 141},
	".svgx":              Icon{"\uf03e", 141},
	".swift":             Icon{"\ue755", 208},
	".t":                 Icon{"\ue769", 141},
	".tex":               Icon{"L", 198},
	".tmp":               Icon{"\uf017", 8},
	".tmpl":              Icon{"\ue612", 75},
	".ts":                Icon{"\ue628", 75},
	".tsx":               Icon{"\ue7ba", 75},
	".twig":              Icon{"\ue61c", 148},
	".vala":              Icon{"\uf7ab", 8},
	".vim":               Icon{"\ue62b", 148},
	".vimrc":             Icon{"\ue62b", 148},
	".vue":               Icon{"\ufd42", 148},
	".wav":               Icon{"\uf001", 148},
	".wave":              Icon{"\uf001", 148},
	".wma":               Icon{"\uf001", 148},
	".wmv":               Icon{"\uf03d", 198},
	".xcplayground":      Icon{"\ue755", 208},
	".xml":               Icon{"\ufabf", 148},
	".xul":               Icon{"\uf269", 208},
	".yaml":              Icon{"\ue615", 198},
	".yml":               Icon{"\ue615", 198},
	".zsh":               Icon{"\ue795", 148},
	".zshrc":             Icon{"\ue615", 8},
	"Dockerfile":         Icon{"\uf308", 75},
	"Gemfile":            Icon{"\ue605", 1},
	"Gruntfile.js":       Icon{"\ue611", 208},
	"Gulpfile.js":        Icon{"\ue610", 1},
	"LICENSE":            Icon{"\ue60a", 3},
	"Procfile":           Icon{"\ue607", 141},
	"TODO":               Icon{"\uf634", 75},
	"bower.json":         Icon{"\ue61a", 208},
	"config.ru":          Icon{"\ue605", 1},
	"docker-compose.yml": Icon{"\uf308", 75},
	"go.mod":             Icon{"\ue627", 75},
	"go.sum":             Icon{"\ue627", 8},
	"gruntfile.js":       Icon{"\ue611", 208},
	"gulpfile.js":        Icon{"\ue610", 1},
	"karma.conf.coffee":  Icon{"\ue622", 148},
	"karma.conf.js":      Icon{"\ue622", 148},
	"karma.conf.ts":      Icon{"\ue622", 148},
}

// GetNerdIcon returns icon by filetype for nerd font
func GetNerdIcon(filename string) Icon {
	baseFilename := filepath.Base(filename)
	if icon, ok := nerdIconsMap[baseFilename]; ok {
		return icon
	}

	ext := filepath.Ext(filename)
	if ext == "" {
		return nerdDefaultIcon
	}

	if icon, ok := nerdIconsMap[ext]; ok {
		return icon
	}
	return nerdDefaultIcon
}
