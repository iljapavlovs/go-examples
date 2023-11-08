package files

type TomlModel struct {
	Version int      `toml:"version"`
	Name    string   `toml:"name"`
	Tags    []string `toml:"tags"`
}
