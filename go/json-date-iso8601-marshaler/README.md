# JSON / Date (ISO 8601) Marshal/Unmarshal Example

The goal of this example is to show how to marshal/unmarshal a
time.Time value as part of a JSON record and to ensure it is stored in
ISO-8601 (~aka [RFC 3339](https://tools.ietf.org/html/rfc3339)) format
using UTC time (aka Zulu).

The solution includes declaring a new type, `Timestamp` and then
implementing `MarshalText()` on that type.  The `Timestamp` also
implements Stringer as a convenience and a constructor method `Now()`

As `time.Time`'s `UnmarshalText()`, as well as `UnmarshalJSON()`,
expect RFC 3339 format the underlying `time.Time` methods work with
out extension.

To demonstrate the code:

```
> go test
```
