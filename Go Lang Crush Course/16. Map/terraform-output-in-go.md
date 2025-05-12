# 🧩 Step-by-Step: Parse Terraform Output in Go

This guide shows how to extract an EC2 instance ID from `terraform output -json` using Go — and draws comparisons with Python for clarity.

---

## ✅ Step 1: Run the Terraform Command

```go
cmd := exec.Command("terraform", "output", "-json")
output, err := cmd.Output()
```

📦 `output` is a byte slice (`[]byte`) containing the raw JSON text.

### 🔍 Example Output JSON

```json
{
  "ec2_instance_id": {
    "sensitive": false,
    "type": "string",
    "value": "i-0abcdef1234567890"
  }
}
```

---

## ✅ Step 2: Parse the JSON

```go
var outputData map[string]interface{}
err = json.Unmarshal(output, &outputData)
```

📘 Now `outputData` becomes a Go map mirroring the structure of the JSON.

```go
outputData["ec2_instance_id"]
// ➜ map[string]interface{}{
//     "sensitive": false,
//     "type": "string",
//     "value": "i-0abcdef1234567890",
// }
```

---

## ✅ Step 3: Extract the EC2 Instance ID

```go
instanceID := outputData["ec2_instance_id"].(map[string]interface{})["value"].(string)
fmt.Println(instanceID) // ➜ i-0abcdef1234567890
```

📌 You're using two type assertions:
1. Assert `ec2_instance_id` is a map
2. Assert `value` is a string

---

## 🔄 Python Equivalent

```python
import json, subprocess

output = subprocess.check_output(["terraform", "output", "-json"])
output_json = json.loads(output)

instance_id = output_json["ec2_instance_id"]["value"]
```

🔸 Same logic, but Python dynamically handles types, so no assertions are needed.

---

## 💡 Go Type Assertion: Why It’s Needed

Go is statically typed. `interface{}` can mean _anything_, so you must **explicitly assert** types:

```go
value := myMap["key"].(ExpectedType)
```

### 📘 Example:

```json
{
  "name": "Nana",
  "age": 30
}
```

Go version with type assertions:

```go
data := map[string]interface{}{
    "name": "Nana",
    "age": 30,
}

name := data["name"].(string) // ✅ must assert
age := data["age"].(int)      // ✅ must assert
```

---

## ✅ No Assertion Needed for Known Types

If your map isn’t `interface{}`:

```go
myMap := map[string]string{
    "name": "Nana",
}

name := myMap["name"] // ✅ No type assertion needed
```

🔑 **Rule of Thumb:**
- ✅ Use type assertion only with `interface{}`.
- ❌ No need for type assertion with `map[string]string`, `map[string]int`, etc.

---

Would you like a version using Go structs for cleaner and safer parsing?