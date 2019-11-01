## ktunnel server

Run the ktunnel server(from remote - usually k8s pod)

### Synopsis

This command would start the tunnel server wait for tunnel clients to bind

```
ktunnel server [flags]
```

### Options

```
      --cert string   tls certificate file
  -h, --help          help for server
      --key string    tls key file
```

### Options inherited from parent commands

```
  -p, --port int   author name for copyright attribution (default 28688)
  -t, --tls        Connection uses TLS if true, else plain TCP
```

### SEE ALSO

* [ktunnel](ktunnel.md)	 - Ktunnel is a network tunneling tool for kubernetes

###### Auto generated by spf13/cobra on 1-Nov-2019