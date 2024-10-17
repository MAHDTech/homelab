# Usage

## Salt SSH

- First, ensure that SSH works.

```bash
# Test SSH
ssh <username>@<hostname>
```

- Next, show the hosts configured in the roster.

```bash
# Show hosts
salt-ssh --hosts
```

- Next, test the connection and deploy the SSH key.

```bash
# Test
salt-ssh '*' test.ping
```

- Next, deploy the SSH key:

```bash
# Deploy the SSH key
salt-ssh '*' saltutil.send_key
```

- Next, apply the states:

```bash
# Apply state
salt-ssh '*' state.apply test=True
```

- Finally, handover to the Salt Master.
