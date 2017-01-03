#!/usr/bin/python
# pylint: disable=C0103
# pylint: disable=C0301

import os
import hashlib
import functools
import shutil
import yaml

def sha256_file(file_path, chunk_size=65336):
    """
    Get the sha 256 checksum of a file.
    :param file_path: path to file
    :type file_path: unicode or str
    :param chunk_size: number of bytes to read in each iteration. Must be > 0.
    :type chunk_size: int
    :return: sha 256 checksum of file
    :rtype : str
    """
    # Read the file in small pieces, so as to prevent failures to read particularly large files.
    # Also ensures memory usage is kept to a minimum. Testing shows default is a pretty good size.
    assert isinstance(chunk_size, int) and chunk_size > 0
    digest = hashlib.sha256()
    with open(file_path, 'rb') as f:
        [digest.update(chunk) for chunk in iter(functools.partial(f.read, chunk_size), '')]
    return digest.hexdigest()

def makedir(directory):
    if not os.path.exists(directory):
        os.makedirs(directory)

builds = os.environ["BUILDS"].split(" ")
print(builds)

devctlDir = os.getcwd().split("/")

if devctlDir[-1] != "devctl":
    for i, e in reversed(list(enumerate(devctlDir))):
        if e == "devctl":
            devctlDir = devctlDir[:i+1]

devctlDir = "/".join(devctlDir)

# create required directories
makedir(os.path.join(devctlDir, "dist", "release"))

newSha = {}

builds = [x.replace("/", "_") for x in builds]

for build in builds:
    print build
    build = build.replace("/", "_")
    tar_file = os.path.join(devctlDir, "dist", "release", "devctl_"+build)
    working_directory = os.path.join(devctlDir, "dist", build, "devctl")
    makedir(working_directory)

    shutil.copyfile(os.path.join(devctlDir, "devctl.sh"), os.path.join(working_directory, "devctl.sh"))
    shutil.move(os.path.join(devctlDir, "dist", "devctl_" + build + "_bin"), os.path.join(working_directory, "devctl"))

    shutil.make_archive(tar_file, "gztar", working_directory)

    newSha[build] = sha256_file(tar_file+".tar.gz")

version = os.environ["BUILD_VERSION"][1:]

home = os.path.expanduser("~")
filename = home+"/devctl.github.io/_data/sha.yml"

f = open(filename, "r")
sha = yaml.safe_load(f)
f.close

for build in builds:
    sha[build][version] = newSha[build]

with open(filename, 'w') as outfile:
  yaml.dump(sha, outfile, default_flow_style=False, indent=2)

filename = home+"/devctl.github.io/_data/version.yml"
f = open(filename, "r")
version_yml = yaml.safe_load(f)
f.close

version_yml["latest"] = version

with open(filename, 'w') as outfile:
  yaml.dump(version_yml, outfile, default_flow_style=False, indent=2)