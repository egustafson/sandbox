Environment Variable Decoration of `map[string]any`
===================================================

This prototype will walk an arbitrarally deep `map[string]any` where `any` may
be a `map[string]any` and identifies leaf `string`s and then attempts to process
the `string` value with a substitution table.  In practice, the substitution
table could be a map of environment variables.

This is a tree decoration demo.