#!jinja|yaml
---

Configure sshd_config:
  file.replace:
    - name: /etc/ssh/sshd_config
    - pattern: '^#Subsystem\s+sftp\s+/usr/libexec/openssh/sftp-server'
    - repl: 'Subsystem      sftp    /usr/libexec/openssh/sftp-server'

Restart sshd:
  service.running:
    - name: sshd
    - enable: True
    - reload: True
    - watch:
      - file: ensure_sftp_subsystem

Update MOTD:
  file.managed:
    - name: /etc/motd
    - contents: |
        ***********************************************
        *  WARNING: This system is managed by Salt   *
        ***********************************************
