#!/usr/bin/env bash
cd `dirname $0`
echo "compiling" $1
wine metaeditor.exe /compile:MQL4/Experts/$1 /log
