package schema

func TestGetSchema(t *testing.T) {
	var want Schema = Schema{
		Name: "minecraft_java"
		Image: "itzg/minecraft-server:latest"
		URL: "https://github.com/itzg/docker-minecraft-server"
		Ratio: "1-2"
		Sizes: Sizes{
			XS Resources `yaml:"xs"`
			S  Resources `yaml:"s"`
			M  Resources `yaml:"m"`
			L  Resources `yaml:"l"`
			XL Resources `yaml:"xl"`
		}
		Network  []Network `yaml:"network"`
		Settings []Setting `yaml:"settings"`
		Volumes  []Volume  `yaml:"volumes"`
		Probes   Probes    `yaml:"probes"`
	}
	got, err := NewContainerEnvVar("EULA", "TRUE")
	if want != got || err != nil {
		t.Fatalf(`NewServerVolume("EULA", "TRUE") = %q, %v, want match for %#q, nil`, got, err, want)
	}
}