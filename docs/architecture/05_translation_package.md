# Translation module

## Module purpose

This module is used to translate the multilingual markdown file into english version.

## Module scope

This module get raw markdown file from input, and translate it into english version with the help of [Niu Translation API](niutrans.com).

## Module structure

When the request translation starts, it will facilitate the text position and start-stop Index of each node in markdown, split the text with random separators, send it to markdown translator, and finally restore the obtained results to the text according to Index.

The pseudocode is as follows:

```go
func TranslateMarkdown(src, from, to) {
  // Step 1: Initialize markdown parser
  md := InitializeMarkdownParser()
  doc := ParseMarkdown(md, src)

  // Step 2: Prepare translation
  translationVec, translationSeg = PrepareTranslation()

  // Step 3: Walk through the AST
  WalkThroughAST(doc, translationVec, translationSeg)

  // Step 4: Translate
  translatedStr = TranslateText(translationVec, from, to)

  // Step 5: Replace text
  result = ReplaceText(src, translatedStr, translationSeg)

  // Step 6: Return result
  return result
}
```

## Module Interface

None.

## Functions

### New

```go
type Engine struct {
        // Has unexported fields.
}
    Engine is the config of translation server

func New(translationAPIKey string, qiniuAccessKey string, qiniuSecretKey string) *Engine
    New create a new TranslateConfig
```

**Example:**

```go
func TestNew(t *testing.T) {
    e := New(mockKey, "", "")
    if e == nil {
        t.Fatal("new engine failed")
    }
}
```

### Text2Audio

```go
func (e *Engine) Text2Audio(content string) (*TTSResponse, error)
    Text2Audio
```

**Example:**

```go
func TestText2Audio(t *testing.T) {
    e := New("", mockAccessKey, mockSecretKey)
    resp, err := e.Text2Audio("hello")
    if err != nil {
        t.Fatal(err)
    }
    t.Log(resp)
}
```

### TranslateBatchPlain

```go
func (e *Engine) TranslateBatchPlain(src []string, from, to language.Tag) ([]string, error)
    TranslateBatchSeq translate a series of sequences of text
```

**Example:**

```go
func TestTranslateBatchPlain(t *testing.T) {
    e := New(mockKey, "", "")
    src := []string{"你好", "hello"}
    ret, err := e.TranslateBatchPlain(src, language.Chinese, language.English)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(ret)
}
```

### TranslateMarkdown

```go
func (e *Engine) TranslateMarkdown(src []byte, from, to language.Tag) (ret []byte, err error)
    TranslateMarkdown translate markdown with bytes

func (e *Engine) TranslateMarkdownText(src string, from, to language.Tag) (ret string, err error)
    TranslateMarkdown translate markdown text
```

**Example:**

```go
func TestTranslateMarkdown(t *testing.T) {
    tests := []struct {
        src  string
        from language.Tag
        to   language.Tag
    }{
        {`# Hello, World!`, language.English, language.Chinese},
    }

	trans := New(mockKey, "", "")
	for _, test := range tests {
		translatedResult, err := trans.TranslateMarkdownText(test.src, test.from, test.to)
		fmt.Println(translatedResult, err)
		if err != nil {
			t.Fatal(err)
		}
	}    
}
```

### TranslatePlain

```go
func (e *Engine) TranslatePlain(src []byte, from, to language.Tag) (ret []byte, err error)
    Translate translate sequence of bytes

func (e *Engine) TranslatePlainText(src string, from, to language.Tag) (ret string, err error)
    Translate translate sequence of text
```

**Example:**

```go
func TestTranslatePlainText(t *testing.T) {
	if mockKey == "" {
		t.Skip("NIUTRANS_API_KEY not set")
	}

	tests := []struct {
		src  string
		from language.Tag
		to   language.Tag
	}{
		{"你好", language.Chinese, language.English},
		{"hello", language.English, language.Chinese},
	}

	trans := New(mockKey, "", "")
	for _, test := range tests {
		to, err := trans.TranslatePlainText(test.src, test.from, test.to)
		fmt.Println(to, err)
		if err != nil {
			t.Fatal(err)
		}
	}
}
```