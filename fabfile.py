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

DEFAULT_GOPATH = '/home/ubuntu/go'
DEFAULT_PATH = os.path.join(DEFAULT_GOPATH, 'src/github.com/aodin/groupthink/')
DEFAULT_DIR = os.path.dirname(os.path.dirname(DEFAULT_PATH))
DEFAULT_VERSION = "1.4.2"

# Output defaults for testing
def main():
    print(DEFAULT_GOPATH)
    print(DEFAULT_PATH)
    print(DEFAULT_DIR)
    print(DEFAULT_VERSION)


def link_service(path=DEFAULT_PATH):
    cmd = os.path.join(path, 'cmd', 'anaheim.conf')
    print(blue(cmd))
    dest = '/etc/init/anaheim.conf'

    with settings(warn_only=True):
        sudo('rm {}'.format(dest))
    sudo('ln -s {} {}'.format(cmd, dest))
    sudo('initctl reload-configuration')


def link_nginx(path=DEFAULT_PATH):
    nginx = os.path.join(path, 'cmd', 'nginx', 'anaheim')
    print(blue(nginx))
    with settings(warn_only=True):
        sudo('rm /etc/nginx/sites-enabled/anaheim')
    sudo('ln -s {} /etc/nginx/sites-enabled/'.format(nginx))


def restart_nginx():
    sudo('nginx -t')
    sudo('service nginx restart')


def build(path=DEFAULT_PATH):
    with cd(path):
        run('godep go build anaheim.go')
        run('grunt build')


def deploy():
    """
    Deploy anaheim.
    """
    install_packages()
    install_go()
    install_go_binaries()

    # Get sensitive user and password info
    smtp_user, smtp_pass, db_pass = prompt_for_passwords()

    install_postgres(db_pass)

    git_clone()
    update_settings(smtp_user, smtp_pass, db_pass)
    goose_up()

    install_npm()
    link_service()
    link_nginx()

    build()
    restart_nginx()

    sudo('service anaheim start')

    # create a link in ~ for fun
    run('ln -s {} ~/.'.format(DEFAULT_PATH))

    
def create_user(first, last, email, is_admin=True):
    password = prompt('password?', validate=r'^.+$')
    path = os.path.join(DEFAULT_PATH, 'cmd')
    with cd(path):
        flags = '-first="{}" -last="{}" -email="{}" -admin={} -p="{}"'.format(
            first,
            last,
            email,
            is_admin,
            password,
        )
        run('godep go run create_user.go {}'.format(flags))


# TODO Other commands
# --------------
# Get upstart logs: cat /var/log/upstart/anaheim.log


def update(goose=False):
    """
    Update anaheim.

    Examples:
    fab -H anahiem update
    fab -H anaheim update:goose=True
    """
    with cd(DEFAULT_PATH):
        # TODO Force over-writes?
        run('git pull origin master')
        build()

        if goose:
            goose_up()

    with settings(warn_only=True):
        sudo('service anaheim restart')


if __name__ == '__main__':
    main()
