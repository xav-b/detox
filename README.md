# Detox

## Philosophy

Command-line driven tools to fetch your digital stuff and store them in a way that makes it easy to:
- Interactively search or analyse
- Backup on personal hardware/NAS

This is intended for terminal power users but the process is designed to be straightforard and safe. One should feel confident enough the data won't be affected, so one can go ahead and trim all the junk on Google drive, Dropbox, ...

Indeed the completely intended side benefit is to reduce my online attack surface/privacy exposure and delete as much as possible, while still being able to archive extensively reliably on devices I own.

## Usage

Not usable at the moment :)

But if you insist: `go run . $HOME/Documents  && cat backup-meta.json`

## Design

Written in Go because:

a) Easy binary distribution
b) Leverage interfaces to plug different sources/exporters
c) Reliable
bonus) 'been a while and that's a cool language to get things done

## Roadmap (WIP)

**Sources**

- [x] Filesystem
- [ ] Gmail
- [ ] Google photos
- [ ] Whatsapp
- [ ] Github/Gitlab
- [ ] [Pocket](https://getpocket.com/)

_Note: I considere services like Google drive or Dropbox as unsafe but
convenient to share files. Anything uploaded there should already be a
duplicated version of what is available locally or backed up. Hence no need to
download it and back it up again._

**Exports**

- [x] JSON lines
- [ ] SQLite
- [ ] Elasticsearch

**Backups**

- [x] `cp -r Backup/ "/mnt/my-disk/my-backup-$(date '+%Y%m%dT%H%M%SZ')"`
- [ ] Incremental backup program ([Restic maybe](https://github.com/restic/restic))
- [ ] S3 / Digital Ocean Space / self-hosted Minio
- [ ] NAS
