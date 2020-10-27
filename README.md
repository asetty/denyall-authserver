# "deny all" External Auth Server

A simple auth server that returns 403 for all requests except those to the health endpoint, "/healthz".

This can be used for the Ambassador [AuthService Plugin](https://www.getambassador.io/docs/latest/topics/running/services/auth-service/), to block requests matching mappings without "bypass_auth: true".


