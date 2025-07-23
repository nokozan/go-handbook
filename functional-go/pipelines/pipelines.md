# Pipelines

A **fan-out pipeline** distributes a single input to multiple processors in parallel.  
Each processor can transform or analyze the input independently.

In Go, we combine **Higher-Order Functions** and **goroutines** to create this pattern.

---

## What is inside (categories)

| Type                  | Description                                  |
|-----------------------|----------------------------------------------|
| `Processor` function  | A simple function: `func(string) string`     |
| `FanOutPipeline`      | HOF that runs multiple processors in parallel|
| `FanInPipeline`       | (Optional extension) Collects outputs into unified result |

## What I can do (benefits, use cases)

- Run **multiple transformations or checks** on a single piece of data
- Increase **parallel throughput** using goroutines
- Decouple transformation logic via **simple, testable functions**

### Real-world use cases
- Process logs with multiple exporters (e.g., Grafana Loki + Elasticsearch)
- Parse blockchain transactions in multiple formats
- Run multiple validators or enrichers on the same data
- Extract multiple insights from an event (e.g., tags, metrics, geo-IP)
