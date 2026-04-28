# Fixing 'bind: address already in use' on macOS

When your Go HTTP server fails to start with `bind: address already in use`, it means another process is already listening on the port you're trying to bind. Here is how to identify and kill that process.

## Step 1: Find the process using the port

### Option A: Using `lsof` (most common on macOS)

```bash
lsof -i :<port>
```

For example, if your server uses port 8080:

```bash
lsof -i :8080
```

Sample output:
```
COMMAND   PID    USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
main     1234 youruser   3u  IPv6 0x...           0t0  TCP *:http-alt (LISTEN)
```

The `PID` column contains the process ID you need.

### Option B: Using `netstat`

```bash
netstat -anv | grep LISTEN | grep <port>
```

For example:
```bash
netstat -anv | grep LISTEN | grep 8080
```

### Option C: Using `ss` (if available)

```bash
ss -tlnp | grep <port>
```

## Step 2: Kill the process

Once you have the PID from the output above:

```bash
kill <PID>
```

For example:
```bash
kill 1234
```

If the process does not terminate, use a stronger signal:

```bash
kill -9 <PID>
```

## One-liner to find and kill in a single command

```bash
lsof -ti :<port> | xargs kill
```

The `-t` flag tells `lsof` to output only the PID, making it easy to pipe directly to `kill`.

If a graceful kill is not enough:

```bash
lsof -ti :<port> | xargs kill -9
```

## Example for port 8080

```bash
# Find and kill gracefully
lsof -ti :8080 | xargs kill

# Find and kill forcefully
lsof -ti :8080 | xargs kill -9
```

## Preventing the issue in Go

You can also configure your Go server to reuse the address by setting `SO_REUSEPORT` or `SO_REUSEADDR` on the listener, but the most practical fix during development is to kill the stale process as shown above.

A common cause is a previous instance of your own server that crashed or was not properly shut down. After killing it, re-running your Go program should succeed.
