package hello

const (
	englishHelloPrefix = "Hello, "
	TurkishHelloPrefix = "Merhaba, "
	SpanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (languagePrefix string) {

	languagePrefix = "English"
	switch language {
	case "Turkish":
		languagePrefix = TurkishHelloPrefix
	case "Spanish":
		languagePrefix = SpanishHelloPrefix
	default:
		languagePrefix = englishHelloPrefix
	}
	return
}
