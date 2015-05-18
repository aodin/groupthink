# -*- coding: utf-8 -*
from __future__ import print_function
import json
import os
import yaml
from StringIO import StringIO

from fabric.api import local, settings, abort, run, cd, env, get
from fabric.colors import blue
from fabric.contrib.console import confirm
from fabric.contrib.files import append
from fabric.operations import sudo, prompt

# Use the local SSH config file
env.use_ssh_config = True

DEFAULT_GOPATH = '/root/go'
DEFAULT_PATH = os.path.join(DEFAULT_GOPATH, 'src/github.com/aodin/groupthink/')
DEFAULT_DIR = os.path.dirname(os.path.dirname(DEFAULT_PATH))
DEFAULT_VERSION = '1.4.2'

# Output defaults for testing
def main():
    print(DEFAULT_GOPATH)
    print(DEFAULT_PATH)
    print(DEFAULT_DIR)
    print(DEFAULT_VERSION)


def restart_nginx():
    sudo('nginx -t')
    sudo('service nginx restart')


def update(fullpath='/root/go/src/github.com/aodin/groupthink'):
    """
    Update groupthink.

    Examples:
    fab -H aoe update
    """
    with cd(fullpath):
        run('git pull origin master')
        run('godep go build groupthink.go')

    with settings(warn_only=True):
        run('restart groupthink')


if __name__ == '__main__':
    main()
