package parser

import (
    "strings"
    "net/url"
    "golang.org/x/net/html"
)

type Scraper struct {
    Title         string
    Language      string
    Author        string
    Description   string
    MetaImage     string
    Generator     string
    Feed          string
    Charset       string
    Links         []string
    Images        []string
    Keywords      []string
    Compatibility map[string]string
}

// findCharset returns the charset given the Content-Type
func findCharset(content string) string {
    if pos := strings.LastIndex(content, "charset="); pos != -1 {
        return content[pos+len("charset="):]
    }
    return ""
}

// mapifyStr converts a string with the format `foo=bar,ro=pi`
// in a map like map[foo:bar ro:pi]
func mapifyStr(content string) map[string]string {
    m := make(map[string]string)
    a := strings.Split(content, ",")
    for i := range a {
        s := strings.Split(a[i], "=")
        m[s[0]] = s[1]
    }
    return m
}

// findAttribute returns the value of the key provided
// within the attributes of the given Node
func findAttribute(n *html.Node, key string) string {
    if a := n.Attr; a != nil {
        for i := range a {
            if a[i].Key == key {
                return a[i].Val
            }
        }
    }
    return ""
}

// addElement adds a string to a given string slice if matches the
// given attribute within the given Node.
// If starts with slash or hashmeans that is a external element, so is appended to
// the root URL
// If it doesn't fit the previous case and it does not have a protocol prefix,
// then it is appended to the root URL as well
func addElement(elems []string, u *url.URL, n *html.Node, attr string) []string {
    if val := findAttribute(n, attr); val != "" {
        if strings.HasPrefix(val, "//") {
            val = u.Scheme + ":" + val
        } else if strings.HasPrefix(val, "/") || strings.HasPrefix(val, "#") {
            val = u.Scheme + "://" + u.Host + val
        } else if !hasProtocolAsPrefix(val) {
            val = u.Scheme + "://" + u.Host + "/" + val
        }

        elems = append(elems, val)
    }
    return elems
}


// newScraper parses the given url and scrapes the site, returning
// a scraper object with all the site info and meta tags
func NewScraper(u *url.URL, content string) (*Scraper, error) {
    var title string
    var language string
    var author string
    var description string
    var generator string
    var feed string
    var metaImage string
    charset := "utf-8"
    links := make([]string, 0)
    images := make([]string, 0)
    keywords := make([]string, 0)
    compatibility := make(map[string]string)

    scrpr := func(n *html.Node) {
        switch n.Data {
        case "html":
            language = findAttribute(n, "lang")
        case "title":
            title = n.FirstChild.Data
        case "a":
            links = addElement(links, u, n, "href")
        case "img":
            images = addElement(images, u, n, "src")
        case "link":
            typ := findAttribute(n, "type")
            switch typ {
            case "application/rss+xml":
                feed = findAttribute(n, "href")
            }
        case "meta":
            name := findAttribute(n, "name")
            switch name {
            case "author":
                author = findAttribute(n, "content")
            case "keywords":
                keywords = strings.Split(findAttribute(n, "content"), ", ")
            case "description":
                description = findAttribute(n, "content")
            case "generator":
                generator = findAttribute(n, "content")
            case "image":
                metaImage = findAttribute(n,"content")
            }

            httpEquiv := findAttribute(n, "http-equiv")
            switch httpEquiv {
            case "Content-Type":
                charset = findCharset(findAttribute(n, "content"))
            case "X-UA-Compatible":
                compatibility = mapifyStr(findAttribute(n, "content"))
            }
        }
    }
    save := &CrawlerSave{}
    save.Init()
    save.Url = u.String()
    tree, err := save.Parse(content)
    if err != nil {
        return nil, err
    }
    tree.Walk(scrpr)

    return &Scraper{title,
        language,
        author,
        description,
        metaImage,
        generator,
        feed,
        charset,
        links,
        images,
        keywords,
        compatibility}, nil
}

// hasProtocolAsPrefix returns true if the value belongs
// to some protocol
func hasProtocolAsPrefix(val string) bool {
    return strings.HasPrefix(val, "http://") ||
        strings.HasPrefix(val, "https://") ||
        strings.HasPrefix(val, "ftp://") ||
        strings.HasPrefix(val, "s3://")
}
