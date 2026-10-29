# NAME

{{ summary . }}

# USAGE

```text
{{ usage . | indent 4 }}
```

{{if .Cmds -}}
# COMMANDS

```text
{{ commands . | indent 4 }}
```

{{ end -}}
{{- if .Long -}}
# DESCRIPTION

{{ long . }}

{{ end }}

