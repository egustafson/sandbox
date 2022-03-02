zerolog-client prototype client lib logging
-------------------------------------------

Loggers, including zerolog, provide a global logger.  Using this
logger is likely to conflict with application (re)configuration of the
same logger when the same log library is used.

This example attempts to build a client logging package that allocates
and manages a zerolog logger that is both isolated from an application
logger, and isolates its configuration from application configuration.

Ref: [zerolog](https://github.com/rs/zerolog)
