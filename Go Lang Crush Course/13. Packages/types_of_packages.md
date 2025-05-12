 Go, packages are typically categorized into three main types based on their origin:
✅ 1. Standard Library Packages

    Description: Built-in packages provided by Go.

    Imported like: import "fmt"

    Examples: fmt, log, os, net/http, strings, io, etc.

    No need to install — they come with Go.

🧩 2. Third-Party (External) Packages

    Description: Packages created and published by the community (usually hosted on GitHub or other public sources).

    Imported like:
    import "github.com/joho/godotenv"

    You need to install them using go get or with a go mod file.

    Example Use Case: github.com/joho/godotenv to load .env files.

    These packages are external dependencies and get downloaded into your go.mod / go.sum files.

🏗️ 3. Custom (Local) Packages

    Description: Packages you write yourself inside your project.

    Imported like:
    import "yourmodule/config"

    Defined in directories like config/, utils/, etc. within your project.

    Organized under a Go module you define with go mod init yourmodule.

🔍 Summary Table
Type	Example	Where It Comes From	Import Path Example
Standard	fmt, os, log	Comes with Go	"fmt"
Third-Party	github.com/joho/godotenv	From GitHub or external repo	"github.com/joho/godotenv"
Custom/Local	config, utils	You define it in your code	"yourmodule/config"
