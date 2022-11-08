## Example

```go
import htmltomd "github.com/XRSec/HTML-TO-MARKDOWN"

body := GetSource(url, "", false)
HtmlToMarkdown(body, output, true)
```

```bash
go run awvs.go -u https://baidu.com -o 123.md
```

```bash
htmltomd -u https://baidu.com -o 123.md
```

## Build

[Github Action](.github/workflows/Go%20Build.yml#L35)

```bash
go build '--ldflags= -s -w -X main.buildTime=2022-04-28/09:26:31 -X main.versionData=0.0.1 -X main.commitId=e5f941ddac24e5177650f6038f38d5935be90921 -X main.author=XRSec'
```
