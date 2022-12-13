# UptimeBot

Log downtimes (based on no response or unexpected HTTP status code) and send webhook alerts if needed.

If the HTTP response code is one of the following then UptimBot follows the HTTP request and compares the final HTTP status code with the expected one provided in the config file.

```
301 (Moved Permanently)
302 (Found)
303 (See Other)
307 (Temporary Redirect)
308 (Permanent Redirect)
```

In most cases your expected status code will be 200 (HTTP Status OK).

**Instructions:** Compile using `make build`. Edit the [config file](/config.yaml) and run as a service.

Usage:

```
  -config string
        Path to your config file (default "./config.yaml")
```
