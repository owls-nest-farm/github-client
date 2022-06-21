---
- organization: owls-nest-farm
  repositories:
    {{ range . }}
    - name: {{ .Name }}
      private: {{ .Private }}
      visibility: {{ .Visibility }}
      tpl_name: {{ .TplName }}
      archived: {{ .Archived }}
    {{ end }}

