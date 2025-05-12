### ✅ Version using a struct
```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os/exec"
)

type TerraformOutput struct {
    EC2InstanceID struct {
        Value string `json:"value"`
    } `json:"ec2_instance_id"`
}

func main() {
    // Run Terraform output -json
    cmd := exec.Command("terraform", "output", "-json")
    output, err := cmd.Output()
    if err != nil {
        log.Fatal(err)
    }

    // Parse JSON into struct
    var tfOutput TerraformOutput
    err = json.Unmarshal(output, &tfOutput)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("EC2 Instance ID:", tfOutput.EC2InstanceID.Value)
}
```
✅ Benefits:
- No type assertions: safer at runtime, avoids panics.
- Compiler-checked fields: ensures correctness at compile time.
- Cleaner code: easier to read and understand.



# 📦 1. Package Declaration

```go
package main
```
This declares the package name as `main`. The `main` package is the starting point for a Go program. When you run a Go program, the `main` function in the `main` package is the entry point.

---

# 📥 2. Import Statements

```go
import (
    "encoding/json"
    "fmt"
    "log"
    "os/exec"
)
```
Here, we're importing the necessary Go packages:

- 📄 `"encoding/json"`: This package is used to work with JSON data. It helps to parse JSON (e.g., from Terraform output) into Go structures and also to encode Go objects into JSON format.
- 🖨️ `"fmt"`: This package provides formatted I/O functions (e.g., printing to the console with `fmt.Println()`).
- 🪵 `"log"`: The log package provides logging functions. In this case, we use `log.Fatal()` to log an error and exit the program if something goes wrong.
- 💻 `"os/exec"`: This package provides functions for running external commands (e.g., running `terraform output -json`).

---

# 🧱 3. Defining the Struct

```go
type TerraformOutput struct {
    EC2InstanceID struct {
        Value string `json:"value"`
    } `json:"ec2_instance_id"`
}
```
- We’re defining a Go struct called `TerraformOutput` that corresponds to the JSON structure returned by Terraform.
- The `EC2InstanceID` field maps to the `"ec2_instance_id"` key, and its nested `Value` field maps to the `"value"` inside it.

---

# 🏁 4. Main Function

```go
func main() {
```
The `main` function is where the program starts executing.

---

# ⚙️ 5. Run Terraform Command

```go
cmd := exec.Command("terraform", "output", "-json")
output, err := cmd.Output()
if err != nil {
    log.Fatal(err)
}
```
- `exec.Command("terraform", "output", "-json")`: Runs the `terraform output -json` command.
- `cmd.Output()`: Executes the command and stores the output.
- If there's an error, `log.Fatal(err)` logs it and stops the program.

---

# 📦 6. Parse JSON into Struct

```go
var tfOutput TerraformOutput
err = json.Unmarshal(output, &tfOutput)
if err != nil {
    log.Fatal(err)
}
```
- Parses the JSON output from Terraform into the Go struct using `json.Unmarshal()`.
- If parsing fails, the program logs the error and exits.

---

# 🖨️ 7. Print EC2 Instance ID

```go
fmt.Println("EC2 Instance ID:", tfOutput.EC2InstanceID.Value)
```
Prints the EC2 instance ID extracted from the Terraform output.

---

# 🧠 Summary

- 📥 Runs the `terraform output -json` command.
- 📦 Parses the JSON into a Go struct.
- 🖨️ Prints the EC2 instance ID from the structured data.

---

# 🤔 Why Use a Nested Struct?

Terraform’s JSON output looks like this:

```json
{
  "ec2_instance_id": {
    "value": "i-0abcdef1234567890"
  }
}
```

To map this structure in Go:

```go
type TerraformOutput struct {
    EC2InstanceID struct {
        Value string `json:"value"`
    } `json:"ec2_instance_id"`
}
```

- `TerraformOutput` represents the top-level JSON object.
- `EC2InstanceID` is a nested struct that models the `"ec2_instance_id"` object.
- Inside it, `Value` maps to the `"value"` key.

✅ Using nested structs allows Go to directly reflect the structure of the JSON and access values like this:

```go
fmt.Println("EC2 Instance ID:", tfOutput.EC2InstanceID.Value)
```

This helps in writing clean and strongly-typed Go code to handle hierarchical data.

---

