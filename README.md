# Go Assessment : WebPage Link Crawler

You are tasked with created a web page link crawler program.
The application is given a `-url` and `-depth` flag and prints the list of Links returned by `CrawlWebpage`.

## Directions

- Crawl the web page starting at the given `-url` for any links within the current website.
- Follow links found on the page to discover any new links.
- Stop following links deeper than `-depth` levels.
  - For example: if the max depth is 2, and the main url is `http://example.com/` and there is a chain of links from `/` -> `/foo` -> `/bar` -> `/baz` -> `/quux`
    - depth 0: crawl `/` and discover the link to `/foo`
    - depth 1: crawl `/foo` and discover `/bar`
    - depth 2: crawl `/bar` and discover `/baz`
    - stop at depth 2, do not crawl `/baz` and therefor will not discover the link to `/quux`
    - Collect the links with their base url to produce a list:
      1. http://example.com/
      2. http://example.com/foo
      3. http://example.com/bar
      4. http://example.com/baz

Your solution should be implemented or called by the body of: 

```go
func CrawlWebpage(rootURL string, maxDepth int) ([]string, error) {
	//TODO: Implement Solution
	return nil, errors.New("solution not implemented")
}
```

However, you are encouraged to create any additional functions, structs, or packages you deem appropriate.

## Assessment Criteria

1. Correctness - Does the program produce the right outputs for the given inputs.
2. Simplicity - Does the program solve the problem without unnecessary complexity?
3. Maintainability - How easy it is to read, understand, and change?
4. Performance - Is the solution implemented optimally.
5. Style - Is the code idiomatic Go? Was care given to code style such as names of variables, functions, types, etc.? See: [Google Go Style-Guide](https://google.github.io/styleguide/go/)
