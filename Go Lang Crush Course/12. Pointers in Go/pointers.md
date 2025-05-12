📂 Mounting C: Drive in WSL (Windows Subsystem for Linux)
When you mount your Windows C: drive into WSL (like /mnt/c/):

You are not copying C drive into Linux.

You are not moving the files into Linux either.

You are simply creating a pointer (shortcut) to C drive.

When you change files inside /mnt/c/Users/..., you are directly modifying the original files in Windows C: drive.

✅ So: No copy, no move, just pointer.

🔥 Same idea with Pointers in Go:
When you use:

go
Copy
Edit
func handler(c *gin.Context)
You are not copying the Context struct (which can be huge).

You are not creating a new Context.

You are simply pointing to the original Context for that HTTP request.

Any change you make to c inside the function will affect the real request/response.

✅ Again: No copy, no move, just pointer.

🧠 Memory/Storage Visual:

Example	Explanation
WSL Mount	/mnt/c/ → Just pointing to Windows C: drive
Pointer *gin.Context	Just pointing to real Context in memory
🔵 Mount/Pointer = lightweight shortcut → Access and modify real resource.

📢 Final words:
✅ Pointer in Go = Mounting C drive in WSL.
✅ Both mean "access and modify the real thing, without copying."