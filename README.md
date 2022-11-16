# GoXML
Short story: I needed an XML parser for another project I was working for and I couldn't find any that
would suit my needs, hence this!

![](assets/image.jpg)

I wanted something to give me the XML in a tree-based structure. So, the final XML object that this
code produces, implements the following interface:

```go
type XML struct {
    Prolog   string
    chlidren []XMLElement
}
```

and the definition for `XMLElement`:
```go
type XMLElement interface {
    GetContent() string
    GetFather() XMLElement
    GetChildren() []XMLElement
    GetTag() string
    GetAttributes() string
}
```

Enjoy!