Here are clean and comprehensive notes on nested structs in Go using Terraform JSON output examples. These notes explain single vs multiple nested structs, json:"..." struct tags, and how json.Unmarshal works.

🧱 What Is a Nested Struct in Go?
A nested struct is a struct defined inside another struct. It's useful for mapping JSON where one key maps to another object.

✅ Single Nested Struct
🧩 JSON Example:
json
Copy
Edit
{
  "ec2_instance_id": {
    "value": "i-0abcdef1234567890"
  }
}
📦 Go Struct (Inline/Nested):
go
Copy
Edit
type TerraformOutput struct {
    EC2InstanceID struct {
        Value string `json:"value"` // maps inner "value" key
    } `json:"ec2_instance_id"` // maps outer "ec2_instance_id" key
}
✅ Benefits:
No extra type name needed for small, one-time-use structures.

Cleaner and concise.

🔍 Access After Unmarshal:
go
Copy
Edit
fmt.Println(output.EC2InstanceID.Value)
✅ Multiple Nested Structs
🧩 JSON Example:
json
Copy
Edit
{
  "ec2_instance_id": {
    "value": "i-0abcdef1234567890"
  },
  "instance_type": {
    "value": "t2.micro"
  },
  "region": {
    "value": "us-east-1"
  }
}
📦 Go Struct (Multiple Nested Structs):
go
Copy
Edit
type TerraformOutput struct {
    EC2InstanceID struct {
        Value string `json:"value"`
    } `json:"ec2_instance_id"`

    InstanceType struct {
        Value string `json:"value"`
    } `json:"instance_type"`

    Region struct {
        Value string `json:"value"`
    } `json:"region"`
}
🔍 Access:
go
Copy
Edit
fmt.Println(output.EC2InstanceID.Value)
fmt.Println(output.InstanceType.Value)
fmt.Println(output.Region.Value)
🔁 Reusable Struct Type
If the nested structure is the same for all keys, use a reusable type:

go
Copy
Edit
type OutputValue struct {
    Value string `json:"value"`
}

type TerraformOutput struct {
    EC2InstanceID OutputValue `json:"ec2_instance_id"`
    InstanceType  OutputValue `json:"instance_type"`
    Region        OutputValue `json:"region"`
}
🧠 Why? Avoids repeating struct { Value string } over and over.

🧠 How json.Unmarshal Works
go
Copy
Edit
jsonData := `{"ec2_instance_id":{"value":"i-0abc"}}`

var output TerraformOutput
err := json.Unmarshal([]byte(jsonData), &output)
🔄 What Happens:
json.Unmarshal reads the JSON string.

Matches JSON keys to Go struct fields using:

Field names.

Struct tags like json:"ec2_instance_id".

Converts the string into Go values and fills the struct.

🏷️ What Is json:"..."?
Struct Tag

Tells Go which JSON key to match for a field.

go
Copy
Edit
Value string `json:"value"` // Maps to "value" in JSON
⛔ If omitted: Go expects the JSON key to be exactly Value, which won't work if your JSON uses lowercase like "value".

✅ When to Use Inline Nested Structs
Use when:

Structure is simple.

Used only once.

You want to keep things concise.

❌ When to Avoid Inline Structs
Avoid if:

The nested struct is large, complex, or reused.

You want testable, modular code.