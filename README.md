# CPU Load Generator

This is a very simple CPU load generator. It generates load by calculating prime numbers (and discarding the results).

Sure, there are many other purpose built load generators. However, many/most seem to have a larger purpose than simply generating CPU load and I wanted a more simple solution.

## Installation

Install one of the released packages or clone this repository and either build your own packages or run directly from the repository root.

### Debian/Ubuntu

```sh
> sudo dpkg -i load-cpu-go-<version>.deb
```

### Red Hat

```sh
> sudo rpm -ivh load-cpu-go-<version>.rpm
```

## Usage

```sh
> load-cpu
System load: 1.33, 1.89, 2.12
System load: 3.50, 2.35, 2.26
```

By default, the tool will generate high CPU load for 10 seconds. If you wish to continue generating load for longer, pass the `-duration <seconds>` flag.

```sh
> load-cpu -duration 60
System load: 1.80, 1.72, 1.85
System load: 6.03, 2.96, 2.27
```
