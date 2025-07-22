# Higher-Order Functions in Go

- Takes one or more **functions as parameters** 
- Returns a **function as its result** 

## What is Inside(Categories of HOF Usage)

| Sub-feature         | Pattern                     | Description                                      |
|---------------------|-----------------------------|--------------------------------------------------|
| `task-wrappers`     | HOF + Decorators            | Retry, Timeout, Hooks — wrap logic dynamically   |
| `closure-state`     | HOF + Closure with memory   | Return functions that retain internal state      |
| `pipeline-fanout`   | HOF + Composition           | Pass data through a series of transforms         |
| `fan-in/fan-out`    | HOF + Concurrency Pattern   | Combine/distribute channels via function chains  |
| `custom-executor`   | HOF + Strategy injection    | Inject runtime behavior into task runners        |

---

## 3. What I can do (benefits, use cases)

-  Add features like retries, hooks, deadlines **without changing original logic**
-  Create composable, reusable wrappers around any unit of work
-  Retain execution context or retry counters using closures
-  Chain execution logic in functional-style pipelines
-  Inject different strategies for execution, timeout, logging, tracing

### Real-world use cases
- Background workers and task processors
- RPC or HTTP client retry logic
- Metrics + logging middleware
- Rate limiters or circuit breakers
- Fan-in / fan-out orchestration for pipelines