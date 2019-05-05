#!/bin/bash

echo -e "\033[0;32mDeploying updates to GitHub...\033[0m"

msg="update files `date`"

git add -A .
git commit -m "$msg"
git push origin master