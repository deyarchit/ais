# Fix "bind: address already in use" — Go HTTP Server on macOS

This error means another process is using the port your server is trying to bind to.

## Step 1: Find the process using the port

```bash
sudo lsof -i :<PORT_NUMBER>
```

Example for port 8080:

```bash
sudo lsof -i :8080
```

Sample output:
```
COMMAND     PID      USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
my_server   12345    user    6u  IPv4 0xabcdef1234       0t0  TCP *:8080 (LISTEN)
```

Note the **PID** (12345 in this example).

## Step 2: Kill the process

```bash
kill <PID>
```

If it doesn't terminate:

```bash
kill -9 <PID>
```

> Use `kill -9` cautiously — it force-kills without cleanup.

## Step 3: Verify the port is free

```bash
sudo lsof -i :8080
```

No output = port is free.

## One-liner (find and kill in one step)

```bash
kill $(lsof -ti :8080)
```

## Additional notes

- **TIME_WAIT state** — Even after closing, a port may stay in TIME_WAIT for a few minutes. This is normal TCP behavior.
- **Go's `net` package** — Go sets `SO_REUSEADDR` automatically (since Go 1.13 on macOS/Darwin), which mitigates most TIME_WAIT issues.
- **`netstat` alternative** — `netstat -an | grep <PORT>` also works but `lsof` is generally preferred on macOS.

*(Sources retrieved via `ais` CLI — see sources.txt)*
