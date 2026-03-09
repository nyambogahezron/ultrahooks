package hooks

// HookTemplate is the shell script template for installed hooks
const HookTemplate = `#!/bin/sh
ultrahooks run %s
`
