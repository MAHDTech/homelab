#!jinja|yaml
---

Install SSH Key:
  ssh_auth.present:
    - user: root
    - source: salt://pki/salt-ssh.rsa.pub
