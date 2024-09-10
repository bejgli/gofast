package sorter

type Config struct {
	Source string
	Rules  []struct {
		Pattern string
		Target  string
	}
}
