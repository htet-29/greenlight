| Feature      | Idempotent                             | Non-Idempotent |
| :---------   | :------------------------------------- | :------------- |
|Definition    | Multiple requests = One request effect | Multiple requests = Multiple effects |
| State Change | Happens once                           | Happens every time |
| HTTP Methods | "GET, PUT, DELETE, HEAD"               | "POST, PATCH*" |
| Safety       | Generally safer for retries            | Dangerous to retry automatically |