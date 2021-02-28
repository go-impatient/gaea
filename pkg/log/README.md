# Log

## Usage

### Structured logging

```
logger := log.NewStdLogger(os.Stdout)
log := log.NewHelper("module_name", logger)
// Levels
log.Info("some log")
log.Infof("format %s", "some log")
log.Infow("field_name", "some log")
```

