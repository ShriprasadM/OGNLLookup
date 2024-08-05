# OGNLLookup
Retrieve value from you Go Objects without JSON parser
1. ``` git clone --single-branch --branch examples_readme git@github.com:ShriprasadM/OGNLLookup.git ```
2. ``` cd OGNLLookup ```
4. ``` make ognl_generate ```
5. This will generate `student_ognl.go` next to `main.go`  
6. Ensure that following statement is added after `package` line in `student_ognl.go`
```go
import (
	"OGNLLookup/examples"
	"fmt"
)
```
7. Now OGNL Library is ready for demo
8. Open main.go and see all compilation errors are resolved
9. ``` make demo ```
10. This will show `English` as output. see main.go for more details on example
