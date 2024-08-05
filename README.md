# OGNLLookup
Retrieve value from you Go Objects without JSON parser

1. ```sh 
   make ognl_generate 
   ```
2. This will generate `student_ognl.go` next to `main.go`  
3. Ensure that following statement is added after `package` line in `student_ognl.go`
```go
import (
	"OGNLLookup/examples"
	"fmt"
)
```
4. Now OGNL Library is ready for demo
4. Open main.go and see all compilation errors are resolved
5. ```sh 
   make demo 
   ```
6. This will show `English` as output. see main.go for more details on example