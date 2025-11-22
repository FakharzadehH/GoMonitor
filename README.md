# GoMonitor

GoMonitor is a lightweight monitoring tool written in Go. It’s designed as a practical project to explore Go architecture patterns, background job execution, and Prometheus metric exposure. While not intended to be a production-grade solution, it provides a clear and extendable foundation for experimenting with how monitoring systems can be built.

## Overview

The core idea behind GoMonitor is simple: run periodic jobs, collect data, and expose useful metrics through a `/metrics` endpoint. These metrics can be scraped by Prometheus, allowing you to visualize or alert on them using your existing monitoring stack.

The project also includes examples of common development and deployment tools such as Docker, docker-compose, and Kubernetes manifests. The focus is on readability and learning rather than completeness or enterprise complexity.

## Features

- Prometheus-friendly metrics endpoint  
- Simple `config.yaml` configuration  
- Example Prometheus config included  
- Dockerfile and docker-compose support  
- Exploratory Kubernetes deployment files  
- Clear folder structure following common Go conventions

## Project Structure

GoMonitor is organized to make the internal logic easy to understand and extend:

```
.
├── main.go
├── internal/ # configuration, helpers, and internal utilities
├── service/ # core monitoring logic
├── job/ # periodic jobs and data collectors
├── repository/ # optional storage layer components
├── k8s/ # basic Kubernetes manifests for experimentation
├── Dockerfile
├── docker-compose.yml
└── config.yaml.example
```

The codebase is intentionally small and approachable, making it useful for learning how to structure a Go project with multiple layers.

## Configuration

Configuration is done through a simple YAML file. A sample file is provided as `config.yaml.example`, which you can copy and adjust as needed. This file allows you to configure things like ports, job intervals, and general runtime behavior.

## Prometheus Integration

GoMonitor exposes its metrics through an HTTP endpoint compatible with Prometheus. You can add it to your Prometheus setup by pointing a scrape job to the service’s port. An example configuration file (`prometheus.yml`) is included to help you get started.

Once Prometheus begins scraping the service, the collected metrics can be visualized using Grafana or any other metrics dashboard you prefer.

## Purpose and Intent

This project sits somewhere between a learning exercise and a practical tool. It showcases:

- Organizing a Go application with multiple layers
- Working with background jobs and intervals
- Integrating Prometheus for metrics
- Using containerization and optional Kubernetes manifests
- Designing a small, clean codebase that’s easy to follow or extend

It’s not meant to be fully polished or feature-complete, but it provides a solid foundation for experimentation or further development.

## License

No explicit license is included. The project is available for personal use, learning, and modification.
