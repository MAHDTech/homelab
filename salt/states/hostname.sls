#!jinja|yaml
---

{% set hostname = salt['pillar.get']('network/hostname', 'unknown-host') %}

Manage file /etc/sysconfig/network:
  file.managed:
    - contents: |
      NETWORKING=yes
      HOSTNAME={{ hostname }}

#Update etc hostname:
#  file.managed:
#    - contents: {{ hostname }}
#
#Update etc hosts:
#  file.replace:
#    - pattern: '127\.0\.0\.1.*'
#    - repl: '127.0.0.1   localhost {{ hostname }}'
#