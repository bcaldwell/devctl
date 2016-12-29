import yaml
import os
print os.environ['HOME']
version = os.environ["BUILD_VERSION"][1:]

home = os.path.expanduser("~")
filename = home+"/devctl.github.io/_data/sha.yml"

f = open(filename, "r")
sha = yaml.safe_load(f)
f.close
print sha

sha["darwin_amd64"][version] = os.environ["DARWIN_AMD64_MD5"]
sha["linux_amd64"][version] = os.environ["LINUX_AMD64_MD5"]

with open(filename, 'w') as outfile:
    yaml.dump(sha, outfile, default_flow_style=True)