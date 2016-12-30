import yaml
import os
print os.environ['HOME']
version = os.environ["BUILD_VERSION"][1:]

home = os.path.expanduser("~")
filename = home+"/devctl.github.io/_data/sha.yml"

f = open(filename, "r")
sha = yaml.safe_load(f)
f.close

sha["darwin_amd64"][version] = os.environ["DARWIN_AMD64_MD5"]
sha["linux_amd64"][version] = os.environ["LINUX_AMD64_MD5"]

with open(filename, 'w') as outfile:
    yaml.dump(sha, outfile, default_flow_style=False, indent=2)

filename = home+"/devctl.github.io/_data/version.yml"
f = open(filename, "r")
version_yml = yaml.safe_load(f)
f.close

version_yml["latest"] = version

with open(filename, 'w') as outfile:
  yaml.dump(version_yml, outfile, default_flow_style=False, indent=2)