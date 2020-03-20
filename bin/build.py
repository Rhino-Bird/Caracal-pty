#!/usr/bin/env python
#-*- coding:utf-8 -*-

import os, time, subprocess


def runCmd(cmd):
    p = subprocess.Popen(cmd, shell = True, stdout = subprocess.PIPE, stderr = subprocess.PIPE)
    stdout = p.communicate()[0].decode('utf-8').strip()
    return stdout


# Get last release tag.
def lastTag():
    return runCmd('git describe --abbrev=0 --tags')


# Get current branch name.
def branch():
    return runCmd('git rev-parse --abbrev-ref HEAD')


# Get last git commit id.
def lastCommitId():
    return runCmd('git log --pretty=format:"%h" -1')


# Assemble build command.
def buildCmd():
    buildFlag = []
    
    ver = lastTag()
    if ver != "":
        buildFlag.append("-X main.Version='{}'".format(ver))

    br = branch()
    if br != "":
        buildFlag.append("-X main.Branch='{}'".format(br))

    cmt = lastCommitId()
    if cmt != "":
        buildFlag.append("-X main.CommitID='{}'".format(cmt))
    

    # current time
    buildFlag.append("-X main.BuildTime='{}".format(time.strftime("%Y-%m-%d_%H:%M:%S")))
    
    return 'go build -ldflags "{}" -o caracal-pty cmd/main.go cmd/parse_init.go cmd/vargs.go cmd/help.go'.format(" ".join(buildFlag))


def main():
    if subprocess.call(buildCmd(), shell = True) != 0:
        print("build faild.")


main()
