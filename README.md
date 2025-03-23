# Cortex
The write-now organize-later terminal notes app

```bash
$ cortex todo -m "I need to do something" -p 0
```

---
Goals
1. Toss notes immediately into some bin
2. At some interval, go through these notes and delete or categorize them
3. Full text search my notes

Features:
1. Interactive, prioritized TODO list
2. Write-now organize-later terminal notes app
3. Key value store for frequently used passwords/tokens

```bash
$ cortex notes new
```
Creates a new file, opens vim

```
cortex notes transform summarize
cortex notes transform summarize uz
```
defaults to the last open note. cleans up, summarizes with llm. Should always
keep the original data

```bash
$ cortex notes clean
```
manually sort through notes, delete where needed, merge, etc.

```
$ cortex templates launch tsconjure 
```
launch template in staging, can copy over to a real dir

```
$ cortex notes search "Inbound message listener"
```

```
$ cortex notes edit
$ cortex notes edit -n uz
```

```bash
$ cortex assets add mydiagram.excalidraw
```
actually this should be part of the notes

```
$ cortex todo new -m "Email michal about something" -p 0
```

```
$ cortex todo list
```

```bash
$ cortex enc get connectivity-token
```
pgp encrypted with yubikey