package htmlgonstants

/*
Character encodings that are to be used for the form submission.

Common values:
    UTF-8 - Character encoding for Unicode
    ISO-8859-1 - Character encoding for the Latin alphabet

In theory, any character encoding can be used, but no browser understands all of them. The more widely a character encoding is used, the better the chance that a browser will understand it.
*/

const (
	CharacterEncodingASCII    = "ASCII"
	CharacterEncodingUTF8     = "UTF-8"
	CharacterEncodingISO88591 = "ISO-8859-1"
)
