# fuhrmannb dev.to blog

This repository contains source code of all blog post of https://dev.to/fuhrmannb/

## Create a new post

After setting your token in dev.to `$DEVTO_TOKEN`:

```shell
dagger call --dev-to-token=env:DEVTO_TOKEN new-post --post-name 'My Article title' export --path=./posts
```

## Sync posts to dev.to

```shell
dagger call sync-posts
```

Posts are automatically synced at each commit in `main` branch.

## License

All posts are licensed under [CC BY-NC-SA 4.0](LICENSE).
