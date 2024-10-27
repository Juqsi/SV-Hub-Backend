package llama

type Prompts string

const (
	PROMPT_KEY_INFOS       Prompts = "Nenn mir bitte die fünf wichtigsten kombinationen von Begriffen nach denen in einer Vektor Datenbank gesucht werden um die Besten Antworten zu finden. Die Begriffe sollten in der Reihenfolge ihrer Wichtigkeit genannt werden und in der folgenden Struktur: [Wort1 Wort2 Wort3 Wort4, Wort5 Wort6 ...]"
	PROMPT_CONCLUSION      Prompts = "comment"
	PROMPT_SIMILAR_MEANING Prompts = "like"
)

func (m Prompts) String() string {
	return string(m)
}
