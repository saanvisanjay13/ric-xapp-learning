#!/bin/bash
echo "Starting backup..."
mkdir -p ~/backups
cp hello.sh ~/backups/hello_backup.sh
echo "Backup complete! Files saved to ~/backups"
ls ~/backups

